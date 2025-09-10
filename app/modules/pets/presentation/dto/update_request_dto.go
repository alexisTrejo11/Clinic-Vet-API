package dto

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
)

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
	CustomerID         *uint    `json:"owner_id,omitempty" validate:"omitempty,gt=0"`
	Allergies          *string  `json:"allergies,omitempty" validate:"omitempty,max=500"`
	CurrentMedications *string  `json:"current_medications,omitempty" validate:"omitempty,max=500"`
	SpecialNeeds       *string  `json:"special_needs,omitempty" validate:"omitempty,max=500"`
	IsActive           *bool    `json:"is_active,omitempty"`
}

func (r UpdatePetRequest) ToUpdatePet(petIDInt uint) *dto.PetUpdateData {
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
