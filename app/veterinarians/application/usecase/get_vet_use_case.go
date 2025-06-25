package vetUsecase

import (
	"context"

	dto "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
)

type GetVetByIdUseCase struct {
}

func (uc *GetVetByIdUseCase) Execute(ctx context.Context, vetId uint) (dto.VetResponse, error) {
	return dto.VetResponse{Id: vetId}, nil
}
