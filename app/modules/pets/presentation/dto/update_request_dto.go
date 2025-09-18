package dto

import (
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/modules/pets/application/dto"
)

// AdminUpdatePetRequest represents the payload for updating an existing pet's information by an admin
// swagger:model AdminUpdatePetRequest
type AdminUpdatePetRequest struct {
	// ID of the customer who owns the pet
	// Required: false
	// Minimum: 1
	// Example: 123
	CustomerID *uint `json:"customer_id,omitempty" validate:"omitempty,gt=0"`

	UpdatePetRequest
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
	// Enum: Male, Female, Unknown
	// Example: Male
	Gender *string `json:"gender,omitempty" validate:"omitempty,oneof=Male Female Unknown"`

	// Weight of the pet in kilograms
	// Required: false
	// Minimum: 0.1
	// Maximum: 1000.0
	// Example: 25.5
	Weight *float64 `json:"weight,omitempty" validate:"omitempty,gt=0,lte=1000"`

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

	// Known allergies of the pet
	// Required: false
	// Maximum length: 500
	// Example: Pollen, Chicken
	Allergies *string `json:"allergies,omitempty" validate:"omitempty,max=500"`

	// Current medications the pet is taking
	// Required: false
	// Maximum length: 500
	// Example: Heartworm prevention, Flea treatment
	CurrentMedications *string `json:"current_medications,omitempty" validate:"omitempty,max=500"`

	// Any special needs or care instructions
	// Required: false
	// Maximum length: 500
	// Example: Requires daily medication, Blind in left eye
	SpecialNeeds *string `json:"special_needs,omitempty" validate:"omitempty,max=500"`

	// Indicates if the pet is currently active in the system
	// Required: false
	// Example: true
	IsActive *bool `json:"is_active,omitempty"`
}

func (r AdminUpdatePetRequest) ToUpdatePet(petIDInt uint) *dto.PetUpdateData {
	petUpdate := &dto.PetUpdateData{
		PetID:              valueobject.NewPetID(petIDInt),
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

	if r.CustomerID != nil {
		customerID := valueobject.NewCustomerID(*r.CustomerID)
		petUpdate.CustomerID = &customerID
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

func (r CustomerUpdatePetRequest) ToUpdatePet(petIDInt uint, customerIDUInt uint) *dto.PetUpdateData {
	petUpdate := &dto.PetUpdateData{
		PetID:              valueobject.NewPetID(petIDInt),
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
	customerID := valueobject.NewCustomerID(customerIDUInt)
	petUpdate.CustomerID = &customerID

	if r.Name != nil {
		petUpdate.Name = r.Name
	}

	if r.Species != nil {
		petUpdate.Species = r.Species
	}

	petUpdate.IsActive = r.IsActive

	return petUpdate
}
