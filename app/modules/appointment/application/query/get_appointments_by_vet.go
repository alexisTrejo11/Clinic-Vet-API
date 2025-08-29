package query

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByVetQuery struct {
	VetID     int `json:"vet_id"`
	pageInput page.PageData
}

func NewGetAppointmentsByVetQuery(vetID, pageNumber, pageSize int) GetAppointmentsByVetQuery {
	return GetAppointmentsByVetQuery{
		VetID: vetID,
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
	appointmentRepo repository.AppointmentRepository
}

func NewGetAppointmentsByVetHandler(appointmentRepo repository.AppointmentRepository) GetAppointmentsByVetHandler {
	return &getAppointmentsByVetHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentsByVetHandler) Handle(ctx context.Context, query GetAppointmentsByVetQuery) (page.Page[[]AppointmentResponse], error) {
	appointmentsPage, err := h.appointmentRepo.ListByVetID(ctx, query.VetID, query.pageInput)
	if err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	responses := mapAppointmentsToResponses(appointmentsPage.Data)
	return page.NewPage(responses, appointmentsPage.Metadata), nil
}
