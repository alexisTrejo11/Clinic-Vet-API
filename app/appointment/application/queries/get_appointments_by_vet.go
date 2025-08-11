package appointmentQuery

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByVetQuery struct {
	VetId    int `json:"vet_id"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func NewGetAppointmentsByVetQuery(vetId, page, pageSize int) GetAppointmentsByVetQuery {
	return GetAppointmentsByVetQuery{
		VetId:    vetId,
		Page:     page,
		PageSize: pageSize,
	}
}

type GetAppointmentsByVetHandler interface {
	Handle(ctx context.Context, query GetAppointmentsByVetQuery) (*page.Page[[]AppointmentResponse], error)
}

type getAppointmentsByVetHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAppointmentsByVetHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAppointmentsByVetHandler {
	return &getAppointmentsByVetHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentsByVetHandler) Handle(ctx context.Context, query GetAppointmentsByVetQuery) (*page.Page[[]AppointmentResponse], error) {
	// Using ListByVetId with empty query string (we'll filter by vet ID in the repo implementation)
	appointments, err := h.appointmentRepo.ListByVetId(ctx, query.VetId, "")
	if err != nil {
		return nil, err
	}

	// Convert to responses
	responses := mapAppointmentsToResponses(appointments)

	// Calculate pagination manually
	totalCount := len(responses)
	pageData := page.PageData{
		PageSize:   query.PageSize,
		PageNumber: query.Page,
	}

	// Apply pagination to responses
	startIndex := (query.Page - 1) * query.PageSize
	endIndex := startIndex + query.PageSize

	if startIndex > totalCount {
		responses = []AppointmentResponse{}
	} else {
		if endIndex > totalCount {
			endIndex = totalCount
		}
		responses = responses[startIndex:endIndex]
	}

	// Calculate metadata
	metadata := page.GetPageMetadata(totalCount, pageData)

	return page.NewPage(responses, *metadata), nil
}
