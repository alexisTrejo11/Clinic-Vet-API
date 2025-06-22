package petDTOs

import (
	enum "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

type PetCreate struct {
	Name               string       `json:"name" validate:"required,min=2,max=100"`                          // Nombre es requerido, entre 2 y 100 caracteres.
	Photo              *string      `json:"photo,omitempty" validate:"omitempty,url"`                        // Si se proporciona, debe ser una URL válida.
	Species            string       `json:"species" validate:"required,min=2,max=50"`                        // Especie es requerida, entre 2 y 50 caracteres.
	Breed              *string      `json:"breed,omitempty" validate:"omitempty,min=2,max=50"`               // Si se proporciona, entre 2 y 50 caracteres.
	Age                *int         `json:"age,omitempty" validate:"omitempty"`                              // Edad opcional, si se da, no debe ser negativa.
	Gender             *enum.Gender `json:"gender,omitempty" validate:"omitempty,oneof=Male Female Unknown"` // Género opcional, si se da, debe ser uno de los valores permitidos.
	Weight             *float64     `json:"weight,omitempty" validate:"omitempty,gt=0,lte=1000"`             // Peso opcional, si se da, mayor que 0 y menor o igual a 1000 kg.
	Color              *string      `json:"color,omitempty" validate:"omitempty,min=2,max=50"`               // Color opcional, si se da, entre 2 y 50 caracteres.
	Microchip          *string      `json:"microchip,omitempty" validate:"omitempty,len=15,numeric"`         // Microchip opcional, si se da, debe ser numérico y tener 15 caracteres (estándar ISO).
	IsNeutered         *bool        `json:"is_neutered,omitempty"`                                           // Booleano opcional, no requiere validación específica aquí.
	OwnerID            uint         `json:"owner_id" validate:"required,gt=0"`                               // ID del propietario es requerido y debe ser mayor que 0.
	Allergies          *string      `json:"allergies,omitempty" validate:"omitempty,max=500"`                // Alergias opcionales, máximo 500 caracteres.
	CurrentMedications *string      `json:"current_medications,omitempty" validate:"omitempty,max=500"`      // Medicaciones actuales opcionales, máximo 500 caracteres.
	SpecialNeeds       *string      `json:"special_needs,omitempty" validate:"omitempty,max=500"`            // Necesidades especiales opcionales, máximo 500 caracteres.
	IsActive           bool         `json:"is_active"`                                                       // Booleano, no requiere validación específica aquí.
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
