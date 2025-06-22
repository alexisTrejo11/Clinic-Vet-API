package petUsecase

import (
	"context"

	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/dtos"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/mapper"
	petRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/repositories"
)

type UpdatePetUseCase struct {
	repository petRepository.PetRepository
}

func NewUpdatePetUseCase(repository petRepository.PetRepository) *UpdatePetUseCase {
	return &UpdatePetUseCase{
		repository: repository,
	}
}

func (uc UpdatePetUseCase) Execute(ctx context.Context, petId uint, petUpdate petDTOs.PetUpdate) (petDTOs.PetResponse, error) {
	pet, err := uc.repository.GetById(ctx, petId)
	if err != nil {
		return petDTOs.PetResponse{}, err
	}

	petMapper.ToDomainFromUpdate(&pet, petUpdate)

	if err := uc.repository.Save(ctx, &pet); err != nil {
		return petDTOs.PetResponse{}, err
	}

	return petMapper.ToResponse(pet), nil
}
