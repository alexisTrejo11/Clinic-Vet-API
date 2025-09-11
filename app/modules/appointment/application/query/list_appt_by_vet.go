package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type FindApptsByEmployeeIDQuery struct {
	vetID     valueobject.EmployeeID
	ctx       context.Context
	pageInput page.PageInput
}

func NewFindApptsByEmployeeIDQuery(ctx context.Context, vetID uint, pageInput page.PageInput) *FindApptsByEmployeeIDQuery {
	return &FindApptsByEmployeeIDQuery{
		ctx:       ctx,
		vetID:     valueobject.NewEmployeeID(vetID),
		pageInput: pageInput,
	}
}

type FindApptsByEmployeeIDHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewFindApptsByEmployeeIDHandler(appointmentRepo repository.AppointmentRepository) cqrs.QueryHandler[(page.Page[ApptResponse])] {
	return &FindApptsByEmployeeIDHandler{appointmentRepo: appointmentRepo}
}

func (h *FindApptsByEmployeeIDHandler) Handle(q cqrs.Query) (page.Page[ApptResponse], error) {
	query := q.(FindApptsByEmployeeIDQuery)
	appointmentsPage, err := h.appointmentRepo.FindByEmployeeID(query.ctx, query.vetID, query.pageInput)
	if err != nil {
		return page.Page[ApptResponse]{}, err
	}

	responses := mapApptsToResponse(appointmentsPage.Items)
	return page.NewPage(responses, appointmentsPage.Metadata), nil
}
