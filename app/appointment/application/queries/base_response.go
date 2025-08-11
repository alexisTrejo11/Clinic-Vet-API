package appointmentQuery

import (
	"time"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
)

type AppointmentResponse struct {
	Id            int                                 `json:"id"`
	OwnerId       int                                 `json:"owner_id"`
	PetId         int                                 `json:"pet_id"`
	VetId         *int                                `json:"vet_id,omitempty"`
	Service       appointmentDomain.ClinicService     `json:"service"`
	ScheduledDate string                              `json:"scheduled_date"`
	Status        appointmentDomain.AppointmentStatus `json:"status"`
	Reason        *string                             `json:"reason,omitempty"`
	Notes         *string                             `json:"notes,omitempty"`
	IsEmergency   bool                                `json:"is_emergency"`
	CreatedAt     string                              `json:"created_at"`
	UpdatedAt     string                              `json:"updated_at"`
}

func NewAppointmentResponse(appointment *appointmentDomain.Appointment) AppointmentResponse {
	var vetId *int
	if appointment.GetVetId() != nil {
		id := appointment.GetVetId().GetValue()
		vetId = &id
	}

	return AppointmentResponse{
		Id:            appointment.GetId().GetValue(),
		OwnerId:       appointment.GetOwnerId(),
		PetId:         appointment.GetPetId().GetValue(),
		VetId:         vetId,
		Service:       appointment.GetService(),
		ScheduledDate: appointment.GetScheduledDate().Format(time.RFC3339),
		Status:        appointment.GetStatus(),
		CreatedAt:     appointment.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt:     appointment.GetUpdatedAt().Format(time.RFC3339),
	}
}

func mapAppointmentsToResponses(appointments []appointmentDomain.Appointment) []AppointmentResponse {
	var responses []AppointmentResponse
	for _, appointment := range appointments {
		responses = append(responses, NewAppointmentResponse(&appointment))
	}
	return responses
}

type AppointmentStatsResponse struct {
	TotalAppointments     int                                         `json:"total_appointments"`
	PendingAppointments   int                                         `json:"pending_appointments"`
	ConfirmedAppointments int                                         `json:"confirmed_appointments"`
	CompletedAppointments int                                         `json:"completed_appointments"`
	CancelledAppointments int                                         `json:"cancelled_appointments"`
	NoShowAppointments    int                                         `json:"no_show_appointments"`
	StatusBreakdown       map[appointmentDomain.AppointmentStatus]int `json:"status_breakdown"`
	ServiceBreakdown      map[appointmentDomain.ClinicService]int     `json:"service_breakdown"`
	EmergencyCount        int                                         `json:"emergency_count"`
	Period                *string                                     `json:"period,omitempty"`
}

func NewAppointmentStatsResponse(
	totalAppointments,
	pendingAppointments,
	confirmedAppointments,
	completedAppointments,
	cancelledAppointments,
	noShowAppointments,
	emergencyCount int,
	statusBreakdown map[appointmentDomain.AppointmentStatus]int,
	serviceBreakdown map[appointmentDomain.ClinicService]int,
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
