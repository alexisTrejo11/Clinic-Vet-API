package usecase

import (
	"context"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/modules/pets/application/dto"
	"clinic-vet-api/app/modules/pets/application/mapper"
)

type UpdatePetUseCase struct {
	petRepository      repository.PetRepository
	customerRepository repository.CustomerRepository
}

func NewUpdatePetUseCase(petRepository repository.PetRepository, customerRepository repository.CustomerRepository) *UpdatePetUseCase {
	return &UpdatePetUseCase{
		petRepository:      petRepository,
		customerRepository: customerRepository,
	}
}

func (uc UpdatePetUseCase) Execute(ctx context.Context, petUpdate dto.PetUpdateData) (dto.PetResponse, error) {
	pet, err := uc.petRepository.FindByID(ctx, petUpdate.PetID)
	if err != nil {
		return dto.PetResponse{}, err
	}

	if petUpdate.CustomerID != nil {
		if err := uc.validate_customer(ctx, *petUpdate.CustomerID); err != nil {
			return dto.PetResponse{}, err
		}
	}

	mapper.ToDomainFromUpdate(ctx, &pet, petUpdate)
	if err := uc.petRepository.Save(ctx, &pet); err != nil {
		return dto.PetResponse{}, err
	}
	return mapper.ToResponse(&pet), nil
}

func (uc UpdatePetUseCase) validate_customer(ctx context.Context, customerID valueobject.CustomerID) error {
	_, err := uc.customerRepository.FindByID(ctx, customerID)
	if err != nil {
		return err
	}

	return nil
}
