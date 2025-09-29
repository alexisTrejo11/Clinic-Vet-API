package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
	"clinic-vet-api/app/shared/mapper"
)

type CancelApptCommand struct {
	appointmentID valueobject.AppointmentID
	employeeID    *valueobject.EmployeeID
	reason        string
}

func NewCancelApptCommand(id uint, employeeID *uint, reason string) (*CancelApptCommand, error) {
	cmd := CancelApptCommand{
		appointmentID: valueobject.NewAppointmentID(id),
		employeeID:    mapper.PtrToEmployeeIDPtr(employeeID),
		reason:        reason,
	}

	if err := cmd.validate(); err != nil {
		return nil, err
	}

	return &cmd, nil
}

func (c *CancelApptCommand) validate() error {
	if c.appointmentID.IsZero() {
		return apperror.CommandDataValidationError("appointment_id", "is required", "CancelApptCommand")
	}

	if c.reason == "" {
		return apperror.CommandDataValidationError("reason", "is required", "CancelApptCommand")
	}

	return nil
}

func (c *CancelApptCommand) EmployeeID() *valueobject.EmployeeID      { return c.employeeID }
func (c *CancelApptCommand) AppointmentID() valueobject.AppointmentID { return c.appointmentID }
func (c *CancelApptCommand) Reason() string                           { return c.reason }
