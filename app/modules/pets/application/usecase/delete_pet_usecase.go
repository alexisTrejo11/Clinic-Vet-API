package usecase

import (
	"context"

	"clinic-vet-api/app/core/domain/entity/pet"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
)

type DeletePetUseCase struct {
	repository repository.PetRepository
}

func NewDeletePetUseCase(repository repository.PetRepository) *DeletePetUseCase {
	return &DeletePetUseCase{
		repository: repository,
	}
}

func (uc *DeletePetUseCase) Execute(cxt context.Context, petID valueobject.PetID, isSoftDelete bool) error {
	pet, err := uc.repository.GetByID(cxt, petID)
	if err != nil {
		return err
	}

	if isSoftDelete {
		if err := uc.softDelete(cxt, pet); err != nil {
			return err
		}
	}

	if err := uc.delete(cxt, petID); err != nil {
		return err
	}

	return nil
}

func (uc *DeletePetUseCase) delete(cxt context.Context, petID valueobject.PetID) error {
	if err := uc.repository.Delete(cxt, petID); err != nil {
		return err
	}
	return nil
}

func (uc *DeletePetUseCase) softDelete(cxt context.Context, pet pet.Pet) error {
	pet.SoftDelete()
	if err := uc.repository.Save(cxt, &pet); err != nil {
		return err
	}
	return nil
}
