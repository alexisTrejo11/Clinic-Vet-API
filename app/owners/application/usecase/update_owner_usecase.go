package ownerUsecase

import (
	"context"

	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerMappers "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/mappers"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type UpdateOwnerUseCase struct {
	ownerRepo ownerDomain.OwnerRepository
}

func NewUpdateOwnerUseCase(ownerRepo ownerDomain.OwnerRepository) *UpdateOwnerUseCase {
	return &UpdateOwnerUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *UpdateOwnerUseCase) Execute(ctx context.Context, id int, dto ownerDTOs.OwnerUpdate) (*ownerDTOs.OwnerResponse, error) {
	owner, err := uc.ownerRepo.GetById(ctx, id)
	if err != nil {
		return nil, ownerDomain.HandleGetByIdError(err, id)
	}

	if dto.PhoneNumber != nil && *dto.PhoneNumber != owner.PhoneNumber() {
		_, err := uc.ownerRepo.GetByPhone(ctx, *dto.PhoneNumber)
		if err == nil {
			return nil, ownerDomain.HandlePhoneConflictError()
		}
		owner.SetPhoneNumber(*dto.PhoneNumber)
	}

	if dto.Photo != nil {
		owner.SetPhoto(*dto.Photo)
	}

	if dto.FirstName != nil || dto.LastName != nil {
		firstName := owner.FullName().FirstName
		if dto.FirstName != nil {
			firstName = *dto.FirstName
		}

		lastName := owner.FullName().LastName
		if dto.LastName != nil {
			lastName = *dto.LastName
		}

		fullName, err := user.NewPersonName(firstName, lastName)
		if err != nil {
			return nil, err
		}
		owner.SetFullName(fullName)
	}

	if dto.Address != nil {
		owner.SetAddress(*dto.Address)
	}

	if err := uc.ownerRepo.Save(ctx, owner); err != nil {
		return nil, err
	}

	return ownerMappers.ToResponse(&owner), nil
}
