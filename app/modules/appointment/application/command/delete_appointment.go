package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type DeleteAppointmentCommand struct {
	appointmentID valueobject.AppointmentID `json:"id" binding:"required"`
}

func NewDeleteAppointmentCommand(appointmentID valueobject.AppointmentID) DeleteAppointmentCommand {
	return DeleteAppointmentCommand{
		appointmentID: appointmentID,
	}
}

type DeleteAppointmentHandler interface {
	Handle(ctx context.Context, command DeleteAppointmentCommand) shared.CommandResult
}

type deleteAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewDeleteAppointmentHandler(appointmentRepo repository.AppointmentRepository) DeleteAppointmentHandler {
	return &deleteAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *deleteAppointmentHandler) Handle(ctx context.Context, command DeleteAppointmentCommand) shared.CommandResult {
	appointment, err := h.appointmentRepo.GetByID(ctx, command.appointmentID)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	if appointment.GetStatus() == enum.StatusCompleted {
		return shared.FailureResult("cannot delete completed appointment", nil)
	}

	if err := h.appointmentRepo.Delete(command.appointmentID); err != nil {
		return shared.FailureResult("failed to delete appointment", err)
	}

	return shared.SuccessResult(appointment.GetID().String(), "appointment deleted successfully")
}
