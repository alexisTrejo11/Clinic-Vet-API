package petController

import petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase/dtos"

func requestToCreatePet(requestData PetInsertRequest) petDTOs.PetCreate {
	return petDTOs.PetCreate{
		Name:               requestData.Name,
		Photo:              requestData.Photo,
		OwnerId:            requestData.OwnerId,
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

func requestToUpdatePet(requestData PetInsertRequest, petId int) petDTOs.PetUpdate {
	petUpdate := petDTOs.PetUpdate{
		PetId:              petId,
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

	if requestData.OwnerId != 0 {
		petUpdate.OwnerID = &requestData.OwnerId
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
