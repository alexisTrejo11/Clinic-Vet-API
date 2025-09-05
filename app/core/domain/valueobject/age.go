// Package valueobject containes all the common valueobject provided to improve entity package
package valueobject

import (
	"errors"
)

type Age struct {
	Years  int
	Months int
}

func NewAge(years, months int) (Age, error) {
	if years < 0 || months < 0 {
		return Age{}, errors.New("age cannot be negative")
	}

	if months >= 12 {
		years += months / 12
		months = months % 12
	}

	if years > 50 {
		return Age{}, errors.New("age seems unrealistic")
	}

	return Age{Years: years, Months: months}, nil
}
