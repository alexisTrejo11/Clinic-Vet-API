package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/mapper"
	"context"
	"errors"
)

type CancelApptCommand struct {
	appointmentID valueobject.AppointmentID
	employeeID    *valueobject.EmployeeID
	reason        string
}

func NewCancelApptCommand(id uint, employeeID *uint, reason string) *CancelApptCommand {
	return &CancelApptCommand{
		appointmentID: valueobject.NewAppointmentID(id),
		employeeID:    mapper.PtrToEmployeeIDPtr(employeeID),
		reason:        reason,
	}
}

func (h *apptCommandHandler) CancelAppointment(ctx context.Context, cmd CancelApptCommand) cqrs.CommandResult {
	if err := cmd.Validate(); err != nil {
		return *cqrs.FailureResult(ErrInvalidCommand, err)
	}

	appointment, err := h.getAppByIDAndEmployeeID(ctx, cmd.appointmentID, cmd.employeeID)
	if err != nil {
		return *cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.Cancel(ctx); err != nil {
		return *cqrs.FailureResult(ErrFailedToCancel, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult(ErrUpdateApptFailed, err)
	}

	return *cqrs.SuccessResult(SuccessApptUpdated)
}

func (c *CancelApptCommand) Validate() error {
	if c.appointmentID.IsZero() {
		return errors.New("appointment ID required")
	}

	if c.reason == "" {
		return errors.New("cancellation reason required")
	}

	return nil
}
