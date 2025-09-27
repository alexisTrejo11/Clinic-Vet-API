package vaccinequery

import (
	med "clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/repository"
	p "clinic-vet-api/app/shared/page"
	"context"
)

type VaccinationQueryHandler interface {
	FindVaccinationByID(ctx context.Context, qry FindVaccinationByIDQuery) (VaccinationResult, error)
	FindVaccinationsByPet(ctx context.Context, qry FindVaccinationsByPetQuery) (p.Page[VaccinationResult], error)
	FindVaccinationsByEmployee(ctx context.Context, qry FindVaccinationsByEmployeeQuery) (p.Page[VaccinationResult], error)
	FindVaccinationsByDateRange(ctx context.Context, qry FindVaccinationsByDateRangeQuery) (p.Page[VaccinationResult], error)
}

type vaccinationQueryHandler struct {
	vaccinationRepo repository.VaccinationRepository
	employeeRepo    repository.EmployeeRepository
	petRepo         repository.PetRepository
}

func NewVaccinationQueryHandler(
	vaccinationRepo repository.VaccinationRepository,
	employeeRepo repository.EmployeeRepository,
	petRepo repository.PetRepository,
) VaccinationQueryHandler {
	return &vaccinationQueryHandler{
		vaccinationRepo: vaccinationRepo,
		employeeRepo:    employeeRepo,
		petRepo:         petRepo,
	}
}

func (h *vaccinationQueryHandler) FindVaccinationByID(ctx context.Context, qry FindVaccinationByIDQuery) (VaccinationResult, error) {
	var vaccination *med.PetVaccination
	var err error

	if qry.OptPetID != nil {
		vaccination, err = h.vaccinationRepo.FindByIDAndPetID(ctx, qry.ID, *qry.OptPetID)
		if err != nil {
			return VaccinationResult{}, err
		}
	} else if qry.OptEmployeeID != nil {
		vaccination, err = h.vaccinationRepo.FindByIDAndEmployeeID(ctx, qry.ID, *qry.OptEmployeeID)
		if err != nil {
			return VaccinationResult{}, err
		}
	} else {
		vaccination, err = h.vaccinationRepo.FindByID(ctx, qry.ID)
		if err != nil {
			return VaccinationResult{}, err
		}
	}

	return toVaccinationResult(*vaccination), nil
}

func (h *vaccinationQueryHandler) FindVaccinationsByPet(ctx context.Context, qry FindVaccinationsByPetQuery) (p.Page[VaccinationResult], error) {
	vaccinePage, err := h.vaccinationRepo.FindByPetID(ctx, qry.PetID, qry.Pagination)
	if err != nil {
		return p.Page[VaccinationResult]{}, err
	}

	return p.MapItems(vaccinePage, toVaccinationResult), nil
}

func (h *vaccinationQueryHandler) FindVaccinationsByEmployee(ctx context.Context, qry FindVaccinationsByEmployeeQuery) (p.Page[VaccinationResult], error) {
	vaccinePage, err := h.vaccinationRepo.FindByEmployeeID(ctx, qry.EmployeeID, qry.Pagination)
	if err != nil {
		return p.Page[VaccinationResult]{}, err
	}

	return p.MapItems(vaccinePage, toVaccinationResult), nil
}

func (h *vaccinationQueryHandler) FindVaccinationsByDateRange(ctx context.Context, qry FindVaccinationsByDateRangeQuery) (p.Page[VaccinationResult], error) {
	vaccinePage, err := h.vaccinationRepo.FindByDateRange(ctx, qry.StartDate, qry.EndDate, qry.Pagination)
	if err != nil {
		return p.Page[VaccinationResult]{}, err
	}

	return p.MapItems(vaccinePage, toVaccinationResult), nil
}
