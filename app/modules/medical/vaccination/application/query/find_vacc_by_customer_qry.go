package query

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	p "clinic-vet-api/app/shared/page"
)

type FindVaccinationsByCustomerQuery struct {
	customerID valueobject.CustomerID
	optPetID   *valueobject.PetID
	pagination p.PaginationRequest
}

func NewFindVaccinationsByCustomerQuery(
	CustomerID uint,
	OptPetID *uint,
	pagination p.PaginationRequest,
) (FindVaccinationsByCustomerQuery, error) {
	cmd := &FindVaccinationsByCustomerQuery{
		customerID: valueobject.NewCustomerID(CustomerID),
		optPetID:   valueobject.NewOptPetID(OptPetID),
		pagination: pagination,
	}

	if err := cmd.Validate(); err != nil {
		return FindVaccinationsByCustomerQuery{}, err
	}

	return *cmd, nil
}

func (q *FindVaccinationsByCustomerQuery) Validate() error {
	return nil
}

func (q FindVaccinationsByCustomerQuery) CustomerID() valueobject.CustomerID { return q.customerID }
func (q FindVaccinationsByCustomerQuery) OptPetID() *valueobject.PetID {
	return q.optPetID
}
func (q FindVaccinationsByCustomerQuery) Pagination() p.PaginationRequest { return q.pagination }
