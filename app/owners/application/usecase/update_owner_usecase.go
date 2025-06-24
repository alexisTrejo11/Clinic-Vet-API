package ownerUsecase

import (
	"context"

	ownerAppErr "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application"
	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerMappers "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/mappers"
	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
)

type UpdateOwnerUseCase struct {
	ownerRepo ownerRepository.OwnerRepository
}

func NewUpdateOwnerUseCase(ownerRepo ownerRepository.OwnerRepository) *UpdateOwnerUseCase {
	return &UpdateOwnerUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *UpdateOwnerUseCase) Execute(ctx context.Context, id uint, dto ownerDTOs.OwnerUpdate) (*ownerDTOs.OwnerResponse, error) {
	owner, err := uc.ownerRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ownerAppErr.HandleGetByIdError(err, id)
	}

	if dto.PhoneNumber != nil && *dto.PhoneNumber != owner.PhoneNumber {
		_, err := uc.ownerRepo.GetByPhone(ctx, *dto.PhoneNumber)
		if err != nil {
			ownerAppErr.HandlePhoneConflictError(err, *dto.PhoneNumber)
		}

	}

	// Update fields
	if dto.Photo != nil {
		owner.Photo = *dto.Photo
	}

	if dto.FirstName != nil {
		owner.FullName.FirstName = *dto.FirstName
	}

	if dto.LastName != nil {
		owner.FullName.FirstName = *dto.LastName
	}

	if dto.PhoneNumber != nil {
		owner.PhoneNumber = *dto.PhoneNumber
	}

	return ownerMappers.ToResponse(owner), nil
}
