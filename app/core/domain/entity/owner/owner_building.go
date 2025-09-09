package owner

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

// OwnerOption defines the functional option type
type OwnerOption func(*Owner) error

func WithPhoto(photo string) OwnerOption {
	return func(o *Owner) error {
		if photo != "" && len(photo) > 500 {
			return domainerr.NewValidationError("owner", "photo", "photo URL too long")
		}
		o.photo = photo
		return nil
	}
}

func WithFullName(fullName valueobject.PersonName) OwnerOption {
	return func(o *Owner) error {
		return o.Person.UpdateName(fullName)
	}
}

func WithGender(gender enum.PersonGender) OwnerOption {
	return func(o *Owner) error {
		return o.Person.UpdateGender(gender)
	}
}

func WithDateOfBirth(dob time.Time) OwnerOption {
	return func(o *Owner) error {
		return o.Person.UpdateDateOfBirth(dob)
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

	// Apply all options
	for _, opt := range opts {
		if err := opt(owner); err != nil {
			return nil, err
		}
	}

	// Final validation
	if err := owner.validate(); err != nil {
		return nil, err
	}

	return owner, nil
}
