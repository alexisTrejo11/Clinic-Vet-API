package query

import (
	"context"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type FindApptStatsQuery struct {
	employeeID *uint
	customerID *uint
	startDate  *time.Time
	endDate    *time.Time
	ctx        context.Context
}

func NewFindApptStatsQuery(employeeID, customerID *uint, startDate, endDate *time.Time) FindApptStatsQuery {
	return FindApptStatsQuery{
		employeeID: employeeID,
		customerID: customerID,
		startDate:  startDate,
		endDate:    endDate,
	}
}

type FindApptStatsHandler struct {
	apptRepo repository.AppointmentRepository
}

func NewFindApptStatsHandler(apptRepo repository.AppointmentRepository) cqrs.QueryHandler[ApptStatsResult] {
	return &FindApptStatsHandler{
		apptRepo: apptRepo,
	}
}

func (h *FindApptStatsHandler) Handle(q cqrs.Query) (ApptStatsResult, error) {
	query := q.(FindApptStatsQuery)

	var appointments []appointment.Appointment
	var err error
	maxPage := page.PageInput{
		PageNumber: 1,
		PageSize:   10000,
	}

	if query.startDate != nil && query.endDate != nil {
		appointmentsPage, dberr := h.apptRepo.FindByDateRange(query.ctx, *query.startDate, *query.endDate, maxPage)
		appointments = appointmentsPage.Items
		err = dberr
	} else {
		appointmentsPage, dberr := h.apptRepo.FindAll(query.ctx, maxPage)
		appointments = appointmentsPage.Items
		err = dberr
	}

	if err != nil {
		return ApptStatsResult{}, err
	}
	// Apply additional filters
	var filteredAppointments []appointment.Appointment
	for _, appointment := range appointments {
		includeAppointment := true

		// Filter by vet ID
		if query.employeeID != nil {
			if appointment.EmployeeID() == nil || appointment.EmployeeID().Value() != *query.employeeID {
				includeAppointment = false
			}
		}

		// Filter by owner ID
		if query.customerID != nil && appointment.CustomerID().Equals(*query.customerID) {
			includeAppointment = false
		}

		if includeAppointment {
			filteredAppointments = append(filteredAppointments, appointment)
		}
	}

	stats := h.calculateStats(filteredAppointments, query)

	return stats, nil
}

func (h *FindApptStatsHandler) calculateStats(appointments []appointment.Appointment, query FindApptStatsQuery) ApptStatsResult {
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
		status := appointment.Status()
		statusBreakdown[status]++

		switch status {
		case enum.AppointmentStatusPending:
			pendingCount++
		case enum.AppointmentStatusCompleted:
			completedCount++
		case enum.AppointmentStatusCancelled:
			cancelledCount++
		case enum.AppointmentStatusNotPresented:
			noShowCount++
		}

		// Count by service
		service := appointment.Service()
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

	return NewApptStatsResult(
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
