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
	name        valueobject.PersonName
	dateOfBirth time.Time
	gender      enum.PersonGender
	createdAt   time.Time
	updatedAt   time.Time
}

func CreatePerson(ctx context.Context, name valueobject.PersonName, dateOfBirth time.Time, gender enum.PersonGender) (*Person, error) {
	person := &Person{
		name:        name,
		dateOfBirth: dateOfBirth,
		gender:      gender,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}

	if err := person.Validate(ctx, OperationCreatePerson); err != nil {
		return nil, err
	}

	return person, nil
}

// TODO: SEPARATE VALIDATIONS FOR EACH FIELD AND ADD IT ON UPDATE METHODS
func (p *Person) Validate(ctx context.Context, operation PersonOperation) error {
	if p.name.FullName() == "" {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonNameRequired, "full_name", "", string(operation))
	}

	if p.dateOfBirth.IsZero() {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonDOBRequired, "date_of_birth", "", string(operation))
	}

	if p.dateOfBirth.After(time.Now()) {
		return domainerr.PersonValidationError(ctx,
			domainerr.PersonDOBFuture, "date_of_birth", p.dateOfBirth.Format(time.RFC3339), string(operation))
	}

	// Validar mayoría de edad (18 años)
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

// Métodos de actualización mejorados
func (p *Person) UpdateName(ctx context.Context, name valueobject.PersonName) error {
	p.name = name
	p.updatedAt = time.Now()
	return nil
}

func (p *Person) UpdateDateOfBirth(ctx context.Context, dob time.Time) error {
	p.dateOfBirth = dob

	p.updatedAt = time.Now()
	return nil
}

func (p *Person) UpdateGender(ctx context.Context, gender enum.PersonGender) error {
	p.gender = gender

	p.updatedAt = time.Now()
	return nil
}

func (p *Person) ID() string {
	return p.id
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

func (p *Person) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Person) UpdatedAt() time.Time {
	return p.updatedAt
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

func (p *Person) SetTimeStamps(createdAt, updatedAt time.Time) {
	p.createdAt = createdAt
	p.updatedAt = updatedAt
}

func (p *Person) SetID(id string) {
	p.id = id
}
