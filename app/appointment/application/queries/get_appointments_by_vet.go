package appointmentQuery

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByVetQuery struct {
	VetId     int           `json:"vet_id"`
	PageInput page.PageData `json:"page_data"`
}

func NewGetAppointmentsByVetQuery(vetId, pageNumber, pageSize int) GetAppointmentsByVetQuery {
	return GetAppointmentsByVetQuery{
		VetId: vetId,
		PageInput: page.PageData{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		},
	}
}

type GetAppointmentsByVetHandler interface {
	Handle(ctx context.Context, query GetAppointmentsByVetQuery) (*page.Page[[]AppointmentResponse], error)
}

type getAppointmentsByVetHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAppointmentsByVetHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAppointmentsByVetHandler {
	return &getAppointmentsByVetHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentsByVetHandler) Handle(ctx context.Context, query GetAppointmentsByVetQuery) (*page.Page[[]AppointmentResponse], error) {
	appointmentsPage, err := h.appointmentRepo.ListByVetId(ctx, query.VetId, query.PageInput)
	if err != nil {
		return nil, err
	}

	responses := mapAppointmentsToResponses(appointmentsPage.Data)
	return page.NewPage(responses, appointmentsPage.Metadata), nil
}
