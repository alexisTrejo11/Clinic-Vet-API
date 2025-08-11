package appointmentQuery

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByPetQuery struct {
	PetId    int `json:"pet_id"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func NewGetAppointmentsByPetQuery(petId, page, pageSize int) GetAppointmentsByPetQuery {
	return GetAppointmentsByPetQuery{
		PetId:    petId,
		Page:     page,
		PageSize: pageSize,
	}
}

type GetAppointmentsByPetHandler interface {
	Handle(ctx context.Context, query GetAppointmentsByPetQuery) (*page.Page[[]AppointmentResponse], error)
}

type getAppointmentsByPetHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAppointmentsByPetHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAppointmentsByPetHandler {
	return &getAppointmentsByPetHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentsByPetHandler) Handle(ctx context.Context, query GetAppointmentsByPetQuery) (*page.Page[[]AppointmentResponse], error) {
	// Using ListByPetId with empty query string (we'll filter by pet ID in the repo implementation)
	appointments, err := h.appointmentRepo.ListByPetId(ctx, query.PetId, "")
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
