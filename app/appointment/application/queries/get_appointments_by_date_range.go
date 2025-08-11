package appointmentQuery

import (
	"context"
	"time"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByDateRangeQuery struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	VetId     *int      `json:"vet_id,omitempty"`
	OwnerId   *int      `json:"owner_id,omitempty"`
	PetId     *int      `json:"pet_id,omitempty"`
	Page      int       `json:"page"`
	PageSize  int       `json:"page_size"`
}

func NewGetAppointmentsByDateRangeQuery(startDate, endDate time.Time, vetId, ownerId, petId *int, page, pageSize int) GetAppointmentsByDateRangeQuery {
	return GetAppointmentsByDateRangeQuery{
		StartDate: startDate,
		EndDate:   endDate,
		VetId:     vetId,
		OwnerId:   ownerId,
		PetId:     petId,
		Page:      page,
		PageSize:  pageSize,
	}
}

type GetAppointmentsByDateRangeHandler interface {
	Handle(ctx context.Context, query GetAppointmentsByDateRangeQuery) (*page.Page[[]AppointmentResponse], error)
}

type getAppointmentsByDateRangeHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAppointmentsByDateRangeHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAppointmentsByDateRangeHandler {
	return &getAppointmentsByDateRangeHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentsByDateRangeHandler) Handle(ctx context.Context, query GetAppointmentsByDateRangeQuery) (*page.Page[[]AppointmentResponse], error) {
	appointments, err := h.appointmentRepo.ListByDateRange(ctx, query.StartDate, query.EndDate)
	if err != nil {
		return nil, err
	}

	// Filter by additional criteria if provided
	var filteredAppointments []appointmentDomain.Appointment
	for _, appointment := range appointments {
		includeAppointment := true

		// Filter by vet ID
		if query.VetId != nil {
			if appointment.GetVetId() == nil || appointment.GetVetId().GetValue() != *query.VetId {
				includeAppointment = false
			}
		}

		// Filter by owner ID
		if query.OwnerId != nil && appointment.GetOwnerId() != *query.OwnerId {
			includeAppointment = false
		}

		// Filter by pet ID
		if query.PetId != nil && appointment.GetPetId().GetValue() != *query.PetId {
			includeAppointment = false
		}

		if includeAppointment {
			filteredAppointments = append(filteredAppointments, appointment)
		}
	}

	// Convert to responses
	responses := mapAppointmentsToResponses(filteredAppointments)

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
