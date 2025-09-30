package token

import (
	infraerr "clinic-vet-api/app/shared/error/infrastructure"
	"fmt"
	"strconv"
	"time"
)

type TokenManager struct {
	factory *TokenFactory
	tokens  map[string]Token // TODO: Add Redis
}

func NewTokenManager(jwtSecret []byte) *TokenManager {
	return &TokenManager{
		factory: NewTokenFactory(jwtSecret),
		tokens:  make(map[string]Token),
	}
}

func (tm *TokenManager) GenerateToken(tokenType TokenType, config TokenConfig) (string, error) {
	token, err := tm.factory.CreateToken(tokenType, config)
	if err != nil {
		return "", infraerr.TokenGenerationError(string(tokenType), err)
	}

	tokenString, err := token.Generate()
	if err != nil {
		return "", infraerr.TokenGenerationError(string(tokenType), err)
	}

	fmt.Println("Generated token:", tokenString) // Debugging line

	key := fmt.Sprintf("%s_%s_%s", config.UserID, tokenType, strconv.FormatInt(time.Now().Unix(), 10))
	tm.tokens[key] = token

	return tokenString, nil
}

func (tm *TokenManager) ValidateToken(tokenString string, tokenType TokenType) (*TokenClaims, error) {
	if tokenType == JWTAccessToken || tokenType == JWTRefreshToken {
		token, err := tm.factory.CreateToken(tokenType, TokenConfig{})
		if err != nil {
			return nil, infraerr.TokenGenerationError(string(tokenType), err)
		}
		return token.Validate(tokenString)
	}

	for _, token := range tm.tokens {
		if token.GetType() == tokenType {
			claims, err := token.Validate(tokenString)
			if err == nil {
				return claims, nil
			}
		}
	}

	return nil, infraerr.InvalidTokenError(string(tokenType))
}

func (tm *TokenManager) CleanupExpiredTokens() {
	for key, token := range tm.tokens {
		if token.IsExpired() {
			delete(tm.tokens, key)
		}
	}
}
