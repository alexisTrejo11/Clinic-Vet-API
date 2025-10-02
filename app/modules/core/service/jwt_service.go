package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
	ExtractToken(authHeader string) (string, error)
	GetUserIDFromToken(tokenString string) (string, error)
	RefreshAccessToken(refreshToken string) (string, error)
	IsTokenExpired(tokenString string) bool
	GetTokenRemainingTime(tokenString string) (time.Duration, error)
}
