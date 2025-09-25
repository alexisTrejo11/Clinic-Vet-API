package pet

import (
	"context"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

// PetOption defines the functional option type
type PetOption func(*Pet)

func WithName(name string) PetOption {
	return func(p *Pet) {
		p.name = name
	}
}

func WithPhoto(photo *string) PetOption {
	return func(p *Pet) {
		p.photo = photo
	}
}

func WithSpecies(species string) PetOption {
	return func(p *Pet) {
		p.species = enum.PetSpecies(species)
	}
}

func WithBreed(breed *string) PetOption {
	return func(p *Pet) {
		p.breed = breed
	}
}

func WithAge(age *int) PetOption {
	return func(p *Pet) {
		p.age = age
	}
}

func WithGender(gender *enum.PetGender) PetOption {
	return func(p *Pet) {
		p.gender = gender
	}
}

func WithColor(color *string) PetOption {
	return func(p *Pet) {
		p.color = color
	}
}

func WithBloodType(bloodType *string) PetOption {
	return func(p *Pet) {
		p.bloodType = bloodType
	}
}

func WithTattoo(tattoo *string) PetOption {
	return func(p *Pet) {
		p.tattoo = tattoo
	}
}

func WithMicrochip(microchip *string) PetOption {
	return func(p *Pet) {
		p.microchip = microchip
	}
}

func WithIsNeutered(isNeutered *bool) PetOption {
	return func(p *Pet) {
		p.isNeutered = isNeutered
	}
}

func WithTimeStamps(createdAt, updatedAt time.Time) PetOption {
	return func(p *Pet) {
		p.SetTimeStamps(createdAt, updatedAt)
	}
}

func WithIsActive(isActive bool) PetOption {
	return func(p *Pet) {
		p.isActive = isActive
	}
}

func (p *Pet) validate(ctx context.Context) error {
	operation := "ValidatePet"

	if p.customerID.IsZero() {
		return CustomerIDRequiredError(ctx, operation)
	}

	if p.name == "" {
		return NameRequiredError(ctx, operation)
	}
	if len(p.name) > 100 {
		return NameTooLongError(ctx, len(p.name), operation)
	}

	if p.species == "" {
		return SpeciesRequiredError(ctx, operation)
	}

	if !p.species.IsValid() {
		return InvalidSpeciesError(ctx, p.species.String())
	}

	if p.age != nil {
		if *p.age < 0 {
			return AgeInvalidError(ctx, *p.age, operation)
		}
		if *p.age > 50 {
			return AgeUnrealisticError(ctx, *p.age, operation)
		}
	}

	if p.gender != nil && !p.gender.IsValid() {
		return GenderInvalidError(ctx, *p.gender, operation)
	}

	if p.breed != nil && len(*p.breed) > 50 {
		return BreedTooLongError(ctx, len(*p.breed), operation)
	}

	if p.microchip != nil && len(*p.microchip) > 50 {
		return MicrochipTooLongError(ctx, len(*p.microchip), operation)
	}

	if p.color != nil && len(*p.color) > 30 {
		return ColorTooLongError(ctx, len(*p.color), operation)
	}

	if p.photo != nil && len(*p.photo) > 500 {
		return PhotoURLTooLongError(ctx, len(*p.photo), operation)
	}

	return nil
}

func CreatePet(
	ctx context.Context,
	customerID valueobject.CustomerID,
	opts ...PetOption,
) (*Pet, error) {
	operation := "CreatePetWithContext"

	if customerID.IsZero() {
		return nil, CustomerIDRequiredError(ctx, operation)
	}

	now := time.Now()
	pet := &Pet{
		Entity:     base.NewEntity(valueobject.PetID{}, &now, &now, 1),
		customerID: customerID,
		isActive:   true,
	}

	for _, opt := range opts {
		opt(pet)
	}

	if err := pet.validate(ctx); err != nil {
		return nil, err
	}
	return pet, nil
}

func NewPet(id valueobject.PetID, customerID valueobject.CustomerID, opts ...PetOption) (*Pet, error) {
	pet := &Pet{
		Entity:     base.NewEntity(id, nil, nil, 0),
		customerID: customerID,
	}

	for _, opt := range opts {
		opt(pet)
	}

	return pet, nil
}
