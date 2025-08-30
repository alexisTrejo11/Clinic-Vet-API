package usecase

import (
	"context"

	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/errors"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/dto"
	mapper "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/mappers"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
)

type UpdateOwnerUseCase struct {
	ownerRepo repository.OwnerRepository
}

func NewUpdateOwnerUseCase(ownerRepo repository.OwnerRepository) *UpdateOwnerUseCase {
	return &UpdateOwnerUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *UpdateOwnerUseCase) Execute(ctx context.Context, id int, updateData dto.OwnerUpdate) (dto.OwnerDetail, error) {
	owner, err := uc.ownerRepo.GetById(ctx, id)
	if err != nil {
		return dto.OwnerDetail{}, domainerr.HandleGetByIdError(err, id)
	}

	if updateData.PhoneNumber != nil && *updateData.PhoneNumber != owner.PhoneNumber() {
		_, err := uc.ownerRepo.GetByPhone(ctx, *updateData.PhoneNumber)
		if err == nil {
			return dto.OwnerDetail{}, domainerr.HandlePhoneConflictError()
		}
		owner.SetPhoneNumber(*updateData.PhoneNumber)
	}

	if updateData.Photo != nil {
		owner.SetPhoto(*updateData.Photo)
	}

	if updateData.FirstName != nil || updateData.LastName != nil {
		firstName := owner.FullName().FirstName
		if updateData.FirstName != nil {
			firstName = *updateData.FirstName
		}

		lastName := owner.FullName().LastName
		if updateData.LastName != nil {
			lastName = *updateData.LastName
		}

		fullName, err := valueObjects.NewPersonName(firstName, lastName)
		if err != nil {
			return dto.OwnerDetail{}, err
		}
		owner.SetFullName(fullName)
	}

	if updateData.Address != nil {
		owner.SetAddress(*updateData.Address)
	}

	if err := uc.ownerRepo.Save(ctx, &owner); err != nil {
		return dto.OwnerDetail{}, err
	}

	return mapper.ToResponse(&owner), nil
}
