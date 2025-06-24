package petUsecase

import (
	"context"

	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
	petAppError "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application"
	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/dtos"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/mapper"
	petRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/repositories"
)

type UpdatePetUseCase struct {
	petRepository   petRepository.PetRepository
	ownerRepository ownerRepository.OwnerRepository
}

func NewUpdatePetUseCase(petRepository petRepository.PetRepository, ownerRepository ownerRepository.OwnerRepository) *UpdatePetUseCase {
	return &UpdatePetUseCase{
		petRepository:   petRepository,
		ownerRepository: ownerRepository,
	}
}

func (uc UpdatePetUseCase) Execute(ctx context.Context, petId uint, petUpdate petDTOs.PetUpdate) (petDTOs.PetResponse, error) {
	pet, err := uc.petRepository.GetById(ctx, petId)
	if err != nil {
		return petDTOs.PetResponse{}, petAppError.HandleGetByIdError(err, petId)
	}

	if petUpdate.OwnerID != nil {
		if err := uc.validate_owner(ctx, petId); err != nil {
			return petDTOs.PetResponse{}, err
		}
	}

	petMapper.ToDomainFromUpdate(&pet, petUpdate)
	if err := uc.petRepository.Save(ctx, &pet); err != nil {
		return petDTOs.PetResponse{}, err
	}

	return petMapper.ToResponse(pet), nil
}

func (uc UpdatePetUseCase) validate_owner(ctx context.Context, owner_id uint) error {
	_, err := uc.ownerRepository.GetByID(ctx, owner_id)
	if err != nil {
		notFounderr := petAppError.HandleGetByIdError(err, owner_id)
		return notFounderr
	}

	return nil
}
