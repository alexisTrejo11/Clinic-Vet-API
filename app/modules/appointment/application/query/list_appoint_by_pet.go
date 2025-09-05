package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListAppointmentsByPetQuery struct {
	petID     valueobject.PetID
	ctx       context.Context
	pageInput page.PageInput
}

func NewListAppointmentsByPetQuery(id int, pageInput page.PageInput) (*ListAppointmentsByPetQuery, error) {
	petID, err := valueobject.NewPetID(id)
	if err != nil {
		return nil, err
	}

	return &ListAppointmentsByPetQuery{
		petID:     petID,
		ctx:       context.Background(),
		pageInput: pageInput,
	}, nil
}

type ListAppointmentsByPetHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewListAppointmentsByPetHandler(appointmentRepo repository.AppointmentRepository) cqrs.QueryHandler[(page.Page[[]AppointmentResponse])] {
	return &ListAppointmentsByPetHandler{appointmentRepo: appointmentRepo}
}

func (h *ListAppointmentsByPetHandler) Handle(q cqrs.Query) (page.Page[[]AppointmentResponse], error) {
	query := q.(ListAppointmentsByPetQuery)
	appointmentsPage, err := h.appointmentRepo.ListByPetID(query.ctx, query.petID, query.pageInput)
	if err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	return page.NewPage(
		mapAppointmentsToResponses(appointmentsPage.Data),
		appointmentsPage.Metadata,
	), nil
}
