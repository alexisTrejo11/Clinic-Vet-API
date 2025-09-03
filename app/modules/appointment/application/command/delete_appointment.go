package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
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
	command := cmd.(DeleteAppointmentCommand)
	appointment, err := h.appointmentRepo.GetByID(command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult("error finding appointment", err)
	}

	if appointment.GetStatus() == enum.StatusCompleted {
		return cqrs.FailureResult("cannot delete completed appointment", nil)
	}

	if err := h.appointmentRepo.Delete(command.appointmentID); err != nil {
		return cqrs.FailureResult("failed to delete appointment", err)
	}

	return cqrs.SuccessResult(appointment.GetID().String(), "appointment deleted successfully")
}
