package appointmentCmd

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type CompleteAppointmentCommand struct {
	Id    int     `json:"id" binding:"required"`
	Notes *string `json:"notes,omitempty"`
}

type CompleteAppointmentHandler interface {
	Handle(ctx context.Context, command CompleteAppointmentCommand) shared.CommandResult
}

type completeAppointmentHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewCompleteAppointmentHandler(appointmentRepo appointmentDomain.AppointmentRepository) CompleteAppointmentHandler {
	return &completeAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *completeAppointmentHandler) Handle(ctx context.Context, command CompleteAppointmentCommand) shared.CommandResult {
	// Get existing appointment
	appointment, err := h.appointmentRepo.GetById(ctx, command.Id)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	// Validate appointment can be completed
	if appointment.GetStatus() == appointmentDomain.StatusCompleted {
		return shared.FailureResult("appointment is already completed", nil)
	}

	if appointment.GetStatus() == appointmentDomain.StatusCancelled {
		return shared.FailureResult("cannot complete cancelled appointment", nil)
	}

	// Complete appointment
	appointment.CompleteAppointment()

	// Save updated appointment
	if err := h.appointmentRepo.Save(ctx, appointment); err != nil {
		return shared.FailureResult("failed to save completed appointment", err)
	}

	return shared.SuccesResult(appointment.GetId().String(), "appointment completed successfully")
}
