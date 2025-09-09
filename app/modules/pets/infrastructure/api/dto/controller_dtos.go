// Package dto contains data transfer objects for the Pets module API.
package dto

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

// @Description Represents the request body for creating a new pet.
type CreatePetRequest struct {
	// The name of the pet. (required, min 2, max 100)
	Name string `json:"name" validate:"required,min=2,max=100"`
	// The unique ID of the pet's owner. (required, greater than 0)
	OwnerID int `json:"owner_id" validate:"required,gt=0"`
	// The species of the pet. (required, min 2, max 50)
	Species string `json:"species" validate:"required,min=2,max=50"`
	// The URL of the pet's photo. (optional, must be a valid URL)
	Photo *string `json:"photo,omitempty" validate:"omitempty,url"`
	// The breed of the pet. (optional, min 2, max 50)
	Breed *string `json:"breed,omitempty" validate:"omitempty,min=2,max=50"`
	// The age of the pet in years. (optional)
	Age *int `json:"age,omitempty" validate:"omitempty"`
	// The gender of the pet. (optional, must be one of "Male", "Female", or "Unknown")
	Gender *string `json:"gender,omitempty" validate:"omitempty,oneof=Male Female Unknown"`
	// The weight of the pet in kilograms. (optional, greater than 0, less than or equal to 1000)
	Weight *float64 `json:"weight,omitempty" validate:"omitempty,gt=0,lte=1000"`
	// The color of the pet's fur/coat. (optional, min 2, max 50)
	Color *string `json:"color,omitempty" validate:"omitempty,min=2,max=50"`
	// The pet's microchip number. (optional, must be 15 digits)
	Microchip *string `json:"microchip,omitempty" validate:"omitempty,len=15,numeric"`
	// Indicates if the pet is neutered. (optional)
	IsNeutered *bool `json:"is_neutered,omitempty"`
	// A list of the pet's known allergies. (optional, max 500)
	Allergies *string `json:"allergies,omitempty" validate:"omitempty,max=500"`
	// The pet's current medications. (optional, max 500)
	CurrentMedications *string `json:"current_medications,omitempty" validate:"omitempty,max=500"`
	// Any special needs the pet may have. (optional, max 500)
	SpecialNeeds *string `json:"special_needs,omitempty" validate:"omitempty,max=500"`
	// Indicates if the pet's record is active. (required)
	IsActive bool `json:"is_active"`
}

func (r *CreatePetRequest) ToCreateData() dto.CreatePetData {
	ownerID, _ := valueobject.NewOwnerID(r.OwnerID)

	return dto.CreatePetData{
		Name:               r.Name,
		Photo:              r.Photo,
		OwnerID:            ownerID,
		Species:            r.Species,
		Breed:              r.Breed,
		Age:                r.Age,
		Gender:             r.Gender,
		Weight:             r.Weight,
		Color:              r.Color,
		Microchip:          r.Microchip,
		IsNeutered:         r.IsNeutered,
		Allergies:          r.Allergies,
		CurrentMedications: r.CurrentMedications,
	}
}

type UpdatePetRequest struct {
	Name               *string  `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Photo              *string  `json:"photo,omitempty" validate:"omitempty,url"`
	Species            *string  `json:"species,omitempty" validate:"omitempty,min=2,max=50"`
	Breed              *string  `json:"breed,omitempty" validate:"omitempty,min=2,max=50"`
	Age                *int     `json:"age,omitempty" validate:"omitempty,min=0"`
	Gender             *string  `json:"gender,omitempty" validate:"omitempty,oneof=Male Female Unknown"`
	Weight             *float64 `json:"weight,omitempty" validate:"omitempty,gt=0,lte=1000"`
	Color              *string  `json:"color,omitempty" validate:"omitempty,min=2,max=50"`
	Microchip          *string  `json:"microchip,omitempty" validate:"omitempty,len=15,numeric"`
	IsNeutered         *bool    `json:"is_neutered,omitempty"`
	OwnerID            *int     `json:"owner_id,omitempty" validate:"omitempty,gt=0"`
	Allergies          *string  `json:"allergies,omitempty" validate:"omitempty,max=500"`
	CurrentMedications *string  `json:"current_medications,omitempty" validate:"omitempty,max=500"`
	SpecialNeeds       *string  `json:"special_needs,omitempty" validate:"omitempty,max=500"`
	IsActive           *bool    `json:"is_active,omitempty"`
}

func (r UpdatePetRequest) ToUpdatePet(petIDInt int) dto.PetUpdate {
	petID, _ := valueobject.NewPetID(petIDInt)
	petUpdate := dto.PetUpdate{
		PetID:              petID,
		Photo:              r.Photo,
		Breed:              r.Breed,
		Age:                r.Age,
		Gender:             r.Gender,
		Weight:             r.Weight,
		Color:              r.Color,
		Microchip:          r.Microchip,
		IsNeutered:         r.IsNeutered,
		Allergies:          r.Allergies,
		CurrentMedications: r.CurrentMedications,
		SpecialNeeds:       r.SpecialNeeds,
	}

	if r.OwnerID != nil {
		ownerID, _ := valueobject.NewOwnerID(*r.OwnerID)
		petUpdate.OwnerID = &ownerID
	}

	if r.Name != nil {
		petUpdate.Name = r.Name
	}

	if r.Species != nil {
		petUpdate.Species = r.Species
	}

	petUpdate.IsActive = r.IsActive

	return petUpdate
}

type PetSearchParams struct {
	page.PageInput
	Filters PetFilters `json:"filters"`
	OrderBy PetOrderBy `json:"order_by" validate:"required,oneof=name species age created_at"`
}

type PetFilters struct {
	Name    *string `json:"name"`
	OwnerID *int    `json:"owner_id"`
	Species *string `json:"species"`
	Breed   *string `json:"breed"`
	MinAge  *int    `json:"min_age" validate:"omitempty,gte=0,lte=100"`
	MaxAge  *int    `json:"max_age" validate:"omitempty,gte=0,lte=100"`
}

type PetOrderBy string
