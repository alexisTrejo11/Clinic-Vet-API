// Package dto contains data transfer objects for the Pets module API.
package dto

import (
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/modules/pets/application/cqrs/command"
)

// AdminCreatePetRequest represents the request payload for creating a new pet by an admin user
// swagger:model AdminCreatePetRequest
type AdminCreatePetRequest struct {
	PetRequestData
	// ID of the customer who owns the pet
	// Required: true
	// Minimum: 1
	// Example: 123
	CustomerID uint `json:"customer_id" validate:"required,gt=0"`

	// Indicates if the pet record is active in the system
	// Required: true
	// Example: true
	IsActive bool `json:"is_active"`
}

// CustomerCreatePetRequest represents the request payload for creating a new pet by a customer user
// swagger:model CustomerCreatePetRequest
type CustomerCreatePetRequest struct {
	PetRequestData
}

type PetRequestData struct {
	// Name of the pet
	// Required: true
	// Minimum length: 2
	// Maximum length: 100
	// Example: Buddy
	Name string `json:"name" validate:"required,min=2,max=100"`

	// Species of the pet (e.g., Dog, Cat, Bird)
	// Required: true
	// Minimum length: 2
	// Maximum length: 50
	// Example: Dog
	Species string `json:"species" validate:"required,min=2,max=50"`

	// URL to the pet's photo
	// Required: false
	// Format: uri
	// Example: https://example.com/pet-photo.jpg
	Photo *string `json:"photo,omitempty" validate:"omitempty,url"`

	// Breed of the pet
	// Required: false
	// Minimum length: 2
	// Maximum length: 50
	// Example: Golden Retriever
	Breed *string `json:"breed,omitempty" validate:"omitempty,min=2,max=50"`

	// Age of the pet in years
	// Required: false
	// Minimum: 0
	// Maximum: 30
	// Example: 5
	Age *int `json:"age,omitempty" validate:"omitempty,min=0,max=30"`

	// Gender of the pet
	// Required: false
	// Enum: Male, Female, Unknown
	// Example: Male
	Gender *string `json:"gender,omitempty" validate:"omitempty,oneof=Male Female Unknown"`

	// Weight of the pet in kilograms
	// Required: false
	// Minimum: 0.1
	// Maximum: 1000.0
	// Example: 25.5
	Weight *float64 `json:"weight,omitempty" validate:"omitempty,gt=0,lte=1000"`

	// Color of the pet's fur or coat
	// Required: false
	// Minimum length: 2
	// Maximum length: 50
	// Example: Golden
	Color *string `json:"color,omitempty" validate:"omitempty,min=2,max=50"`

	// Microchip identification number (15 digits)
	// Required: false
	// Pattern: ^[0-9]{15}$
	// Example: 123456789012345
	Microchip *string `json:"microchip,omitempty" validate:"omitempty,len=15,numeric"`

	// Indicates if the pet is neutered or spayed
	// Required: false
	// Example: true
	IsNeutered *bool `json:"is_neutered,omitempty"`

	// Known allergies of the pet
	// Required: false
	// Maximum length: 500
	// Example: Pollen, Chicken, Wheat
	Allergies *string `json:"allergies,omitempty" validate:"omitempty,max=500"`

	// Current medications the pet is taking
	// Required: false
	// Maximum length: 500
	// Example: Heartworm prevention, Flea treatment
	CurrentMedications *string `json:"current_medications,omitempty" validate:"omitempty,max=500"`

	// Special needs or care instructions for the pet
	// Required: false
	// Maximum length: 500
	// Example: Requires daily medication, Blind in left eye
	SpecialNeeds *string `json:"special_needs,omitempty" validate:"omitempty,max=500"`
}

func (r *PetRequestData) ToCommand(customerID uint, isActive bool) command.CreatePetCommand {
	cmd := &command.CreatePetCommand{
		Name:               r.Name,
		Photo:              r.Photo,
		CustomerID:         valueobject.NewCustomerID(customerID),
		Species:            r.Species,
		Breed:              r.Breed,
		Age:                r.Age,
		IsNeutered:         r.IsNeutered,
		Weight:             r.Weight,
		Color:              r.Color,
		Microchip:          r.Microchip,
		Allergies:          r.Allergies,
		CurrentMedications: r.CurrentMedications,
		SpecialNeeds:       r.SpecialNeeds,
		IsActive:           isActive,
	}
	return *cmd
}
