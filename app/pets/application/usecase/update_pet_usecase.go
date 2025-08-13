package petUsecase

import (
	"context"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/dtos"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/mapper"
	petRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/repositories"
)

type UpdatePetUseCase struct {
	petRepository   petRepository.PetRepository
	ownerRepository ownerDomain.OwnerRepository
}

func NewUpdatePetUseCase(petRepository petRepository.PetRepository, ownerRepository ownerDomain.OwnerRepository) *UpdatePetUseCase {
	return &UpdatePetUseCase{
		petRepository:   petRepository,
		ownerRepository: ownerRepository,
	}
}

func (uc UpdatePetUseCase) Execute(ctx context.Context, petId int, petUpdate petDTOs.PetUpdate) (petDTOs.PetResponse, error) {
	pet, err := uc.petRepository.GetById(ctx, petId)
	if err != nil {
		return petDTOs.PetResponse{}, err
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

func (uc UpdatePetUseCase) validate_owner(ctx context.Context, owner_id int) error {
	_, err := uc.ownerRepository.GetById(ctx, owner_id)
	if err != nil {
		return err
	}

	return nil
}
