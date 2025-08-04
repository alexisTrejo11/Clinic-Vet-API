package petDomain

import (
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type Gender string

const (
	GenderMale    Gender = "Male"
	GenderFemale  Gender = "Female"
	GenderUnknown Gender = "Unknown"
)

func (g Gender) IsValid() bool {
	switch g {
	case GenderMale, GenderFemale, GenderUnknown:
		return true
	}
	return false
}

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

type PetId struct {
	id shared.IntegerId
}

func NewPetId(value any) (PetId, error) {
	id, err := shared.NewIntegerId(value)
	if err != nil {
		return PetId{}, err
	}

	return PetId{id: id}, nil
}

func (p PetId) GetValue() int {
	return p.id.GetValue()
}

func (p PetId) String() string {
	return p.id.String()
}

func (p PetId) Equals(other PetId) bool {
	return p.id.GetValue() == other.id.GetValue()
}
