// Package usecase implements the use case for creating a new pet.
package usecase

import (
	"context"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/modules/pets/application/dto"
	"clinic-vet-api/app/modules/pets/application/mapper"
	apperror "clinic-vet-api/app/shared/error/application"
)

type CreatePetUseCase struct {
	petRepository   repository.PetRepository
	ownerRepository repository.CustomerRepository
}

func NewCreatePetUseCase(
	petRepository repository.PetRepository,
	ownerRepository repository.CustomerRepository,
) *CreatePetUseCase {
	return &CreatePetUseCase{
		petRepository:   petRepository,
		ownerRepository: ownerRepository,
	}
}

func (uc CreatePetUseCase) Execute(ctx context.Context, petCreate dto.CreatePetData) (dto.PetResponse, error) {
	if err := uc.validateCustomer(ctx, petCreate.CustomerID); err != nil {
		return dto.PetResponse{}, err
	}

	newPet, err := mapper.ToDomainFromCreate(petCreate)
	if err != nil {
		return dto.PetResponse{}, err
	}

	if err := uc.petRepository.Save(ctx, newPet); err != nil {
		return dto.PetResponse{}, err
	}

	return mapper.ToResponse(newPet), nil
}

func (uc CreatePetUseCase) validateCustomer(ctx context.Context, customerID valueobject.CustomerID) error {
	exists, err := uc.ownerRepository.ExistsByID(ctx, customerID)
	if err != nil {
		return err
	}

	if !exists {
		apperror.EntityValidationError("customer", "id", customerID.String())
	}

	return nil
}
