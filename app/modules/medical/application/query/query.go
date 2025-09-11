// Package query contains query definitions for medical history operations.
package query

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	p "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetMedHistByIDQuery struct {
	ID  valueobject.MedHistoryID
	CTX context.Context
}

type GetMedHistBySpecQuery struct {
	Spec specification.MedicalHistorySpecification
	CTX  context.Context
}

type GetAllMedHistQuery struct {
	PageInput p.PageInput
	CTX       context.Context
}

type GetMedHistByEmployeeIDQuery struct {
	EmployeeID valueobject.EmployeeID
	PageInput  p.PageInput
	CTX        context.Context
}

type GetMedHistByPetIDQuery struct {
	PetID     valueobject.PetID
	PageInput p.PageInput
	CTX       context.Context
}

type GetMedHistByCustomerIDQuery struct {
	CustomerID valueobject.CustomerID
	PageInput  p.PageInput
	CTX        context.Context
}

type GetRecentMedHistByPetIDQuery struct {
	PetID valueobject.PetID
	Limit int
	CTX   context.Context
}

type GetMedHistByDateRangeQuery struct {
	StartDate time.Time
	EndDate   time.Time
	PageInput p.PageInput
	CTX       context.Context
}

type GetMedHistByPetAndDateRangeQuery struct {
	PetID     valueobject.PetID
	StartDate time.Time
	EndDate   time.Time
	CTX       context.Context
}

type GetMedHistByDiagnosisQuery struct {
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
