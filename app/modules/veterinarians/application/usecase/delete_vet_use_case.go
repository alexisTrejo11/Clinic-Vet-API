package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
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
	vet, err := uc.vetRepository.GetByID(ctx, vetId)
	if err != nil {
		return err
	}
	if err = uc.vetRepository.SoftDelete(ctx, vet.GetID()); err != nil {
		return err
	}

	return nil
}
