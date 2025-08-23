package userDomain

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type Address struct {
	Street              string
	City                string
	State               string
	ZipCode             string
	Country             Country
	BuildingType        BuildingType
	BuildingOuterNumber string
	BuildingInnerNumber *string
}

type BuildingType string
type Country string

const (
	USA    Country = "USA"
	Mexico Country = "Mexico"
	Canada Country = "Canada"
)

func (c Country) IsValid() bool {
	switch c {
	case USA, Mexico, Canada:
		return true
	}
	return false
}

const (
	BuildingTypeHouse     BuildingType = "house"
	BuildingTypeApartment BuildingType = "apartment"
	BuildingTypeOffice    BuildingType = "office"
	BuildingTypeOther     BuildingType = "other"
)

type TwoFactorAuth struct {
	IsEnabled bool
	Secret    string
}

type UserId struct {
	id shared.IntegerId
}

func NilUserId() UserId {
	return UserId{id: shared.NilIntegerId()}
}

func NewUserId(id any) (UserId, error) {
	userId, err := shared.NewIntegerId(id)
	if err != nil {
		return UserId{}, fmt.Errorf("invalid UserId: %w", err)
	}
	return UserId{id: userId}, nil
}

func (u UserId) GetValue() int {
	return u.id.GetValue()
}

func (u UserId) String() string {
	return u.id.String()
}

func (u UserId) Equals(other UserId) bool {
	return u.id.Equals(other.id)
}

type Email struct {
	value string
}

func NewEmail(emailStr string) (Email, error) {
	if emailStr == "" {
		return Email{}, errors.New("email cannot be empty")
	}

	email := Email{value: strings.ToLower(emailStr)}
	if !email.isValid() {
		return Email{}, errors.New("invalid email format")
	}

	return email, nil
}

func (e Email) isValid() bool {
	// Simple regex for email validation
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(e.value)
}

func (e Email) Value() string {
	return e.value
}

func (e Email) String() string {
	return e.value
}

type PhoneNumber struct {
	value string
}

func NewPhoneNumber(phone string) (PhoneNumber, error) {
	if phone == "" {
		return PhoneNumber{}, errors.New("phone number cannot be empty")
	}

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
