package valueobject

import (
	domainerr "clinic-vet-api/app/modules/core/error"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type (
	BuildingType string
	Country      string
)

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

func NewEmailNoErr(emailStr string) Email {
	e := &Email{value: strings.ToLower(emailStr)}
	return *e
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
		return PhoneNumber{}, domainerr.InvalidFieldFormat(context.Background(), "phone", "", "phone number cannot be empty", "create phone number")
	}

	cleaned := regexp.MustCompile(`[^\d+]`).ReplaceAllString(phone, "")

	if len(cleaned) < 10 {
		return PhoneNumber{}, domainerr.InvalidFieldFormat(context.Background(), "phone", "too short", "phone number must have at least 10 digits", "create phone number")
	}

	return PhoneNumber{value: cleaned}, nil
}

func NewPhoneNumberNoErr(phone string) PhoneNumber {
	p, _ := NewPhoneNumber(phone)
	return p
}

func (p PhoneNumber) Value() string {
	return p.value
}

func (p PhoneNumber) String() string {
	return p.value
}

type PersonName struct {
	FirstName string
	LastName  string
}

func NewPersonName(firstName, lastName string) (PersonName, error) {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	if firstName == "" {
		return PersonName{}, domainerr.InvalidFieldValue(context.Background(), "first_name", "empty", "first name cannot be empty", "create name")
	}

	if lastName == "" {
		return PersonName{}, domainerr.InvalidFieldValue(context.Background(), "last_name", "empty", "last name cannot be empty", "create name")
	}

	return PersonName{
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}

func (n PersonName) FullName() string {
	return fmt.Sprintf("%s %s", n.FirstName, n.LastName)
}
