package services

import (
	"context"

	dtos "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
)

type PetService struct {
	petRepository repository.PetRepositoryInterface
}

func NewPetService(petRepository repository.PetRepositoryInterface) *PetService {
	return &PetService{
		petRepository: petRepository,
	}
}

func (ps *PetService) CreatePet(petInsertDTO dtos.PetInsertDTO) error {
	params := mappers.MapPetInsertDTOToCreatePetParams(petInsertDTO)

	_, err := ps.petRepository.CreatePet(context.Background(), params)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PetService) GetPetById(petID int32) (*dtos.PetDTO, error) {
	pet, err := ps.petRepository.GetPetById(context.Background(), petID)
	if err != nil {
		return nil, err
	}

	petDTO := mappers.MapPetToPetDTO(pet)

	return &petDTO, nil
}

func (ps *PetService) UpdatePet(petUpdateDTO dtos.PetUpdateDTO) error {
	params := mappers.MapPetToPetUpdateDTO(petUpdateDTO)

	err := ps.petRepository.UpdatePetById(context.Background(), params)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PetService) DeletePetbyId(petID int32) error {
	err := ps.petRepository.DeletePetById(context.Background(), petID)
	if err != nil {
		return err
	}

	return nil
}
