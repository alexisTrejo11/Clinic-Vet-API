package controller

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
)

func requestToCreatePet(requestData PetInsertRequest) dto.PetCreate {
	ownerID, _ := valueobject.NewOwnerID(requestData.OwnerID)

	return dto.PetCreate{
		Name:               requestData.Name,
		Photo:              requestData.Photo,
		OwnerID:            ownerID,
		Species:            requestData.Species,
		Breed:              requestData.Breed,
		Age:                requestData.Age,
		Gender:             requestData.Gender,
		Weight:             requestData.Weight,
		Color:              requestData.Color,
		Microchip:          requestData.Microchip,
		IsNeutered:         requestData.IsNeutered,
		Allergies:          requestData.Allergies,
		CurrentMedications: requestData.CurrentMedications,
	}
}

func requestToUpdatePet(requestData PetInsertRequest, petIDInt int) dto.PetUpdate {
	petID, _ := valueobject.NewPetID(petIDInt)
	petUpdate := dto.PetUpdate{
		PetID:              petID,
		Photo:              requestData.Photo,
		Breed:              requestData.Breed,
		Age:                requestData.Age,
		Gender:             requestData.Gender,
		Weight:             requestData.Weight,
		Color:              requestData.Color,
		Microchip:          requestData.Microchip,
		IsNeutered:         requestData.IsNeutered,
		Allergies:          requestData.Allergies,
		CurrentMedications: requestData.CurrentMedications,
		SpecialNeeds:       requestData.SpecialNeeds,
	}

	if requestData.OwnerID != 0 {
		ownerID, _ := valueobject.NewOwnerID(requestData.OwnerID)
		petUpdate.OwnerID = &ownerID
	}

	if requestData.Name != "" {
		petUpdate.Name = &requestData.Name
	}

	if requestData.Species != "" {
		petUpdate.Species = &requestData.Species
	}

	petUpdate.IsActive = &requestData.IsActive

	return petUpdate
}
