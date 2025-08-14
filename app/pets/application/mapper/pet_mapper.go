package petMapper

import (
	"time"

	dtos "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/dtos"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

func ToDomainFromCreate(dto dtos.PetCreate) *petDomain.Pet {
	builder := petDomain.NewPetBuilder().
		WithName(dto.Name).
		WithSpecies(dto.Species).
		WithOwnerID(dto.OwnerID).
		WithIsActive(dto.IsActive).
		WithCreatedAt(time.Now()).
		WithUpdatedAt(time.Now())

	if dto.Photo != nil {
		builder.WithPhoto(dto.Photo)
	}
	if dto.Breed != nil {
		builder.WithBreed(dto.Breed)
	}
	if dto.Age != nil {
		builder.WithAge(dto.Age)
	}
	if dto.Gender != nil {
		domainGender := petDomain.Gender(*dto.Gender)
		builder.WithGender(&domainGender)
	}
	if dto.Weight != nil {
		builder.WithWeight(dto.Weight)
	}
	if dto.Color != nil {
		builder.WithColor(dto.Color)
	}
	if dto.Microchip != nil {
		builder.WithMicrochip(dto.Microchip)
	}
	if dto.IsNeutered != nil {
		builder.WithIsNeutered(dto.IsNeutered)
	}
	if dto.Allergies != nil {
		builder.WithAllergies(dto.Allergies)
	}
	if dto.CurrentMedications != nil {
		builder.WithCurrentMedications(dto.CurrentMedications)
	}
	if dto.SpecialNeeds != nil {
		builder.WithSpecialNeeds(dto.SpecialNeeds)
	}

	return builder.Build()
}

func ToDomainFromUpdate(pet *petDomain.Pet, dto dtos.PetUpdate) {
	if dto.Name != nil {
		pet.SetName(*dto.Name)
	}
	if dto.Photo != nil {
		pet.SetPhoto(dto.Photo)
	}
	if dto.Species != nil {
		pet.SetSpecies(*dto.Species)
	}
	if dto.Breed != nil {
		pet.SetBreed(dto.Breed)
	}
	if dto.Age != nil {
		pet.SetAge(dto.Age)
	}
	if dto.Gender != nil {
		domainGender := petDomain.Gender(*dto.Gender)
		pet.SetGender(&domainGender)
	}
	if dto.Weight != nil {
		pet.SetWeight(dto.Weight)
	}
	if dto.Color != nil {
		pet.SetColor(dto.Color)
	}
	if dto.Microchip != nil {
		pet.SetMicrochip(dto.Microchip)
	}
	if dto.IsNeutered != nil {
		pet.SetIsNeutered(dto.IsNeutered)
	}
	if dto.OwnerID != nil {
		pet.SetOwnerID(*dto.OwnerID)
	}
	if dto.Allergies != nil {
		pet.SetAllergies(dto.Allergies)
	}
	if dto.CurrentMedications != nil {
		pet.SetCurrentMedications(dto.CurrentMedications)
	}
	if dto.SpecialNeeds != nil {
		pet.SetSpecialNeeds(dto.SpecialNeeds)
	}
	if dto.IsActive != nil {
		pet.SetIsActive(*dto.IsActive)
	}

	pet.SetUpdatedAt(time.Now())
}

func ToResponse(pet *petDomain.Pet) dtos.PetResponse {
	response := dtos.PetResponse{
		ID:                 pet.GetID(),
		Name:               pet.GetName(),
		Photo:              pet.GetPhoto(),
		Species:            pet.GetSpecies(),
		Breed:              pet.GetBreed(),
		Weight:             pet.GetWeight(),
		Color:              pet.GetColor(),
		Microchip:          pet.GetMicrochip(),
		IsNeutered:         pet.GetIsNeutered(),
		OwnerID:            pet.GetOwnerID(),
		Allergies:          pet.GetAllergies(),
		CurrentMedications: pet.GetCurrentMedications(),
		SpecialNeeds:       pet.GetSpecialNeeds(),
		IsActive:           pet.GetIsActive(),
		CreatedAt:          pet.GetCreatedAt(),
		UpdatedAt:          pet.GetUpdatedAt(),
	}

	if pet.GetAge() != nil {
		response.Age = pet.GetAge()
	}
	if pet.GetGender() != nil {
		dtoGender := petDomain.Gender(*pet.GetGender())
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
		dtos[i] = ToResponse(&pet)
	}
	return dtos
}
