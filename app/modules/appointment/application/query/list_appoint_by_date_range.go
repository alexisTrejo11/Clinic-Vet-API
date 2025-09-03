package query

import (
	"context"
	"time"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListAppointmentsByDateRangeQuery struct {
	startDate time.Time
	endDate   time.Time
	ctx       context.Context
	pageInput page.PageData
}

func NewListAppointmentsByDateRangeQuery(startDate, endDate time.Time, pageInput page.PageData) *ListAppointmentsByDateRangeQuery {
	return &ListAppointmentsByDateRangeQuery{
		startDate: startDate,
		endDate:   endDate,
		pageInput: pageInput,
	}
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
