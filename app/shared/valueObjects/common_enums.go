package valueObjects

import (
	"errors"
	"fmt"
	"strings"
)

type PersonName struct {
	FirstName string
	LastName  string
}

type Age struct {
	years  int
	months int
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
