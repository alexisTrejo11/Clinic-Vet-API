package query

import (
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
)

type FindApptsByCustomerIDQuery struct {
	customerID valueobject.CustomerID
	petID      *valueobject.PetID
	pagination specification.Pagination
}

func NewFindApptsByCustomerIDQuery(pagination page.PaginationRequest, customerId uint, petID *uint, status *string) FindApptsByCustomerIDQuery {
	var petIDvo *valueobject.PetID
	if petID != nil {
		val := valueobject.NewPetID(*petID)
		petIDvo = &val
	}

	return FindApptsByCustomerIDQuery{
		customerID: valueobject.NewCustomerID(customerId),
		pagination: pagination.ToSpecPagination(),
		petID:      petIDvo,
	}
}

func (q FindApptsByCustomerIDQuery) CustomerID() valueobject.CustomerID   { return q.customerID }
func (q FindApptsByCustomerIDQuery) PetID() *valueobject.PetID            { return q.petID }
func (q FindApptsByCustomerIDQuery) Pagination() specification.Pagination { return q.pagination }
