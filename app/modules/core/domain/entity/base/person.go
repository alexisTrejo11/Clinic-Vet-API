// domain/person.go
package base

import (
	"context"
	"time"

	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	domainerr "clinic-vet-api/app/modules/core/error"
)

// PersonOperation representa operaciones comunes sobre Person
type PersonOperation string

const (
	OperationCreatePerson PersonOperation = "CREATE_PERSON"
	OperationUpdatePerson PersonOperation = "UPDATE_PERSON"
	OperationGetPerson    PersonOperation = "GET_PERSON"
)

type Person struct {
	id          string
	firstName   string
	lastName    string
	dateOfBirth time.Time
	gender      enum.PersonGender
}

func (p *Person) SetName(firstName, lastName string) {
	p.firstName = firstName
	p.lastName = lastName
}

func (p *Person) SetFirstName(firstName string) {
	p.firstName = firstName
}

func (p *Person) SetLastName(lastName string) {
	p.lastName = lastName
}

func (p *Person) SetDateOfBirth(dob time.Time) {
	p.dateOfBirth = dob
}

func (p *Person) SetGender(gender enum.PersonGender) {
	p.gender = gender
}

func CreatePerson(ctx context.Context, firstName, lastName string, dateOfBirth time.Time, gender enum.PersonGender) (*Person, error) {
	person := &Person{
		firstName:   firstName,
		lastName:    lastName,
		dateOfBirth: dateOfBirth,
		gender:      gender,
	}

	if err := person.Validate(ctx); err != nil {
		return nil, err
	}

	return person, nil
}

// TODO: SEPARATE VALIDATIONS FOR EACH FIELD AND ADD IT ON UPDATE METHODS
func (p *Person) Validate(ctx context.Context) error {
	operation := "Validating Person"
	if p.firstName == "" {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonNameRequired, "first_name", "", string(operation))
	}

	if p.lastName == "" {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonNameRequired, "last_name", "", string(operation))
	}

	if p.dateOfBirth.IsZero() {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonDOBRequired, "date_of_birth", "", string(operation))
	}

	if p.dateOfBirth.After(time.Now()) {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonDOBFuture, "date_of_birth", p.dateOfBirth.Format(time.RFC3339), string(operation))
	}

	minAgeDate := time.Now().AddDate(-18, 0, 0)
	if p.dateOfBirth.After(minAgeDate) {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonUnderage, "date_of_birth", p.dateOfBirth.Format(time.RFC3339), string(operation))
	}

	if !p.gender.IsValid() {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonGenderInvalid, "gender", string(p.gender), string(operation))
	}

	return nil
}

func (p *Person) UpdateName(ctx context.Context, firstName, lastName string) error {
	if firstName == "" {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonNameRequired, "first_name", "", "UPDATE_PERSON")
	}

	if lastName == "" {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonNameRequired, "last_name", "", "UPDATE_PERSON")
	}
	p.firstName = firstName
	p.lastName = lastName

	return nil
}

func (p *Person) UpdateDateOfBirth(ctx context.Context, dob time.Time) error {
	p.dateOfBirth = dob
	if dob.IsZero() {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonDOBRequired, "date_of_birth", "", "UPDATE_PERSON")
	}

	if dob.After(time.Now()) {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonDOBFuture, "date_of_birth", dob.Format(time.RFC3339), "UPDATE_PERSON")
	}

	minAgeDate := time.Now().AddDate(-18, 0, 0)
	if dob.After(minAgeDate) {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonUnderage, "date_of_birth", dob.Format(time.RFC3339), "UPDATE_PERSON")
	}
	return nil
}

func (p *Person) UpdateGender(ctx context.Context, gender enum.PersonGender) error {
	if !gender.IsValid() {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonGenderInvalid, "gender", string(gender), "UPDATE_PERSON")
	}

	p.gender = gender

	return nil
}

func (p *Person) ID() string {
	return p.id
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) LastName() string {
	return p.lastName
}

func (p *Person) FullName() valueobject.PersonName {
	return valueobject.NewPersonName(p.firstName, p.lastName)
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

func (p *Person) IsAdult() bool {
	return p.Age() >= 18
}

func (p *Person) SetID(id string) {
	p.id = id
}
