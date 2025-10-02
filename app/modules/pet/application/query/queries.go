package query

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
	"clinic-vet-api/app/shared/page"
)

type FindPetsByCustomerIDQuery struct {
	customerID valueobject.CustomerID
	pagination page.PaginationRequest
}

func (q *FindPetsByCustomerIDQuery) Validate() error {
	if q.customerID.IsZero() {
		return apperror.ValidationError("customer ID cannot be zero")
	}

	if !q.pagination.IsValid() {
		return apperror.ValidationError("invalid pagination parameters")
	}

	return nil
}

func NewFindPetsByCustomerIDQuery(customerID uint, pagination page.PaginationRequest) *FindPetsByCustomerIDQuery {
	return &FindPetsByCustomerIDQuery{
		customerID: valueobject.NewCustomerID(customerID),
		pagination: pagination,
	}
}

type FindPetByIDQuery struct {
	petID      valueobject.PetID
	customerID *valueobject.CustomerID
}

func NewFindPetByIDQuery(petID uint, customerID *uint) *FindPetByIDQuery {
	var custID *valueobject.CustomerID
	if customerID != nil {
		c := valueobject.NewCustomerID(*customerID)
		custID = &c
	}
	return &FindPetByIDQuery{
		petID:      valueobject.NewPetID(petID),
		customerID: custID,
	}
}

type FindPetBySpecificationQuery struct {
	specification specification.PetSpecification
	pagination    page.PaginationRequest
}

func NewFindPetBySpecificationQuery(spec specification.PetSpecification) *FindPetBySpecificationQuery {
	return &FindPetBySpecificationQuery{
		specification: spec,
		pagination:    page.NewPaginationRequest(),
	}
}

type FindPetsBySpeciesQuery struct {
	PetSpecies enum.PetSpecies
	pagination page.PaginationRequest
}

func NewFindPetsBySpeciesQuery(petSpecies string, pagination page.PaginationRequest) *FindPetsBySpeciesQuery {
	return &FindPetsBySpeciesQuery{
		PetSpecies: enum.PetSpecies(petSpecies),
		pagination: pagination,
	}
}
