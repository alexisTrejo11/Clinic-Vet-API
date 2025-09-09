package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type SearchApptsQuery struct {
	spec specification.ApptSearchSpecification
	ctx  context.Context
}

func NewSearchApptsQuery(ctx context.Context, spec specification.ApptSearchSpecification) *SearchApptsQuery {
	return &SearchApptsQuery{
		ctx:  ctx,
		spec: spec,
	}
}

type SearchApptsHandler struct {
	apptRepo repository.AppointmentRepository
}

func NewSearchApptsHandler(apptRepo repository.AppointmentRepository) cqrs.QueryHandler[(page.Page[[]ApptResponse])] {
	return &SearchApptsHandler{apptRepo: apptRepo}
}

func (h *SearchApptsHandler) Handle(q cqrs.Query) (page.Page[[]ApptResponse], error) {
	query := q.(SearchApptsQuery)

	appointmentPage, err := h.apptRepo.Search(query.ctx, query.spec)
	if err != nil {
		return page.Page[[]ApptResponse]{}, err
	}

	responses := mapApptsToResponse(appointmentPage.Data)
	return page.NewPage(responses, appointmentPage.Metadata), nil
}
