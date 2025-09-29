package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/mapper"
	"time"
)

type RescheduleApptCommand struct {
	appointmentID valueobject.AppointmentID
	employeeID    *valueobject.EmployeeID
	datetime      time.Time
}

func NewRescheduleApptCommand(
	appointIDInt uint, employeeID *uint, dateTime time.Time, reason *string,
) (RescheduleApptCommand, error) {
	cmd := &RescheduleApptCommand{
		appointmentID: valueobject.NewAppointmentID(appointIDInt),
		employeeID:    mapper.PtrToEmployeeIDPtr(employeeID),
		datetime:      dateTime,
	}

	if err := cmd.validate(); err != nil {
		return RescheduleApptCommand{}, err
	}

	return *cmd, nil
}

func (c *RescheduleApptCommand) validate() error {
	if c.appointmentID.IsZero() {
		return updateApptCmdErr("appointmentID", "is required")
	}

	if c.employeeID != nil && c.employeeID.IsZero() {
		return updateApptCmdErr("employeeID", "cannot be zero if provided")
	}

	if c.datetime.IsZero() {
		return updateApptCmdErr("datetime", "is required")
	}

	if c.datetime.Before(time.Now()) {
		return updateApptCmdErr("datetime", "must be in the future")
	}

	return nil
}

func (c *RescheduleApptCommand) AppointmentID() valueobject.AppointmentID { return c.appointmentID }
func (c *RescheduleApptCommand) EmployeeID() *valueobject.EmployeeID      { return c.employeeID }
func (c *RescheduleApptCommand) DateTime() time.Time                      { return c.datetime }
