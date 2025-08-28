package petUsecase

import (
	"context"

	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase/dtos"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase/mapper"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

type ListPetsUseCase struct {
	repository petDomain.PetRepository
}

func NewListPetsUseCase(repository petDomain.PetRepository) *ListPetsUseCase {
	return &ListPetsUseCase{
		repository: repository,
	}
}

func (uc *ListPetsUseCase) Execute(ctx context.Context) ([]petDTOs.PetResponse, error) {
	petList, err := uc.repository.List(ctx)
	if err != nil {
		return []petDTOs.PetResponse{}, err
	}

	return petMapper.ToResponseList(petList), nil
}

type ListPetsByOwnerIdUseCase struct {
	repository petDomain.PetRepository
}

func NewListPetByOwnerAndIdUseCase(repository petDomain.PetRepository) *ListPetsByOwnerIdUseCase {
	return &ListPetsByOwnerIdUseCase{
		repository: repository,
	}
}

func (uc *ListPetsByOwnerIdUseCase) Execute(ctx context.Context, ownerId int) ([]petDTOs.PetResponse, error) {
	petList, err := uc.repository.ListByOwnerId(ctx, ownerId)
	if err != nil {
		return []petDTOs.PetResponse{}, err
	}

	return petMapper.ToResponseList(petList), nil
}
