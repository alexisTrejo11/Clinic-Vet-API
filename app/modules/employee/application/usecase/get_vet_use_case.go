package usecase

import (
	"context"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/modules/veterinarians/application/dto"
	"clinic-vet-api/app/modules/veterinarians/application/mapper"
)

type GetVetByIDUseCase struct {
	repository repository.CustomerRepository
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

