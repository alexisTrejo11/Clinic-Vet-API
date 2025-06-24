package ownerUsecase

import (
	"context"

	ownerAppErr "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application"
	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerMappers "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/mappers"
	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
)

type CreateOwnerUseCase struct {
	ownerRepo ownerRepository.OwnerRepository
}

func NewCreateOwnerUseCase(ownerRepo ownerRepository.OwnerRepository) *CreateOwnerUseCase {
	return &CreateOwnerUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *CreateOwnerUseCase) Execute(ctx context.Context, dto ownerDTOs.OwnerCreate) (*ownerDTOs.OwnerResponse, error) {
	_, err := uc.ownerRepo.GetByPhone(ctx, dto.PhoneNumber)
	if err != nil {
		return nil, ownerAppErr.HandlePhoneConflictError(err, dto.PhoneNumber)
	}

	new_owner := ownerMappers.FromRequestCreate(dto)
	if err := uc.ownerRepo.Save(ctx, &new_owner); err != nil {
		return nil, err
	}

	return ownerMappers.ToResponse(new_owner), nil
}
