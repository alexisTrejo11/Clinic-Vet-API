// Package pet defines the Pet entity and its related business logic.
package pet

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

type Pet struct {
	base.Entity[valueobject.PetID]
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
	customerID         valueobject.CustomerID
	allergies          *string
	currentMedications *string
	specialNeeds       *string
	isActive           bool
	deletedAt          *time.Time
}

func (p *Pet) ID() valueobject.PetID {
	return p.Entity.ID()
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

func (p *Pet) CustomerID() valueobject.CustomerID {
	return p.customerID
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

func (p *Pet) IsDeleted() bool {
	return p.deletedAt != nil
}

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
