package query

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type FindApptByIDQuery struct {
	appointmentID valueobject.AppointmentID
	ctx           context.Context
}

func NewFindApptByIDQuery(ctx context.Context, id uint) *FindApptByIDQuery {
	return &FindApptByIDQuery{appointmentID: valueobject.NewAppointmentID(id), ctx: ctx}
}

type FindApptByIDAndCustomerIDQuery struct {
	apptID     valueobject.AppointmentID
	customerID valueobject.CustomerID
	ctx        context.Context
}

func NewFindApptByIDAndCustomerIDQuery(ctx context.Context, apptID uint, customerID uint) *FindApptByIDAndCustomerIDQuery {
	return &FindApptByIDAndCustomerIDQuery{ctx: ctx, apptID: valueobject.NewAppointmentID(apptID), customerID: valueobject.NewCustomerID(customerID)}
}

type FindApptByIDAndEmployeeIDQuery struct {
	apptID     valueobject.AppointmentID
	employeeID valueobject.EmployeeID
	ctx        context.Context
}

func NewFindApptByIDAndEmployeeIDQuery(ctx context.Context, apptID uint, employeeID uint) *FindApptByIDAndEmployeeIDQuery {
	return &FindApptByIDAndEmployeeIDQuery{apptID: valueobject.NewAppointmentID(apptID), employeeID: valueobject.NewEmployeeID(employeeID)}
}

type FindApptsByDateRangeQuery struct {
	startDate time.Time
	endDate   time.Time
	ctx       context.Context
	pageInput page.PageInput
}

func NewFindApptsByDateRangeQuery(ctx context.Context, startDate, endDate time.Time, pageInput page.PageInput) (FindApptsByDateRangeQuery, error) {
	qry := &FindApptsByDateRangeQuery{
		startDate: startDate,
		endDate:   endDate,
		pageInput: pageInput,
	}

	if startDate.IsZero() {
		return FindApptsByDateRangeQuery{}, apperror.FieldValidationError("startDate", "zero", "startDate can't be zero")
	}

	if endDate.IsZero() {
		return FindApptsByDateRangeQuery{}, apperror.FieldValidationError("endDate", "zero", "endDate can't be zero")
	}

	if startDate.Before(endDate) {
		return FindApptsByDateRangeQuery{}, apperror.FieldValidationError("date-range", "", "startDate can't be before endDate")
	}

	return *qry, nil
}

type FindApptsByCustomerIDQuery struct {
	ownerID   valueobject.CustomerID
	ctx       context.Context
	pageInput page.PageInput
}

func NewFindApptsByCustomerIDQuery(ctx context.Context, pageInput page.PageInput, customerId uint, petID *uint, status *string) *FindApptsByCustomerIDQuery {
	return &FindApptsByCustomerIDQuery{ownerID: valueobject.NewCustomerID(customerId), pageInput: pageInput, ctx: ctx}
}

type FindApptsByPetQuery struct {
	petID     valueobject.PetID
	ctx       context.Context
	pageInput page.PageInput
}

func NewFindApptsByPetQuery(ctx context.Context, vetID uint, pagination page.PageInput) *FindApptsByPetQuery {
	return &FindApptsByPetQuery{petID: valueobject.NewPetID(vetID), ctx: ctx, pageInput: pagination}
}

type FindApptsByEmployeeIDQuery struct {
	vetID     valueobject.EmployeeID
	ctx       context.Context
	pageInput page.PageInput
}

func NewFindApptsByEmployeeIDQuery(ctx context.Context, vetID uint, pageInput page.PageInput) *FindApptsByEmployeeIDQuery {
	return &FindApptsByEmployeeIDQuery{ctx: ctx, vetID: valueobject.NewEmployeeID(vetID), pageInput: pageInput}
}

type FindApptsBySpecQuery struct {
	spec specification.ApptSearchSpecification
	ctx  context.Context
}

func NewFindApptsBySpecQuery(ctx context.Context, spec specification.ApptSearchSpecification) *FindApptsBySpecQuery {
	return &FindApptsBySpecQuery{ctx: ctx, spec: spec}
}
