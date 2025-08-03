package petUsecase

import (
	"context"

	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/dtos"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/mapper"
	petRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/repositories"
)

type ListPetsUseCase struct {
	repository petRepository.PetRepository
}

func NewListPetsUseCase(repository petRepository.PetRepository) *ListPetsUseCase {
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
	repository petRepository.PetRepository
}

func NewListPetByOwnerAndIdUseCase(repository petRepository.PetRepository) *ListPetsByOwnerIdUseCase {
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
