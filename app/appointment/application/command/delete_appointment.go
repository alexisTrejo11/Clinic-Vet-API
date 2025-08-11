package appointmentCmd

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type DeleteAppointmentCommand struct {
	AppointmentId int `json:"id" binding:"required"`
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
	// Check if appointment exists
	appointment, err := h.appointmentRepo.GetById(ctx, command.AppointmentId)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	// Validate if appointment can be deleted
	if appointment.GetStatus() == appointmentDomain.StatusCompleted {
		return shared.FailureResult("cannot delete completed appointment", nil)
	}

	if err := h.appointmentRepo.Delete(command.AppointmentId); err != nil {
		return shared.FailureResult("failed to delete appointment", err)
	}

	return shared.SuccesResult(appointment.GetId().String(), "appointment deleted successfully")
}
