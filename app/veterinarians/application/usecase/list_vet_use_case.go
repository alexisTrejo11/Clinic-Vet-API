package vetUsecase

import (
	"context"

	dto "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
)

type ListVetUseCase struct {
}

func (uc *ListVetUseCase) Execute(ctx context.Context, limit, offset int) ([]dto.VetResponse, error) {
	return []dto.VetResponse{}, nil
}
