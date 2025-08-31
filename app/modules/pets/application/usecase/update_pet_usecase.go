package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/mapper"
)

type UpdatePetUseCase struct {
	petRepository   repository.PetRepository
	ownerRepository repository.OwnerRepository
}

func NewUpdatePetUseCase(petRepository repository.PetRepository, ownerRepository repository.OwnerRepository) *UpdatePetUseCase {
	return &UpdatePetUseCase{
		petRepository:   petRepository,
		ownerRepository: ownerRepository,
	}
}

func (uc UpdatePetUseCase) Execute(ctx context.Context, petUpdate dto.PetUpdate) (dto.PetResponse, error) {
	pet, err := uc.petRepository.GetByID(ctx, petUpdate.PetID)
	if err != nil {
		return dto.PetResponse{}, err
	}

	if petUpdate.OwnerID != nil {
		if err := uc.validate_owner(ctx, *petUpdate.OwnerID); err != nil {
			return dto.PetResponse{}, err
		}
	}

	mapper.ToDomainFromUpdate(&pet, petUpdate)
	if err := uc.petRepository.Save(ctx, &pet); err != nil {
		return dto.PetResponse{}, err
	}
	return mapper.ToResponse(&pet), nil
}

func (uc UpdatePetUseCase) validate_owner(ctx context.Context, ownerID valueobject.OwnerID) error {
	_, err := uc.ownerRepository.GetByID(ctx, ownerID)
	if err != nil {
		return err
	}

	return nil
}
