package usecase

import (
	"context"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/modules/pets/application/dto"
	"clinic-vet-api/app/modules/pets/application/mapper"
)

type GetPetByIDUseCase struct {
	repository repository.PetRepository
}

func NewGetPetByIDUseCase(repository repository.PetRepository) *GetPetByIDUseCase {
	return &GetPetByIDUseCase{
		repository: repository,
	}
}

func (uc *GetPetByIDUseCase) Execute(ctx context.Context, petID valueobject.PetID) (dto.PetResponse, error) {
	pet, err := uc.repository.GetByID(ctx, petID)
	if err != nil {
		return dto.PetResponse{}, err
	}

	return mapper.ToResponse(&pet), nil
}
