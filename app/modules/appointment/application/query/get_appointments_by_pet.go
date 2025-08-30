package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByPetQuery struct {
	PetID     valueobject.PetID `json:"pet_id"`
	PageInput page.PageData
}

func NewGetAppointmentsByPetQuery(petID valueobject.PetID, pageNumber, pageSize int) GetAppointmentsByPetQuery {
	return GetAppointmentsByPetQuery{
		PetID: petID,
		PageInput: page.PageData{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		},
	}
}

type GetAppointmentsByPetHandler interface {
	Handle(ctx context.Context, query GetAppointmentsByPetQuery) (page.Page[[]AppointmentResponse], error)
}

type getAppointmentsByPetHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewGetAppointmentsByPetHandler(appointmentRepo repository.AppointmentRepository) GetAppointmentsByPetHandler {
	return &getAppointmentsByPetHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentsByPetHandler) Handle(ctx context.Context, query GetAppointmentsByPetQuery) (page.Page[[]AppointmentResponse], error) {
	appointmentsPage, err := h.appointmentRepo.ListByPetID(ctx, query.PetID, query.PageInput)
	if err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	return page.NewPage(
		mapAppointmentsToResponses(appointmentsPage.Data),
		appointmentsPage.Metadata,
	), nil
}
