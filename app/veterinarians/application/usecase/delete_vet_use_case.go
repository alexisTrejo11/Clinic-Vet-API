package vetUsecase

import (
	"context"
	"strconv"

	vetaApplication "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application"
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

func (uc *DeleteVetUseCase) Execute(ctx context.Context, vetId int) error {
	exists, err := uc.vetRepository.Exists(ctx, vetId)
	if err != nil {
		return vetaApplication.VetDBErr("search", err)
	}

	if !exists {
		return vetaApplication.VetNotFoundErr("id", strconv.Itoa(vetId))
	}

	if err := uc.vetRepository.Delete(ctx, vetId); err != nil {
		return vetaApplication.VetDBErr("delete", err)
	}

	return nil
}
