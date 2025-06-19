package services

import (
	"context"

	dtos "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
)

type VeterinarianService interface {
	CreateVeterinarian(vetInsertDTO dtos.VetInsertDTO) error
	GetVeterinarianById(veterinarianID int32) (*dtos.VetDTO, error)
	UpdateVeterinarian(vetUpdateDTO dtos.VetUpdateDTO) error
	DeleteVeterinarianbyId(veterinarianID int32) error
	ValidateExistingVet(veterinarianID int32) bool
}

type VeterinarianServiceImpl struct {
	VeterinarianRepository repository.VeterinarianRepository
}

func NewVeterinarianService(VeterinarianRepository repository.VeterinarianRepository) *VeterinarianServiceImpl {
	return &VeterinarianServiceImpl{
		VeterinarianRepository: VeterinarianRepository,
	}
}

func (vs *VeterinarianServiceImpl) CreateVeterinarian(vetInsertDTO dtos.VetInsertDTO) error {
	params := mappers.MapVetInsertDtoToVetInsertParams(vetInsertDTO)

	_, err := vs.VeterinarianRepository.CreateVeterinarian(context.Background(), params)
	if err != nil {
		return err
	}

	return nil
}

func (vs *VeterinarianServiceImpl) GetVeterinarianById(veterinarianID int32) (*dtos.VetDTO, error) {
	veterinarian, err := vs.VeterinarianRepository.GetVeterinarianByID(context.Background(), veterinarianID)
	if err != nil {
		return nil, err
	}

	vetDTO := mappers.MapSqlcEntityToDTO(veterinarian)

	return &vetDTO, nil
}

func (vs *VeterinarianServiceImpl) UpdateVeterinarian(vetUpdateDTO dtos.VetUpdateDTO) error {
	veterinarian, _ := vs.VeterinarianRepository.GetVeterinarianByID(context.Background(), vetUpdateDTO.Id)

	params := mappers.MapVetUpdateDtoToEntity(&vetUpdateDTO, veterinarian)

	err := vs.VeterinarianRepository.UpdateVeterinarian(context.Background(), params)
	if err != nil {
		return err
	}

	return nil
}

func (vs *VeterinarianServiceImpl) DeleteVeterinarianbyId(veterinarianID int32) error {
	err := vs.VeterinarianRepository.DeleteVeterinarian(context.Background(), veterinarianID)
	if err != nil {
		return err
	}
	return nil
}

func (vs VeterinarianServiceImpl) ValidateExistingVet(veterinarianID int32) bool {
	isExistingVet := vs.VeterinarianRepository.ValidateExistingVeterinarian(context.Background(), veterinarianID)
	return isExistingVet
}
