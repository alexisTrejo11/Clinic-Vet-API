package token

import (
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

type ActivationTokenImpl struct {
	config    TokenConfig
	expiresAt time.Time
	token     string
}

func (a *ActivationTokenImpl) Generate() (string, error) {
	// Generar código de activación de 8 dígitos
	max := big.NewInt(99999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", fmt.Errorf("error generating activation token: %w", err)
	}

	a.token = fmt.Sprintf("%08d", n.Int64())
	a.expiresAt = time.Now().Add(a.config.ExpiresIn)

	return a.token, nil
}

func (a *ActivationTokenImpl) Validate(token string) (any, error) {
	if a.IsExpired() {
		return nil, fmt.Errorf("activation token expired")
	}

	if token != a.token {
		return nil, fmt.Errorf("invalid activation token")
	}

	return &TokenClaims{
		UserID:    a.config.UserID,
		TokenType: vo.ActivationToken,
		IssuedAt:  time.Now().Add(-a.config.ExpiresIn),
		ExpiresAt: a.expiresAt,
		Purpose:   a.config.Purpose,
	}, nil
}

func (a *ActivationTokenImpl) GetType() vo.TokenType   { return vo.ActivationToken }
func (a *ActivationTokenImpl) GetExpiresAt() time.Time { return a.expiresAt }
func (a *ActivationTokenImpl) IsExpired() bool         { return time.Now().After(a.expiresAt) }
