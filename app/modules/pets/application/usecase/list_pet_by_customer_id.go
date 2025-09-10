package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/mapper"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListPetByCustomerID struct {
	petRepo repository.PetRepository
}

func NewListPetsByCustomerIDUseCase(petRepo repository.PetRepository) *ListPetByCustomerID {
	return &ListPetByCustomerID{
		petRepo: petRepo,
	}
}

func (uc *ListPetByCustomerID) Execute(ctx context.Context, customerID valueobject.CustomerID) (page.Page[[]dto.PetResponse], error) {
	petPage, err := uc.petRepo.ListByCustomerID(ctx, customerID)
	if err != nil {
		return page.Page[[]dto.PetResponse]{}, err
	}

	return mapper.ToResponsesPage(petPage), nil
}
