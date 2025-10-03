package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/mapper"
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
	customerID, petID uint,
	employeeID *uint,
	service string,
	dateTime time.Time,
	status string,
	notes *string,
) (CreateApptCommand, error) {
	cmd := CreateApptCommand{
		customerID: valueobject.NewCustomerID(customerID),
		petID:      valueobject.NewPetID(petID),
		employeeID: mapper.PtrToEmployeeIDPtr(employeeID),
		datetime:   dateTime,
		notes:      notes,
		status:     enum.AppointmentStatus(status),
		service:    enum.ClinicService(service),
	}

	if err := cmd.Validate(); err != nil {
		return CreateApptCommand{}, err
	}

	return cmd, nil
}

func (c *CreateApptCommand) Validate() error {
	if !c.service.IsValid() {
		return createApptCommand("service", "invalid value")
	}

	if !c.status.IsValid() {
		return createApptCommand("status", "invalid value")
	}

	if c.customerID.IsZero() {
		return createApptCommand("customerID", "is required")
	}

	if c.petID.IsZero() {
		return createApptCommand("petID", "is required")
	}

	if c.datetime.IsZero() {
		return createApptCommand("datetime", "is required")
	}

	if c.datetime.Before(time.Now()) {
		return createApptCommand("datetime", "cannot be in the past")
	}

	return nil
}

func (cmd *CreateApptCommand) ToEntity() appointment.Appointment {
	return *appointment.NewAppointmentBuilder().
		WithPetID(cmd.PetID()).
		WithCustomerID(cmd.CustomerID()).
		WithEmployeeID(cmd.EmployeeID()).
		WithService(cmd.Service()).
		WithScheduledDate(cmd.DateTime()).
		WithNotes(cmd.Notes()).
		WithStatus(cmd.Status()).
		Build()
}

func (c *CreateApptCommand) CustomerID() valueobject.CustomerID  { return c.customerID }
func (c *CreateApptCommand) PetID() valueobject.PetID            { return c.petID }
func (c *CreateApptCommand) Service() enum.ClinicService         { return c.service }
func (c *CreateApptCommand) DateTime() time.Time                 { return c.datetime }
func (c *CreateApptCommand) Status() enum.AppointmentStatus      { return c.status }
func (c *CreateApptCommand) EmployeeID() *valueobject.EmployeeID { return c.employeeID }
func (c *CreateApptCommand) Notes() *string                      { return c.notes }
