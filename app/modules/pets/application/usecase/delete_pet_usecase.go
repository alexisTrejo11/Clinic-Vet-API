package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
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

func (uc *DeletePetUseCase) softDelete(cxt context.Context, pet entity.Pet) error {
	pet.SoftDelete()
	if err := uc.repository.Save(cxt, &pet); err != nil {
		return err
	}
	return nil
}
