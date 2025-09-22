package query

import (
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
	"context"
)

type PetQueryHandler interface {
	FindPetsByCustomerID(ctx context.Context, query FindPetsByCustomerIDQuery) (page.Page[PetResult], error)
	FindPetByID(ctx context.Context, query FindPetByIDQuery) (PetResult, error)
	FindPetBySpecification(ctx context.Context, query FindPetBySpecificationQuery) (page.Page[PetResult], error)
	FindPetsBySpecies(ctx context.Context, query FindPetsBySpeciesQuery) (page.Page[PetResult], error)
}

type petQueryHandler struct {
	petRepository      repository.PetRepository
	customerRepository repository.CustomerRepository
}

func NewPetQueryHandler(petRepo repository.PetRepository, customerRepo repository.CustomerRepository) PetQueryHandler {
	return &petQueryHandler{
		petRepository:      petRepo,
		customerRepository: customerRepo,
	}
}

func (h *petQueryHandler) FindPetsByCustomerID(ctx context.Context, query FindPetsByCustomerIDQuery) (page.Page[PetResult], error) {
	petsPage, err := h.petRepository.FindByCustomerID(ctx, query.customerID, query.pagination)
	if err != nil {
		return page.Page[PetResult]{}, err
	}

	results := entitiesToResults(petsPage.Items)
	return page.NewPage(results, petsPage.Metadata), nil
}

func (h *petQueryHandler) FindPetByID(ctx context.Context, query FindPetByIDQuery) (PetResult, error) {
	if query.customerID != nil {
		pet, err := h.petRepository.FindByIDAndCustomerID(ctx, query.petID, *query.customerID)
		if err != nil {
			return PetResult{}, err
		}
		return entityToResult(pet), nil
	}

	pet, err := h.petRepository.FindByID(ctx, query.petID)
	if err != nil {
		return PetResult{}, err
	}

	return entityToResult(pet), nil
}

func (h *petQueryHandler) FindPetBySpecification(ctx context.Context, query FindPetBySpecificationQuery) (page.Page[PetResult], error) {
	petsPage, err := h.petRepository.FindBySpecification(ctx, query.specification)
	if err != nil {
		return page.Page[PetResult]{}, err
	}

	results := entitiesToResults(petsPage.Items)
	return page.NewPage(results, petsPage.Metadata), nil
}

func (h *petQueryHandler) FindPetsBySpecies(ctx context.Context, query FindPetsBySpeciesQuery) (page.Page[PetResult], error) {
	petsPage, err := h.petRepository.FindBySpecies(ctx, query.PetSpecies, query.pagination)
	if err != nil {
		return page.Page[PetResult]{}, err
	}

	results := entitiesToResults(petsPage.Items)
	return page.NewPage(results, petsPage.Metadata), nil
}
