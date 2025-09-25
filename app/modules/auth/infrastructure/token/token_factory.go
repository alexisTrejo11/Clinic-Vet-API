package token

import (
	"fmt"
	"time"
)

type TokenFactory struct {
	jwtSecret []byte
}

func NewTokenFactory(jwtSecret []byte) *TokenFactory {
	return &TokenFactory{
		jwtSecret: jwtSecret,
	}
}

func (f *TokenFactory) CreateToken(tokenType TokenType, config TokenConfig) (Token, error) {
	if config.Secret == nil {
		config.Secret = f.jwtSecret
	}

	switch tokenType {
	case JWTAccessToken:
		if config.ExpiresIn == 0 {
			config.ExpiresIn = 15 * time.Minute // Default: 15 min
		}
		return &JWTAccessTokenImpl{config: config}, nil

	case JWTRefreshToken:
		if config.ExpiresIn == 0 {
			config.ExpiresIn = 24 * time.Hour * 7 // Default: 7 days
		}
		return &JWTRefreshTokenImpl{config: config}, nil

	case TwoFAToken:
		if config.ExpiresIn == 0 {
			config.ExpiresIn = 5 * time.Minute // Default: 5 min
		}
		return &TwoFATokenImpl{config: config}, nil

	case OAuth2SecretToken:
		if config.ExpiresIn == 0 {
			config.ExpiresIn = 24 * time.Hour // Default: 24 hours
		}
		return &OAuth2SecretTokenImpl{config: config}, nil

	case ActivationToken:
		if config.ExpiresIn == 0 {
			config.ExpiresIn = 24 * time.Hour // Default: 24 horas
		}
		return &ActivationTokenImpl{config: config}, nil

	default:
		return nil, fmt.Errorf("unsupported token type: %s", tokenType)
	}
}
