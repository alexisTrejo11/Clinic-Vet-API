package petDTOs

import (
	enum "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

type PetCreate struct {
	Name               string       `json:"name" validate:"required,min=2,max=100"`
	Photo              *string      `json:"photo,omitempty" validate:"omitempty,url"`
	Species            string       `json:"species" validate:"required,min=2,max=50"`
	Breed              *string      `json:"breed,omitempty" validate:"omitempty,min=2,max=50"`
	Age                *int         `json:"age,omitempty" validate:"omitempty"`
	Gender             *enum.Gender `json:"gender,omitempty" validate:"omitempty,oneof=Male Female Unknown"`
	Weight             *float64     `json:"weight,omitempty" validate:"omitempty,gt=0,lte=1000"`
	Color              *string      `json:"color,omitempty" validate:"omitempty,min=2,max=50"`
	Microchip          *string      `json:"microchip,omitempty" validate:"omitempty,len=15,numeric"`
	IsNeutered         *bool        `json:"is_neutered,omitempty"`
	OwnerID            uint         `json:"owner_id" validate:"required,gt=0"`
	Allergies          *string      `json:"allergies,omitempty" validate:"omitempty,max=500"`
	CurrentMedications *string      `json:"current_medications,omitempty" validate:"omitempty,max=500"`
	SpecialNeeds       *string      `json:"special_needs,omitempty" validate:"omitempty,max=500"`
	IsActive           bool         `json:"is_active"`
}

type PetUpdate struct {
	Name               *string      `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Photo              *string      `json:"photo,omitempty" validate:"omitempty,url"`
	Species            *string      `json:"species,omitempty" validate:"omitempty,min=2,max=50"`
	Breed              *string      `json:"breed,omitempty" validate:"omitempty,min=2,max=50"`
	Age                *int         `json:"age,omitempty" validate:"omitempty,min=0"`
	Gender             *enum.Gender `json:"gender,omitempty" validate:"omitempty,oneof=Male Female Unknown"`
	Weight             *float64     `json:"weight,omitempty" validate:"omitempty,gt=0,lte=1000"`
	Color              *string      `json:"color,omitempty" validate:"omitempty,min=2,max=50"`
	Microchip          *string      `json:"microchip,omitempty" validate:"omitempty,len=15,numeric"`
	IsNeutered         *bool        `json:"is_neutered,omitempty"`
	OwnerID            *uint        `json:"owner_id,omitempty" validate:"omitempty,gt=0"`
	Allergies          *string      `json:"allergies,omitempty" validate:"omitempty,max=500"`
	CurrentMedications *string      `json:"current_medications,omitempty" validate:"omitempty,max=500"`
	SpecialNeeds       *string      `json:"special_needs,omitempty" validate:"omitempty,max=500"`
	IsActive           *bool        `json:"is_active,omitempty"`
}
