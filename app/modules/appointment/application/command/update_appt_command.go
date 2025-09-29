package command

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type UpdateApptCommand struct {
	appointmentID valueobject.AppointmentID
	status        *enum.AppointmentStatus
	notes         *string
	service       *enum.ClinicService
}

func NewUpdateApptCommand(appointIDInt uint, status *string, notes *string, service *string) (UpdateApptCommand, error) {
	cmd := &UpdateApptCommand{
		appointmentID: valueobject.NewAppointmentID(appointIDInt),
		service:       enum.ClinicServicePtr(service),
		status:        enum.AppointmentStatusPtr(status),
		notes:         notes,
	}

	if err := cmd.Validate(); err != nil {
		return UpdateApptCommand{}, err
	}

	return *cmd, nil

}

func (c *UpdateApptCommand) Validate() error {
	if c.appointmentID.IsZero() {
		return updateApptCmdErr("appointmentID", "is required")
	}

	if c.status != nil {
		if !c.status.IsValid() {
			return updateApptCmdErr("status", "invalid value")
		}
	}

	if c.service != nil {
		if !c.service.IsValid() {
			return updateApptCmdErr("service", "invalid value")
		}
	}

	return nil
}

func (c *UpdateApptCommand) AppointmentID() valueobject.AppointmentID { return c.appointmentID }
func (c *UpdateApptCommand) Status() *enum.AppointmentStatus          { return c.status }
func (c *UpdateApptCommand) Notes() *string                           { return c.notes }
func (c *UpdateApptCommand) Service() *enum.ClinicService             { return c.service }
