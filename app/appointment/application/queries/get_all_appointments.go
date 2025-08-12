package appointmentQuery

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAllAppointmentsQuery struct {
	pageInput page.PageData `json:"page_data"`
}

func NewGetAllAppointmentsQuery(pageNumber, pageSize int) GetAllAppointmentsQuery {
	return GetAllAppointmentsQuery{
		pageInput: page.PageData{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		},
	}
}

type GetAllAppointmentsHandler interface {
	Handle(ctx context.Context, query GetAllAppointmentsQuery) (*page.Page[[]AppointmentResponse], error)
}

type getAllAppointmentsHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAllAppointmentsHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAllAppointmentsHandler {
	return &getAllAppointmentsHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAllAppointmentsHandler) Handle(ctx context.Context, query GetAllAppointmentsQuery) (*page.Page[[]AppointmentResponse], error) {
	appointmentPage, err := h.appointmentRepo.ListAll(ctx, query.pageInput)
	if err != nil {
		return nil, err
	}

	if appointmentPage != nil {
		if len(appointmentPage.Data) == 0 {
			return page.NewPage([]AppointmentResponse{}, appointmentPage.Metadata), nil
		}
	}

	responses := mapAppointmentsToResponses(appointmentPage.Data)
	return page.NewPage(responses, appointmentPage.Metadata), nil
}
