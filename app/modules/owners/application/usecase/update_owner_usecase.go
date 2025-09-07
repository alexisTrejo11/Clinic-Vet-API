package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/dto"
	mapper "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/mappers"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type UpdateOwnerUseCase struct {
	ownerRepo repository.OwnerRepository
}

func NewUpdateOwnerUseCase(ownerRepo repository.OwnerRepository) *UpdateOwnerUseCase {
	return &UpdateOwnerUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *UpdateOwnerUseCase) Execute(ctx context.Context, id valueobject.OwnerID, updateData dto.OwnerUpdate) (dto.OwnerDetail, error) {
	owner, err := uc.ownerRepo.GetByID(ctx, id)
	if err != nil {
		return dto.OwnerDetail{}, nil
	}

	if updateData.PhoneNumber != nil && *updateData.PhoneNumber != owner.PhoneNumber() {
		exists, err := uc.ownerRepo.ExistsByPhone(ctx, *updateData.PhoneNumber)
		if err != nil {
			return dto.OwnerDetail{}, err
		}
		if exists {
			return dto.OwnerDetail{}, apperror.ConflictError("phone Number", "phone Number already taken")
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

		fullName, _ := valueobject.NewPersonName(firstName, lastName)
		if err != nil {
			return dto.OwnerDetail{}, apperror.FieldValidationError("full-name", firstName+" "+lastName, err.Error())
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
