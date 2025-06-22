package petUsecase

import (
	"context"

	petAppError "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application"
	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/dtos"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/mapper"
	petRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/repositories"
)

type GetPetByIdUseCase struct {
	repository petRepository.PetRepository
}

func NewGetPetByIdUseCase(repository petRepository.PetRepository) *GetPetByIdUseCase {
	return &GetPetByIdUseCase{
		repository: repository,
	}
}

func (uc *GetPetByIdUseCase) Execute(ctx context.Context, petId uint) (petDTOs.PetResponse, error) {
	pet, err := uc.repository.GetById(ctx, petId)
	if err != nil {
		return petDTOs.PetResponse{}, petAppError.HandleGetByIdError(err, petId)
	}

	return petMapper.ToResponse(pet), nil
}
