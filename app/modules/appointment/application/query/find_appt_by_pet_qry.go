package query

import (
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
)

type FindApptsByPetQuery struct {
	petID      valueobject.PetID
	pagination specification.Pagination
}

func NewFindApptsByPetQuery(employeeID uint, pagination page.PaginationRequest) FindApptsByPetQuery {
	return FindApptsByPetQuery{petID: valueobject.NewPetID(employeeID), pagination: pagination.ToSpecPagination()}
}

func (q FindApptsByPetQuery) PetID() valueobject.PetID             { return q.petID }
func (q FindApptsByPetQuery) Pagination() specification.Pagination { return q.pagination }
