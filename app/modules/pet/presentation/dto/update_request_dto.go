package dto

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/pet/application/command"
)

// AdminUpdatePetRequest represents the payload for updating an existing pet's information by an admin
// swagger:model AdminUpdatePetRequest
type AdminUpdatePetExtraFields struct {
	// ID of the customer who owns the pet
	// Required: false
	// Minimum: 1
	// Example: 123
	CustomerID *uint `json:"customer_id,omitempty" validate:"omitempty,gt=0"`

	// IsActive indicates if the pet record is active in the system
	// Required: false
	// Example: true
	IsActive *bool `json:"is_active,omitempty"`
}

// CustomerUpdatePetRequest represents the payload for updating an existing pet's information by a customer
// swagger:model CustomerUpdatePetRequest
type CustomerUpdatePetRequest struct {
	UpdatePetRequest
}

// UpdatePetRequest represents the payload for updating an existing pet's information
// swagger:model UpdatePetRequest
type UpdatePetRequest struct {
	// Pet's name
	// Required: false
	// Minimum length: 2
	// Maximum length: 100
	// Example: Max
	Name *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`

	// URL to pet's photo
	// Required: false
	// Format: uri
	// Example: https://example.com/pet-photo.jpg
	Photo *string `json:"photo,omitempty" validate:"omitempty,url"`

	// Species of the pet (e.g., Dog, Cat, Bird)
	// Required: false
	// Minimum length: 2
	// Maximum length: 50
	// Example: Dog
	Species *string `json:"species,omitempty" validate:"omitempty,min=2,max=50"`

	// Breed of the pet
	// Required: false
	// Minimum length: 2
	// Maximum length: 50
	// Example: Golden Retriever
	Breed *string `json:"breed,omitempty" validate:"omitempty,min=2,max=50"`

	// Age of the pet in years
	// Required: false
	// Minimum: 0
	// Example: 5
	Age *int `json:"age,omitempty" validate:"omitempty,min=0"`

	// Gender of the pet
	// Required: false
	// Enum: male, female, unknown
	// Example: male
	Gender *string `json:"gender" validate:"min=2,oneof=male,female,unknown"`

	// Color of the pet
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

	// Indicates if the pet is neutered/spayed
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

func (r *UpdatePetRequest) ToCommand(petIDInt uint, customerIDUInt *uint, isActive *bool) command.UpdatePetCommand {
	var customerID *valueobject.CustomerID
	if customerIDUInt != nil {
		cid := valueobject.NewCustomerID(*customerIDUInt)
		customerID = &cid
	}

	var petGender *enum.PetGender
	if r.Gender != nil {
		petGenderValue := enum.PetGender(*r.Gender)
		petGender = &petGenderValue
	}

	var species *enum.PetSpecies
	if r.Species != nil {
		speciesValue := enum.PetSpecies(*r.Species)
		species = &speciesValue
	}

	cmd := &command.UpdatePetCommand{
		Name:       r.Name,
		PetID:      valueobject.NewPetID(petIDInt),
		Photo:      r.Photo,
		CustomerID: customerID,
		Breed:      r.Breed,
		Age:        r.Age,
		Gender:     petGender,
		Color:      r.Color,
		Microchip:  r.Microchip,
		IsNeutered: r.IsNeutered,
		Tattoo:     r.Tattoo,
		BloodType:  r.BloodType,
		Species:    species,
		IsActive:   isActive,
	}

	return *cmd
}
