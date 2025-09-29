package command

import "clinic-vet-api/app/modules/core/domain/valueobject"

type DeleteApptCommand struct {
	appointmentID valueobject.AppointmentID
}

func NewDeleteApptCommand(id uint) (DeleteApptCommand, error) {
	cmd := &DeleteApptCommand{
		appointmentID: valueobject.NewAppointmentID(id),
	}

	if err := cmd.validate(); err != nil {
		return DeleteApptCommand{}, err
	}

	return *cmd, nil
}

func (c *DeleteApptCommand) validate() error {
	if c.appointmentID.IsZero() {
		return deleteApptCmdErr("appointmentID", "is required")
	}
	return nil
}

func (c *DeleteApptCommand) AppointmentID() valueobject.AppointmentID { return c.appointmentID }
