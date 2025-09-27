// Package pet defines the Pet entity and its related business logic.
package pet

import (
	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type Pet struct {
	base.Entity[valueobject.PetID]
	name       string
	species    enum.PetSpecies
	gender     enum.PetGender
	photo      *string
	breed      *string
	age        *int
	color      *string
	microchip  *string
	tattoo     *string
	bloodType  *string
	isNeutered *bool
	customerID valueobject.CustomerID
	isActive   bool
}

type PetBuilder struct{ pet *Pet }

func NewPetBuilder() *PetBuilder {
	return &PetBuilder{pet: &Pet{}}
}

func (b *PetBuilder) WithID(id valueobject.PetID) *PetBuilder {
	b.pet.Entity.SetID(id)
	return b
}

func (b *PetBuilder) WithCustomerID(customerID valueobject.CustomerID) *PetBuilder {
	b.pet.customerID = customerID
	return b
}

func (b *PetBuilder) WithName(name string) *PetBuilder {
	b.pet.name = name
	return b
}

func (b *PetBuilder) WithPhoto(photo *string) *PetBuilder {
	b.pet.photo = photo
	return b
}

func (b *PetBuilder) WithSpecies(species enum.PetSpecies) *PetBuilder {
	b.pet.species = species
	return b
}

func (b *PetBuilder) WithBreed(breed *string) *PetBuilder {
	b.pet.breed = breed
	return b
}

func (b *PetBuilder) WithAge(age *int) *PetBuilder {
	b.pet.age = age
	return b
}

func (b *PetBuilder) WithGender(gender enum.PetGender) *PetBuilder {
	b.pet.gender = gender
	return b
}

func (b *PetBuilder) WithColor(color *string) *PetBuilder {
	b.pet.color = color
	return b
}

func (b *PetBuilder) WithBloodType(bloodType *string) *PetBuilder {
	b.pet.bloodType = bloodType
	return b
}

func (b *PetBuilder) WithTattoo(tattoo *string) *PetBuilder {
	b.pet.tattoo = tattoo
	return b
}

func (b *PetBuilder) WithMicrochip(microchip *string) *PetBuilder {
	b.pet.microchip = microchip
	return b
}

func (b *PetBuilder) WithIsNeutered(isNeutered *bool) *PetBuilder {
	b.pet.isNeutered = isNeutered
	return b
}

func (b *PetBuilder) WithTimeStamps(createdAt, updatedAt time.Time) *PetBuilder {
	b.pet.SetTimeStamps(createdAt, updatedAt)
	return b
}

func (b *PetBuilder) WithIsActive(isActive bool) *PetBuilder {
	b.pet.isActive = isActive
	return b
}

func (b *PetBuilder) Build() *Pet {
	return b.pet
}

func (p *Pet) ID() valueobject.PetID              { return p.Entity.ID() }
func (p *Pet) Name() string                       { return p.name }
func (p *Pet) Photo() *string                     { return p.photo }
func (p *Pet) Species() enum.PetSpecies           { return p.species }
func (p *Pet) Breed() *string                     { return p.breed }
func (p *Pet) Age() *int                          { return p.age }
func (p *Pet) Gender() enum.PetGender             { return p.gender }
func (p *Pet) Color() *string                     { return p.color }
func (p *Pet) Microchip() *string                 { return p.microchip }
func (p *Pet) IsNeutered() *bool                  { return p.isNeutered }
func (p *Pet) CustomerID() valueobject.CustomerID { return p.customerID }
func (p *Pet) BloodType() *string                 { return p.bloodType }
func (p *Pet) Tattoo() *string                    { return p.tattoo }
func (p *Pet) IsActive() bool                     { return p.isActive }
