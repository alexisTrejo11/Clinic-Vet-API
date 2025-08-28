package valueobject

import (
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
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

type PetID struct {
	id shared.IntegerId
}

func NewPetID(value any) (PetID, error) {
	id, err := shared.NewIntegerId(value)
	if err != nil {
		return PetID{}, err
	}

	return PetID{id: id}, nil
}

func (p PetID) GetValue() int {
	return p.id.GetValue()
}

func (p PetID) String() string {
	return p.id.String()
}

func (p PetID) Equals(other PetID) bool {
	return p.id.GetValue() == other.id.GetValue()
}
