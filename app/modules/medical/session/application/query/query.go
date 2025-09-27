// Package query contains query definitions for medical history operations.
package query

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	p "clinic-vet-api/app/shared/page"
)

// FindMedSessionByIDQuery represents a query to find a medical history by its ID, with optional Filter only one opt parameters will be considered
type FindMedSessionByIDQuery struct {
	ID            valueobject.MedSessionID
	optCustomerID *valueobject.CustomerID
	optPetID      *valueobject.PetID
	optEmployeeID *valueobject.EmployeeID
}

func NewFindMedSessionByIDQuery(id uint) *FindMedSessionByIDQuery {
	return &FindMedSessionByIDQuery{
		ID: valueobject.NewMedSessionID(id),
	}
}

func FindMedSessionByIDQueryWithCustomerID(id uint, customerID uint) *FindMedSessionByIDQuery {
	custID := valueobject.NewCustomerID(customerID)

	return &FindMedSessionByIDQuery{
		ID: valueobject.NewMedSessionID(id), optCustomerID: &custID,
	}
}

func FindMedSessionByIDQueryWithPetID(id uint, petID uint) *FindMedSessionByIDQuery {
	pID := valueobject.NewPetID(petID)
	return &FindMedSessionByIDQuery{
		ID:       valueobject.NewMedSessionID(id),
		optPetID: &pID,
	}
}

func FindMedSessionByIDQueryWithEmployeeID(id uint, employeeID uint) *FindMedSessionByIDQuery {
	empID := valueobject.NewEmployeeID(employeeID)
	return &FindMedSessionByIDQuery{
		ID:            valueobject.NewMedSessionID(id),
		optEmployeeID: &empID,
	}
}

type FindMedSessionBySpecQuery struct {
	Spec specification.MedicalSessionSpecification
}

type FindAllMedSessionQuery struct {
	PaginationRequest p.PaginationRequest
}

type FindMedSessionByEmployeeIDQuery struct {
	EmployeeID        valueobject.EmployeeID
	PaginationRequest p.PaginationRequest
}

type FindMedSessionByPetIDQuery struct {
	petID             valueobject.PetID
	optCustomerID     *valueobject.CustomerID
	PaginationRequest p.PaginationRequest
}

func NewFindMedSessionByPetIDQuery(petID uint, optCustomerID *uint, PaginationRequest p.PaginationRequest) *FindMedSessionByPetIDQuery {
	var optCustID *valueobject.CustomerID
	if optCustomerID != nil {
		val := valueobject.NewCustomerID(*optCustomerID)
		optCustID = &val
	}

	return &FindMedSessionByPetIDQuery{
		petID:             valueobject.NewPetID(petID),
		optCustomerID:     optCustID,
		PaginationRequest: PaginationRequest,
	}
}

type FindMedSessionByCustomerIDQuery struct {
	CustomerID        valueobject.CustomerID
	PaginationRequest p.PaginationRequest
}

type FindRecentMedSessionByPetIDQuery struct {
	PetID valueobject.PetID
	Limit int
}

type FindMedSessionByDateRangeQuery struct {
	StartDate         time.Time
	EndDate           time.Time
	PaginationRequest p.PaginationRequest
}

type FindMedSessionByPetAndDateRangeQuery struct {
	PetID     valueobject.PetID
	StartDate time.Time
	EndDate   time.Time
}

type FindMedSessionByDiagnosisQuery struct {
	Diagnosis         string
	PaginationRequest p.PaginationRequest
}

type ExistsMedSessionByIDQuery struct {
	ID valueobject.MedSessionID
}

type ExistsMedSessionByPetAndDateQuery struct {
	PetID valueobject.PetID
	Date  time.Time
}
