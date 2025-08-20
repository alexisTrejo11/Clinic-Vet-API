package petDTOs

import (
	"time"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

type PetResponse struct {
	ID                 int               `json:"id"`
	Name               string            `json:"name"`
	Photo              *string           `json:"photo,omitempty"`
	Species            string            `json:"species"`
	Breed              *string           `json:"breed,omitempty"`
	Age                *int              `json:"age,omitempty"`
	Gender             *petDomain.Gender `json:"gender,omitempty"`
	Weight             *float64          `json:"weight,omitempty"` // kg
	Color              *string           `json:"color,omitempty"`
	Microchip          *string           `json:"microchip,omitempty"`
	IsNeutered         *bool             `json:"is_neutered,omitempty"`
	OwnerID            int               `json:"owner_id"`
	Allergies          *string           `json:"allergies,omitempty"`
	CurrentMedications *string           `json:"current_medications,omitempty"`
	SpecialNeeds       *string           `json:"special_needs,omitempty"`
	IsActive           bool              `json:"is_active"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
}
