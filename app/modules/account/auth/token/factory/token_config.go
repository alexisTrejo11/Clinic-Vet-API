package token

import (
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID    string       `json:"user_id"`
	TokenType vo.TokenType `json:"token_type"`
	IssuedAt  time.Time    `json:"iat"`
	ExpiresAt time.Time    `json:"exp"`
	Purpose   string       `json:"purpose,omitempty"`
	jwt.RegisteredClaims
}

type TokenConfig struct {
	UserID    string
	Purpose   string
	ExpiresIn time.Duration
	Secret    []byte
}
