package employee

import (
	"context"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type EmployeeOption func(*Employee)

func WithName(name valueobject.PersonName) EmployeeOption {
	return func(v *Employee) {
		v.Person.UpdateName(context.Background(), name)
	}
}

func WithPhoto(photo string) EmployeeOption {
	return func(v *Employee) {
		v.photo = photo
	}
}

func WithLicenseNumber(licenseNumber string) EmployeeOption {
	return func(v *Employee) {
		v.licenseNumber = licenseNumber
	}
}

func WithSpecialty(specialty enum.VetSpecialty) EmployeeOption {
	return func(v *Employee) {
		v.specialty = specialty
	}
}

func WithYearsExperience(years int) EmployeeOption {
	return func(v *Employee) {
		v.yearsExperience = years
	}
}

func WithIsActive(isActive bool) EmployeeOption {
	return func(v *Employee) {
		v.isActive = isActive
	}
}

func WithUserID(userID *valueobject.UserID) EmployeeOption {
	return func(v *Employee) {
		v.userID = userID
	}
}

func WithSchedule(schedule *valueobject.Schedule) EmployeeOption {
	return func(v *Employee) {
		v.schedule = schedule
	}
}

func WithTimestamps(createdAt, updatedAt time.Time) EmployeeOption {
	return func(v *Employee) {
		v.SetTimeStamps(createdAt, updatedAt)
	}
}

func NewEmployee(id valueobject.EmployeeID, opts ...EmployeeOption) *Employee {
	vet := &Employee{Entity: base.NewEntity(id, nil, nil, 1)}

	for _, opt := range opts {
		opt(vet)
	}

	return vet
}

func CreateEmployee(
	ctx context.Context,
	name valueobject.PersonName,
	gender enum.PersonGender,
	dateOfBirth time.Time,
	opts ...EmployeeOption,
) (*Employee, error) {
	person, err := base.CreatePerson(ctx, name, dateOfBirth, gender)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	vet := &Employee{
		Entity:   base.NewEntity(valueobject.EmployeeID{}, &now, &now, 1),
		isActive: true,
		Person:   *person,
	}

	for _, opt := range opts {
		opt(vet)
	}

	if err := vet.validate(ctx); err != nil {
		return nil, err
	}

	return vet, nil
}
