package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type FindApptsByPetQuery struct {
	petID     valueobject.PetID
	ctx       context.Context
	pageInput page.PageInput
}

func NewFindApptsByPetQuery(ctx context.Context, vetID uint, pagination page.PageInput) *FindApptsByPetQuery {
	return &FindApptsByPetQuery{
		petID:     valueobject.NewPetID(vetID),
		ctx:       ctx,
		pageInput: pagination,
	}
}

type FindApptsByPetHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewFindApptsByPetHandler(appointmentRepo repository.AppointmentRepository) cqrs.QueryHandler[(page.Page[ApptResponse])] {
	return &FindApptsByPetHandler{appointmentRepo: appointmentRepo}
}

func (h *FindApptsByPetHandler) Handle(q cqrs.Query) (page.Page[ApptResponse], error) {
	query := q.(FindApptsByPetQuery)
	appointmentsPage, err := h.appointmentRepo.FindByPetID(query.ctx, query.petID, query.pageInput)
	if err != nil {
		return page.Page[ApptResponse]{}, err
	}

	return page.NewPage(
		mapApptsToResponse(appointmentsPage.Items),
		appointmentsPage.Metadata,
	), nil
}
