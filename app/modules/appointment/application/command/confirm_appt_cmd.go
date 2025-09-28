package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
	"errors"
)

type ConfirmApptCommand struct {
	id         valueobject.AppointmentID
	employeeID valueobject.EmployeeID
}

func NewConfirmAppointmentCommand(appointIDInt, vetIDInt uint) *ConfirmApptCommand {
	return &ConfirmApptCommand{
		id:         valueobject.NewAppointmentID(appointIDInt),
		employeeID: valueobject.NewEmployeeID(vetIDInt),
	}
}

func (h *apptCommandHandler) ConfirmAppointment(ctx context.Context, cmd ConfirmApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(ctx, cmd.id)
	if err != nil {
		return *cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.Confirm(ctx, cmd.employeeID); err != nil {
		return *cqrs.FailureResult("failed to confirm appointment", err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult("failed to save confirmed appointment", err)
	}

	return *cqrs.SuccessResult("appointment confirmed successfully")
}

func (c *ConfirmApptCommand) Validate() error {
	if c.id.IsZero() {
		return errors.New("appointment ID required")
	}

	if c.employeeID.IsZero() {
		return errors.New("employee ID required")
	}

	return nil
}
