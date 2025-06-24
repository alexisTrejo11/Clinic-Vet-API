package petUsecase

import (
	"context"

	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
	petAppError "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application"
	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/dtos"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/mapper"
	petRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/repositories"
)

type CreatePetUseCase struct {
	petRepository   petRepository.PetRepository
	ownerRepository ownerRepository.OwnerRepository
}

func NewCreatePetUseCase(petRepository petRepository.PetRepository, ownerRepository ownerRepository.OwnerRepository) *CreatePetUseCase {
	return &CreatePetUseCase{
		petRepository:   petRepository,
		ownerRepository: ownerRepository,
	}
}

func (uc CreatePetUseCase) Execute(ctx context.Context, petCreate petDTOs.PetCreate) (petDTOs.PetResponse, error) {
	if err := uc.validate_owner(ctx, petCreate.OwnerID); err != nil {
		return petDTOs.PetResponse{}, err
	}

	newPet := petMapper.ToDomainFromCreate(petCreate)
	if err := uc.petRepository.Save(ctx, &newPet); err != nil {
		return petDTOs.PetResponse{}, err
	}

	return petMapper.ToResponse(newPet), nil
}

func (uc CreatePetUseCase) validate_owner(ctx context.Context, owner_id uint) error {
	_, err := uc.ownerRepository.GetByID(ctx, owner_id)
	if err != nil {
		return petAppError.HandleGetByIdError(err, owner_id)

	}

	return nil
}
