package query

import (
	"context"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentStatsQuery struct {
	vetID     *int
	ownerID   *int
	startDate *time.Time
	endDate   *time.Time
	ctx       context.Context
}

func NewGetAppointmentStatsQuery(vetID, ownerID *int, startDate, endDate *time.Time) GetAppointmentStatsQuery {
	return GetAppointmentStatsQuery{
		vetID:     vetID,
		ownerID:   ownerID,
		startDate: startDate,
		endDate:   endDate,
	}
}

type GetAppointmentStatsHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewGetAppointmentStatsHandler(appointmentRepo repository.AppointmentRepository) cqrs.QueryHandler[AppointmentStatsResponse] {
	return &GetAppointmentStatsHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *GetAppointmentStatsHandler) Handle(q cqrs.Query) (AppointmentStatsResponse, error) {
	query := q.(GetAppointmentStatsQuery)

	var appointments []entity.Appointment
	var err error
	maxPage := page.PageInput{
		PageNumber: 1,
		PageSize:   10000,
	}

	// Get appointments based on filters
	if query.startDate != nil && query.endDate != nil {
		appointmentsPage, dberr := h.appointmentRepo.ListByDateRange(query.ctx, *query.startDate, *query.endDate, maxPage)
		appointments = appointmentsPage.Data
		err = dberr
	} else {
		appointmentsPage, dberr := h.appointmentRepo.ListAll(query.ctx, maxPage)
		appointments = appointmentsPage.Data
		err = dberr
	}

	if err != nil {
		return AppointmentStatsResponse{}, err
	}
	// Apply additional filters
	var filteredAppointments []entity.Appointment
	for _, appointment := range appointments {
		includeAppointment := true

		// Filter by vet ID
		if query.vetID != nil {
			if appointment.GetVetID() == nil || appointment.GetVetID().GetValue() != *query.vetID {
				includeAppointment = false
			}
		}

		// Filter by owner ID
		if query.ownerID != nil && appointment.GetOwnerID().Equals(*query.ownerID) {
			includeAppointment = false
		}

		if includeAppointment {
			filteredAppointments = append(filteredAppointments, appointment)
		}
	}

	// Calculate statistics
	stats := h.calculateStats(filteredAppointments, query)

	return stats, nil
}

func (h *GetAppointmentStatsHandler) calculateStats(appointments []entity.Appointment, query GetAppointmentStatsQuery) AppointmentStatsResponse {
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
	if query.startDate != nil && query.endDate != nil {
		periodStr := fmt.Sprintf("%s to %s",
			query.startDate.Format("2006-01-02"),
			query.endDate.Format("2006-01-02"))
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
