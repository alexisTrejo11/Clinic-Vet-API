package valueobject

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	domainerr "clinic-vet-api/app/modules/core/error"
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
	if !email.IsValid() {
		return Email{}, errors.New("invalid email format")
	}

	return email, nil
}

func (e *Email) Validate() error {
	if e.value == "" {
		return errors.New("email cannot be empty")
	}
	if !e.IsValid() {
		return errors.New("invalid email format")
	}
	return nil
}

func NewEmailNoErr(emailStr string) Email {
	e := &Email{value: strings.ToLower(emailStr)}
	return *e
}

func (e Email) IsValid() bool {
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
	Value string
}

func EmptyPhoneNumber() PhoneNumber {
	return PhoneNumber{Value: ""}
}

func NewPhoneNumber(phone string) (PhoneNumber, error) {
	if phone == "" {
		return PhoneNumber{}, domainerr.InvalidFieldFormat(context.Background(), "phone", "", "phone number cannot be empty", "create phone number")
	}

	cleaned := regexp.MustCompile(`[^\d+]`).ReplaceAllString(phone, "")

	if len(cleaned) < 10 {
		return PhoneNumber{}, domainerr.InvalidFieldFormat(context.Background(), "phone", "too short", "phone number must have at least 10 digits", "create phone number")
	}

	return PhoneNumber{Value: cleaned}, nil
}

func (p PhoneNumber) Validate() error {
	if p.Value == "" {
		return domainerr.InvalidFieldFormat(context.Background(), "phone", "", "phone number cannot be empty", "validate phone number")
	}

	cleaned := regexp.MustCompile(`[^\d+]`).ReplaceAllString(p.Value, "")

	if len(cleaned) < 10 {
		return domainerr.InvalidFieldFormat(context.Background(), "phone", "too short", "phone number must have at least 10 digits", "validate phone number")
	}

	return nil
}

func NewPhoneNumberNoErr(phone string) PhoneNumber {
	p, _ := NewPhoneNumber(phone)
	return p
}

func NewOptPhoneNumber(phone *string) *PhoneNumber {
	if phone == nil {
		return nil
	}

	return &PhoneNumber{Value: *phone}
}

func (p PhoneNumber) String() string {
	return p.Value
}

type PersonName struct {
	firstName string
	lastName  string
}

func NewPersonName(firstName, lastName string) PersonName {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	return PersonName{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (n PersonName) IsValid() bool {
	return n.firstName != "" && n.lastName != ""
}

func (n PersonName) FullName() string {
	return fmt.Sprintf("%s %s", n.firstName, n.lastName)
}

func (n PersonName) FirstName() string {
	return n.firstName
}

func (n PersonName) LastName() string {
	return n.lastName
}

func NewOptPersonName(firstName, lastName *string) *PersonName {
	name := &PersonName{}
	if firstName != nil {
		name.firstName = *firstName
	}
	if lastName != nil {
		name.lastName = *lastName
	}
	return name
}
