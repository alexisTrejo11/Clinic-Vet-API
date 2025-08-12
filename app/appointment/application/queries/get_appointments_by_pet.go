package appointmentQuery

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByPetQuery struct {
	PetId     int `json:"pet_id"`
	PageInput page.PageData
}

func NewGetAppointmentsByPetQuery(petId, pageNumber, pageSize int) GetAppointmentsByPetQuery {
	return GetAppointmentsByPetQuery{
		PetId: petId,
		PageInput: page.PageData{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		},
	}
}

type GetAppointmentsByPetHandler interface {
	Handle(ctx context.Context, query GetAppointmentsByPetQuery) (*page.Page[[]AppointmentResponse], error)
}

type getAppointmentsByPetHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAppointmentsByPetHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAppointmentsByPetHandler {
	return &getAppointmentsByPetHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentsByPetHandler) Handle(ctx context.Context, query GetAppointmentsByPetQuery) (*page.Page[[]AppointmentResponse], error) {
	appointmentsPage, err := h.appointmentRepo.ListByPetId(ctx, query.PetId, query.PageInput)
	if err != nil {
		return nil, err
	}

	return page.NewPage(
		mapAppointmentsToResponses(appointmentsPage.Data),
		appointmentsPage.Metadata,
	), nil
}
