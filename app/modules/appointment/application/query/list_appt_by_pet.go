package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListApptsByPetQuery struct {
	petID     valueobject.PetID
	ctx       context.Context
	pageInput page.PageInput
}

func NewListApptsByPetQuery(ctx context.Context, vetID uint, pagination page.PageInput) *ListApptsByPetQuery {
	return &ListApptsByPetQuery{
		petID:     valueobject.NewPetID(vetID),
		ctx:       ctx,
		pageInput: pagination,
	}
}

type ListApptsByPetHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewListApptsByPetHandler(appointmentRepo repository.AppointmentRepository) cqrs.QueryHandler[(page.Page[[]ApptResponse])] {
	return &ListApptsByPetHandler{appointmentRepo: appointmentRepo}
}

func (h *ListApptsByPetHandler) Handle(q cqrs.Query) (page.Page[[]ApptResponse], error) {
	query := q.(ListApptsByPetQuery)
	appointmentsPage, err := h.appointmentRepo.ListByPetID(query.ctx, query.petID, query.pageInput)
	if err != nil {
		return page.Page[[]ApptResponse]{}, err
	}

	return page.NewPage(
		mapApptsToResponse(appointmentsPage.Data),
		appointmentsPage.Metadata,
	), nil
}
