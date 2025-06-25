package vetUsecase

import (
	"context"

	vetDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
)

type UpdateVetUseCase struct {
}

func (uc *UpdateVetUseCase) Execute(ctx context.Context, vetId uint, vetCreateData vetDtos.VetUpdate) (vetDtos.VetResponse, error) {
	return vetDtos.VetResponse{}, nil
}
