package vetUsecase

import (
	"context"

	dto "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
)

type CreateVetUseCase struct {
}

func (uc *CreateVetUseCase) Execute(ctx context.Context, vetCreateData dto.VetCreate) (dto.VetResponse, error) {
	return dto.VetResponse{}, nil
}
