package usecase

import (
	"context"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/modules/pets/application/dto"
	"clinic-vet-api/app/modules/pets/application/mapper"
	"clinic-vet-api/app/shared/page"
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
