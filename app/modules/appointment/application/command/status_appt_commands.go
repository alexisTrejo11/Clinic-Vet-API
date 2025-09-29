package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/mapper"
)

type NotAttendApptCommand struct {
	appointmentID valueobject.AppointmentID
	employeeID    *valueobject.EmployeeID
}

func NewNotAttendApptCommand(id uint, employeeIDUint *uint) (NotAttendApptCommand, error) {
	cmd := &NotAttendApptCommand{
		appointmentID: valueobject.NewAppointmentID(id),
		employeeID:    mapper.PtrToEmployeeIDPtr(employeeIDUint),
	}
	if err := cmd.Validate(); err != nil {
		return NotAttendApptCommand{}, err
	}
	return *cmd, nil
}

func (c *NotAttendApptCommand) Validate() error {
	if c.appointmentID.IsZero() {
		return notAttendApptCmdErr("appointmentID", "is required")
	}

	if c.employeeID != nil && c.employeeID.IsZero() {
		return notAttendApptCmdErr("employeeID", "cannot be zero if provided")
	}

	return nil
}

type CompleteApptCommand struct {
	id         valueobject.AppointmentID
	employeeID *valueobject.EmployeeID
}

func NewCompleteApptCommand(id uint, employeeID *uint) (CompleteApptCommand, error) {
	cmd := &CompleteApptCommand{
		id:         valueobject.NewAppointmentID(id),
		employeeID: mapper.PtrToEmployeeIDPtr(employeeID),
	}
	if err := cmd.validate(); err != nil {
		return CompleteApptCommand{}, err
	}
	return *cmd, nil
}

func (c *CompleteApptCommand) validate() error {
	if c.id.IsZero() {
		return completeApptCmdErr("appointmentID", "is required")
	}

	if c.employeeID != nil && c.employeeID.IsZero() {
		return completeApptCmdErr("employeeID", "cannot be zero if provided")
	}

	return nil
}

type ConfirmApptCommand struct {
	id         valueobject.AppointmentID
	employeeID valueobject.EmployeeID
}

func NewConfirmAppointmentCommand(appointIDInt, vetIDInt uint) (ConfirmApptCommand, error) {
	cmd := &ConfirmApptCommand{
		id:         valueobject.NewAppointmentID(appointIDInt),
		employeeID: valueobject.NewEmployeeID(vetIDInt),
	}

	if err := cmd.validate(); err != nil {
		return ConfirmApptCommand{}, err
	}

	return *cmd, nil
}

func (c *ConfirmApptCommand) validate() error {
	if c.id.IsZero() {
		return confirmApptCmdErr("appointmentID", "is required")
	}

	if c.employeeID.IsZero() {
		return confirmApptCmdErr("employeeID", "is required")
	}

	return nil
}

func (c *NotAttendApptCommand) AppointmentID() valueobject.AppointmentID { return c.appointmentID }
func (c *NotAttendApptCommand) EmployeeID() *valueobject.EmployeeID      { return c.employeeID }

func (c *CompleteApptCommand) ID() valueobject.AppointmentID       { return c.id }
func (c *CompleteApptCommand) EmployeeID() *valueobject.EmployeeID { return c.employeeID }

func (c *ConfirmApptCommand) ID() valueobject.AppointmentID      { return c.id }
func (c *ConfirmApptCommand) EmployeeID() valueobject.EmployeeID { return c.employeeID }
