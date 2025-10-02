package handler

import (
	med "clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	q "clinic-vet-api/app/modules/medical/vaccination/application/query"
	p "clinic-vet-api/app/shared/page"
	"context"
)

type VaccinationQueryHandler struct {
	vaccinationRepo repository.VaccinationRepository
	employeeRepo    repository.EmployeeRepository
	petRepo         repository.PetRepository
}

func NewVaccinationQueryHandler(
	vaccinationRepo repository.VaccinationRepository,
	employeeRepo repository.EmployeeRepository,
	petRepo repository.PetRepository,
) *VaccinationQueryHandler {
	return &VaccinationQueryHandler{
		vaccinationRepo: vaccinationRepo,
		employeeRepo:    employeeRepo,
		petRepo:         petRepo,
	}
}

func (h *VaccinationQueryHandler) HandleVaccinationByID(
	ctx context.Context,
	qry q.FindVaccinationByIDQuery,
) (VaccinationResult, error) {
	var vaccination *med.PetVaccination
	var err error

	if qry.OptPetID() != nil {
		vaccination, err = h.vaccinationRepo.FindByIDAndPetID(ctx, qry.ID(), *qry.OptPetID())
		if err != nil {
			return VaccinationResult{}, err
		}
	} else if qry.OptEmployeeID() != nil {
		vaccination, err = h.vaccinationRepo.FindByIDAndEmployeeID(ctx, qry.ID(), *qry.OptEmployeeID())
		if err != nil {
			return VaccinationResult{}, err
		}
	} else {
		vaccination, err = h.vaccinationRepo.FindByID(ctx, qry.ID())
		if err != nil {
			return VaccinationResult{}, err
		}
	}

	return toVaccinationResult(*vaccination), nil
}

func (h *VaccinationQueryHandler) HandleVaccinationsByPet(
	ctx context.Context,
	qry q.FindVaccinationsByPetQuery,
) (p.Page[VaccinationResult], error) {
	vaccinePage, err := h.vaccinationRepo.FindByPetID(ctx, qry.PetID(), qry.Pagination())
	if err != nil {
		return p.Page[VaccinationResult]{}, err
	}

	return p.MapItems(vaccinePage, toVaccinationResult), nil
}

func (h *VaccinationQueryHandler) HandleVaccinationsByEmployee(
	ctx context.Context,
	qry q.FindVaccinationsByEmployeeQuery,
) (p.Page[VaccinationResult], error) {
	vaccinePage, err := h.vaccinationRepo.FindByEmployeeID(ctx, qry.EmployeeID(), qry.Pagination())
	if err != nil {
		return p.Page[VaccinationResult]{}, err
	}

	return p.MapItems(vaccinePage, toVaccinationResult), nil
}

func (h *VaccinationQueryHandler) HandleVaccinationsByDateRange(
	ctx context.Context,
	qry q.FindVaccinationsByDateRangeQuery,
) (p.Page[VaccinationResult], error) {
	vaccinePage, err := h.vaccinationRepo.FindByDateRange(ctx, qry.StartDate(), qry.EndDate(), qry.Pagination())
	if err != nil {
		return p.Page[VaccinationResult]{}, err
	}

	return p.MapItems(vaccinePage, toVaccinationResult), nil
}

func (h *VaccinationQueryHandler) HandleVaccinationsByCustomer(ctx context.Context, qry q.FindVaccinationsByCustomerQuery) (p.Page[VaccinationResult], error) {
	pets, err := h.petRepo.FindAllByCustomerID(ctx, qry.CustomerID())
	if err != nil {
		return p.Page[VaccinationResult]{}, err
	} else if len(pets) == 0 {
		return p.Page[VaccinationResult]{}, nil
	}

	if qry.OptPetID() != nil {
		// Return empty insdtead of forbbiden error
		if !isPetOwnedByCustomer(pets, qry.CustomerID()) {
			return p.Page[VaccinationResult]{}, nil
		}

		vaccinePage, err := h.vaccinationRepo.FindByPetID(ctx, *qry.OptPetID(), qry.Pagination())
		if err != nil {
			return p.Page[VaccinationResult]{}, err
		}

		return p.MapItems(vaccinePage, toVaccinationResult), nil
	}

	vaccinePage, err := h.vaccinationRepo.FindByPetIDs(ctx, extractPetIDs(pets), qry.Pagination())
	if err != nil {
		return p.Page[VaccinationResult]{}, err
	}

	return p.MapItems(vaccinePage, toVaccinationResult), nil
}

func extractPetIDs(pets []pet.Pet) []valueobject.PetID {
	petIDs := make([]valueobject.PetID, len(pets))
	for i, pet := range pets {
		petIDs[i] = pet.ID()
	}
	return petIDs
}

func isPetOwnedByCustomer(pets []pet.Pet, customerID valueobject.CustomerID) bool {
	for _, pet := range pets {
		if pet.CustomerID() == customerID {
			return true
		}
	}
	return false
}
