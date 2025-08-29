package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/errors"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type DeleteVetUseCase struct {
	vetRepository repository.VetRepository
}

func NewDeleteVetUseCase(vetRepository repository.VetRepository) *DeleteVetUseCase {
	return &DeleteVetUseCase{
		vetRepository: vetRepository,
	}
}

func (uc *DeleteVetUseCase) Execute(ctx context.Context, vetId valueobject.VetID) error {
	exists, err := uc.vetRepository.Exists(ctx, vetId)
	if err != nil {
		return err
	}

	if !exists {
		return domainerr.NewEntityNotFoundError("veterinarians", vetId.String())
	}

	if err = uc.vetRepository.SoftDelete(ctx, vetId); err != nil {
		return err
	}

	return nil
}
