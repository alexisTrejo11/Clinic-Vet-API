package usecase

import (
	"context"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/modules/pets/application/dto"
	"clinic-vet-api/app/modules/pets/application/mapper"
	"clinic-vet-api/app/shared/page"
)

type FindPetByCustomerID struct {
	petRepo repository.PetRepository
}

func NewFindPetsByCustomerIDUseCase(petRepo repository.PetRepository) *FindPetByCustomerID {
	return &FindPetByCustomerID{
		petRepo: petRepo,
	}
}

func (uc *FindPetByCustomerID) Execute(ctx context.Context, customerID valueobject.CustomerID, pageInput page.PageInput) (page.Page[dto.PetResponse], error) {
	petPage, err := uc.petRepo.FindByCustomerID(ctx, customerID, pageInput)
	if err != nil {
		return page.Page[dto.PetResponse]{}, err
	}

	return mapper.ToResponsesPage(petPage), nil
}
