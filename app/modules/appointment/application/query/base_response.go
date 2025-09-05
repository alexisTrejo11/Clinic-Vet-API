// Package query contains all the query implementations for appointment query operations
package query

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
)

type AppointmentResponse struct {
	ID            int                    `json:"id"`
	OwnerID       int                    `json:"owner_id"`
	PetID         int                    `json:"pet_id"`
	VetID         *int                   `json:"vet_id,omitempty"`
	Service       enum.ClinicService     `json:"service"`
	ScheduledDate string                 `json:"scheduled_date"`
	Status        enum.AppointmentStatus `json:"status"`
	Reason        *string                `json:"reason,omitempty"`
	Notes         *string                `json:"notes,omitempty"`
	IsEmergency   bool                   `json:"is_emergency"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
}

func NewAppointmentResponse(appointment *entity.Appointment) AppointmentResponse {
	var vetID *int
	if appointment.GetVetID() != nil {
		id := appointment.GetVetID().GetValue()
		vetID = &id
	}

	return AppointmentResponse{
		ID:            appointment.GetID().GetValue(),
		OwnerID:       appointment.GetOwnerID().GetValue(),
		PetID:         appointment.GetPetID().GetValue(),
		VetID:         vetID,
		Service:       appointment.GetService(),
		ScheduledDate: appointment.GetScheduledDate().Format(time.RFC3339),
		Status:        appointment.GetStatus(),
		CreatedAt:     appointment.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt:     appointment.GetUpdatedAt().Format(time.RFC3339),
	}
}

func mapAppointmentsToResponses(appointments []entity.Appointment) []AppointmentResponse {
	responses := make([]AppointmentResponse, 0, len(appointments))
	for _, appointment := range appointments {
		responses = append(responses, NewAppointmentResponse(&appointment))
	}
	return responses
}

type AppointmentStatsResponse struct {
	TotalAppointments     int                            `json:"total_appointments"`
	PendingAppointments   int                            `json:"pending_appointments"`
	ConfirmedAppointments int                            `json:"confirmed_appointments"`
	CompletedAppointments int                            `json:"completed_appointments"`
	CancelledAppointments int                            `json:"cancelled_appointments"`
	NoShowAppointments    int                            `json:"no_show_appointments"`
	StatusBreakdown       map[enum.AppointmentStatus]int `json:"status_breakdown"`
	ServiceBreakdown      map[enum.ClinicService]int     `json:"service_breakdown"`
	EmergencyCount        int                            `json:"emergency_count"`
	Period                *string                        `json:"period,omitempty"`
}

func NewAppointmentStatsResponse(
	totalAppointments,
	pendingAppointments,
	confirmedAppointments,
	completedAppointments,
	cancelledAppointments,
	noShowAppointments,
	emergencyCount int,
	statusBreakdown map[enum.AppointmentStatus]int,
	serviceBreakdown map[enum.ClinicService]int,
	period *string,
) AppointmentStatsResponse {
	return AppointmentStatsResponse{
		TotalAppointments:     totalAppointments,
		PendingAppointments:   pendingAppointments,
		ConfirmedAppointments: confirmedAppointments,
		CompletedAppointments: completedAppointments,
		CancelledAppointments: cancelledAppointments,
		NoShowAppointments:    noShowAppointments,
		StatusBreakdown:       statusBreakdown,
		ServiceBreakdown:      serviceBreakdown,
		EmergencyCount:        emergencyCount,
		Period:                period,
	}
}
