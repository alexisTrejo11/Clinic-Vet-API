package mapper

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
)

func ToDomainFromCreate(dto dto.PetCreate) *entity.Pet {
	builder := entity.NewPetBuilder().
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
		domainGender := enum.PetGender(*dto.Gender)
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

func ToDomainFromUpdate(pet *entity.Pet, dto dto.PetUpdate) {
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
		domainGender := enum.PetGender(*dto.Gender)
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

func ToResponse(pet *entity.Pet) dto.PetResponse {
	response := dto.PetResponse{
		ID:                 pet.GetID().GetValue(),
		Name:               pet.GetName(),
		Photo:              pet.GetPhoto(),
		Species:            pet.GetSpecies(),
		Breed:              pet.GetBreed(),
		Weight:             pet.GetWeight(),
		Color:              pet.GetColor(),
		Microchip:          pet.GetMicrochip(),
		IsNeutered:         pet.GetIsNeutered(),
		OwnerID:            pet.GetOwnerID().GetValue(),
		Allergies:          pet.GetAllergies(),
		CurrentMedications: pet.GetCurrentMedications(),
		SpecialNeeds:       pet.GetSpecialNeeds(),
		IsActive:           pet.GetIsActive(),
		CreatedAt:          pet.GetCreatedAt().Format("2005-10-01 20:00:00"),
		UpdatedAt:          pet.GetUpdatedAt().Format("2005-10-01 20:00:00"),
	}

	if pet.GetAge() != nil {
		response.Age = pet.GetAge()
	}
	if pet.GetGender() != nil {
		response.Gender = (*string)(pet.GetGender())
	}

	return response
}

func ToResponseList(pets []entity.Pet) []dto.PetResponse {
	if pets == nil {
		return []dto.PetResponse{}
	}
	dtos := make([]dto.PetResponse, len(pets))
	for i, pet := range pets {
		dtos[i] = ToResponse(&pet)
	}
	return dtos
}
