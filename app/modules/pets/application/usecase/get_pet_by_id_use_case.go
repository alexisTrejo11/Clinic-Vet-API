package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/mapper"
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
