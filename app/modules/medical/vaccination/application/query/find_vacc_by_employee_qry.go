package query

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	p "clinic-vet-api/app/shared/page"
)

type FindVaccinationsByEmployeeQuery struct {
	employeeID valueobject.EmployeeID
	pagination p.PaginationRequest
}

func NewFindVaccinationsByEmployeeQuery(employeeID uint, pagination p.PaginationRequest) (FindVaccinationsByEmployeeQuery, error) {
	return FindVaccinationsByEmployeeQuery{
		employeeID: valueobject.NewEmployeeID(employeeID),
		pagination: pagination,
	}, nil
}

func (q FindVaccinationsByEmployeeQuery) EmployeeID() valueobject.EmployeeID { return q.employeeID }
func (q FindVaccinationsByEmployeeQuery) Pagination() p.PaginationRequest    { return q.pagination }

func (q FindVaccinationsByEmployeeQuery) Validate() error {
	return nil
}
