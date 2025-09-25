package token

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

type TwoFATokenImpl struct {
	config    TokenConfig
	expiresAt time.Time
	token     string
}

func (t *TwoFATokenImpl) Generate() (string, error) {
	// 6 Digits
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", fmt.Errorf("error generating 2FA token: %w", err)
	}

	t.token = fmt.Sprintf("%06d", n.Int64())
	t.expiresAt = time.Now().Add(t.config.ExpiresIn)

	return t.token, nil
}

func (t *TwoFATokenImpl) Validate(token string) (*TokenClaims, error) {
	if t.IsExpired() {
		return nil, fmt.Errorf("2FA token expired")
	}

	if token != t.token {
		return nil, fmt.Errorf("invalid 2FA token")
	}

	return &TokenClaims{
		UserID:    t.config.UserID,
		TokenType: TwoFAToken,
		IssuedAt:  time.Now().Add(-t.config.ExpiresIn),
		ExpiresAt: t.expiresAt,
		Purpose:   t.config.Purpose,
	}, nil
}

func (t *TwoFATokenImpl) GetType() TokenType      { return TwoFAToken }
func (t *TwoFATokenImpl) GetExpiresAt() time.Time { return t.expiresAt }
func (t *TwoFATokenImpl) IsExpired() bool         { return time.Now().After(t.expiresAt) }
