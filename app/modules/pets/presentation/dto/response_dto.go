package dto

import "clinic-vet-api/app/modules/pets/application/cqrs/query"

// @Description Represents the response structure for a pet.
type PetResponse struct {
	// The unique ID of the pet.
	ID uint `json:"id"`
	// The name of the pet.
	Name string `json:"name"`
	// The URL of the pet's photo.
	Photo *string `json:"photo,omitempty"`
	// The species of the pet.
	Species string `json:"species"`
	// The breed of the pet.
	Breed *string `json:"breed,omitempty"`
	// The age of the pet in years.
	Age *int `json:"age,omitempty"`
	// The gender of the pet.
	Gender string `json:"gender,omitempty"`
	// The weight of the pet in kilograms.
	Weight *float64 `json:"weight,omitempty"`
	// The color of the pet's fur/coat.
	Color *string `json:"color,omitempty"`
	// The pet's microchip number.
	Microchip *string `json:"microchip,omitempty"`
	// Indicates if the pet is neutered.
	IsNeutered *bool `json:"is_neutered,omitempty"`
	// A list of the pet's known allergies.
	Allergies *string `json:"allergies,omitempty"`
	// The pet's current medications.
	CurrentMedications *string `json:"current_medications,omitempty"`
	// Any special needs the pet may have.
	SpecialNeeds *string `json:"special_needs,omitempty"`
	// Indicates if the pet's record is active.
	IsActive bool `json:"is_active"`
	// The date and time when the pet's record was created.
	CreatedAt string `json:"created_at"`
	// The date and time when the pet's record was last updated.
	UpdatedAt string `json:"updated_at"`
}

func ToResponse(result query.PetResult) *PetResponse {
	return &PetResponse{
		ID:                 result.ID,
		Name:               result.Name,
		Photo:              result.Photo,
		Species:            result.Species,
		Breed:              result.Breed,
		Age:                result.Age,
		Allergies:          result.Allergies,
		CurrentMedications: result.CurrentMedications,
		SpecialNeeds:       result.SpecialNeeds,
		Color:              result.Color,
		Microchip:          result.Microchip,
		IsNeutered:         result.IsNeutered,
		IsActive:           result.IsActive,
		CreatedAt:          result.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:          result.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToResponseList(results []query.PetResult) []*PetResponse {
	var pets []*PetResponse
	for _, result := range results {
		pet := ToResponse(result)
		pets = append(pets, pet)
	}
	return pets
}
