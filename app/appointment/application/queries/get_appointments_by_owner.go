package appointmentQuery

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByOwnerQuery struct {
	OwnerId   int `json:"owner_id"`
	PageInput page.PageData
}

func NewGetAppointmentsByOwnerQuery(ownerId, pageNumber, pageSize int) GetAppointmentsByOwnerQuery {
	return GetAppointmentsByOwnerQuery{
		OwnerId: ownerId,
		PageInput: page.PageData{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		},
	}
}

type GetAppointmentsByOwnerHandler interface {
	Handle(ctx context.Context, query GetAppointmentsByOwnerQuery) (page.Page[[]AppointmentResponse], error)
}

type getAppointmentsByOwnerHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAppointmentsByOwnerHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAppointmentsByOwnerHandler {
	return &getAppointmentsByOwnerHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentsByOwnerHandler) Handle(ctx context.Context, query GetAppointmentsByOwnerQuery) (page.Page[[]AppointmentResponse], error) {
	appointmentsPage, err := h.appointmentRepo.ListByOwnerId(ctx, query.OwnerId, query.PageInput)
	if err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	return *page.NewPage(
		mapAppointmentsToResponses(appointmentsPage.Data),
		appointmentsPage.Metadata,
	), nil
}
