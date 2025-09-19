package pet

import (
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	domainerr "clinic-vet-api/app/core/error"
	"context"
	"fmt"
	"time"
)

type PetUpdateOptions struct {
	Name         *string
	Species      *string
	Breed        *string
	Age          *int
	Gender       *enum.PetGender
	Weight       *float64
	Photo        *string
	Allergies    *string
	Medications  *string
	SpecialNeeds *string
	IsNeutered   *bool
	Microchip    *string
	CustomerID   *valueobject.CustomerID
	IsActive     *bool
	DeletedAt    *time.Time
	Color        *string
}

type PetUpdateOption func(*PetUpdateOptions)

func UpdateName(name string) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.Name = &name
	}
}

func UpdateSpecies(species string) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.Species = &species
	}
}

func UpdateBreed(breed *string) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.Breed = breed
	}
}

func UpdateAge(age *int) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.Age = age
	}
}

func UpdateGender(gender *enum.PetGender) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.Gender = gender
	}
}

func UpdateWeight(weight *float64) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.Weight = weight
	}
}

func UpdatePhoto(photo *string) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.Photo = photo
	}
}

func UpdateMedicalInfo(allergies, medications, specialNeeds *string) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.Allergies = allergies
		opts.Medications = medications
		opts.SpecialNeeds = specialNeeds
	}
}

func UpdateIsNeutered(isNeutered *bool) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.IsNeutered = isNeutered
	}
}

func UpdateMicrochip(microchip *string) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.Microchip = microchip
	}
}

func UpdateCustomerID(customerID valueobject.CustomerID) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.CustomerID = &customerID
	}
}

func UpdateIsActive(isActive bool) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.IsActive = &isActive
	}
}

func UpdateDeletedAt(deletedAt time.Time) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.DeletedAt = &deletedAt
	}
}

func UpdateColor(color *string) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.Color = color
	}
}

func UpdateAllergies(allergies *string) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.Allergies = allergies
	}
}

func UpdateMedications(medications *string) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.Medications = medications
	}
}

func UpdateSpecialNeeds(specialNeeds *string) PetUpdateOption {
	return func(opts *PetUpdateOptions) {
		opts.SpecialNeeds = specialNeeds
	}
}

func (p *Pet) Update(ctx context.Context, options ...PetUpdateOption) error {
	opts := &PetUpdateOptions{}
	operation := "UpdatePet" // Definir la operación para logging

	for _, option := range options {
		option(opts)
	}

	if opts.Name != nil {
		if *opts.Name == "" {
			return domainerr.ValidationError(ctx, "name", "", "name cannot be empty", operation)
		}
		if len(*opts.Name) > 100 {
			return domainerr.ValidationError(ctx, "name", *opts.Name, "name too long", operation)
		}
		p.name = *opts.Name
	}

	if opts.Species != nil {
		if *opts.Species == "" {
			return domainerr.ValidationError(ctx, "species", "", "species cannot be empty", operation)
		}
		petSpecies, err := enum.ParsePetSpecies(*opts.Species)
		if err != nil {
			return domainerr.ValidationError(ctx, "species", *opts.Species, "invalid species", operation)
		}
		p.species = petSpecies
	}

	if opts.Breed != nil {
		if *opts.Breed != "" && len(*opts.Breed) > 50 {
			return domainerr.ValidationError(ctx, "breed", *opts.Breed, "breed too long", operation)
		}
		p.breed = opts.Breed
	}

	if opts.Age != nil {
		if *opts.Age < 0 {
			return domainerr.ValidationError(ctx, "age", fmt.Sprintf("%d", *opts.Age), "age cannot be negative", operation)
		}
		if *opts.Age > 50 {
			return domainerr.ValidationError(ctx, "age", fmt.Sprintf("%d", *opts.Age), "age seems unrealistic", operation)
		}
		p.age = opts.Age
	}

	if opts.Gender != nil {
		if !opts.Gender.IsValid() {
			return domainerr.ValidationError(ctx, "gender", opts.Gender.String(), "invalid gender", operation)
		}
		p.gender = opts.Gender
	}

	if opts.Weight != nil {
		if *opts.Weight <= 0 {
			return domainerr.ValidationError(ctx, "weight", fmt.Sprintf("%f", *opts.Weight), "weight must be positive", operation)
		}
		if *opts.Weight > 1000 {
			return domainerr.ValidationError(ctx, "weight", fmt.Sprintf("%f", *opts.Weight), "weight seems unrealistic", operation)
		}
		p.weight = opts.Weight
	}

	if opts.Photo != nil {
		if *opts.Photo != "" && len(*opts.Photo) > 500 {
			return domainerr.ValidationError(ctx, "photo", *opts.Photo, "photo URL too long", operation)
		}
		p.photo = opts.Photo
	}

	if opts.Allergies != nil {
		if *opts.Allergies != "" && len(*opts.Allergies) > 500 {
			return domainerr.ValidationError(ctx, "allergies", *opts.Allergies, "allergies too long", operation)
		}
		p.allergies = opts.Allergies
	}

	if opts.Medications != nil {
		if *opts.Medications != "" && len(*opts.Medications) > 500 {
			return domainerr.ValidationError(ctx, "medications", *opts.Medications, "medications too long", operation)
		}
		p.currentMedications = opts.Medications
	}

	if opts.SpecialNeeds != nil {
		if *opts.SpecialNeeds != "" && len(*opts.SpecialNeeds) > 500 {
			return domainerr.ValidationError(ctx, "specialNeeds", *opts.SpecialNeeds, "special needs too long", operation)
		}
		p.specialNeeds = opts.SpecialNeeds
	}

	if opts.IsNeutered != nil {
		p.isNeutered = opts.IsNeutered
	}

	if opts.Microchip != nil {
		if *opts.Microchip != "" && len(*opts.Microchip) > 50 {
			return domainerr.ValidationError(ctx, "microchip", *opts.Microchip, "microchip too long", operation)
		}
		p.microchip = opts.Microchip
	}

	if opts.CustomerID != nil {
		if opts.CustomerID.IsZero() {
			return domainerr.ValidationError(ctx, "customerID", opts.CustomerID.String(), "customerID cannot be zero", operation)
		}
		p.customerID = *opts.CustomerID
	}

	if opts.IsActive != nil {
		p.isActive = *opts.IsActive
	}

	if opts.DeletedAt != nil {
		p.deletedAt = opts.DeletedAt
	}

	if opts.Color != nil {
		if *opts.Color != "" && len(*opts.Color) > 50 {
			return domainerr.ValidationError(ctx, "color", *opts.Color, "color too long", operation)
		}
		p.color = opts.Color
	}

	return nil
}
