package repository

import (
	token "clinic-vet-api/app/modules/account/auth/token/factory"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"context"
)

type TokenRepository interface {
	GenerateToken(ctx context.Context, tokenType vo.TokenType, config token.TokenConfig) (string, error)
	ValidateToken(ctx context.Context, tokenString string, tokenType vo.TokenType) (*token.TokenClaims, error)
	InvalidateToken(ctx context.Context, userID string, tokenType vo.TokenType, tokenString string) error
	InvalidateAllUserTokens(ctx context.Context, userID string, tokenType vo.TokenType) error
}
