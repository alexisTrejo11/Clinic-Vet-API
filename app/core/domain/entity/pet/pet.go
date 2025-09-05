package pet

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

type Pet struct {
	base.Entity
	name               string
	photo              *string
	species            string
	breed              *string
	age                *int
	gender             *enum.PetGender
	weight             *float64
	color              *string
	microchip          *string
	isNeutered         *bool
	ownerID            valueobject.OwnerID
	allergies          *string
	currentMedications *string
	specialNeeds       *string
	isActive           bool
	deletedAt          *time.Time
}

// PetOption defines the functional option type
type PetOption func(*Pet) error

// Functional options
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
	if id.GetValue() == 0 {
		return nil, domainerr.NewValidationError("pet", "id", "pet ID is required")
	}
	if ownerID.GetValue() == 0 {
		return nil, domainerr.NewValidationError("pet", "ownerID", "owner ID is required")
	}

	pet := &Pet{
		Entity:   base.NewEntity(id),
		ownerID:  ownerID,
		isActive: true, // Default to active
	}

	// Apply all options
	for _, opt := range opts {
		if err := opt(pet); err != nil {
			return nil, err
		}
	}

	// Final validation
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

// Getters
func (p *Pet) ID() valueobject.PetID {
	return p.ID()
}

func (p *Pet) Name() string {
	return p.name
}

func (p *Pet) Photo() *string {
	return p.photo
}

func (p *Pet) Species() string {
	return p.species
}

func (p *Pet) Breed() *string {
	return p.breed
}

func (p *Pet) Age() *int {
	return p.age
}

func (p *Pet) Gender() *enum.PetGender {
	return p.gender
}

func (p *Pet) Weight() *float64 {
	return p.weight
}

func (p *Pet) Color() *string {
	return p.color
}

func (p *Pet) Microchip() *string {
	return p.microchip
}

func (p *Pet) IsNeutered() *bool {
	return p.isNeutered
}

func (p *Pet) OwnerID() valueobject.OwnerID {
	return p.ownerID
}

func (p *Pet) Allergies() *string {
	return p.allergies
}

func (p *Pet) CurrentMedications() *string {
	return p.currentMedications
}

func (p *Pet) SpecialNeeds() *string {
	return p.specialNeeds
}

func (p *Pet) IsActive() bool {
	return p.isActive
}

func (p *Pet) DeletedAt() *time.Time {
	return p.deletedAt
}

// Business logic methods
func (p *Pet) UpdateName(newName string) error {
	if newName == "" {
		return domainerr.NewValidationError("pet", "name", "name cannot be empty")
	}
	if len(newName) > 100 {
		return domainerr.NewValidationError("pet", "name", "name too long")
	}
	p.name = newName
	p.IncrementVersion()
	return nil
}

func (p *Pet) UpdateSpecies(newSpecies string) error {
	if newSpecies == "" {
		return domainerr.NewValidationError("pet", "species", "species cannot be empty")
	}
	if len(newSpecies) > 50 {
		return domainerr.NewValidationError("pet", "species", "species too long")
	}
	p.species = newSpecies
	p.IncrementVersion()
	return nil
}

func (p *Pet) UpdateBreed(newBreed *string) error {
	if newBreed != nil && len(*newBreed) > 50 {
		return domainerr.NewValidationError("pet", "breed", "breed too long")
	}
	p.breed = newBreed
	p.IncrementVersion()
	return nil
}

func (p *Pet) UpdateAge(newAge *int) error {
	if newAge != nil && *newAge < 0 {
		return domainerr.NewValidationError("pet", "age", "age cannot be negative")
	}
	if newAge != nil && *newAge > 50 {
		return domainerr.NewValidationError("pet", "age", "age seems unrealistic")
	}
	p.age = newAge
	p.IncrementVersion()
	return nil
}

func (p *Pet) UpdateGender(newGender *enum.PetGender) error {
	if newGender != nil && !newGender.IsValid() {
		return domainerr.NewValidationError("pet", "gender", "invalid gender")
	}
	p.gender = newGender
	p.IncrementVersion()
	return nil
}

func (p *Pet) UpdateWeight(newWeight *float64) error {
	if newWeight != nil && *newWeight <= 0 {
		return domainerr.NewValidationError("pet", "weight", "weight must be positive")
	}
	if newWeight != nil && *newWeight > 1000 {
		return domainerr.NewValidationError("pet", "weight", "weight seems unrealistic")
	}
	p.weight = newWeight
	p.IncrementVersion()
	return nil
}

func (p *Pet) UpdatePhoto(newPhoto *string) error {
	if newPhoto != nil && len(*newPhoto) > 500 {
		return domainerr.NewValidationError("pet", "photo", "photo URL too long")
	}
	p.photo = newPhoto
	p.IncrementVersion()
	return nil
}

func (p *Pet) UpdateMedicalInfo(allergies, medications, specialNeeds *string) error {
	if allergies != nil && len(*allergies) > 500 {
		return domainerr.NewValidationError("pet", "allergies", "allergies too long")
	}
	if medications != nil && len(*medications) > 500 {
		return domainerr.NewValidationError("pet", "medications", "medications too long")
	}
	if specialNeeds != nil && len(*specialNeeds) > 500 {
		return domainerr.NewValidationError("pet", "specialNeeds", "special needs too long")
	}

	p.allergies = allergies
	p.currentMedications = medications
	p.specialNeeds = specialNeeds
	p.IncrementVersion()
	return nil
}

func (p *Pet) UpdateNeuteredStatus(isNeutered *bool) error {
	p.isNeutered = isNeutered
	p.IncrementVersion()
	return nil
}

func (p *Pet) UpdateMicrochip(microchip *string) error {
	if microchip != nil && len(*microchip) > 50 {
		return domainerr.NewValidationError("pet", "microchip", "microchip too long")
	}
	p.microchip = microchip
	p.IncrementVersion()
	return nil
}

func (p *Pet) Activate() error {
	if p.isActive {
		return nil // Already active
	}
	p.isActive = true
	p.deletedAt = nil
	p.IncrementVersion()
	return nil
}

func (p *Pet) Deactivate() error {
	if !p.isActive {
		return nil // Already inactive
	}
	p.isActive = false
	p.IncrementVersion()
	return nil
}

func (p *Pet) SoftDelete() error {
	if p.deletedAt != nil {
		return nil // Already deleted
	}
	now := time.Now()
	p.isActive = false
	p.deletedAt = &now
	p.IncrementVersion()
	return nil
}

func (p *Pet) Restore() error {
	if p.deletedAt == nil {
		return nil // Not deleted
	}
	p.isActive = true
	p.deletedAt = nil
	p.IncrementVersion()
	return nil
}

func (p *Pet) IsDeleted() bool {
	return p.deletedAt != nil
}

func (p *Pet) RequiresVaccination() bool {
	// Logic to determine if pet needs vaccination based on age and species
	if p.age == nil {
		return false
	}

	// Puppies/kittens need more frequent vaccinations
	if *p.age < 1 {
		return true
	}

	// Adult pets need annual vaccinations
	return true
}

func (p *Pet) LifeStage() string {
	if p.age == nil {
		return "unknown"
	}

	age := *p.age
	switch {
	case age < 1:
		return "baby"
	case age < 3:
		return "young"
	case age < 8:
		return "adult"
	default:
		return "senior"
	}
}

func (p *Pet) HasMedicalConditions() bool {
	return (p.allergies != nil && *p.allergies != "") ||
		(p.currentMedications != nil && *p.currentMedications != "") ||
		(p.specialNeeds != nil && *p.specialNeeds != "")
}
