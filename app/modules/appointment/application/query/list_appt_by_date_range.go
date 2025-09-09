package query

import (
	"context"
	"errors"
	"time"

	repo "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	p "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListApptByDateRangeQuery struct {
	startDate time.Time
	endDate   time.Time
	ctx       context.Context
	pageInput p.PageInput
}

func NewListApptByDateRangeQuery(startDate, endDate time.Time, pageInput p.PageInput) (ListApptByDateRangeQuery, error) {
	qry := &ListApptByDateRangeQuery{
		startDate: startDate,
		endDate:   endDate,
		pageInput: pageInput,
	}

	if startDate.IsZero() {
		return ListApptByDateRangeQuery{}, apperror.FieldValidationError("startDate", "zero", "startDate can't be zero")
	}

	if endDate.IsZero() {
		return ListApptByDateRangeQuery{}, apperror.FieldValidationError("endDate", "zero", "endDate can't be zero")
	}

	if startDate.Before(endDate) {
		return ListApptByDateRangeQuery{}, apperror.FieldValidationError("date-range", "", "startDate can't be before endDate")
	}

	return *qry, nil
}

type ListAppointmentByDateRangeHandler struct {
	appointmentRepo repo.AppointmentRepository
}

func NewListAppointmentsByDateRangeHandler(appointmentRepo repo.AppointmentRepository) cqrs.QueryHandler[p.Page[[]ApptResponse]] {
	return &ListAppointmentByDateRangeHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *ListAppointmentByDateRangeHandler) Handle(q cqrs.Query) (p.Page[[]ApptResponse], error) {
	query, valid := q.(ListApptByDateRangeQuery)
	if !valid {
		return p.Page[[]ApptResponse]{}, errors.New("invalid query type")
	}
	appointmentsPage, err := h.appointmentRepo.ListByDateRange(query.ctx, query.startDate, query.endDate, query.pageInput)
	if err != nil {
		return p.Page[[]ApptResponse]{}, err
	}

	return p.NewPage(
		mapApptsToResponse(appointmentsPage.Data),
		appointmentsPage.Metadata,
	), nil
}
