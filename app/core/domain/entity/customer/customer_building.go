package customer

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

// CustomerOption defines the functional option type
type CustomerOption func(*Customer) error

func WithPhoto(photo string) CustomerOption {
	return func(o *Customer) error {
		o.photo = photo
		return nil
	}
}

func WithFullName(fullName valueobject.PersonName) CustomerOption {
	return func(o *Customer) error {
		return o.UpdateName(fullName)
	}
}

func WithGender(gender enum.PersonGender) CustomerOption {
	return func(o *Customer) error {
		return o.Person.UpdateGender(gender)
	}
}

func WithDateOfBirth(dob time.Time) CustomerOption {
	return func(o *Customer) error {
		return o.UpdateDateOfBirth(dob)
	}
}

func WithPhoneNumber(phoneNumber string) CustomerOption {
	return func(o *Customer) error {
		if err := validatePhoneNumber(phoneNumber); err != nil {
			return err
		}
		o.phoneNumber = phoneNumber
		return nil
	}
}

func WithUserID(userID *valueobject.UserID) CustomerOption {
	return func(o *Customer) error {
		o.userID = userID
		return nil
	}
}

func WithIsActive(isActive bool) CustomerOption {
	return func(o *Customer) error {
		o.isActive = isActive
		return nil
	}
}

func WithPets(pets []pet.Pet) CustomerOption {
	return func(o *Customer) error {
		o.pets = pets
		return nil
	}
}

func WithTimestamp(createdAt, updatedAt time.Time) CustomerOption {
	return func(o *Customer) error {
		o.SetTimeStamps(createdAt, updatedAt)
		return nil
	}
}

// NewCustomer creates a new Owner with functional options
func NewCustomer(
	id valueobject.CustomerID,
	opts ...CustomerOption,
) (*Customer, error) {
	customer := &Customer{
		Entity:   base.NewEntity(id, time.Now(), time.Now(), 1),
		isActive: true, // Default to active
		pets:     []pet.Pet{},
	}

	for _, opt := range opts {
		if err := opt(customer); err != nil {
			return nil, err
		}
	}

	if err := customer.validate(); err != nil {
		return nil, err
	}

	return customer, nil
}

func CreateCustomer(opts ...CustomerOption) (*Customer, error) {
	customer := &Customer{
		Entity:   base.CreateEntity(valueobject.CustomerID{}),
		isActive: true, // Default to active
		pets:     []pet.Pet{},
	}

	for _, opt := range opts {
		if err := opt(customer); err != nil {
			return nil, err
		}
	}

	if err := customer.validate(); err != nil {
		return nil, err
	}

	return customer, nil
}
