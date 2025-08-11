package appointmentQuery

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAllAppointmentsQuery struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func NewGetAllAppointmentsQuery(page, pageSize int) GetAllAppointmentsQuery {
	return GetAllAppointmentsQuery{
		Page:     page,
		PageSize: pageSize,
	}
}

type GetAllAppointmentsHandler interface {
	Handle(ctx context.Context, query GetAllAppointmentsQuery) (*page.Page[[]AppointmentResponse], error)
}

type getAllAppointmentsHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAllAppointmentsHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAllAppointmentsHandler {
	return &getAllAppointmentsHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAllAppointmentsHandler) Handle(ctx context.Context, query GetAllAppointmentsQuery) (*page.Page[[]AppointmentResponse], error) {
	// Get all appointments (in a real implementation, this would use pagination from the repository)
	appointments, err := h.appointmentRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	// Convert to responses
	responses := mapAppointmentsToResponses(appointments)

	// Calculate pagination manually (in a real implementation, this would be done in the repository)
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
