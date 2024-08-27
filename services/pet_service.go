package services

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
)

type PetService struct {
	petRepository repository.PetRepository
}

func NewPetService(petRepository repository.PetRepository) *PetService {
	return &PetService{
		petRepository: petRepository,
	}
}

func (ps *PetService) CreatePet(petInsertDTO DTOs.PetInsertDTO, ownerId int32) error {
	params := mappers.MapPetInsertDTOToCreatePetParams(petInsertDTO, ownerId)

	_, err := ps.petRepository.CreatePet(params)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PetService) GetPetById(petID int32) (*DTOs.PetDTO, error) {
	pet, err := ps.petRepository.GetPetById(petID)
	if err != nil {
		return nil, err
	}

	petDTO := mappers.MapPetToPetDTO(*pet)

	return &petDTO, nil
}

func (ps *PetService) GetPetsByOwnerID(ownerID int32) ([]DTOs.PetDTO, error) {
	petList, err := ps.petRepository.GetPetByOwnerID(ownerID)
	if err != nil {
		return nil, err
	}

	var petListDTO []DTOs.PetDTO
	for _, pet := range petList {
		petDTO := mappers.MapPetToPetDTO(pet)
		petListDTO = append(petListDTO, petDTO)
	}

	return petListDTO, nil
}

func (ps *PetService) UpdatePet(petUpdateDTO DTOs.PetUpdateDTO, ownerID int32) error {
	params := mappers.MapPetToPetUpdateDTO(petUpdateDTO, ownerID)

	err := ps.petRepository.UpdatePetById(params)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PetService) DeletePetById(petID int32) error {
	err := ps.petRepository.DeletePetById(petID)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PetService) ValidPetOwner(petID, ownerID int32) bool {
	pet, _ := ps.petRepository.GetPetById(petID)

	if pet.OwnerID != ownerID {
		return false
	}

	return true
}
