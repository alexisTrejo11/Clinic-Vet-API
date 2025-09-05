package query

import (
	"context"
	"time"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListAppointmentsByDateRangeQuery struct {
	startDate time.Time
	endDate   time.Time
	ctx       context.Context
	pageInput page.PageInput
}

func NewListAppointmentsByDateRangeQuery(startDate, endDate time.Time, pageInput page.PageInput) (ListAppointmentsByDateRangeQuery, error) {
	qry := &ListAppointmentsByDateRangeQuery{
		startDate: startDate,
		endDate:   endDate,
		pageInput: pageInput,
	}

	if startDate.IsZero() {
		return ListAppointmentsByDateRangeQuery{}, apperror.FieldValidationError("startDate", "zero", "startDate can't be zero")
	}

	if endDate.IsZero() {
		return ListAppointmentsByDateRangeQuery{}, apperror.FieldValidationError("endDate", "zero", "endDate can't be zero")
	}

	if startDate.Before(endDate) {
		return ListAppointmentsByDateRangeQuery{}, apperror.FieldValidationError("date-range", "", "startDate can't be before endDate")
	}

	return *qry, nil
}

type ListAppointmentsByDateRangeHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewListAppointmentsByDateRangeHandler(appointmentRepo repository.AppointmentRepository) cqrs.QueryHandler[page.Page[[]AppointmentResponse]] {
	return &ListAppointmentsByDateRangeHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *ListAppointmentsByDateRangeHandler) Handle(q cqrs.Query) (page.Page[[]AppointmentResponse], error) {
	query := q.(ListAppointmentsByDateRangeQuery)
	appointmentsPage, err := h.appointmentRepo.ListByDateRange(query.ctx, query.startDate, query.endDate, query.pageInput)
	if err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	return page.NewPage(
		mapAppointmentsToResponses(appointmentsPage.Data),
		appointmentsPage.Metadata,
	), nil
}
