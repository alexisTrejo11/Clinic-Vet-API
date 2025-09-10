package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/mapper"
)

type UpdatePetUseCase struct {
	petRepository   repository.PetRepository
	ownerRepository repository.CustomerRepository
}

func NewUpdatePetUseCase(petRepository repository.PetRepository, ownerRepository repository.CustomerRepository) *UpdatePetUseCase {
	return &UpdatePetUseCase{
		petRepository:   petRepository,
		ownerRepository: ownerRepository,
	}
}

func (uc UpdatePetUseCase) Execute(ctx context.Context, petUpdate dto.PetUpdateData) (dto.PetResponse, error) {
	pet, err := uc.petRepository.GetByID(ctx, petUpdate.PetID)
	if err != nil {
		return dto.PetResponse{}, err
	}

	if petUpdate.CustomerID != nil {
		if err := uc.validate_owner(ctx, *petUpdate.CustomerID); err != nil {
			return dto.PetResponse{}, err
		}
	}

	mapper.ToDomainFromUpdate(&pet, petUpdate)
	if err := uc.petRepository.Save(ctx, &pet); err != nil {
		return dto.PetResponse{}, err
	}
	return mapper.ToResponse(&pet), nil
}

func (uc UpdatePetUseCase) validate_owner(ctx context.Context, customerID valueobject.CustomerID) error {
	_, err := uc.ownerRepository.GetByID(ctx, customerID)
	if err != nil {
		return err
	}

	return nil
}
