package base

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

type Person struct {
	name        valueobject.PersonName
	dateOfBirth time.Time
	gender      enum.PersonGender
}

func (p *Person) Name() valueobject.PersonName {
	return p.name
}

func (p *Person) DateOfBirth() time.Time {
	return p.dateOfBirth
}

func (p *Person) Gender() enum.PersonGender {
	return p.gender
}

func (p *Person) Age() int {
	now := time.Now()
	age := now.Year() - p.dateOfBirth.Year()

	if now.YearDay() < p.dateOfBirth.YearDay() {
		age--
	}
	return age
}

func NewPerson(name valueobject.PersonName, dateOfBirth time.Time, gender enum.PersonGender) Person {
	return Person{
		name:        name,
		dateOfBirth: dateOfBirth,
		gender:      gender,
	}
}

func (p *Person) UpdateName(name valueobject.PersonName) error {
	if name.FullName() == "" {
		return domainerr.NewValidationError("owner", "full name", "full name is required")
	}

	p.name = name

	return nil
}

func (p *Person) UpdateDateOfBirth(dob time.Time) error {
	if dob.IsZero() {
		return domainerr.NewValidationError("owner", "date-of-birth", "date of birth is required")
	}
	if dob.After(time.Now()) {
		return domainerr.NewValidationError("owner", "date-of-birth", "date of birth cannot be in the future")
	}
	minAgeDate := time.Now().AddDate(-18, 0, 0)
	if dob.After(minAgeDate) {
		return domainerr.NewValidationError("owner", "date-of-birth", "owner must be at least 18 years old")
	}

	p.dateOfBirth = dob
	return nil
}

func (p *Person) UpdateGender(gender enum.PersonGender) error {
	if !gender.IsValid() {
		return domainerr.NewValidationError("owner", "gender", "invalid gender")
	}
	p.gender = gender
	return nil
}
