package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/jwt"
	jwtLib "github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secretKey            string
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
	issuer               string
}

const (
	DefaultAccessTokenDuration  = time.Hour
	DefaultRefreshTokenDuration = 7 * 24 * time.Hour
	DefaultIssuer               = "clinic-vet-api"
)

// NewJWTService crea una nueva instancia del servicio JWT
func NewJWTService(secretKey string) jwt.JWTService {
	return &jwtService{
		secretKey:            secretKey,
		accessTokenDuration:  DefaultAccessTokenDuration,
		refreshTokenDuration: DefaultRefreshTokenDuration,
		issuer:               DefaultIssuer,
	}
}

func NewJWTServiceWithConfig(secretKey, issuer string, accessDuration, refreshDuration time.Duration) jwt.JWTService {
	return &jwtService{
		secretKey:            secretKey,
		accessTokenDuration:  accessDuration,
		refreshTokenDuration: refreshDuration,
		issuer:               issuer,
	}
}

func (s *jwtService) GenerateAccessToken(userID string) (string, error) {
	if userID == "" {
		return "", errors.New("user ID cannot be empty")
	}

	now := time.Now()
	claims := &jwt.Claims{
		UserID: userID,
		Type:   "access",
		RegisteredClaims: jwtLib.RegisteredClaims{
			ExpiresAt: jwtLib.NewNumericDate(now.Add(s.accessTokenDuration)),
			IssuedAt:  jwtLib.NewNumericDate(now),
			NotBefore: jwtLib.NewNumericDate(now),
			Issuer:    s.issuer,
			Subject:   userID,
		},
	}

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}

func (s *jwtService) GenerateRefreshToken(userID string) (string, error) {
	if userID == "" {
		return "", errors.New("user ID cannot be empty")
	}

	now := time.Now()
	claims := &jwt.Claims{
		UserID: userID,
		Type:   "refresh",
		RegisteredClaims: jwtLib.RegisteredClaims{
			ExpiresAt: jwtLib.NewNumericDate(now.Add(s.refreshTokenDuration)),
			IssuedAt:  jwtLib.NewNumericDate(now),
			NotBefore: jwtLib.NewNumericDate(now),
			Issuer:    s.issuer,
			Subject:   userID,
		},
	}

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", fmt.Errorf("error signing refresh token: %w", err)
	}

	return tokenString, nil
}

func (s *jwtService) ExtractToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("authorization header is empty")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		return "", errors.New("authorization header format must be 'Bearer {token}'")
	}

	if strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("authorization header must start with 'Bearer'")
	}

	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", errors.New("token cannot be empty")
	}

	return token, nil
}

func (s *jwtService) ValidateToken(tokenString string) (*jwt.Claims, error) {
	if tokenString == "" {
		return nil, errors.New("token string cannot be empty")
	}

	token, err := jwtLib.ParseWithClaims(tokenString, &jwt.Claims{}, func(token *jwtLib.Token) (any, error) {
		// Verificar que el método de firma sea HMAC
		if _, ok := token.Method.(*jwtLib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected token signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		/*
			if ve, ok := err.(*jwtLib.ValidationErrior); ok {
				switch {
				case ve.Errors&jwtLib.ValidationErrorMalformed != 0:
					return nil, errors.New("malformed token")
				case ve.Errors&jwtLib.ValidationErrorExpired != 0:
					return nil, errors.New("token has expired")
				case ve.Errors&jwtLib.ValidationErrorNotValidYet != 0:
					return nil, errors.New("token not valid yet")
				case ve.Errors&jwtLib.ValidationErrorSignatureInvalid != 0:
					return nil, errors.New("invalid token signature")
				default:
					return nil, fmt.Errorf("token validation error: %w", err)
				}
			}
		*/
		return nil, fmt.Errorf("could not parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*jwt.Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	if claims.UserID == "" {
		return nil, errors.New("token missing user ID")
	}

	if claims.Type != "access" && claims.Type != "refresh" {
		return nil, errors.New("invalid token type")
	}

	if s.issuer != "" && claims.Issuer != s.issuer {
		return nil, errors.New("invalid token issuer")
	}

	return claims, nil
}

func (s *jwtService) GetUserIDFromToken(tokenString string) (string, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	if claims.UserID == "" {
		return "", errors.New("user ID not found in token")
	}

	return claims.UserID, nil
}

func (s *jwtService) RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := s.ValidateToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	if claims.Type != "refresh" {
		return "", errors.New("provided token is not a refresh token")
	}

	// Generar nuevo access token
	return s.GenerateAccessToken(claims.UserID)
}

func (s *jwtService) IsTokenExpired(tokenString string) bool {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return true
	}

	return claims.ExpiresAt.Time.Before(time.Now())
}

func (s *jwtService) GetTokenRemainingTime(tokenString string) (time.Duration, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}

	remaining := time.Until(claims.ExpiresAt.Time)
	if remaining < 0 {
		return 0, errors.New("token has expired")
	}

	return remaining, nil
}

func (s *jwtService) ValidateAccessToken(tokenString string) (*jwt.Claims, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.Type != "access" {
		return nil, errors.New("token is not an access token")
	}

	return claims, nil
}

func (s *jwtService) ValidateRefreshToken(tokenString string) (*jwt.Claims, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.Type != "refresh" {
		return nil, errors.New("token is not a refresh token")
	}

	return claims, nil
}

func (s *jwtService) GenerateTokenPair(userID string) (accessToken, refreshToken string, err error) {
	accessToken, err = s.GenerateAccessToken(userID)
	if err != nil {
		return "", "", fmt.Errorf("error generating access token: %w", err)
	}

	refreshToken, err = s.GenerateRefreshToken(userID)
	if err != nil {
		return "", "", fmt.Errorf("error generating refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}
