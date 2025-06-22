package petDomain

import "errors"

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
