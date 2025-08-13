package appointmentQuery

import (
	"context"
	"time"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByDateRangeQuery struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	pageInput page.PageData
}

func NewGetAppointmentsByDateRangeQuery(startDate, endDate time.Time, pageNumber, pageSize int) GetAppointmentsByDateRangeQuery {
	return GetAppointmentsByDateRangeQuery{
		StartDate: startDate,
		EndDate:   endDate,
		pageInput: page.PageData{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		},
	}
}

type GetAppointmentsByDateRangeHandler interface {
	Handle(ctx context.Context, query GetAppointmentsByDateRangeQuery) (page.Page[[]AppointmentResponse], error)
}

type getAppointmentsByDateRangeHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAppointmentsByDateRangeHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAppointmentsByDateRangeHandler {
	return &getAppointmentsByDateRangeHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentsByDateRangeHandler) Handle(ctx context.Context, query GetAppointmentsByDateRangeQuery) (page.Page[[]AppointmentResponse], error) {
	appointmentsPage, err := h.appointmentRepo.ListByDateRange(ctx, query.StartDate, query.EndDate, query.pageInput)
	if err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	return *page.NewPage(
		mapAppointmentsToResponses(appointmentsPage.Data),
		appointmentsPage.Metadata,
	), nil
}
