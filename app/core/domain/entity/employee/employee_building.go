package employee

import (
	"context"
	"errors"
	"fmt"
	"time"

	"clinic-vet-api/app/core/domain/entity/base"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	domainerr "clinic-vet-api/app/core/error"
)

type EmployeeOption func(*Employee) error

func WithName(name valueobject.PersonName) EmployeeOption {
	return func(v *Employee) error {
		return v.Person.UpdateName(context.Background(), name)
	}
}

func WithPhoto(photo string) EmployeeOption {
	return func(v *Employee) error {
		v.photo = photo
		return nil
	}
}

func WithLicenseNumber(licenseNumber string) EmployeeOption {
	return func(v *Employee) error {
		v.licenseNumber = licenseNumber
		return nil
	}
}

func WithSpecialty(specialty enum.VetSpecialty) EmployeeOption {
	return func(v *Employee) error {
		v.specialty = specialty
		return nil
	}
}

func WithYearsExperience(years int) EmployeeOption {
	return func(v *Employee) error {
		if err := validateYearsExperience(years); err != nil {
			return err
		}
		v.yearsExperience = years
		return nil
	}
}

func WithIsActive(isActive bool) EmployeeOption {
	return func(v *Employee) error {
		v.isActive = isActive
		return nil
	}
}

func WithUserID(userID *valueobject.UserID) EmployeeOption {
	return func(v *Employee) error {
		v.userID = userID
		return nil
	}
}

func WithSchedule(schedule *valueobject.Schedule) EmployeeOption {
	return func(v *Employee) error {
		v.schedule = schedule
		return nil
	}
}

func WithTimestamps(createdAt, updatedAt time.Time) EmployeeOption {
	return func(v *Employee) error {
		if createdAt.IsZero() || updatedAt.IsZero() {
			return errors.New("createdAt and updatedAt cannot be zero")
		}
		v.SetTimeStamps(createdAt, updatedAt)
		return nil
	}
}

func NewEmployee(
	id valueobject.EmployeeID,
	opts ...EmployeeOption,
) (*Employee, error) {
	vet := &Employee{
		Entity:   base.NewEntity(id, time.Now(), time.Now(), 1),
		isActive: true, // Default to active
	}

	for _, opt := range opts {
		if err := opt(vet); err != nil {
			return nil, fmt.Errorf("invalid veterinarian option: %w", err)
		}
	}

	return vet, nil
}

func CreateEmployee(ctx context.Context, name valueobject.PersonName, gender enum.PersonGender, dateOfBirth time.Time, opts ...EmployeeOption) (*Employee, error) {
	person, err := base.CreatePerson(ctx, name, dateOfBirth, gender)
	if err != nil {
		return nil, fmt.Errorf("invalid person data: %w", err)
	}

	vet := &Employee{
		Entity:   base.NewEntity(valueobject.EmployeeID{}, time.Now(), time.Now(), 1),
		isActive: true, // Default to active
		Person:   *person,
	}

	for _, opt := range opts {
		if err := opt(vet); err != nil {
			return nil, domainerr.WrapError(ctx, err, "invalid veterinarian option", "employee", "createEmployee", "create employee")
		}
	}

	if err := vet.validate(ctx); err != nil {
		return nil, err
	}

	return vet, nil
}
