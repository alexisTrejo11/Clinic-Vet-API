package petUsecase

import (
	"context"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase/dtos"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase/mapper"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

type UpdatePetUseCase struct {
	petRepository   petDomain.PetRepository
	ownerRepository ownerDomain.OwnerRepository
}

func NewUpdatePetUseCase(petRepository petDomain.PetRepository, ownerRepository ownerDomain.OwnerRepository) *UpdatePetUseCase {
	return &UpdatePetUseCase{
		petRepository:   petRepository,
		ownerRepository: ownerRepository,
	}
}

func (uc UpdatePetUseCase) Execute(ctx context.Context, petUpdate petDTOs.PetUpdate) (petDTOs.PetResponse, error) {
	pet, err := uc.petRepository.GetById(ctx, petUpdate.PetId)
	if err != nil {
		return petDTOs.PetResponse{}, err
	}

	if petUpdate.OwnerID != nil {
		if err := uc.validate_owner(ctx, *petUpdate.OwnerID); err != nil {
			return petDTOs.PetResponse{}, err
		}
	}

	petMapper.ToDomainFromUpdate(&pet, petUpdate)
	if err := uc.petRepository.Save(ctx, &pet); err != nil {
		return petDTOs.PetResponse{}, err
	}

	return petMapper.ToResponse(&pet), nil
}

func (uc UpdatePetUseCase) validate_owner(ctx context.Context, owner_id int) error {
	_, err := uc.ownerRepository.GetById(ctx, owner_id)
	if err != nil {
		return err
	}

	return nil
}
