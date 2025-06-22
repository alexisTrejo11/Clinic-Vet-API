package petUsecase

import (
	"context"

	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/dtos"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/mapper"
	petRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/repositories"
)

type CreatePetUseCase struct {
	repository petRepository.PetRepository
}

func NewCreatePetUseCase(repository petRepository.PetRepository) *CreatePetUseCase {
	return &CreatePetUseCase{
		repository: repository,
	}
}

func (uc CreatePetUseCase) Execute(ctx context.Context, petCreate petDTOs.PetCreate) (petDTOs.PetResponse, error) {
	newPet := petMapper.ToDomainFromCreate(petCreate)

	if err := uc.repository.Save(ctx, &newPet); err != nil {
		return petDTOs.PetResponse{}, err
	}

	return petMapper.ToResponse(newPet), nil
}
