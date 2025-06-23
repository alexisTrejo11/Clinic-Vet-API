package userValueObjects

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
