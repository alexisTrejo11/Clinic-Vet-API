package query

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
)

type FindPetsByCustomerIDQuery struct {
	customerID valueobject.CustomerID
	pagination page.PageInput
}

func (q *FindPetsByCustomerIDQuery) Validate() error {
	if q.customerID.IsZero() {
		return nil
	}

	if err := q.pagination.Validate(); err != nil {
		return err
	}

	return nil
}

func NewFindPetsByCustomerIDQuery(customerID uint, pagination page.PageInput) *FindPetsByCustomerIDQuery {
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
	pagination    page.PageInput
}

func NewFindPetBySpecificationQuery(spec specification.PetSpecification) *FindPetBySpecificationQuery {
	return &FindPetBySpecificationQuery{
		specification: spec,
	}
}

type FindPetsBySpeciesQuery struct {
	PetSpecies enum.PetSpecies
	pagination page.PageInput
}

func NewFindPetsBySpeciesQuery(petSpecies string, pagination page.PageInput) *FindPetsBySpeciesQuery {
	return &FindPetsBySpeciesQuery{
		PetSpecies: enum.PetSpecies(petSpecies),
		pagination: pagination,
	}
}
