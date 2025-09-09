package pet

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

// PetOption defines the functional option type
type PetOption func(*Pet) error

func WithName(name string) PetOption {
	return func(p *Pet) error {
		if name == "" {
			return domainerr.NewValidationError("pet", "name", "name is required")
		}
		if len(name) > 100 {
			return domainerr.NewValidationError("pet", "name", "name too long")
		}
		p.name = name
		return nil
	}
}

func WithPhoto(photo *string) PetOption {
	return func(p *Pet) error {
		if photo != nil && len(*photo) > 500 {
			return domainerr.NewValidationError("pet", "photo", "photo URL too long")
		}
		p.photo = photo
		return nil
	}
}

func WithSpecies(species string) PetOption {
	return func(p *Pet) error {
		if species == "" {
			return domainerr.NewValidationError("pet", "species", "species is required")
		}
		if len(species) > 50 {
			return domainerr.NewValidationError("pet", "species", "species too long")
		}
		p.species = species
		return nil
	}
}

func WithBreed(breed *string) PetOption {
	return func(p *Pet) error {
		if breed != nil && len(*breed) > 50 {
			return domainerr.NewValidationError("pet", "breed", "breed too long")
		}
		p.breed = breed
		return nil
	}
}

func WithAge(age *int) PetOption {
	return func(p *Pet) error {
		if age != nil && *age < 0 {
			return domainerr.NewValidationError("pet", "age", "age cannot be negative")
		}
		if age != nil && *age > 50 {
			return domainerr.NewValidationError("pet", "age", "age seems unrealistic")
		}
		p.age = age
		return nil
	}
}

func WithGender(gender *enum.PetGender) PetOption {
	return func(p *Pet) error {
		if gender != nil && !gender.IsValid() {
			return domainerr.NewValidationError("pet", "gender", "invalid gender")
		}
		p.gender = gender
		return nil
	}
}

func WithWeight(weight *float64) PetOption {
	return func(p *Pet) error {
		if weight != nil && *weight <= 0 {
			return domainerr.NewValidationError("pet", "weight", "weight must be positive")
		}
		if weight != nil && *weight > 1000 {
			return domainerr.NewValidationError("pet", "weight", "weight seems unrealistic")
		}
		p.weight = weight
		return nil
	}
}

func WithColor(color *string) PetOption {
	return func(p *Pet) error {
		if color != nil && len(*color) > 30 {
			return domainerr.NewValidationError("pet", "color", "color too long")
		}
		p.color = color
		return nil
	}
}

func WithMicrochip(microchip *string) PetOption {
	return func(p *Pet) error {
		if microchip != nil && len(*microchip) > 50 {
			return domainerr.NewValidationError("pet", "microchip", "microchip too long")
		}
		p.microchip = microchip
		return nil
	}
}

func WithIsNeutered(isNeutered *bool) PetOption {
	return func(p *Pet) error {
		p.isNeutered = isNeutered
		return nil
	}
}

func WithAllergies(allergies *string) PetOption {
	return func(p *Pet) error {
		if allergies != nil && len(*allergies) > 500 {
			return domainerr.NewValidationError("pet", "allergies", "allergies too long")
		}
		p.allergies = allergies
		return nil
	}
}

func WithCurrentMedications(medications *string) PetOption {
	return func(p *Pet) error {
		if medications != nil && len(*medications) > 500 {
			return domainerr.NewValidationError("pet", "medications", "medications too long")
		}
		p.currentMedications = medications
		return nil
	}
}

func WithTimeStamps(createdAt, updatedAt time.Time) PetOption {
	return func(p *Pet) error {
		if createdAt.IsZero() || updatedAt.IsZero() {
			return domainerr.NewValidationError("pet", "timestamps", "createdAt and updatedAt are required")
		}
		p.SetTimeStamps(createdAt, updatedAt)
		return nil
	}
}

func WithSpecialNeeds(specialNeeds *string) PetOption {
	return func(p *Pet) error {
		if specialNeeds != nil && len(*specialNeeds) > 500 {
			return domainerr.NewValidationError("pet", "specialNeeds", "special needs too long")
		}
		p.specialNeeds = specialNeeds
		return nil
	}
}

func WithIsActive(isActive bool) PetOption {
	return func(p *Pet) error {
		p.isActive = isActive
		return nil
	}
}

// NewPet creates a new Pet with functional options
func NewPet(
	id valueobject.PetID,
	ownerID valueobject.OwnerID,
	opts ...PetOption,
) (*Pet, error) {
	if id.IsZero() {
		return nil, domainerr.NewValidationError("pet", "id", "pet ID is required")
	}
	if id.IsZero() {
		return nil, domainerr.NewValidationError("pet", "owner-ID", "owner ID is required")
	}

	pet := &Pet{
		Entity:   base.NewEntity(id, time.Now(), time.Now(), 1),
		ownerID:  ownerID,
		isActive: true, // Default to active
	}

	// Apply all options
	for _, opt := range opts {
		if err := opt(pet); err != nil {
			return nil, err
		}
	}

	if err := pet.validate(); err != nil {
		return nil, err
	}
	return pet, nil
}

// Validation
func (p *Pet) validate() error {
	if p.name == "" {
		return domainerr.NewValidationError("pet", "name", "name is required")
	}
	if p.species == "" {
		return domainerr.NewValidationError("pet", "species", "species is required")
	}
	return nil
}
