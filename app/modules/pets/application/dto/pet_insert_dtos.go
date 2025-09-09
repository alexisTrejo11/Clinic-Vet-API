// Package dto contains Data Transfer Objects for the pets module.
package dto

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type CreatePetData struct {
	Name               string
	Photo              *string
	Species            string
	Breed              *string
	Age                *int
	Gender             *string
	Weight             *float64
	Color              *string
	Microchip          *string
	IsNeutered         *bool
	OwnerID            valueobject.OwnerID
	Allergies          *string
	CurrentMedications *string
	SpecialNeeds       *string
	IsActive           bool
}

type PetUpdate struct {
	PetID              valueobject.PetID    `json:"pet_id"`
	Name               *string              `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Photo              *string              `json:"photo,omitempty" validate:"omitempty,url"`
	Species            *string              `json:"species,omitempty" validate:"omitempty,min=2,max=50"`
	Breed              *string              `json:"breed,omitempty" validate:"omitempty,min=2,max=50"`
	Age                *int                 `json:"age,omitempty" validate:"omitempty,min=0"`
	Gender             *string              `json:"gender,omitempty" validate:"omitempty,oneof=Male Female Unknown"`
	Weight             *float64             `json:"weight,omitempty" validate:"omitempty,gt=0,lte=1000"`
	Color              *string              `json:"color,omitempty" validate:"omitempty,min=2,max=50"`
	Microchip          *string              `json:"microchip,omitempty" validate:"omitempty,len=15,numeric"`
	IsNeutered         *bool                `json:"is_neutered,omitempty"`
	OwnerID            *valueobject.OwnerID `json:"owner_id,omitempty" validate:"omitempty,gt=0"`
	Allergies          *string              `json:"allergies,omitempty" validate:"omitempty,max=500"`
	CurrentMedications *string              `json:"current_medications,omitempty" validate:"omitempty,max=500"`
	SpecialNeeds       *string              `json:"special_needs,omitempty" validate:"omitempty,max=500"`
	IsActive           *bool                `json:"is_active,omitempty"`
}
