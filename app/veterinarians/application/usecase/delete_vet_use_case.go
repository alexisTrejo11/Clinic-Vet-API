package vetUsecase

import (
	"context"
	"fmt"

	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
)

type DeleteVetUseCase struct {
	vetRepository vetRepo.VeterinarianRepository
}

func NewDeleteVetUseCase(vetRepository vetRepo.VeterinarianRepository) *DeleteVetUseCase {
	return &DeleteVetUseCase{
		vetRepository: vetRepository,
	}
}

func (uc *DeleteVetUseCase) Execute(ctx context.Context, vetId uint, isSoftDelete bool) error {
	exists, err := uc.vetRepository.Exists(ctx, vetId)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("vet %v not found", vetId)
	}
	// TODO: Hard Delete Validation
	if !isSoftDelete {
	}

	if err := uc.vetRepository.Delete(ctx, vetId, isSoftDelete); err != nil {
		return err
	}

	return nil
}
