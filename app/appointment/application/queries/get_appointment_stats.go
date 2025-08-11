package appointmentQuery

import (
	"context"
	"fmt"
	"time"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
)

type GetAppointmentStatsQuery struct {
	VetId     *int       `json:"vet_id,omitempty"`
	OwnerId   *int       `json:"owner_id,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

func NewGetAppointmentStatsQuery(vetId, ownerId *int, startDate, endDate *time.Time) GetAppointmentStatsQuery {
	return GetAppointmentStatsQuery{
		VetId:     vetId,
		OwnerId:   ownerId,
		StartDate: startDate,
		EndDate:   endDate,
	}
}

type GetAppointmentStatsHandler interface {
	Handle(ctx context.Context, query GetAppointmentStatsQuery) (*AppointmentStatsResponse, error)
}

type getAppointmentStatsHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAppointmentStatsHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAppointmentStatsHandler {
	return &getAppointmentStatsHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentStatsHandler) Handle(ctx context.Context, query GetAppointmentStatsQuery) (*AppointmentStatsResponse, error) {
	var appointments []appointmentDomain.Appointment
	var err error

	// Get appointments based on filters
	if query.StartDate != nil && query.EndDate != nil {
		appointments, err = h.appointmentRepo.ListByDateRange(ctx, *query.StartDate, *query.EndDate)
	} else {
		appointments, err = h.appointmentRepo.ListAll(ctx)
	}

	if err != nil {
		return nil, err
	}

	// Apply additional filters
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

		if includeAppointment {
			filteredAppointments = append(filteredAppointments, appointment)
		}
	}

	// Calculate statistics
	stats := h.calculateStats(filteredAppointments, query)

	return &stats, nil
}

func (h *getAppointmentStatsHandler) calculateStats(appointments []appointmentDomain.Appointment, query GetAppointmentStatsQuery) AppointmentStatsResponse {
	totalAppointments := len(appointments)
	pendingCount := 0
	confirmedCount := 0
	completedCount := 0
	cancelledCount := 0
	noShowCount := 0
	emergencyCount := 0

	statusBreakdown := make(map[appointmentDomain.AppointmentStatus]int)
	serviceBreakdown := make(map[appointmentDomain.ClinicService]int)

	for _, appointment := range appointments {
		// Count by status
		status := appointment.GetStatus()
		statusBreakdown[status]++

		switch status {
		case appointmentDomain.StatusPending:
			pendingCount++
		case appointmentDomain.StatusCompleted:
			completedCount++
		case appointmentDomain.StatusCancelled:
			cancelledCount++
		case appointmentDomain.StatusNotPresented:
			noShowCount++
		}

		// Count by service
		service := appointment.GetService()
		serviceBreakdown[service]++

		// Count emergency appointments
		// Note: IsEmergency is not available in the current domain model
		// You might need to add this field to the Appointment struct
	}

	// Generate period string
	var period *string
	if query.StartDate != nil && query.EndDate != nil {
		periodStr := fmt.Sprintf("%s to %s",
			query.StartDate.Format("2006-01-02"),
			query.EndDate.Format("2006-01-02"))
		period = &periodStr
	}

	return NewAppointmentStatsResponse(
		totalAppointments,
		pendingCount,
		confirmedCount,
		completedCount,
		cancelledCount,
		noShowCount,
		emergencyCount,
		statusBreakdown,
		serviceBreakdown,
		period,
	)
}
