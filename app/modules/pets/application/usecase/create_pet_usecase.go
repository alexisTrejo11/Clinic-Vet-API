package petUsecase

import (
	"context"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	petApplicationError "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application"
	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase/dtos"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase/mapper"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

type CreatePetUseCase struct {
	petRepository   petDomain.PetRepository
	ownerRepository ownerDomain.OwnerRepository
}

func NewCreatePetUseCase(
	petRepository petDomain.PetRepository,
	ownerRepository ownerDomain.OwnerRepository,
) *CreatePetUseCase {
	return &CreatePetUseCase{
		petRepository:   petRepository,
		ownerRepository: ownerRepository,
	}
}

func (uc CreatePetUseCase) Execute(ctx context.Context, petCreate petDTOs.PetCreate) (petDTOs.PetResponse, error) {
	if err := uc.validateOwner(ctx, petCreate.OwnerId); err != nil {
		return petDTOs.PetResponse{}, err
	}

	newPet := petMapper.ToDomainFromCreate(petCreate)
	if err := uc.petRepository.Save(ctx, newPet); err != nil {
		return petDTOs.PetResponse{}, err
	}

	return petMapper.ToResponse(newPet), nil
}

func (uc CreatePetUseCase) validateOwner(ctx context.Context, ownerId int) error {
	exists, err := uc.ownerRepository.ExistsByID(ctx, ownerId)
	if err != nil {
		return err
	}

	if !exists {
		petApplicationError.OwnerNotFoundError(ownerId)
	}

	return nil
}
