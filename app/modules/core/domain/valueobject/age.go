// Package valueobject containes all the common valueobject provided to improve entity package
package valueobject

import (
	domainerr "clinic-vet-api/app/modules/core/error"
	"context"
)

type Age struct {
	Years  int
	Months int
}

func NewAge(years, months int) (Age, error) {
	if years < 0 || months < 0 {
		return Age{}, domainerr.InvalidFieldFormat(context.Background(), "age", "negative int", "years and months must be non-negative", "created age: "+string(years)+" years, "+string(months)+" months")
	}

	if months >= 12 {
		years += months / 12
		months = months % 12
	}

	if years > 50 {
		return Age{}, domainerr.InvalidFieldFormat(context.Background(), "age", "unrealistic age", "years must be less than or equal to 50", "created age: "+string(years)+" years, "+string(months)+" months")
	}

	return Age{Years: years, Months: months}, nil
}
