package valueObjects

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type PhoneNumber struct {
	value string
}

type PersonName struct {
	FirstName string
	LastName  string
}

type Age struct {
	years  int
	months int
}

type Email struct {
	value string
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

func NewAge(years, months int) (Age, error) {
	if years < 0 || months < 0 {
		return Age{}, errors.New("age cannot be negative")
	}

	if months >= 12 {
		years += months / 12
		months = months % 12
	}

	if years > 50 { // Límite razonable para longevidad animal
		return Age{}, errors.New("age seems unrealistic")
	}

	return Age{years: years, months: months}, nil
}

func (a Age) Years() int {
	return a.years
}

func (a Age) Months() int {
	return a.months
}

func (a Age) TotalMonths() int {
	return a.years*12 + a.months
}

func (a Age) String() string {
	if a.years == 0 {
		return fmt.Sprintf("%d meses", a.months)
	}
	if a.months == 0 {
		return fmt.Sprintf("%d años", a.years)
	}
	return fmt.Sprintf("%d años y %d meses", a.years, a.months)
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

type Gender string

const (
	MALE         Gender = "male"
	Female       Gender = "female"
	NotSpecified Gender = "not_specified"
)

func NewGender(value string) Gender {
	switch value {
	case "male":
		return MALE
	case "female":
		return Female
	case "not_specified":
		return NotSpecified
	case "":
		return NotSpecified
	default:
		return NotSpecified
	}
}

func (g Gender) String() string {
	return string(g)
}
