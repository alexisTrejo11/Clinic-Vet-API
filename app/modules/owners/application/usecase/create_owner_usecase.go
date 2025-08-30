package usecase

import (
	"context"

	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/errors"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/dto"
	mapper "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/mappers"
)

type CreateOwnerUseCase struct {
	ownerRepo repository.OwnerRepository
}

func NewCreateOwnerUseCase(ownerRepo repository.OwnerRepository) *CreateOwnerUseCase {
	return &CreateOwnerUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *CreateOwnerUseCase) Execute(ctx context.Context, createData dto.OwnerCreate) (dto.OwnerDetail, error) {
	if exists, err := uc.ownerRepo.ExistsByPhone(ctx, createData.PhoneNumber); err != nil {
		return dto.OwnerDetail{}, err
	} else if exists {
		return dto.OwnerDetail{}, domainerr.HandlePhoneConflictError()
	}

	owner := mapper.FromRequestCreate(createData)
	if err := uc.ownerRepo.Save(ctx, owner); err != nil {
		return dto.OwnerDetail{}, err
	}

	return mapper.ToResponse(owner), nil
}
