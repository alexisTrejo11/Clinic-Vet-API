package appointmentQuery

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByOwnerQuery struct {
	OwnerId  int `json:"owner_id"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func NewGetAppointmentsByOwnerQuery(ownerId, page, pageSize int) GetAppointmentsByOwnerQuery {
	return GetAppointmentsByOwnerQuery{
		OwnerId:  ownerId,
		Page:     page,
		PageSize: pageSize,
	}
}

type GetAppointmentsByOwnerHandler interface {
	Handle(ctx context.Context, query GetAppointmentsByOwnerQuery) (*page.Page[[]AppointmentResponse], error)
}

type getAppointmentsByOwnerHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAppointmentsByOwnerHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAppointmentsByOwnerHandler {
	return &getAppointmentsByOwnerHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentsByOwnerHandler) Handle(ctx context.Context, query GetAppointmentsByOwnerQuery) (*page.Page[[]AppointmentResponse], error) {
	appointments, err := h.appointmentRepo.ListByOwnerId(ctx, query.OwnerId)
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
