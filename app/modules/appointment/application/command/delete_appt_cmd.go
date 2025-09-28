package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	apperror "clinic-vet-api/app/shared/error/application"
	"context"
	"errors"
)

type DeleteApptCommand struct {
	appointmentID valueobject.AppointmentID
}

func NewDeleteApptCommand(id uint) *DeleteApptCommand {
	return &DeleteApptCommand{
		appointmentID: valueobject.NewAppointmentID(id),
	}
}

func (c *DeleteApptCommand) Validate() error {
	if c.appointmentID.IsZero() {
		return errors.New("appointment ID required")
	}
	return nil
}

func (h *apptCommandHandler) DeleteAppointment(ctx context.Context, cmd DeleteApptCommand) cqrs.CommandResult {
	if err := cmd.Validate(); err != nil {
		return *cqrs.FailureResult(ErrInvalidCommand, err)
	}

	if exists, err := h.apptRepository.ExistsByID(ctx, cmd.appointmentID); err != nil {
		return *cqrs.FailureResult(ErrFailedToCheckExistence, err)
	} else if !exists {
		return *cqrs.FailureResult(apperror.EntityNotFoundValidationError("Appointment", "id", cmd.appointmentID.String()).Error(), nil)
	}

	if err := h.apptRepository.Delete(ctx, cmd.appointmentID, false); err != nil {
		return *cqrs.FailureResult(ErrFailedToDelete, err)
	}

	return *cqrs.SuccessResult(SuccessApptDeleted)
}
