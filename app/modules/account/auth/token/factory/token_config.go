package token

import (
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID     string       `json:"user_id"`
	TokenType  vo.TokenType `json:"token_type"`
	CreatedAt  time.Time    `json:"created_at"`
	ValidUntil time.Time    `json:"valid_until"`
	jwt.RegisteredClaims
}

type TokenConfig struct {
	UserID    string
	Issuer    string
	ExpiresIn time.Duration
	Secret    []byte
}
