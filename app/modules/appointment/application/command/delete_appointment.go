package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type DeleteApptCommand struct {
	appointmentID valueobject.AppointmentID
	ctx           context.Context
}

func NewDeleteApptCommand(id uint, ctx context.Context) *DeleteApptCommand {
	return &DeleteApptCommand{
		appointmentID: valueobject.NewAppointmentID(id),
		ctx:           ctx,
	}
}

type DeleteApptHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewDeleteApptHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &DeleteApptHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *DeleteApptHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, ok := cmd.(*DeleteApptCommand)
	if !ok {
		return cqrs.FailureResult(ErrInvalidCommandType, nil)
	}

	appointment, err := h.appointmentRepo.FindByID(command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult(ErrFindingAppt, err)
	}
	if err := h.appointmentRepo.Delete(command.ctx, command.appointmentID); err != nil {
		return cqrs.FailureResult(ErrFailedToDelete, err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), SuccessApptDeleted)
}
