// Package pet defines the Pet entity and its related business logic.
package pet

import (
	"time"

	"clinic-vet-api/app/core/domain/entity/base"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
)

type Pet struct {
	base.Entity[valueobject.PetID]
	name               string
	photo              *string
	species            enum.PetSpecies
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

func (p *Pet) Species() enum.PetSpecies {
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
