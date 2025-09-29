package query

import (
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
)

type FindApptsByEmployeeIDQuery struct {
	employeeID valueobject.EmployeeID
	pagination specification.Pagination
}

func NewFindApptsByEmployeeIDQuery(employeeID uint, pagination page.PaginationRequest) *FindApptsByEmployeeIDQuery {
	return &FindApptsByEmployeeIDQuery{
		employeeID: valueobject.NewEmployeeID(employeeID),
		pagination: pagination.ToSpecPagination(),
	}
}

func (q *FindApptsByEmployeeIDQuery) EmployeeID() valueobject.EmployeeID {
	return q.employeeID
}

func (q *FindApptsByEmployeeIDQuery) Pagination() specification.Pagination {
	return q.pagination
}
