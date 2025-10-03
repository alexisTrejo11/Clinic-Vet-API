package query

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	p "clinic-vet-api/app/shared/page"
)

type FindVaccinationsByPetQuery struct {
	petID         valueobject.PetID
	optEmployeeID *valueobject.EmployeeID
	pagination    p.PaginationRequest
}

func NewFindVaccinationsByPetQuery(
	petID uint,
	optEmployeeID *uint,
	pagination p.PaginationRequest,
) (FindVaccinationsByPetQuery, error) {
	cmd := FindVaccinationsByPetQuery{
		petID:         valueobject.NewPetID(petID),
		optEmployeeID: valueobject.NewOptEmployeeID(optEmployeeID),
		pagination:    pagination,
	}

	if err := cmd.Validate(); err != nil {
		return FindVaccinationsByPetQuery{}, err
	}

	return cmd, nil
}

func (q FindVaccinationsByPetQuery) Validate() error {
	return nil
}

func (q FindVaccinationsByPetQuery) PetID() valueobject.PetID               { return q.petID }
func (q FindVaccinationsByPetQuery) OptEmployeeID() *valueobject.EmployeeID { return q.optEmployeeID }
func (q FindVaccinationsByPetQuery) Pagination() p.PaginationRequest        { return q.pagination }
