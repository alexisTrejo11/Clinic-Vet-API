// Package query contains query definitions for medical history operations.
package query

import (
	"context"
	"time"

	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
	p "clinic-vet-api/app/shared/page"
)

type FindMedHistByIDQuery struct {
	ID  valueobject.MedHistoryID
	CTX context.Context
}

func NewFindMedHistByIDQuery(id uint, ctx context.Context) *FindMedHistByIDQuery {
	return &FindMedHistByIDQuery{
		ID:  valueobject.NewMedHistoryID(id),
		CTX: ctx,
	}
}

type FindMedHistBySpecQuery struct {
	Spec specification.MedicalHistorySpecification
	CTX  context.Context
}

type FindAllMedHistQuery struct {
	PageInput p.PageInput
	CTX       context.Context
}

type FindMedHistByEmployeeIDQuery struct {
	EmployeeID valueobject.EmployeeID
	PageInput  p.PageInput
	CTX        context.Context
}

type FindMedHistByPetIDQuery struct {
	PetID     valueobject.PetID
	PageInput p.PageInput
	CTX       context.Context
}

type FindMedHistByCustomerIDQuery struct {
	CustomerID valueobject.CustomerID
	PageInput  p.PageInput
	CTX        context.Context
}

type FindRecentMedHistByPetIDQuery struct {
	PetID valueobject.PetID
	Limit int
	CTX   context.Context
}

type FindMedHistByDateRangeQuery struct {
	StartDate time.Time
	EndDate   time.Time
	PageInput p.PageInput
	CTX       context.Context
}

type FindMedHistByPetAndDateRangeQuery struct {
	PetID     valueobject.PetID
	StartDate time.Time
	EndDate   time.Time
	CTX       context.Context
}

type FindMedHistByDiagnosisQuery struct {
	Diagnosis string
	PageInput p.PageInput
	CTX       context.Context
}

type ExistsMedHistByIDQuery struct {
	ID  valueobject.MedHistoryID
	CTX context.Context
}

type ExistsMedHistByPetAndDateQuery struct {
	PetID valueobject.PetID
	Date  time.Time
	CTX   context.Context
}
