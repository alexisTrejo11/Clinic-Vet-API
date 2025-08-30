package usecase

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/dto"
	mapper "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/mappers"
)

type GetOwnerByIDUseCase struct {
	ownerRepo repository.OwnerRepository
}

func NewGetOwnerByIDUseCase(ownerRepo repository.OwnerRepository) *GetOwnerByIDUseCase {
	return &GetOwnerByIDUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *GetOwnerByIDUseCase) Execute(ctx context.Context, id int) (dto.OwnerDetail, error) {
	owner, err := uc.ownerRepo.GetById(ctx, id)
	if err != nil {
		return dto.OwnerDetail{}, err
	}

	return mapper.ToResponse(&owner), nil
}
