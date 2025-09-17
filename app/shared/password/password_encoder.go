// Package password contains all the logic for managing passwrod enconding
package password

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordEncoder interface {
	HashPassword(password string) (string, error)
	CheckPassword(hashedPassword, password string) bool
}

type PasswordEncoderImpl struct{}

func NewPasswordEncoder() PasswordEncoder {
	return &PasswordEncoderImpl{}
}

func (p *PasswordEncoderImpl) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (p *PasswordEncoderImpl) CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
