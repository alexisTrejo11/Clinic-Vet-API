package query

import (
	"context"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentStatsQuery struct {
	VetID     *int       `json:"vet_id,omitempty"`
	OwnerID   *int       `json:"owner_id,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

func NewGetAppointmentStatsQuery(vetID, ownerID *int, startDate, endDate *time.Time) GetAppointmentStatsQuery {
	return GetAppointmentStatsQuery{
		VetID:     vetID,
		OwnerID:   ownerID,
		StartDate: startDate,
		EndDate:   endDate,
	}
}

type GetAppointmentStatsHandler interface {
	Handle(ctx context.Context, query GetAppointmentStatsQuery) (*AppointmentStatsResponse, error)
}

type getAppointmentStatsHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewGetAppointmentStatsHandler(appointmentRepo repository.AppointmentRepository) GetAppointmentStatsHandler {
	return &getAppointmentStatsHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentStatsHandler) Handle(ctx context.Context, query GetAppointmentStatsQuery) (*AppointmentStatsResponse, error) {
	var appointments []entity.Appointment
	var err error
	maxPage := page.PageData{
		PageNumber: 1,
		PageSize:   10000,
	}

	// Get appointments based on filters
	if query.StartDate != nil && query.EndDate != nil {
		appointmentsPage, dberr := h.appointmentRepo.ListByDateRange(ctx, *query.StartDate, *query.EndDate, maxPage)
		appointments = appointmentsPage.Data
		err = dberr
	} else {
		appointmentsPage, dberr := h.appointmentRepo.ListAll(ctx, maxPage)
		appointments = appointmentsPage.Data
		err = dberr
	}

	if err != nil {
		return nil, err
	}

	// Apply additional filters
	var filteredAppointments []entity.Appointment
	for _, appointment := range appointments {
		includeAppointment := true

		// Filter by vet ID
		if query.VetID != nil {
			if appointment.GetVetID() == nil || appointment.GetVetID().GetValue() != *query.VetID {
				includeAppointment = false
			}
		}

		// Filter by owner ID
		if query.OwnerID != nil && appointment.GetOwnerID() != *query.OwnerID {
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

func (h *getAppointmentStatsHandler) calculateStats(appointments []entity.Appointment, query GetAppointmentStatsQuery) AppointmentStatsResponse {
	totalAppointments := len(appointments)
	pendingCount := 0
	confirmedCount := 0
	completedCount := 0
	cancelledCount := 0
	noShowCount := 0
	emergencyCount := 0

	statusBreakdown := make(map[enum.AppointmentStatus]int)
	serviceBreakdown := make(map[enum.ClinicService]int)

	for _, appointment := range appointments {
		// Count by status
		status := appointment.GetStatus()
		statusBreakdown[status]++

		switch status {
		case enum.StatusPending:
			pendingCount++
		case enum.StatusCompleted:
			completedCount++
		case enum.StatusCancelled:
			cancelledCount++
		case enum.StatusNotPresented:
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
