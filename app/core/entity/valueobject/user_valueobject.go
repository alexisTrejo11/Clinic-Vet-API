package valueobject

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
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

type UserID struct {
	id shared.IntegerId
}

func NilUserID() UserID {
	return UserID{id: shared.NilIntegerId()}
}

func NewUserID(id any) (UserID, error) {
	userID, err := shared.NewIntegerId(id)
	if err != nil {
		return UserID{}, fmt.Errorf("invalid UserID: %w", err)
	}
	return UserID{id: userID}, nil
}

func (u UserID) GetValue() int {
	return u.id.GetValue()
}

func (u UserID) String() string {
	return u.id.String()
}

func (u UserID) Equals(other UserID) bool {
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

type PersonName struct {
	FirstName string
	LastName  string
}

func NewPersonName(firstName, lastName string) (PersonName, error) {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	if firstName == "" {
		return PersonName{}, errors.New("first name cannot be empty")
	}

	if lastName == "" {
		return PersonName{}, errors.New("last name cannot be empty")
	}

	return PersonName{
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}

func (n PersonName) FullName() string {
	return fmt.Sprintf("%s %s", n.FirstName, n.LastName)
}
