package appointmentCmd

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type DeleteAppointmentCommand struct {
	appointmentId int `json:"id" binding:"required"`
}

func NewDeleteAppointmentCommand(appointmentId int) DeleteAppointmentCommand {
	return DeleteAppointmentCommand{
		appointmentId: appointmentId,
	}
}

type DeleteAppointmentHandler interface {
	Handle(ctx context.Context, command DeleteAppointmentCommand) shared.CommandResult
}

type deleteAppointmentHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewDeleteAppointmentHandler(appointmentRepo appointmentDomain.AppointmentRepository) DeleteAppointmentHandler {
	return &deleteAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *deleteAppointmentHandler) Handle(ctx context.Context, command DeleteAppointmentCommand) shared.CommandResult {
	appointment, err := h.appointmentRepo.GetById(ctx, command.appointmentId)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	if appointment.GetStatus() == appointmentDomain.StatusCompleted {
		return shared.FailureResult("cannot delete completed appointment", nil)
	}

	if err := h.appointmentRepo.Delete(command.appointmentId); err != nil {
		return shared.FailureResult("failed to delete appointment", err)
	}

	return shared.SuccessResult(appointment.GetId().String(), "appointment deleted successfully")
}
