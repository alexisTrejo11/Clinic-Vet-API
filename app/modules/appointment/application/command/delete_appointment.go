package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type DeleteAppointmentCommand struct {
	appointmentID valueobject.AppointmentID
	ctx           context.Context
}

func NewDeleteAppointmentCommand(id int, ctx context.Context) (*DeleteAppointmentCommand, error) {
	appointmentID, err := valueobject.NewAppointmentID(id)
	if err != nil {
		return nil, err
	}

	return &DeleteAppointmentCommand{
		appointmentID: appointmentID,
		ctx:           ctx,
	}, nil
}

type DeleteAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewDeleteAppointmentHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &DeleteAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *DeleteAppointmentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, ok := cmd.(*DeleteAppointmentCommand)
	if !ok {
		return cqrs.FailureResult(ErrInvalidCommandType, nil)
	}

	appointment, err := h.appointmentRepo.GetByID(command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult(ErrFindingAppointment, err)
	}
	if err := h.appointmentRepo.Delete(command.appointmentID); err != nil {
		return cqrs.FailureResult(ErrFailedToDelete, err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), SuccessAppointmentDeleted)
}
