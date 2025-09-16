package usecase

import (
	"context"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/modules/pets/application/dto"
	"clinic-vet-api/app/modules/pets/application/mapper"
)

type FindPetByIDUseCase struct {
	repository repository.PetRepository
}

func NewFindPetByIDUseCase(repository repository.PetRepository) *FindPetByIDUseCase {
	return &FindPetByIDUseCase{
		repository: repository,
	}
}

func (uc *FindPetByIDUseCase) Execute(ctx context.Context, petID valueobject.PetID) (dto.PetResponse, error) {
	pet, err := uc.repository.FindByID(ctx, petID)
	if err != nil {
		return dto.PetResponse{}, err
	}

	return mapper.ToResponse(&pet), nil
}
