package appointment

import (
	"context"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	domainerr "clinic-vet-api/app/modules/core/error"
)

// AppointmentOption defines the functional option type with context
type AppointmentOption func(context.Context, *Appointment) error

func WithService(service enum.ClinicService) AppointmentOption {
	return func(ctx context.Context, a *Appointment) error {
		a.service = service
		return nil
	}
}

func WithScheduledDate(date time.Time) AppointmentOption {
	return func(ctx context.Context, a *Appointment) error {
		a.scheduledDate = date
		return nil
	}
}

func WithStatus(status enum.AppointmentStatus) AppointmentOption {
	return func(ctx context.Context, a *Appointment) error {
		a.status = status
		return nil
	}
}

func WithNotes(notes *string) AppointmentOption {
	return func(ctx context.Context, a *Appointment) error {
		a.notes = notes
		return nil
	}
}

func WithEmployeeID(employeeID *valueobject.EmployeeID) AppointmentOption {
	return func(ctx context.Context, a *Appointment) error {
		a.employeeID = employeeID
		return nil
	}
}

func WithPetID(petID valueobject.PetID) AppointmentOption {
	return func(ctx context.Context, a *Appointment) error {
		a.petID = petID
		return nil
	}
}

func WithCustomerID(customerID valueobject.CustomerID) AppointmentOption {
	return func(ctx context.Context, a *Appointment) error {
		a.customerID = customerID
		return nil
	}
}

func WithTimestamps(createdAt, updatedAt time.Time) AppointmentOption {
	return func(ctx context.Context, a *Appointment) error {
		a.Entity.SetTimeStamps(createdAt, updatedAt)
		return nil
	}
}

// NewAppointment creates a new Appointment with functional options (legacy version)
func NewAppointment(
	id valueobject.AppointmentID,
	petID valueobject.PetID,
	customerID valueobject.CustomerID,
	opts ...AppointmentOption,
) (*Appointment, error) {
	ctx := context.Background()
	appointment, err := NewAppointmentWithContext(ctx, id, petID, customerID, opts...)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

// NewAppointmentWithContext creates a new Appointment with context and functional options
func NewAppointmentWithContext(
	ctx context.Context,
	id valueobject.AppointmentID,
	petID valueobject.PetID,
	customerID valueobject.CustomerID,
	opts ...AppointmentOption,
) (*Appointment, error) {
	appointment := &Appointment{
		Entity:     base.NewEntity(id, time.Now(), time.Now(), 1),
		petID:      petID,
		customerID: customerID,
		status:     enum.AppointmentStatusPending, // Default status
	}

	for _, opt := range opts {
		if err := opt(ctx, appointment); err != nil {
			return nil, err
		}
	}

	return appointment, nil
}

// CreateAppointment creates a new appointment with auto-generated ID
func CreateAppointment(
	ctx context.Context,
	petID valueobject.PetID,
	customerID valueobject.CustomerID,
	opts ...AppointmentOption,
) (*Appointment, error) {
	operation := "CreateAppointment"

	// Validate required fields
	if petID.IsZero() {
		return nil, domainerr.ValidationError(ctx, "appointment", "pet_id", "pet ID is required", operation)
	}
	if customerID.IsZero() {
		return nil, domainerr.ValidationError(ctx, "appointment", "customer_id", "customer ID is required", operation)
	}

	appointment := &Appointment{
		Entity:     base.CreateEntity(valueobject.AppointmentID{}),
		petID:      petID,
		customerID: customerID,
		status:     enum.AppointmentStatusPending, // Default status
	}

	for _, opt := range opts {
		if err := opt(ctx, appointment); err != nil {
			return nil, err
		}
	}

	if err := appointment.validate(ctx); err != nil {
		return nil, err
	}

	return appointment, nil
}

func (a *Appointment) validate(ctx context.Context) error {
	operation := "AppointmentValidation"

	if a.service == "" {
		return domainerr.ValidationError(ctx, "appointment", "service", "service is required", operation)
	}

	if a.scheduledDate.IsZero() {
		return domainerr.AppointmentScheduleDateZeroErr(ctx)
	}

	if err := validateScheduledDate(ctx, a.scheduledDate); err != nil {
		return err
	}

	if a.notes != nil && len(*a.notes) > 1000 {
		return domainerr.ValidationError(ctx, "appointment", "notes", "notes too long (max 1000 characters)", operation)
	}

	if !a.status.IsValid() {
		return domainerr.ValidationError(ctx, "appointment", "status", "invalid appointment status", operation)
	}

	if !a.service.IsValid() {
		return domainerr.ValidationError(ctx, "appointment", "service", "invalid clinic service", operation)
	}

	if a.reason != "" && !a.reason.IsValid() {
		return domainerr.ValidationError(ctx, "appointment", "reason", "invalid visit reason", operation)
	}

	return nil
}

// UpdateScheduledDate updates the appointment date with validation
func (a *Appointment) UpdateScheduledDate(ctx context.Context, newDate time.Time) error {
	if err := validateScheduledDate(ctx, newDate); err != nil {
		return err
	}

	a.scheduledDate = newDate
	return nil
}

// UpdateStatus updates the appointment status with validation
func (a *Appointment) UpdateStatus(ctx context.Context, newStatus enum.AppointmentStatus) error {
	if !newStatus.IsValid() {
		return domainerr.ValidationError(ctx, "appointment", "status", "invalid appointment status", "UpdateStatus")
	}

	a.status = newStatus
	return nil
}

// UpdateNotes updates the appointment notes with validation
func (a *Appointment) UpdateNotes(ctx context.Context, notes *string) error {
	if notes != nil && len(*notes) > 1000 {
		return domainerr.ValidationError(ctx, "appointment", "notes", "notes too long (max 1000 characters)", "UpdateNotes")
	}

	a.notes = notes
	return nil
}
