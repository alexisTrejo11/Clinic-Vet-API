package query

import (
	"clinic-vet-api/app/shared/page"
)

type FindActiveEmployeesQuery struct {
	pagination page.PaginationRequest
}

func NewFindActiveEmployeesQuery(pagination page.PaginationRequest) (FindActiveEmployeesQuery, error) {
	return FindActiveEmployeesQuery{pagination: pagination}, nil
}

func (q FindActiveEmployeesQuery) Pagination() page.PaginationRequest {
	return q.pagination
}
