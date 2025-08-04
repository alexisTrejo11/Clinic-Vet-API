package petMapper

import (
	"time"

	dtos "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/dtos"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

func ToDomainFromCreate(dto dtos.PetCreate) petDomain.Pet {
	pet := petDomain.Pet{
		Name:      dto.Name,
		Species:   dto.Species,
		OwnerId:   dto.OwnerID,
		IsActive:  dto.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if dto.Photo != nil {
		pet.Photo = dto.Photo
	}
	if dto.Breed != nil {
		pet.Breed = dto.Breed
	}
	if dto.Age != nil {
		domainAge := *dto.Age
		pet.Age = &domainAge
	}
	if dto.Gender != nil {
		domainGender := petDomain.Gender(*dto.Gender)
		pet.Gender = &domainGender
	}
	if dto.Weight != nil {
		pet.Weight = dto.Weight
	}
	if dto.Color != nil {
		pet.Color = dto.Color
	}
	if dto.Microchip != nil {
		pet.Microchip = dto.Microchip
	}
	if dto.IsNeutered != nil {
		pet.IsNeutered = dto.IsNeutered
	}
	if dto.Allergies != nil {
		pet.Allergies = dto.Allergies
	}
	if dto.CurrentMedications != nil {
		pet.CurrentMedications = dto.CurrentMedications
	}
	if dto.SpecialNeeds != nil {
		pet.SpecialNeeds = dto.SpecialNeeds
	}

	return pet
}

func ToDomainFromUpdate(pet *petDomain.Pet, dto dtos.PetUpdate) {
	if dto.Name != nil {
		pet.Name = *dto.Name
	}
	if dto.Photo != nil {
		pet.Photo = dto.Photo
	}
	if dto.Species != nil {
		pet.Species = *dto.Species
	}
	if dto.Breed != nil {
		pet.Breed = dto.Breed
	}
	if dto.Age != nil {
		pet.Age = dto.Age
	}
	if dto.Gender != nil {
		domainGender := petDomain.Gender(*dto.Gender)
		pet.Gender = &domainGender
	}
	if dto.Weight != nil {
		pet.Weight = dto.Weight
	}
	if dto.Color != nil {
		pet.Color = dto.Color
	}
	if dto.Microchip != nil {
		pet.Microchip = dto.Microchip
	}
	if dto.IsNeutered != nil {
		pet.IsNeutered = dto.IsNeutered
	}
	if dto.OwnerID != nil {
		pet.OwnerId = *dto.OwnerID
	}
	if dto.Allergies != nil {
		pet.Allergies = dto.Allergies
	}
	if dto.CurrentMedications != nil {
		pet.CurrentMedications = dto.CurrentMedications
	}
	if dto.SpecialNeeds != nil {
		pet.SpecialNeeds = dto.SpecialNeeds
	}
	if dto.IsActive != nil {
		pet.IsActive = *dto.IsActive
	}
	pet.UpdatedAt = time.Now()
}

func ToResponse(pet petDomain.Pet) dtos.PetResponse {
	response := dtos.PetResponse{
		ID:                 pet.Id,
		Name:               pet.Name,
		Photo:              pet.Photo,
		Species:            pet.Species,
		Breed:              pet.Breed,
		Weight:             pet.Weight,
		Color:              pet.Color,
		Microchip:          pet.Microchip,
		IsNeutered:         pet.IsNeutered,
		OwnerID:            pet.OwnerId,
		Allergies:          pet.Allergies,
		CurrentMedications: pet.CurrentMedications,
		SpecialNeeds:       pet.SpecialNeeds,
		IsActive:           pet.IsActive,
		CreatedAt:          pet.CreatedAt,
		UpdatedAt:          pet.UpdatedAt,
	}

	if pet.Age != nil {
		response.Age = pet.Age
	}
	if pet.Gender != nil {
		dtoGender := petDomain.Gender(*pet.Gender)
		response.Gender = &dtoGender
	}

	return response
}

func ToResponseList(pets []petDomain.Pet) []dtos.PetResponse {
	if pets == nil {
		return []dtos.PetResponse{}
	}
	dtos := make([]dtos.PetResponse, len(pets))
	for i, pet := range pets {
		dtos[i] = ToResponse(pet)
	}
	return dtos
}
