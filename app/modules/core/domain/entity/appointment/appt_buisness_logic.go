// appointment.go
package appointment

import (
	"context"
	"slices"
	"time"

	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

const (
	MinAllowedDaysToSchedule = 3
	MaxAllowedDaysToSchedule = 30
)

func (a *Appointment) Update(
	ctx context.Context,
	notes *string,
	employeeID *valueobject.EmployeeID,
	service *enum.ClinicService,
) error {
	operation := "UpdateAppointment"

	if notes != nil {
		if len(*notes) > 1000 {
			return NotesTooLongError(ctx, operation)
		}
		a.notes = notes
	}

	if employeeID != nil {
		a.employeeID = employeeID
	}

	if service != nil {
		if !service.IsValid() {
			return InvalidServiceError(ctx, *service, operation)
		}
		a.service = *service
	}

	a.IncrementVersion()
	return nil
}

func (a *Appointment) Reschedule(ctx context.Context, newDate time.Time) error {
	operation := "RescheduleAppointment"

	if err := validateScheduledDate(ctx, newDate); err != nil {
		return err
	}

	if !a.canBeRescheduled() {
		return CannotRescheduleError(ctx, a.status, operation)
	}

	a.scheduledDate = newDate
	a.status = enum.AppointmentStatusRescheduled
	a.IncrementVersion()

	return nil
}

func (a *Appointment) Cancel(ctx context.Context) error {
	operation := "CancelAppointment"

	if !a.canTransitionTo(enum.AppointmentStatusCancelled) {
		return CannotCancelError(ctx, a.status, operation)
	}

	a.status = enum.AppointmentStatusCancelled
	a.IncrementVersion()
	return nil
}

func (a *Appointment) Complete(ctx context.Context) error {
	operation := "CompleteAppointment"

	if !a.canTransitionTo(enum.AppointmentStatusCompleted) {
		return CannotCompleteError(ctx, a.status, operation)
	}

	a.status = enum.AppointmentStatusCompleted
	a.IncrementVersion()
	return nil
}

func (a *Appointment) MarkAsNotPresented(ctx context.Context) error {
	operation := "MarkAppointmentNotPresented"

	if !a.canTransitionTo(enum.AppointmentStatusNotPresented) {
		return CannotMarkNotPresentedError(ctx, a.status, operation)
	}

	a.status = enum.AppointmentStatusNotPresented
	a.IncrementVersion()
	return nil
}

func (a *Appointment) Confirm(ctx context.Context, employeeID valueobject.EmployeeID) error {
	operation := "ConfirmAppointment"

	if !a.canTransitionTo(enum.AppointmentStatusConfirmed) {
		return CannotConfirmError(ctx, a.status, operation)
	}

	a.employeeID = &employeeID
	a.status = enum.AppointmentStatusConfirmed
	a.IncrementVersion()
	return nil
}

func (a *Appointment) canBeRescheduled() bool {
	reschedulableStatuses := []enum.AppointmentStatus{
		enum.AppointmentStatusPending,
		enum.AppointmentStatusConfirmed,
		enum.AppointmentStatusRescheduled,
	}

	return slices.Contains(reschedulableStatuses, a.status)
}

func (a *Appointment) canTransitionTo(newStatus enum.AppointmentStatus) bool {
	validTransitions := map[enum.AppointmentStatus][]enum.AppointmentStatus{
		enum.AppointmentStatusPending: {
			enum.AppointmentStatusConfirmed,
			enum.AppointmentStatusCancelled,
			enum.AppointmentStatusRescheduled,
		},
		enum.AppointmentStatusConfirmed: {
			enum.AppointmentStatusCompleted,
			enum.AppointmentStatusCancelled,
			enum.AppointmentStatusNotPresented,
			enum.AppointmentStatusRescheduled,
		},
		enum.AppointmentStatusRescheduled: {
			enum.AppointmentStatusConfirmed,
			enum.AppointmentStatusCancelled,
		},
		enum.AppointmentStatusCompleted:    {},
		enum.AppointmentStatusCancelled:    {},
		enum.AppointmentStatusNotPresented: {},
	}

	allowedTransitions, exists := validTransitions[a.status]
	if !exists {
		return false
	}

	return slices.Contains(allowedTransitions, newStatus)
}

func (a *Appointment) IsUpcoming() bool {
	now := time.Now()
	return a.status == enum.AppointmentStatusConfirmed &&
		a.scheduledDate.After(now) &&
		a.scheduledDate.Before(now.Add(24*time.Hour))
}

func (a *Appointment) RequiresReminder() bool {
	now := time.Now()
	reminderTime := a.scheduledDate.Add(-24 * time.Hour)
	return a.status == enum.AppointmentStatusConfirmed &&
		now.After(reminderTime) &&
		a.scheduledDate.After(now)
}

func (a *Appointment) CanBeReviewed() bool {
	return a.status == enum.AppointmentStatusCompleted &&
		a.scheduledDate.After(time.Now().Add(-30*24*time.Hour)) // Within 30 days
}

func (a *Appointment) CanBeDeleted(ctx context.Context) error {
	operation := "DeleteAppointment"

	if a.Status() == enum.AppointmentStatusCompleted {
		return CannotDeleteError(ctx, a.status, operation)
	}
	return nil
}

func validateScheduledDate(ctx context.Context, date time.Time) error {
	operation := "ValidateScheduledDate"

	if date.IsZero() {
		return ScheduledDateInvalidError(ctx, "scheduled date cannot be zero", operation)
	}

	now := time.Now()
	if date.Before(now) {
		return ScheduledDateInvalidError(ctx, "scheduled date cannot be in the past", operation)
	}

	minDate := now.AddDate(0, 0, MinAllowedDaysToSchedule)
	if date.Before(minDate) {
		return ScheduledDateInvalidError(ctx, "appointments must be scheduled at least 3 days in advance", operation)
	}

	maxDate := now.AddDate(0, 0, MaxAllowedDaysToSchedule)
	if date.After(maxDate) {
		return ScheduledDateInvalidError(ctx, "appointments cannot be scheduled more than 30 days in advance", operation)
	}

	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return ScheduledDateInvalidError(ctx, "appointments cannot be scheduled on weekends", operation)
	}

	hour := date.Hour()
	if hour < 9 || hour >= 18 {
		return ScheduledDateInvalidError(ctx, "appointments can only be scheduled during business hours (9 AM to 6 PM)", operation)
	}

	return nil
}
