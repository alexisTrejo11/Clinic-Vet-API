package userValueObjects

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

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

type Email struct {
	value string
}

func NewEmail(email string) (Email, error) {
	if email == "" {
		return Email{}, errors.New("email cannot be empty")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return Email{}, errors.New("invalid email format")
	}

	return Email{value: strings.ToLower(email)}, nil
}

func (e Email) Value() string {
	return e.value
}

func (e Email) String() string {
	return e.value
}
