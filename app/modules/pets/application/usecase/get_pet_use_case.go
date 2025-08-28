package petUsecase

import (
	"context"

	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase/dtos"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase/mapper"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

type GetPetByIdUseCase struct {
	repository petDomain.PetRepository
}

func NewGetPetByIdUseCase(repository petDomain.PetRepository) *GetPetByIdUseCase {
	return &GetPetByIdUseCase{
		repository: repository,
	}
}

func (uc *GetPetByIdUseCase) Execute(ctx context.Context, petId int) (petDTOs.PetResponse, error) {
	pet, err := uc.repository.GetById(ctx, petId)

	if err != nil {
		return petDTOs.PetResponse{}, err
	}

	return petMapper.ToResponse(&pet), nil
}
