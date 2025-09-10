package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListApptsByEmployeeIDQuery struct {
	vetID     valueobject.EmployeeID
	ctx       context.Context
	pageInput page.PageInput
}

func NewListApptsByEmployeeIDQuery(ctx context.Context, vetID uint, pageInput page.PageInput) *ListApptsByEmployeeIDQuery {
	return &ListApptsByEmployeeIDQuery{
		ctx:       ctx,
		vetID:     valueobject.NewEmployeeID(vetID),
		pageInput: pageInput,
	}
}

type ListApptsByEmployeeIDHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewListApptsByEmployeeIDHandler(appointmentRepo repository.AppointmentRepository) cqrs.QueryHandler[(page.Page[[]ApptResponse])] {
	return &ListApptsByEmployeeIDHandler{appointmentRepo: appointmentRepo}
}

func (h *ListApptsByEmployeeIDHandler) Handle(q cqrs.Query) (page.Page[[]ApptResponse], error) {
	query := q.(ListApptsByEmployeeIDQuery)
	appointmentsPage, err := h.appointmentRepo.ListByEmployeeID(query.ctx, query.vetID, query.pageInput)
	if err != nil {
		return page.Page[[]ApptResponse]{}, err
	}

	responses := mapApptsToResponse(appointmentsPage.Data)
	return page.NewPage(responses, appointmentsPage.Metadata), nil
}
