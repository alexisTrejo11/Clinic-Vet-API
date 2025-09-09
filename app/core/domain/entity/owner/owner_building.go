package owner

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

// OwnerOption defines the functional option type
type OwnerOption func(*Owner) error

func WithPhoto(photo string) OwnerOption {
	return func(o *Owner) error {
		o.photo = photo
		return nil
	}
}

func WithFullName(fullName valueobject.PersonName) OwnerOption {
	return func(o *Owner) error {
		return o.UpdateName(fullName)
	}
}

func WithGender(gender enum.PersonGender) OwnerOption {
	return func(o *Owner) error {
		return o.Person.UpdateGender(gender)
	}
}

func WithDateOfBirth(dob time.Time) OwnerOption {
	return func(o *Owner) error {
		return o.UpdateDateOfBirth(dob)
	}
}

func WithPhoneNumber(phoneNumber string) OwnerOption {
	return func(o *Owner) error {
		if err := validatePhoneNumber(phoneNumber); err != nil {
			return err
		}
		o.phoneNumber = phoneNumber
		return nil
	}
}

func WithUserID(userID *valueobject.UserID) OwnerOption {
	return func(o *Owner) error {
		o.userID = userID
		return nil
	}
}

func WithIsActive(isActive bool) OwnerOption {
	return func(o *Owner) error {
		o.isActive = isActive
		return nil
	}
}

func WithPets(pets []pet.Pet) OwnerOption {
	return func(o *Owner) error {
		o.pets = pets
		return nil
	}
}

func WithTimestamp(createdAt, updatedAt time.Time) OwnerOption {
	return func(o *Owner) error {
		o.SetTimeStamps(createdAt, updatedAt)
		return nil
	}
}

// NewOwner creates a new Owner with functional options
func NewOwner(
	id valueobject.OwnerID,
	opts ...OwnerOption,
) (*Owner, error) {
	owner := &Owner{
		Entity:   base.NewEntity(id, time.Now(), time.Now(), 1),
		isActive: true, // Default to active
		pets:     []pet.Pet{},
	}

	for _, opt := range opts {
		if err := opt(owner); err != nil {
			return nil, err
		}
	}

	if err := owner.validate(); err != nil {
		return nil, err
	}

	return owner, nil
}

func CreateOwner(opts ...OwnerOption) (*Owner, error) {
	owner := &Owner{
		Entity:   base.CreateEntity(valueobject.OwnerID{}),
		isActive: true, // Default to active
		pets:     []pet.Pet{},
	}

	for _, opt := range opts {
		if err := opt(owner); err != nil {
			return nil, err
		}
	}

	if err := owner.validate(); err != nil {
		return nil, err
	}

	return owner, nil
}
