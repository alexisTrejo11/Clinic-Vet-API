package pet

import (
	"context"
	"time"

	"clinic-vet-api/app/core/domain/entity/base"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
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
		p.species = enum.PetSpecies(species)
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

	if p.weight != nil {
		if *p.weight <= 0 {
			return WeightInvalidError(ctx, *p.weight, operation)
		}
		if *p.weight > 1000 {
			return WeightUnrealisticError(ctx, *p.weight, operation)
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

	if p.currentMedications != nil && len(*p.currentMedications) > 500 {
		return MedicationsTooLongError(ctx, len(*p.currentMedications), operation)
	}

	if p.allergies != nil && len(*p.allergies) > 500 {
		return AllergiesTooLongError(ctx, len(*p.allergies), operation)
	}

	if p.specialNeeds != nil && len(*p.specialNeeds) > 500 {
		return SpecialNeedsTooLongError(ctx, len(*p.specialNeeds), operation)
	}

	return nil
}

func NewPetWithContext(
	ctx context.Context,
	id valueobject.PetID,
	customerID valueobject.CustomerID,
	opts ...PetOption,
) (*Pet, error) {

	pet := &Pet{
		Entity:     base.NewEntity(id, time.Now(), time.Now(), 1),
		customerID: customerID,
		isActive:   true,
	}

	for _, opt := range opts {
		if err := opt(pet); err != nil {
			return nil, err
		}
	}
	return pet, nil
}

func CreatePetWithContext(
	ctx context.Context,
	customerID valueobject.CustomerID,
	opts ...PetOption,
) (*Pet, error) {
	operation := "CreatePetWithContext"

	if customerID.IsZero() {
		return nil, CustomerIDRequiredError(ctx, operation)
	}

	pet := &Pet{
		Entity:     base.NewEntity(valueobject.PetID{}, time.Now(), time.Now(), 1),
		customerID: customerID,
		isActive:   true,
	}

	for _, opt := range opts {
		if err := opt(pet); err != nil {
			return nil, err
		}
	}

	if err := pet.validate(ctx); err != nil {
		return nil, err
	}
	return pet, nil
}

func NewPet(id valueobject.PetID, customerID valueobject.CustomerID, opts ...PetOption) (*Pet, error) {
	ctx := context.Background()
	return NewPetWithContext(ctx, id, customerID, opts...)
}

func CreatePet(customerID valueobject.CustomerID, opts ...PetOption) (*Pet, error) {
	ctx := context.Background()
	return CreatePetWithContext(ctx, customerID, opts...)
}
