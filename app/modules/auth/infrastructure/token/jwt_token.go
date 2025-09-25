package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAccessTokenImpl struct {
	config    TokenConfig
	expiresAt time.Time
}

func (j *JWTAccessTokenImpl) Generate() (string, error) {
	now := time.Now()
	j.expiresAt = now.Add(j.config.ExpiresIn)

	claims := TokenClaims{
		UserID:    j.config.UserID,
		TokenType: JWTAccessToken,
		IssuedAt:  now,
		ExpiresAt: j.expiresAt,
		Purpose:   j.config.Purpose,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(j.expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   j.config.UserID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.config.Secret)
}

func (j *JWTAccessTokenImpl) Validate(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.config.Secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func (j *JWTAccessTokenImpl) GetType() TokenType      { return JWTAccessToken }
func (j *JWTAccessTokenImpl) GetExpiresAt() time.Time { return j.expiresAt }
func (j *JWTAccessTokenImpl) IsExpired() bool         { return time.Now().After(j.expiresAt) }

type JWTRefreshTokenImpl struct {
	config    TokenConfig
	expiresAt time.Time
}

func (r *JWTRefreshTokenImpl) Generate() (string, error) {
	now := time.Now()
	r.expiresAt = now.Add(r.config.ExpiresIn)

	claims := TokenClaims{
		UserID:    r.config.UserID,
		TokenType: JWTRefreshToken,
		IssuedAt:  now,
		ExpiresAt: r.expiresAt,
		Purpose:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(r.expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   r.config.UserID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(r.config.Secret)
}

func (r *JWTRefreshTokenImpl) Validate(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return r.config.Secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing refresh token: %w", err)
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid refresh token")
}

func (r *JWTRefreshTokenImpl) GetType() TokenType      { return JWTRefreshToken }
func (r *JWTRefreshTokenImpl) GetExpiresAt() time.Time { return r.expiresAt }
func (r *JWTRefreshTokenImpl) IsExpired() bool         { return time.Now().After(r.expiresAt) }
