package appointmentQuery

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByVetQuery struct {
	VetId     int `json:"vet_id"`
	pageInput page.PageData
}

func NewGetAppointmentsByVetQuery(vetId, pageNumber, pageSize int) GetAppointmentsByVetQuery {
	return GetAppointmentsByVetQuery{
		VetId: vetId,
		pageInput: page.PageData{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		},
	}
}

type GetAppointmentsByVetHandler interface {
	Handle(ctx context.Context, query GetAppointmentsByVetQuery) (page.Page[[]AppointmentResponse], error)
}

type getAppointmentsByVetHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAppointmentsByVetHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAppointmentsByVetHandler {
	return &getAppointmentsByVetHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentsByVetHandler) Handle(ctx context.Context, query GetAppointmentsByVetQuery) (page.Page[[]AppointmentResponse], error) {
	appointmentsPage, err := h.appointmentRepo.ListByVetId(ctx, query.VetId, query.pageInput)
	if err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	responses := mapAppointmentsToResponses(appointmentsPage.Data)
	return *page.NewPage(responses, appointmentsPage.Metadata), nil
}
