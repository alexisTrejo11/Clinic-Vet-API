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

type FindApptsByDateRangeQuery struct {
	startDate time.Time
	endDate   time.Time
	ctx       context.Context
	pageInput p.PageInput
}

func NewFindApptsByDateRangeQuery(startDate, endDate time.Time, pageInput p.PageInput) (FindApptsByDateRangeQuery, error) {
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

type FindApptsByDateRangeHandler struct {
	appointmentRepo repo.AppointmentRepository
}

func NewFindApptsByDateRangeHandler(appointmentRepo repo.AppointmentRepository) cqrs.QueryHandler[p.Page[ApptResponse]] {
	return &FindApptsByDateRangeHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *FindApptsByDateRangeHandler) Handle(q cqrs.Query) (p.Page[ApptResponse], error) {
	query, valid := q.(FindApptsByDateRangeQuery)
	if !valid {
		return p.Page[ApptResponse]{}, errors.New("invalid query type")
	}
	appointmentsPage, err := h.appointmentRepo.FindByDateRange(query.ctx, query.startDate, query.endDate, query.pageInput)
	if err != nil {
		return p.Page[ApptResponse]{}, err
	}

	return p.NewPage(
		mapApptsToResponse(appointmentsPage.Items),
		appointmentsPage.Metadata,
	), nil
}
