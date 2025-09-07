package appointment

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

// AppointmentOption defines the functional option type
type AppointmentOption func(*Appointment) error

func WithService(service enum.ClinicService) AppointmentOption {
	return func(a *Appointment) error {
		if !service.IsValid() {
			return domainerr.NewValidationError("appointment", "service", "invalid clinic service")
		}
		a.service = service
		return nil
	}
}

func WithScheduledDate(date time.Time) AppointmentOption {
	return func(a *Appointment) error {
		if err := validateScheduledDate(date); err != nil {
			return err
		}
		a.scheduledDate = date
		return nil
	}
}

func WithStatus(status enum.AppointmentStatus) AppointmentOption {
	return func(a *Appointment) error {
		if !status.IsValid() {
			return domainerr.NewValidationError("appointment", "status", "invalid appointment status")
		}
		a.status = status
		return nil
	}
}

func WithReason(reason enum.VisitReason) AppointmentOption {
	return func(a *Appointment) error {
		if !reason.IsValid() {
			return domainerr.NewValidationError("appointment", "reason", "invalid visit reason")
		}
		a.reason = reason
		return nil
	}
}

func WithNotes(notes *string) AppointmentOption {
	return func(a *Appointment) error {
		if notes != nil && len(*notes) > 1000 {
			return domainerr.NewValidationError("appointment", "notes", "notes too long")
		}
		a.notes = notes
		return nil
	}
}

func WithVetID(vetID *valueobject.VetID) AppointmentOption {
	return func(a *Appointment) error {
		a.vetID = vetID
		return nil
	}
}

// NewAppointment creates a new Appointment with functional options
func NewAppointment(
	id valueobject.AppointmentID,
	petID valueobject.PetID,
	ownerID valueobject.OwnerID,
	opts ...AppointmentOption,
) (*Appointment, error) {
	if id.IsZero() {
		return nil, domainerr.NewValidationError("appointment", "id", "appointment ID is required")
	}
	if petID.IsZero() {
		return nil, domainerr.NewValidationError("appointment", "pet-ID", "pet ID is required")
	}
	if ownerID.IsZero() {
		return nil, domainerr.NewValidationError("appointment", "owner-ID", "owner ID is required")
	}

	appointment := &Appointment{
		Entity:  base.NewEntity(id),
		petID:   petID,
		ownerID: ownerID,
		status:  enum.AppointmentStatusPending, // Default status
	}

	for _, opt := range opts {
		if err := opt(appointment); err != nil {
			return nil, err
		}
	}

	if err := appointment.validate(); err != nil {
		return nil, err
	}

	return appointment, nil
}

func validateScheduledDate(date time.Time) error {
	if date.IsZero() {
		return domainerr.AppointmentScheduleDateZeroErr()
	}

	now := time.Now()
	if date.Before(now) {
		return domainerr.AppointmentScheduleDateRuleErr("scheduled date cannot be in the past")
	}

	minDate := now.AddDate(0, 0, MinAllowedDaysToSchedule)
	if date.Before(minDate) {
		return domainerr.AppointmentScheduleDateRuleErr("appointments must be scheduled at least 3 days in advance")
	}

	maxDate := now.AddDate(0, 0, MaxAllowedDaysToSchedule)
	if date.After(maxDate) {
		return domainerr.AppointmentScheduleDateRuleErr("appointments cannot be scheduled more than 30 days in advance")
	}

	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return domainerr.AppointmentScheduleDateRuleErr("appointments cannot be scheduled on weekends")
	}

	return nil
}

func (a *Appointment) validate() error {
	if a.service == "" {
		return domainerr.NewValidationError("appointment", "service", "service is required")
	}
	if a.scheduledDate.IsZero() {
		return domainerr.AppointmentScheduleDateZeroErr()
	}
	if err := validateScheduledDate(a.scheduledDate); err != nil {
		return err
	}
	return nil
}
