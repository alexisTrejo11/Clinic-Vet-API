package query

import (
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"time"
)

type PetResult struct {
	ID                 uint
	Name               string
	Photo              *string
	Species            string
	Breed              *string
	Age                *int
	Weight             *float64
	Color              *string
	Microchip          *string
	IsNeutered         *bool
	PetSpecies         string
	CustomerID         uint
	Allergies          *string
	CurrentMedications *string
	SpecialNeeds       *string
	IsActive           bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func entityToResult(p pet.Pet) PetResult {
	return PetResult{
		ID:         p.ID().Value(),
		Name:       p.Name(),
		Photo:      p.Photo(),
		Species:    p.Species().DisplayName(),
		Breed:      p.Breed(),
		Age:        p.Age(),
		Color:      p.Color(),
		Microchip:  p.Microchip(),
		IsNeutered: p.IsNeutered(),
		PetSpecies: p.Species().String(),
		CustomerID: p.CustomerID().Value(),
		IsActive:   p.IsActive(),
		CreatedAt:  p.CreatedAt(),
		UpdatedAt:  p.UpdatedAt(),
	}
}

func entitiesToResults(pets []pet.Pet) []PetResult {
	results := make([]PetResult, len(pets))
	for i, p := range pets {
		results[i] = entityToResult(p)
	}
	return results
}
