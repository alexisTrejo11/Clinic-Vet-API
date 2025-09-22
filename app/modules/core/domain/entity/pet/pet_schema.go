// Package pet defines the Pet entity and its related business logic.
package pet

import (
	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type Pet struct {
	base.Entity[valueobject.PetID]
	name       string
	photo      *string
	species    enum.PetSpecies
	breed      *string
	age        *int
	gender     *enum.PetGender
	color      *string
	microchip  *string
	tattoo     *string
	bloodType  *string
	isNeutered *bool
	customerID valueobject.CustomerID
	isActive   bool
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

func (p *Pet) IsActive() bool {
	return p.isActive
}

func (p *Pet) Tattoo() *string {
	return p.tattoo
}

func (p *Pet) BloodType() *string {
	return p.bloodType
}
