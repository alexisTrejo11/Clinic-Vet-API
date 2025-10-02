package token

import (
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type OAuth2SecretTokenImpl struct {
	config    TokenConfig
	expiresAt time.Time
}

func (o *OAuth2SecretTokenImpl) Generate() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("error generating OAuth2 secret: %w", err)
	}

	o.expiresAt = time.Now().Add(o.config.ExpiresIn)

	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:]), nil
}

func (o *OAuth2SecretTokenImpl) Validate(token string) (any, error) {
	if o.IsExpired() {
		return nil, fmt.Errorf("OAuth2 secret expired")
	}

	// TODO: Validate on db
	return &TokenClaims{
		UserID:    o.config.UserID,
		TokenType: vo.OAuth2SecretToken,
		IssuedAt:  time.Now().Add(-o.config.ExpiresIn),
		ExpiresAt: o.expiresAt,
		Purpose:   o.config.Purpose,
	}, nil
}

func (o *OAuth2SecretTokenImpl) GetType() vo.TokenType   { return vo.OAuth2SecretToken }
func (o *OAuth2SecretTokenImpl) GetExpiresAt() time.Time { return o.expiresAt }
func (o *OAuth2SecretTokenImpl) IsExpired() bool         { return time.Now().After(o.expiresAt) }
