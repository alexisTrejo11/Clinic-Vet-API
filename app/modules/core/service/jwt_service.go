package service

import (
	token "clinic-vet-api/app/modules/account/auth/token/factory"
	"time"
)

type JWTService interface {
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateToken(tokenString string) (*token.TokenClaims, error)
	ExtractToken(authHeader string) (string, error)
	GetUserIDFromToken(tokenString string) (string, error)
	RefreshAccessToken(refreshToken string) (string, error)
	IsTokenExpired(tokenString string) bool
	GetTokenRemainingTime(tokenString string) (time.Duration, error)
}
