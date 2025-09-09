// Package mapper contains functions to map between different representations of Pet data.
package mapper

import (
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
)

func ToDomainFromCreate(dto dto.CreatePetData) (*pet.Pet, error) {
	// Convertir y validar los IDs de value objects

	ownerID, err := valueobject.NewOwnerID(dto.OwnerID.Value())
	if err != nil {
		return nil, fmt.Errorf("invalid owner ID: %w", err)
	}

	// Crear options básicas
	opts := []pet.PetOption{
		pet.WithName(dto.Name),
		pet.WithSpecies(dto.Species),
		pet.WithIsActive(dto.IsActive),
	}

	// Agregar options para campos opcionales
	if dto.Photo != nil {
		opts = append(opts, pet.WithPhoto(dto.Photo))
	}

	if dto.Breed != nil {
		opts = append(opts, pet.WithBreed(dto.Breed))
	}

	if dto.Age != nil {
		opts = append(opts, pet.WithAge(dto.Age))
	}

	if dto.Gender != nil {
		gender, err := enum.ParsePetGender(*dto.Gender)
		if err != nil {
			return nil, fmt.Errorf("invalid gender: %w", err)
		}
		opts = append(opts, pet.WithGender(&gender))
	}

	if dto.Weight != nil {
		opts = append(opts, pet.WithWeight(dto.Weight))
	}

	if dto.Color != nil {
		opts = append(opts, pet.WithColor(dto.Color))
	}

	if dto.Microchip != nil {
		opts = append(opts, pet.WithMicrochip(dto.Microchip))
	}

	if dto.IsNeutered != nil {
		opts = append(opts, pet.WithIsNeutered(dto.IsNeutered))
	}

	if dto.Allergies != nil {
		opts = append(opts, pet.WithAllergies(dto.Allergies))
	}

	if dto.CurrentMedications != nil {
		opts = append(opts, pet.WithCurrentMedications(dto.CurrentMedications))
	}

	if dto.SpecialNeeds != nil {
		opts = append(opts, pet.WithSpecialNeeds(dto.SpecialNeeds))
	}

	// Crear la entidad Pet
	petEntity, err := pet.NewPet(
		valueobject.PetID{}, // ID will be generated inside NewPet
		ownerID,
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create pet: %w", err)
	}

	return petEntity, nil
}

// TODO: implement ToDomainFromUpdate
func ToDomainFromUpdate(pet *pet.Pet, dto dto.PetUpdate) {
}

func ToResponse(pet *pet.Pet) dto.PetResponse {
	response := dto.PetResponse{
		ID:                 pet.ID().Value(),
		Name:               pet.Name(),
		Photo:              pet.Photo(),
		Species:            pet.Species(),
		Breed:              pet.Breed(),
		Weight:             pet.Weight(),
		Color:              pet.Color(),
		Microchip:          pet.Microchip(),
		IsNeutered:         pet.IsNeutered(),
		OwnerID:            pet.OwnerID().Value(),
		Allergies:          pet.Allergies(),
		CurrentMedications: pet.CurrentMedications(),
		SpecialNeeds:       pet.SpecialNeeds(),
		IsActive:           pet.IsActive(),
		CreatedAt:          pet.CreatedAt().Format(time.RFC822),
		UpdatedAt:          pet.UpdatedAt().Format(time.RFC822),
	}

	if pet.Age() != nil {
		response.Age = pet.Age()
	}
	if pet.Gender() != nil {
		response.Gender = pet.Gender().DisplayName()
	}

	return response
}

func ToResponseList(pets []pet.Pet) []dto.PetResponse {
	if pets == nil {
		return []dto.PetResponse{}
	}
	dtos := make([]dto.PetResponse, len(pets))
	for i, pet := range pets {
		dtos[i] = ToResponse(&pet)
	}
	return dtos
}
