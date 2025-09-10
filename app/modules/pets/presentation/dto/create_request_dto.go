// Package dto contains data transfer objects for the Pets module API.
package dto

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
)

// @Description Represents the request body for creating a  new pet.
type CreatePetRequest struct {
	// The name of the pet. (required, min 2, max 100)
	Name string `json:"name" validate:"required,min=2,max=100"`
	// The unique ID of the pet's owner. (required, greater than 0)
	CustomerID uint `json:"owner_id" validate:"required,gt=0"`
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
	return dto.CreatePetData{
		Name:               r.Name,
		Photo:              r.Photo,
		CustomerID:         valueobject.NewCustomerID(r.CustomerID),
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
