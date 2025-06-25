package shared

import (
	"errors"
	"fmt"
	"strings"
)

type PersonName struct {
	firstName string
	lastName  string
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
		firstName: firstName,
		lastName:  lastName,
	}, nil
}

func (n PersonName) FirstName() string {
	return n.firstName
}

func (n PersonName) LastName() string {
	return n.lastName
}

func (n PersonName) FullName() string {
	return fmt.Sprintf("%s %s", n.firstName, n.lastName)
}

type Money struct {
	amount   float64
	currency string
}

func NewMoney(amount float64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, errors.New("amount cannot be negative")
	}

	if currency == "" {
		currency = "MXN"
	}

	return Money{
		amount:   amount,
		currency: strings.ToUpper(currency),
	}, nil
}
