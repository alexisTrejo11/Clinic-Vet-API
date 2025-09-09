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
		return dto.OwnerDetail{}, err
	}

	if updateData.Photo != nil {
		if err := owner.UpdatePhoto(*updateData.Photo); err != nil {
			return dto.OwnerDetail{}, apperror.FieldValidationError("photo", "", err.Error())
		}
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

		fullName, err := valueobject.NewPersonName(firstName, lastName)
		if err != nil {
			return dto.OwnerDetail{}, apperror.FieldValidationError("name", "", err.Error())
		}

		if err := owner.UpdateName(fullName); err != nil {
			return dto.OwnerDetail{}, apperror.FieldValidationError("name", "", err.Error())
		}
	}

	if err := uc.ownerRepo.Save(ctx, &owner); err != nil {
		return dto.OwnerDetail{}, err
	}

	return mapper.ToResponse(&owner), nil
}
