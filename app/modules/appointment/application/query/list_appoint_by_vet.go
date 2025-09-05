package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListAppointmentsByVetQuery struct {
	vetID     valueobject.VetID
	ctx       context.Context
	pageInput page.PageInput
}

func NewListAppointmentsByVetQuery(id int, ctx context.Context, pageInput page.PageInput) (*ListAppointmentsByVetQuery, error) {
	vetID, err := valueobject.NewVetID(id)
	if err != nil {
		return nil, err
	}

	return &ListAppointmentsByVetQuery{
		ctx:       ctx,
		vetID:     vetID,
		pageInput: pageInput,
	}, nil
}

type ListAppointmentsByVetHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewListAppointmentsByVetHandler(appointmentRepo repository.AppointmentRepository) cqrs.QueryHandler[(page.Page[[]AppointmentResponse])] {
	return &ListAppointmentsByVetHandler{appointmentRepo: appointmentRepo}
}

func (h *ListAppointmentsByVetHandler) Handle(q cqrs.Query) (page.Page[[]AppointmentResponse], error) {
	query := q.(ListAppointmentsByVetQuery)
	appointmentsPage, err := h.appointmentRepo.ListByVetID(query.ctx, query.vetID, query.pageInput)
	if err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	responses := mapAppointmentsToResponses(appointmentsPage.Data)
	return page.NewPage(responses, appointmentsPage.Metadata), nil
}
