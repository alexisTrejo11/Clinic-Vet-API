package userValueObjects

import (
	"errors"
	"regexp"
)

type PhoneNumber struct {
	value string
}

func NewPhoneNumber(phone string) (PhoneNumber, error) {
	if phone == "" {
		return PhoneNumber{}, errors.New("phone number cannot be empty")
	}

	// Remover espacios y caracteres especiales
	cleaned := regexp.MustCompile(`[^\d+]`).ReplaceAllString(phone, "")

	if len(cleaned) < 10 {
		return PhoneNumber{}, errors.New("phone number too short")
	}

	return PhoneNumber{value: cleaned}, nil
}

func (p PhoneNumber) Value() string {
	return p.value
}

func (p PhoneNumber) String() string {
	return p.value
}
