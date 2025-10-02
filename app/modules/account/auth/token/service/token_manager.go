package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	token "clinic-vet-api/app/modules/account/auth/token/factory"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	infraerr "clinic-vet-api/app/shared/error/infrastructure"

	"github.com/redis/go-redis/v9"
)

type TokenManager struct {
	factory     *token.TokenFactory
	redisClient *redis.Client
}

// StoredToken represents a token stored in Redis
type StoredToken struct {
	TokenString string       `json:"token_string"`
	TokenType   vo.TokenType `json:"token_type"`
	UserID      string       `json:"user_id"`
	ExpiresAt   time.Time    `json:"expires_at"`
}

func NewTokenManager(jwtSecret []byte, redisClient *redis.Client) *TokenManager {
	return &TokenManager{
		factory:     token.NewTokenFactory(jwtSecret),
		redisClient: redisClient,
	}
}

// tokenKey returns the Redis key for a single token
// Example: "token:user123_OTP_1696089600"
func (tm *TokenManager) tokenKey(userID string, tokenType vo.TokenType, timestamp int64) string {
	return fmt.Sprintf("token:%s_%s_%d", userID, tokenType, timestamp)
}

// userTokensKey returns the Redis key for a set of tokens for a specific user
// Example: "user_tokens:user123:OTP"
func (tm *TokenManager) userTokensKey(userID string, tokenType vo.TokenType) string {
	return fmt.Sprintf("user_tokens:%s:%s", userID, tokenType)
}

func (tm *TokenManager) GenerateToken(ctx context.Context, tokenType vo.TokenType, config token.TokenConfig) (string, error) {
	tokenInterface, err := tm.factory.CreateToken(tokenType, config)
	if err != nil {
		return "", infraerr.TokenGenerationError(string(tokenType), err)
	}

	tokenString, err := tokenInterface.Generate()
	if err != nil {
		return "", infraerr.TokenGenerationError(string(tokenType), err)
	}

	fmt.Println("Generated token:", tokenString) // Debugging line

	// Only store non-JWT tokens in Redis (JWT tokens are stateless)
	if tokenType != vo.JWTAccessToken && tokenType != vo.JWTRefreshToken {
		if err := tm.storeTokenInRedis(ctx, tokenString, tokenType, config, tokenInterface); err != nil {
			return "", err
		}
	}

	return tokenString, nil
}

func (tm *TokenManager) storeTokenInRedis(ctx context.Context, tokenString string, tokenType vo.TokenType, config token.TokenConfig, token vo.Token) error {
	timestamp := time.Now().Unix()
	key := tm.tokenKey(config.UserID, tokenType, timestamp)

	storedToken := StoredToken{
		TokenString: tokenString,
		TokenType:   tokenType,
		UserID:      config.UserID,
		ExpiresAt:   time.Now().Add(config.ExpiresIn),
	}

	tokenJSON, err := json.Marshal(storedToken)
	if err != nil {
		return fmt.Errorf("failed to marshal token to JSON: %w", err)
	}

	_, err = tm.redisClient.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		// Store the token data with an expiration time
		pipe.Set(ctx, key, tokenJSON, config.ExpiresIn)

		// Add the token key to the set of user tokens for this type
		pipe.SAdd(ctx, tm.userTokensKey(config.UserID, tokenType), key)

		// Set expiration on the user tokens set (cleanup)
		pipe.Expire(ctx, tm.userTokensKey(config.UserID, tokenType), config.ExpiresIn*2)
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to store token in Redis: %w", err)
	}

	return nil
}

func (tm *TokenManager) ValidateToken(ctx context.Context, tokenString string, tokenType vo.TokenType) (*token.TokenClaims, error) {
	// JWT tokens are stateless, validate directly
	if tokenType == vo.JWTAccessToken || tokenType == vo.JWTRefreshToken {
		tokenInterface, err := tm.factory.CreateToken(tokenType, token.TokenConfig{})
		if err != nil {
			return nil, infraerr.TokenGenerationError(string(tokenType), err)
		}
		claims, err := tokenInterface.Validate(tokenString)
		if err != nil {
			return nil, err
		}

		tokenClaims, ok := claims.(*token.TokenClaims)
		if !ok {
			return nil, fmt.Errorf("failed to assert claims to *token.TokenClaims")
		}
		return tokenClaims, nil
	}

	// For non-JWT tokens, check Redis
	return tm.validateTokenFromRedis(ctx, tokenString, tokenType)
}

func (tm *TokenManager) validateTokenFromRedis(ctx context.Context, tokenString string, tokenType vo.TokenType) (*token.TokenClaims, error) {
	// We need to scan through user tokens to find the matching token
	// This is a limitation - ideally we'd have the userID in the validation request
	// For now, we'll create a token and validate it directly
	tokenInterface, err := tm.factory.CreateToken(tokenType, token.TokenConfig{})
	if err != nil {
		return nil, infraerr.TokenGenerationError(string(tokenType), err)
	}

	claims, err := tokenInterface.Validate(tokenString)
	if err != nil {
		return nil, err
	}

	tokenClaims, ok := claims.(*token.TokenClaims)
	if !ok {
		return nil, fmt.Errorf("failed to assert claims to *token.TokenClaims")
	}

	// Verify token exists in Redis
	if err := tm.verifyTokenInRedis(ctx, tokenClaims.UserID, tokenType, tokenString); err != nil {
		return nil, err
	}

	return tokenClaims, nil
}

func (tm *TokenManager) verifyTokenInRedis(ctx context.Context, userID string, tokenType vo.TokenType, tokenString string) error {
	userTokensKey := tm.userTokensKey(userID, tokenType)
	tokenKeys, err := tm.redisClient.SMembers(ctx, userTokensKey).Result()
	if err != nil {
		return fmt.Errorf("failed to get token keys from Redis: %w", err)
	}

	if len(tokenKeys) == 0 {
		return infraerr.InvalidTokenError(string(tokenType))
	}

	// Check if any stored token matches
	for _, key := range tokenKeys {
		tokenJSON, err := tm.redisClient.Get(ctx, key).Bytes()
		if err != nil {
			continue // Token may have expired
		}

		var storedToken StoredToken
		if err := json.Unmarshal(tokenJSON, &storedToken); err != nil {
			continue
		}

		if storedToken.TokenString == tokenString {
			return nil // Token found and valid
		}
	}

	return infraerr.InvalidTokenError(string(tokenType))
}

func (tm *TokenManager) InvalidateToken(ctx context.Context, userID string, tokenType vo.TokenType, tokenString string) error {
	userTokensKey := tm.userTokensKey(userID, tokenType)
	tokenKeys, err := tm.redisClient.SMembers(ctx, userTokensKey).Result()
	if err != nil {
		return fmt.Errorf("failed to get token keys: %w", err)
	}

	// Find and delete the matching token
	for _, key := range tokenKeys {
		tokenJSON, err := tm.redisClient.Get(ctx, key).Bytes()
		if err != nil {
			continue
		}

		var storedToken StoredToken
		if err := json.Unmarshal(tokenJSON, &storedToken); err != nil {
			continue
		}

		if storedToken.TokenString == tokenString {
			_, err := tm.redisClient.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				pipe.Del(ctx, key)
				pipe.SRem(ctx, userTokensKey, key)
				return nil
			})
			return err
		}
	}

	return fmt.Errorf("token not found")
}

func (tm *TokenManager) InvalidateAllUserTokens(ctx context.Context, userID string, tokenType vo.TokenType) error {
	userTokensKey := tm.userTokensKey(userID, tokenType)
	tokenKeys, err := tm.redisClient.SMembers(ctx, userTokensKey).Result()
	if err != nil {
		return fmt.Errorf("failed to get token keys: %w", err)
	}

	_, err = tm.redisClient.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, key := range tokenKeys {
			pipe.Del(ctx, key)
		}
		pipe.Del(ctx, userTokensKey)
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to delete all tokens for user: %w", err)
	}

	return nil
}
