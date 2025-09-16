package usecase

import (
	"context"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/modules/pets/application/dto"
	"clinic-vet-api/app/modules/pets/application/mapper"
)

type FindPetByIDAndCustomerUseCase struct {
	petRepo repository.PetRepository
}

func NewFindPetByIDAndCustomerIDUseCase(petRepo repository.PetRepository) *FindPetByIDAndCustomerUseCase {
	return &FindPetByIDAndCustomerUseCase{
		petRepo: petRepo,
	}
}

func (uc *FindPetByIDAndCustomerUseCase) Execute(ctx context.Context, petID valueobject.PetID, customerID valueobject.CustomerID) (dto.PetResponse, error) {
	pet, err := uc.petRepo.FindByIDAndCustomerID(ctx, petID, customerID)
	if err != nil {
		return dto.PetResponse{}, err
	}
	return mapper.ToResponse(&pet), nil
}
