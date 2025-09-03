package query

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type SearchAppointmentsQuery struct {
	pageInput      page.PageData
	ctx            context.Context
	searchCriteria map[string]any
}

func NewSearchAppointmentsQuery(pageNumber, pageSize int) *SearchAppointmentsQuery {
	return &SearchAppointmentsQuery{
		pageInput: page.PageData{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		},
	}
}

type SearchAppointmentsHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewSearchAppointmentsHandler(appointmentRepo repository.AppointmentRepository) cqrs.QueryHandler[(page.Page[[]AppointmentResponse])] {
	return &SearchAppointmentsHandler{appointmentRepo: appointmentRepo}
}

func (h *SearchAppointmentsHandler) Handle(q cqrs.Query) (page.Page[[]AppointmentResponse], error) {
	query := q.(SearchAppointmentsQuery)

	appointmentPage, err := h.appointmentRepo.Search(query.ctx, query.pageInput, query.searchCriteria)
	if err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	responses := mapAppointmentsToResponses(appointmentPage.Data)
	return page.NewPage(responses, appointmentPage.Metadata), nil
}
