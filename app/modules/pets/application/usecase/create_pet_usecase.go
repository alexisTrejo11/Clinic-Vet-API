package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	petApplicationError "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/mapper"
)

type CreatePetUseCase struct {
	petRepository   repository.PetRepository
	ownerRepository repository.OwnerRepository
}

func NewCreatePetUseCase(
	petRepository repository.PetRepository,
	ownerRepository repository.OwnerRepository,
) *CreatePetUseCase {
	return &CreatePetUseCase{
		petRepository:   petRepository,
		ownerRepository: ownerRepository,
	}
}

func (uc CreatePetUseCase) Execute(ctx context.Context, petCreate dto.PetCreate) (dto.PetResponse, error) {
	if err := uc.validateOwner(ctx, petCreate.OwnerID); err != nil {
		return dto.PetResponse{}, err
	}

	newPet := mapper.ToDomainFromCreate(petCreate)
	if err := uc.petRepository.Save(ctx, newPet); err != nil {
		return dto.PetResponse{}, err
	}

	return mapper.ToResponse(newPet), nil
}

func (uc CreatePetUseCase) validateOwner(ctx context.Context, ownerID valueobject.OwnerID) error {
	exists, err := uc.ownerRepository.ExistsByID(ctx, ownerID)
	if err != nil {
		return err
	}

	if !exists {
		petApplicationError.OwnerNotFoundError(ownerID.GetValue())
	}

	return nil
}
