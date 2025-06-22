package petUsecase

import (
	"context"

	petRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/repositories"
)

type DeletePetUseCase struct {
	repository petRepository.PetRepository
}

func NewDeletePetUseCase(repository petRepository.PetRepository) *DeletePetUseCase {
	return &DeletePetUseCase{
		repository: repository,
	}
}

func (uc DeletePetUseCase) Execute(cxt context.Context, petId uint) error {
	if _, err := uc.repository.GetById(cxt, petId); err != nil {
		return err
	}

	if err := uc.repository.Delete(cxt, petId); err != nil {
		return err
	}

	return nil
}
