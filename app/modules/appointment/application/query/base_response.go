// Package query contains all the query implementations for appointment query operations
package query

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
)

type ApptResponse struct {
	ID            uint                   `json:"id"`
	OwnerID       uint                   `json:"owner_id"`
	PetID         uint                   `json:"pet_id"`
	VetID         *uint                  `json:"vet_id,omitempty"`
	Service       enum.ClinicService     `json:"service"`
	ScheduledDate string                 `json:"scheduled_date"`
	Status        enum.AppointmentStatus `json:"status"`
	Reason        *string                `json:"reason,omitempty"`
	Notes         *string                `json:"notes,omitempty"`
	IsEmergency   bool                   `json:"is_emergency"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
}

func NewApptResponse(appointment *appointment.Appointment) ApptResponse {
	var vetID *uint
	if appointment.VetID() != nil {
		id := appointment.VetID().Value()
		vetID = &id
	}

	return ApptResponse{
		ID:            appointment.ID().Value(),
		OwnerID:       appointment.OwnerID().Value(),
		PetID:         appointment.PetID().Value(),
		VetID:         vetID,
		Service:       appointment.Service(),
		ScheduledDate: appointment.ScheduledDate().Format(time.RFC3339),
		Status:        appointment.Status(),
		CreatedAt:     appointment.CreatedAt().Format(time.RFC3339),
		UpdatedAt:     appointment.UpdatedAt().Format(time.RFC3339),
	}
}

func mapApptsToResponse(appointments []appointment.Appointment) []ApptResponse {
	responses := make([]ApptResponse, 0, len(appointments))
	for _, appointment := range appointments {
		responses = append(responses, NewApptResponse(&appointment))
	}
	return responses
}

type ApptStatsResponse struct {
	TotalAppts       int                            `json:"total_appointments"`
	PendingAppts     int                            `json:"pending_appointments"`
	ConfirmedAppts   int                            `json:"confirmed_appointments"`
	CompletedAppts   int                            `json:"completed_appointments"`
	CancelledAppts   int                            `json:"cancelled_appointments"`
	NoShowAppts      int                            `json:"no_show_appointments"`
	StatusBreakdown  map[enum.AppointmentStatus]int `json:"status_breakdown"`
	ServiceBreakdown map[enum.ClinicService]int     `json:"service_breakdown"`
	EmergencyCount   int                            `json:"emergency_count"`
	Period           *string                        `json:"period,omitempty"`
}

func NewApptStatsResponse(
	totalAppts,
	pendingAppts,
	confirmedAppts,
	completedAppts,
	cancelledAppts,
	noShowAppts,
	emergencyCount int,
	statusBreakdown map[enum.AppointmentStatus]int,
	serviceBreakdown map[enum.ClinicService]int,
	period *string,
) ApptStatsResponse {
	return ApptStatsResponse{
		TotalAppts:       totalAppts,
		PendingAppts:     pendingAppts,
		ConfirmedAppts:   confirmedAppts,
		CompletedAppts:   completedAppts,
		CancelledAppts:   cancelledAppts,
		NoShowAppts:      noShowAppts,
		StatusBreakdown:  statusBreakdown,
		ServiceBreakdown: serviceBreakdown,
		EmergencyCount:   emergencyCount,
		Period:           period,
	}
}
