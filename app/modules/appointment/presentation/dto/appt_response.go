package dto

import (
	"time"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/modules/appointment/application/query"
	"clinic-vet-api/app/shared/page"
)

// AppointmentResponse represents an appointment response
type AppointmentResponse struct {
	ID            uint    `json:"id"`
	PetID         uint    `json:"pet_id"`
	CustomerID    uint    `json:"owner_id"`
	EmployeeID    uint    `json:"vet_id,omitempty"`
	Service       string  `json:"service"`
	Datetime      string  `json:"date_time"`
	ScheduledDate string  `json:"scheduled_date"`
	Status        string  `json:"status"`
	Reason        *string `json:"reason"`
	Notes         *string `json:"notes,omitempty"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

// AppointmentDetail represents a detailed appointment with related entities
type AppointmentDetail struct {
	ID            int                    `json:"id"`
	Pet           *PetSummary            `json:"pet,omitempty"`
	Owner         *OwnerSummary          `json:"owner,omitempty"`
	Veterinarian  *VetSummary            `json:"veterinarian,omitempty"`
	Service       enum.ClinicService     `json:"service"`
	ScheduledDate time.Time              `json:"scheduled_date"`
	Status        enum.AppointmentStatus `json:"status"`
	Notes         *string                `json:"notes,omitempty"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

// PetSummary represents a pet summary for appointment details
type PetSummary struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Species string  `json:"species"`
	Breed   *string `json:"breed,omitempty"`
}

// OwnerSummary represents an owner summary for appointment details
type OwnerSummary struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

// VetSummary represents a veterinarian summary for appointment details
type VetSummary struct {
	ID             int     `json:"id"`
	FullName       string  `json:"full_name"`
	Specialization *string `json:"specialization,omitempty"`
}

// AppointmentListResponse represents a paginated list of appointments
type AppointmentListResponse struct {
	Data     []AppointmentResponse `json:"data"`
	Metadata page.PageMetadata     `json:"metadata"`
}

// AppointmentStats represents appointment statistics
type AppointmentStats struct {
	TotalAppointments     int                            `json:"total_appointments"`
	AppointmentsByStatus  map[enum.AppointmentStatus]int `json:"appointments_by_status"`
	AppointmentsByService map[enum.ClinicService]int     `json:"appointments_by_service"`
	UpcomingAppointments  int                            `json:"upcoming_appointments"`
	OverdueAppointments   int                            `json:"overdue_appointments"`
}

type ResponseMapper struct{}

func NewResponseMapper() *ResponseMapper {
	return &ResponseMapper{}
}

func (m *ResponseMapper) FromResult(result query.ApptResult) *AppointmentResponse {
	return &AppointmentResponse{
		ID:         result.ID.Value(),
		PetID:      result.PetID.Value(),
		CustomerID: result.CustomerID.Value(),
		EmployeeID: result.EmployeeID.Value(),
		Service:    result.Service.DisplayName(),
		Datetime:   result.ScheduledDate.Format(time.RFC822),
		Reason:     result.Reason,
		Notes:      result.Notes,
		Status:     result.Status.DisplayName(),
		CreatedAt:  result.CreatedAt.Format(time.RFC822),
		UpdatedAt:  result.UpdatedAt.Format(time.RFC822),
	}
}

func (m *ResponseMapper) FromResults(results []query.ApptResult) []AppointmentResponse {
	if results == nil {
		return []AppointmentResponse{}
	}

	if len(results) == 0 {
		return []AppointmentResponse{}
	}

	responses := make([]AppointmentResponse, len(results))
	for i, result := range results {
		responses[i] = *m.FromResult(result)
	}
	return responses
}
