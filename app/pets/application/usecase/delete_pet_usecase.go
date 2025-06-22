package petUsecase

import (
	"context"

	petAppError "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application"
	petRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/repositories"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

type DeletePetUseCase struct {
	repository petRepository.PetRepository
}

func NewDeletePetUseCase(repository petRepository.PetRepository) *DeletePetUseCase {
	return &DeletePetUseCase{
		repository: repository,
	}
}

func (uc *DeletePetUseCase) Execute(cxt context.Context, petId uint, isSoftDelete bool) error {
	pet, err := uc.repository.GetById(cxt, petId)
	if err != nil {
		return petAppError.HandleGetByIdError(err, petId)
	}

	if isSoftDelete {
		if err := uc.softDelete(cxt, pet); err != nil {
			return err
		}
	} else {
		if err := uc.delete(cxt, petId); err != nil {
			return err
		}
	}

	return nil
}

func (uc *DeletePetUseCase) delete(cxt context.Context, petId uint) error {
	if err := uc.repository.Delete(cxt, petId); err != nil {
		return err
	}
	return nil
}

func (uc *DeletePetUseCase) softDelete(cxt context.Context, pet petDomain.Pet) error {
	pet.SoftDelete()
	if err := uc.repository.Save(cxt, &pet); err != nil {
		return err
	}
	return nil
}
