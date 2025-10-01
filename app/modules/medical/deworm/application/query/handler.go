package query

import (
	"context"

	med "clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	p "clinic-vet-api/app/shared/page"
)

type DewormQueryHandler struct {
	dewormRepo   repository.DewormRepository
	employeeRepo repository.EmployeeRepository
	petRepo      repository.PetRepository
}

func NewDewormQueryHandler(
	dewormRepo repository.DewormRepository,
	employeeRepo repository.EmployeeRepository,
	petRepo repository.PetRepository,
) *DewormQueryHandler {
	return &DewormQueryHandler{
		dewormRepo:   dewormRepo,
		employeeRepo: employeeRepo,
		petRepo:      petRepo,
	}
}

func (h *DewormQueryHandler) HandleByIDQuery(ctx context.Context, qry FindDewormByIDQuery) (DewormResult, error) {
	var deworm *med.PetDeworming
	var err error

	if qry.OptPetID != nil {
		deworm, err = h.dewormRepo.FindByIDAndPetID(ctx, qry.ID, *qry.OptPetID)
		if err != nil {
			return DewormResult{}, err
		}
	} else if qry.OptEmployeeID != nil {
		deworm, err = h.dewormRepo.FindByIDAndEmployeeID(ctx, qry.ID, *qry.OptEmployeeID)
		if err != nil {
			return DewormResult{}, err
		}
	} else {
		deworm, err = h.dewormRepo.FindByID(ctx, qry.ID)
		if err != nil {
			return DewormResult{}, err
		}
	}

	return toDewormResult(*deworm), nil
}

func (h *DewormQueryHandler) HandleByPetQuery(ctx context.Context, qry FindDewormsByPetQuery) (p.Page[DewormResult], error) {
	dewormPage, err := h.dewormRepo.FindByPetID(ctx, qry.PetID, qry.Pagination)
	if err != nil {
		return p.Page[DewormResult]{}, err
	}

	return p.MapItems(dewormPage, toDewormResult), nil
}

func (h *DewormQueryHandler) HandleByEmployeeQuery(ctx context.Context, qry FindDewormsByEmployeeQuery) (p.Page[DewormResult], error) {
	dewormPage, err := h.dewormRepo.FindByEmployeeID(ctx, qry.EmployeeID, qry.Pagination)
	if err != nil {
		return p.Page[DewormResult]{}, err
	}

	return p.MapItems(dewormPage, toDewormResult), nil
}

func (h *DewormQueryHandler) HandleByDateRangeQuery(ctx context.Context, qry FindDewormsByDateRangeQuery) (p.Page[DewormResult], error) {
	dewormPage, err := h.dewormRepo.FindByDateRange(ctx, qry.StartDate, qry.EndDate, qry.Pagination)
	if err != nil {
		return p.Page[DewormResult]{}, err
	}

	return p.MapItems(dewormPage, toDewormResult), nil
}

func (h *DewormQueryHandler) HandleByCustomerQuery(ctx context.Context, qry FindDewormsByCustomerQuery) (p.Page[DewormResult], error) {
	pets, err := h.petRepo.FindAllByCustomerID(ctx, qry.CustomerID)
	if err != nil {
		return p.Page[DewormResult]{}, err
	}

	if len(pets) == 0 {
		return p.Page[DewormResult]{}, nil
	}

	if qry.OptPetID != nil {
		found := false
		for _, pet := range pets {
			if pet.ID().Equals(qry.OptPetID.Value()) {
				found = true
				break
			}
		}

		// Invalid Ownership request, return empty result
		if !found {
			return p.Page[DewormResult]{}, nil
		}

		dewormPage, err := h.dewormRepo.FindByPetID(ctx, *qry.OptPetID, qry.Pagination)
		if err != nil {
			return p.Page[DewormResult]{}, err
		}

		return p.MapItems(dewormPage, toDewormResult), nil
	}

	petIDs := make([]valueobject.PetID, len(pets))
	for i, pet := range pets {
		petIDs[i] = pet.ID()
	}

	dewormPage, err := h.dewormRepo.FindByPetIDs(ctx, petIDs, qry.Pagination)
	if err != nil {
		return p.Page[DewormResult]{}, err
	}
	return p.MapItems(dewormPage, toDewormResult), nil
}
