package ownerUsecase

import (
	"context"

	ownerAppErr "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application"
	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerMappers "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/mappers"
	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
)

type GetOwnerByIdUseCase struct {
	ownerRepo ownerRepository.OwnerRepository
}

func NewGetOwnerByIdUseCase(ownerRepo ownerRepository.OwnerRepository) *GetOwnerByIdUseCase {
	return &GetOwnerByIdUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *GetOwnerByIdUseCase) Execute(ctx context.Context, id uint, includePets bool) (*ownerDTOs.OwnerResponse, error) {
	owner, err := uc.ownerRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ownerAppErr.HandleGetByIdError(err, id)
	}

	OwnerResponse := ownerMappers.ToResponse(owner)
	if includePets {
	}

	return OwnerResponse, nil
}
