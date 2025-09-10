package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/mapper"
)

type GetPetByIDAndCustomerUseCase struct {
	petRepo repository.PetRepository
}

func NewGetPetByIDAndCustomerIDUseCase(petRepo repository.PetRepository) *GetPetByIDAndCustomerUseCase {
	return &GetPetByIDAndCustomerUseCase{
		petRepo: petRepo,
	}
}

func (uc *GetPetByIDAndCustomerUseCase) Execute(ctx context.Context, petID valueobject.PetID, customerID valueobject.CustomerID) (dto.PetResponse, error) {
	pet, err := uc.petRepo.GetByIDAndCustomerID(ctx, petID, customerID)
	if err != nil {
		return dto.PetResponse{}, err
	}
	return mapper.ToResponse(&pet), nil
}
