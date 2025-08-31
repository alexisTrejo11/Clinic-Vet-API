package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/mapper"
)

type ListPetsUseCase struct {
	repository repository.PetRepository
}

func NewListPetsUseCase(repository repository.PetRepository) *ListPetsUseCase {
	return &ListPetsUseCase{
		repository: repository,
	}
}

func (uc *ListPetsUseCase) Execute(ctx context.Context) ([]dto.PetResponse, error) {
	petList, err := uc.repository.List(ctx)
	if err != nil {
		return []dto.PetResponse{}, err
	}

	return mapper.ToResponseList(petList), nil
}

type ListPetsByOwnerIDUseCase struct {
	repository repository.PetRepository
}

func NewListPetByOwnerAndIDUseCase(repository repository.PetRepository) *ListPetsByOwnerIDUseCase {
	return &ListPetsByOwnerIDUseCase{
		repository: repository,
	}
}

func (uc *ListPetsByOwnerIDUseCase) Execute(ctx context.Context, ownerID valueobject.OwnerID) ([]dto.PetResponse, error) {
	petList, err := uc.repository.ListByOwnerID(ctx, ownerID)
	if err != nil {
		return []dto.PetResponse{}, err
	}

	return mapper.ToResponseList(petList), nil
}
