package appointment

import (
	"slices"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

const (
	MinAllowedDaysToSchedule = 3
	MaxAllowedDaysToSchedule = 30
)

func (a *Appointment) Update(notes *string, vetID *valueobject.VetID, service *enum.ClinicService, reason *string) error {
	if notes != nil {
		if len(*notes) > 1000 {
			return domainerr.NewValidationError("appointment", "notes", "notes too long")
		}
		a.notes = notes
	}

	if vetID != nil {
		a.vetID = vetID
	}

	if service != nil {
		if !service.IsValid() {
			return domainerr.NewValidationError("appointment", "service", "invalid clinic service")
		}
		a.service = *service
	}

	if reason != nil {
		parsedReason, err := enum.ParseVisitReason(*reason)
		if err != nil {
			return domainerr.NewValidationError("appointment", "reason", "invalid visit reason")
		}
		a.reason = parsedReason
	}

	a.IncrementVersion()
	return nil
}

func (a *Appointment) Reschedule(newDate time.Time) error {
	if err := validateScheduledDate(newDate); err != nil {
		return err
	}

	if !a.canBeRescheduled() {
		return domainerr.NewBusinessRuleError("appointment", "reschedule", "appointment cannot be rescheduled in current status")
	}

	a.scheduledDate = newDate
	a.status = enum.AppointmentStatusRescheduled
	a.IncrementVersion()

	return nil
}

func (a *Appointment) Cancel() error {
	if !a.canTransitionTo(enum.AppointmentStatusCancelled) {
		return domainerr.NewBusinessRuleError("appointment", "cancel", "appointment cannot be cancelled in current status")
	}

	a.status = enum.AppointmentStatusCancelled
	a.IncrementVersion()
	return nil
}

func (a *Appointment) Complete() error {
	if !a.canTransitionTo(enum.AppointmentStatusCompleted) {
		return domainerr.NewBusinessRuleError("appointment", "complete", "appointment cannot be completed in current status")
	}

	a.status = enum.AppointmentStatusCompleted
	a.IncrementVersion()
	return nil
}

func (a *Appointment) MarkAsNotPresented() error {
	if !a.canTransitionTo(enum.AppointmentStatusNotPresented) {
		return domainerr.NewBusinessRuleError("appointment", "not_presented", "appointment cannot be marked as not presented in current status")
	}

	a.status = enum.AppointmentStatusNotPresented
	a.IncrementVersion()
	return nil
}

func (a *Appointment) Confirm(vetID valueobject.VetID) error {
	if !a.canTransitionTo(enum.AppointmentStatusConfirmed) {
		return domainerr.NewBusinessRuleError("appointment", "confirm", "appointment cannot be confirmed in current status")
	}

	a.vetID = &vetID
	a.status = enum.AppointmentStatusConfirmed
	a.IncrementVersion()
	return nil
}

// Private helper methods for business rules
func (a *Appointment) canBeRescheduled() bool {
	reschedulableStatuses := []enum.AppointmentStatus{
		enum.AppointmentStatusPending,
		enum.AppointmentStatusConfirmed,
		enum.AppointmentStatusRescheduled,
	}

	return slices.Contains(reschedulableStatuses, a.status)
}

func (a *Appointment) canTransitionTo(newStatus enum.AppointmentStatus) bool {
	// Define valid status transitions
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

func (a *Appointment) CanBeDeleted() error {
	if a.Status() == enum.AppointmentStatusCompleted {
		return domainerr.NewBusinessRuleError("appointment", "delete", "completed appointments cannot be deleted")
	}
	return nil
}
