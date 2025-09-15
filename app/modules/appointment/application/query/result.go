// Package query contains all the query implementations for appointment query operations
package query

import (
	"time"

	"clinic-vet-api/app/core/domain/entity/appointment"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
)

type ApptResult struct {
	ID            valueobject.AppointmentID
	CustomerID    valueobject.CustomerID
	PetID         valueobject.PetID
	EmployeeID    *valueobject.EmployeeID
	Service       enum.ClinicService
	ScheduledDate time.Time
	Status        enum.AppointmentStatus
	Reason        *string
	Notes         *string
	IsEmergency   bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewApptResult(appointment *appointment.Appointment) ApptResult {
	return ApptResult{
		ID:            appointment.ID(),
		CustomerID:    appointment.CustomerID(),
		PetID:         appointment.PetID(),
		EmployeeID:    appointment.EmployeeID(),
		Service:       appointment.Service(),
		ScheduledDate: appointment.ScheduledDate(),
		Status:        appointment.Status(),
		CreatedAt:     appointment.CreatedAt(),
		UpdatedAt:     appointment.UpdatedAt(),
	}
}

func mapApptsToResult(appointments []appointment.Appointment) []ApptResult {
	responses := make([]ApptResult, 0, len(appointments))
	for _, appointment := range appointments {
		responses = append(responses, NewApptResult(&appointment))
	}
	return responses
}

type ApptStatsResult struct {
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

func NewApptStatsResult(
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
) ApptStatsResult {
	return ApptStatsResult{
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
