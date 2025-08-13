package ownerUsecase

import (
	"context"

	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerMappers "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/mappers"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
)

type CreateOwnerUseCase struct {
	ownerRepo ownerDomain.OwnerRepository
}

func NewCreateOwnerUseCase(ownerRepo ownerDomain.OwnerRepository) *CreateOwnerUseCase {
	return &CreateOwnerUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *CreateOwnerUseCase) Execute(ctx context.Context, dto ownerDTOs.OwnerCreate) (*ownerDTOs.OwnerResponse, error) {
	_, err := uc.ownerRepo.GetByPhone(ctx, dto.PhoneNumber)
	if err == nil {
		return nil, ownerDomain.HandlePhoneConflictError()
	}

	new_owner := ownerMappers.FromRequestCreate(dto)
	if err := uc.ownerRepo.Save(ctx, *new_owner); err != nil {
		return nil, err
	}

	return ownerMappers.ToResponse(new_owner), nil
}
