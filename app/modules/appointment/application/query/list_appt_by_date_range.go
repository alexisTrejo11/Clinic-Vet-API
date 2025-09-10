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

type ListApptsByDateRangeQuery struct {
	startDate time.Time
	endDate   time.Time
	ctx       context.Context
	pageInput p.PageInput
}

func NewListApptsByDateRangeQuery(startDate, endDate time.Time, pageInput p.PageInput) (ListApptsByDateRangeQuery, error) {
	qry := &ListApptsByDateRangeQuery{
		startDate: startDate,
		endDate:   endDate,
		pageInput: pageInput,
	}

	if startDate.IsZero() {
		return ListApptsByDateRangeQuery{}, apperror.FieldValidationError("startDate", "zero", "startDate can't be zero")
	}

	if endDate.IsZero() {
		return ListApptsByDateRangeQuery{}, apperror.FieldValidationError("endDate", "zero", "endDate can't be zero")
	}

	if startDate.Before(endDate) {
		return ListApptsByDateRangeQuery{}, apperror.FieldValidationError("date-range", "", "startDate can't be before endDate")
	}

	return *qry, nil
}

type ListApptsByDateRangeHandler struct {
	appointmentRepo repo.AppointmentRepository
}

func NewListApptsByDateRangeHandler(appointmentRepo repo.AppointmentRepository) cqrs.QueryHandler[p.Page[[]ApptResponse]] {
	return &ListApptsByDateRangeHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *ListApptsByDateRangeHandler) Handle(q cqrs.Query) (p.Page[[]ApptResponse], error) {
	query, valid := q.(ListApptsByDateRangeQuery)
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
