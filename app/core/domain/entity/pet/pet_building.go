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
		p.name = name
		return nil
	}
}

func WithPhoto(photo *string) PetOption {
	return func(p *Pet) error {
		p.photo = photo
		return nil
	}
}

func WithSpecies(species string) PetOption {
	return func(p *Pet) error {
		p.species = species
		return nil
	}
}

func WithBreed(breed *string) PetOption {
	return func(p *Pet) error {
		p.breed = breed
		return nil
	}
}

func WithAge(age *int) PetOption {
	return func(p *Pet) error {
		p.age = age
		return nil
	}
}

func WithGender(gender *enum.PetGender) PetOption {
	return func(p *Pet) error {
		p.gender = gender
		return nil
	}
}

func WithWeight(weight *float64) PetOption {
	return func(p *Pet) error {
		p.weight = weight
		return nil
	}
}

func WithColor(color *string) PetOption {
	return func(p *Pet) error {
		p.color = color
		return nil
	}
}

func WithMicrochip(microchip *string) PetOption {
	return func(p *Pet) error {
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
		p.allergies = allergies
		return nil
	}
}

func WithCurrentMedications(medications *string) PetOption {
	return func(p *Pet) error {
		p.currentMedications = medications
		return nil
	}
}

func WithTimeStamps(createdAt, updatedAt time.Time) PetOption {
	return func(p *Pet) error {
		p.SetTimeStamps(createdAt, updatedAt)
		return nil
	}
}

func WithSpecialNeeds(specialNeeds *string) PetOption {
	return func(p *Pet) error {
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
func NewPet(id valueobject.PetID, customerID valueobject.CustomerID, opts ...PetOption) (*Pet, error) {
	pet := &Pet{
		Entity:     base.NewEntity(id, time.Now(), time.Now(), 1),
		customerID: customerID,
		isActive:   true, // Default to active
	}

	for _, opt := range opts {
		if err := opt(pet); err != nil {
			return nil, err
		}
	}
	return pet, nil
}

func CreatePet(customerID valueobject.CustomerID, opts ...PetOption) (*Pet, error) {
	pet := &Pet{
		Entity:     base.NewEntity(valueobject.PetID{}, time.Now(), time.Now(), 1),
		customerID: customerID,
		isActive:   true, // Default to active
	}

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
	if p.currentMedications != nil && len(*p.currentMedications) > 500 {
		return domainerr.NewValidationError("pet", "medications", "medications too long")
	}

	if p.age != nil && *p.age < 0 {
		return domainerr.NewValidationError("pet", "age", "age cannot be negative")
	}
	if p.age != nil && *p.age > 50 {
		return domainerr.NewValidationError("pet", "age", "age seems unrealistic")
	}

	if p.breed != nil && len(*p.breed) > 50 {
		return domainerr.NewValidationError("pet", "breed", "breed too long")
	}

	if p.microchip != nil && len(*p.microchip) > 50 {
		return domainerr.NewValidationError("pet", "microchip", "microchip too long")
	}

	if p.specialNeeds != nil && len(*p.specialNeeds) > 500 {
		return domainerr.NewValidationError("pet", "specialNeeds", "special needs too long")
	}

	if p.allergies != nil && len(*p.allergies) > 500 {
		return domainerr.NewValidationError("pet", "allergies", "allergies too long")
	}

	if p.color != nil && len(*p.color) > 30 {
		return domainerr.NewValidationError("pet", "color", "color too long")
	}

	if p.name == "" {
		return domainerr.NewValidationError("pet", "name", "name is required")
	}

	if p.weight != nil && *p.weight <= 0 {
		return domainerr.NewValidationError("pet", "weight", "weight must be positive")
	}
	if p.weight != nil && *p.weight > 1000 {
		return domainerr.NewValidationError("pet", "weight", "weight seems unrealistic")
	}

	if p.gender != nil && !p.gender.IsValid() {
		return domainerr.NewValidationError("pet", "gender", "invalid gender")
	}

	if p.species == "" {
		return domainerr.NewValidationError("pet", "species", "species is required")
	}

	if p.species == "" {
		return domainerr.NewValidationError("pet", "species", "species is required")
	}
	if len(p.species) > 50 {
		return domainerr.NewValidationError("pet", "species", "species too long")
	}

	if p.photo != nil && len(*p.photo) > 500 {
		return domainerr.NewValidationError("pet", "photo", "photo URL too long")
	}

	if p.name == "" {
		return domainerr.NewValidationError("pet", "name", "name is required")
	}
	if len(p.name) > 100 {
		return domainerr.NewValidationError("pet", "name", "name too long")
	}

	return nil
}
