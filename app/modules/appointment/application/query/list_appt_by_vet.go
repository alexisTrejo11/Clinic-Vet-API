package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListApptByVetQuery struct {
	vetID     valueobject.VetID
	ctx       context.Context
	pageInput page.PageInput
}

func NewListApptByVetQuery(ctx context.Context, vetID uint, pageInput page.PageInput) *ListApptByVetQuery {
	return &ListApptByVetQuery{
		ctx:       ctx,
		vetID:     valueobject.NewVetID(vetID),
		pageInput: pageInput,
	}
}

type ListApptByVetHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewListApptsByVetHandler(appointmentRepo repository.AppointmentRepository) cqrs.QueryHandler[(page.Page[[]ApptResponse])] {
	return &ListApptByVetHandler{appointmentRepo: appointmentRepo}
}

func (h *ListApptByVetHandler) Handle(q cqrs.Query) (page.Page[[]ApptResponse], error) {
	query := q.(ListApptByVetQuery)
	appointmentsPage, err := h.appointmentRepo.ListByVetID(query.ctx, query.vetID, query.pageInput)
	if err != nil {
		return page.Page[[]ApptResponse]{}, err
	}

	responses := mapApptsToResponse(appointmentsPage.Data)
	return page.NewPage(responses, appointmentsPage.Metadata), nil
}
