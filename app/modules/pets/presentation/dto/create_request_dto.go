// Package dto contains data transfer objects for the Pets module API.
package dto

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/pets/application/command"
)

// AdminCreatePetRequest represents the request payload for creating a new pet by an admin user
// swagger:model AdminCreatePetRequest
type AdminPetCreateExtraFields struct {
	PetRequestData
	// ID of the customer who owns the pet
	// Required: false
	// Minimum: 1
	// Example: 123
	CustomerID *uint `json:"customer_id" validate:"required,gt=0"`

	// Indicates if the pet record is active in the system
	// Required: true
	// Example: true
	IsActive bool `json:"is_active" validate:"required"`
}

// PetRequestData encapsulates the common fields for creating or updating a pet
// swagger:model PetRequestData
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
	// Required: true
	// Enum: male, female, unknown
	// Example: male
	Gender string `json:"gender" validate:"required,oneof=male female unknown"`

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

	// Tattoo identification (if applicable)
	// Required: false
	// Minimum length: 2
	// Maximum length: 20
	// Example: A12345
	Tattoo *string `json:"tattoo,omitempty"`

	// Blood type of the pet (if known)
	// Required: false
	// Minimum length: 1
	// Maximum length: 3
	// Example: A+
	BloodType *string `json:"blood_type,omitempty"`
}

func (r *PetRequestData) ToCommand(customerID uint, isActive bool) command.CreatePetCommand {
	cmd := &command.CreatePetCommand{
		Name:       r.Name,
		CustomerID: valueobject.NewCustomerID(customerID),
		Species:    enum.PetSpecies(r.Species),
		Gender:     enum.PetGender(r.Gender),
		Photo:      r.Photo,
		Breed:      r.Breed,
		Age:        r.Age,
		IsNeutered: r.IsNeutered,
		Color:      r.Color,
		Microchip:  r.Microchip,
		IsActive:   isActive,
	}
	return *cmd
}
