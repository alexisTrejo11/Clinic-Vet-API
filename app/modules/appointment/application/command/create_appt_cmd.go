package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/mapper"
	"context"
	"errors"
	"time"
)

type CreateApptCommand struct {
	customerID valueobject.CustomerID
	petID      valueobject.PetID
	service    enum.ClinicService
	datetime   time.Time
	status     enum.AppointmentStatus
	employeeID *valueobject.EmployeeID
	notes      *string
}

func NewCreateApptCommand(
	customerID, petID uint, employeeID *uint, service string,
	dateTime time.Time, status string, notes *string,
) *CreateApptCommand {
	return &CreateApptCommand{
		customerID: valueobject.NewCustomerID(customerID),
		petID:      valueobject.NewPetID(petID),
		employeeID: mapper.PtrToEmployeeIDPtr(employeeID),
		datetime:   dateTime,
		notes:      notes,
		status:     enum.AppointmentStatus(status),
		service:    enum.ClinicService(service),
	}
}

func (h *apptCommandHandler) CreateAppointment(ctx context.Context, cmd CreateApptCommand) cqrs.CommandResult {
	if err := cmd.validate(); err != nil {
		return *cqrs.FailureResult("invalid command", err)
	}

	appointment := commandToEntity(cmd)
	if err := appointment.ValidatePersistence(ctx); err != nil {
		return *cqrs.FailureResult("appointment validation failed", err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult("failed to save appointment", err)
	}

	return *cqrs.SuccessCreateResult(appointment.ID().String(), "appointment created successfully")
}

func (c *CreateApptCommand) validate() error {
	_, err := enum.ParseClinicService(c.service.String())
	if err != nil {
		return err
	}

	_, err = enum.ParseAppointmentStatus(c.status.String())
	if err != nil {
		return err
	}

	if c.customerID.IsZero() {
		return errors.New("customer ID required")
	}

	if c.petID.IsZero() {
		return errors.New("pet ID required")
	}

	if c.datetime.IsZero() {
		return errors.New("appointment date and time required")
	}

	return nil
}

func commandToEntity(command CreateApptCommand) appointment.Appointment {
	appointment := appointment.NewAppointmentBuilder().
		WithPetID(command.petID).
		WithCustomerID(command.customerID).
		WithEmployeeID(command.employeeID).
		WithService(command.service).
		WithScheduledDate(command.datetime).
		WithNotes(command.notes).
		WithStatus(command.status).
		Build()

	return *appointment
}
