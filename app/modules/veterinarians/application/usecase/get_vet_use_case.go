package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/mapper"
)

type GetVetByIDUseCase struct {
	repository repository.VetRepository
}

func NewGetVetByIDUseCase(repository repository.VetRepository) *GetVetByIDUseCase {
	return &GetVetByIDUseCase{
		repository: repository,
	}
}

func (uc *GetVetByIDUseCase) Execute(ctx context.Context, vetID valueobject.VetID) (dto.VetResponse, error) {
	veterinarian, err := uc.repository.GetByID(ctx, vetID)
	if err != nil {
		return dto.VetResponse{}, err
	}

	return mapper.ToResponse(&veterinarian), nil
}
