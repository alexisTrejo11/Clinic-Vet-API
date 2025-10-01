package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	JWTAccessToken    TokenType = "jwt_access"
	JWTRefreshToken   TokenType = "jwt_refresh"
	TwoFAToken        TokenType = "2fa"
	OAuth2SecretToken TokenType = "oauth2_secret"
	ActivationToken   TokenType = "activation_token"
	VerificationToken TokenType = "verification_token"
)

type Token interface {
	Generate() (string, error)
	Validate(token string) (*TokenClaims, error)
	GetType() TokenType
	IsExpired() bool
}

type TokenClaims struct {
	UserID    string    `json:"user_id"`
	TokenType TokenType `json:"token_type"`
	IssuedAt  time.Time `json:"iat"`
	ExpiresAt time.Time `json:"exp"`
	Purpose   string    `json:"purpose,omitempty"`
	jwt.RegisteredClaims
}

type TokenConfig struct {
	UserID    string
	Purpose   string
	ExpiresIn time.Duration
	Secret    []byte
}
