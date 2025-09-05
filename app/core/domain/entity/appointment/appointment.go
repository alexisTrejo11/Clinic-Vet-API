package appointment

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

const (
	MinAllowedDaysToSchedule = 3
	MaxAllowedDaysToSchedule = 30
)

type Appointment struct {
	base.Entity
	service       enum.ClinicService
	scheduledDate time.Time
	status        enum.AppointmentStatus
	reason        string
	notes         *string
	ownerID       valueobject.OwnerID
	vetID         *valueobject.VetID
	petID         valueobject.PetID
}

// AppointmentOption defines the functional option type
type AppointmentOption func(*Appointment) error

// Functional options
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

func WithReason(reason string) AppointmentOption {
	return func(a *Appointment) error {
		if len(reason) > 500 {
			return domainerr.NewValidationError("appointment", "reason", "reason too long")
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
	if id.GetValue() == 0 {
		return nil, domainerr.NewValidationError("appointment", "id", "appointment ID is required")
	}
	if petID.GetValue() == 0 {
		return nil, domainerr.NewValidationError("appointment", "petID", "pet ID is required")
	}
	if ownerID.GetValue() == 0 {
		return nil, domainerr.NewValidationError("appointment", "ownerID", "owner ID is required")
	}

	appointment := &Appointment{
		Entity:  base.NewEntity(id),
		petID:   petID,
		ownerID: ownerID,
		status:  enum.AppointmentStatusPending, // Default status
	}

	// Apply all options
	for _, opt := range opts {
		if err := opt(appointment); err != nil {
			return nil, err
		}
	}

	// Final validation
	if err := appointment.validate(); err != nil {
		return nil, err
	}

	return appointment, nil
}

// Validation helpers
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

// Getters
func (a *Appointment) ID() valueobject.AppointmentID {
	return a.ID()
}

func (a *Appointment) PetID() valueobject.PetID {
	return a.petID
}

func (a *Appointment) OwnerID() valueobject.OwnerID {
	return a.ownerID
}

func (a *Appointment) VetID() *valueobject.VetID {
	return a.vetID
}

func (a *Appointment) Service() enum.ClinicService {
	return a.service
}

func (a *Appointment) ScheduledDate() time.Time {
	return a.scheduledDate
}

func (a *Appointment) Status() enum.AppointmentStatus {
	return a.status
}

func (a *Appointment) Reason() string {
	return a.reason
}

func (a *Appointment) Notes() *string {
	return a.notes
}

// Business logic methods
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
		if len(*reason) > 500 {
			return domainerr.NewValidationError("appointment", "reason", "reason too long")
		}
		a.reason = *reason
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

	for _, status := range reschedulableStatuses {
		if a.status == status {
			return true
		}
	}
	return false
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

	for _, allowedStatus := range allowedTransitions {
		if allowedStatus == newStatus {
			return true
		}
	}

	return false
}

// Additional business logic methods
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
