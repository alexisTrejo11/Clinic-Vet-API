package appointmentCmd

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type CancelAppointmentCommand struct {
	AppointmentId int    `json:"id" binding:"required"`
	Reason        string `json:"reason" binding:"required"`
}

type CancelAppointmentHandler interface {
	Handle(ctx context.Context, command CancelAppointmentCommand) shared.CommandResult
}

type cancelAppointmentHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewCancelAppointmentHandler(appointmentRepo appointmentDomain.AppointmentRepository) CancelAppointmentHandler {
	return &cancelAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *cancelAppointmentHandler) Handle(ctx context.Context, command CancelAppointmentCommand) shared.CommandResult {
	// Get existing appointment
	appointment, err := h.appointmentRepo.GetById(ctx, command.AppointmentId)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	// Validate appointment can be cancelled
	if appointment.GetStatus() == appointmentDomain.StatusCompleted {
		return shared.FailureResult("cannot cancel completed appointment", nil)
	}

	if appointment.GetStatus() == appointmentDomain.StatusCancelled {
		return shared.FailureResult("appointment is already cancelled", nil)
	}

	// Cancel appointment
	if err := appointment.Cancel(); err != nil {
		return shared.FailureResult("failed to cancel appointment", err)
	}

	// Save updated appointment
	if err := h.appointmentRepo.Save(ctx, appointment); err != nil {
		return shared.FailureResult("failed to save cancelled appointment", err)
	}

	return shared.SuccesResult(appointment.GetId().String(), "appointment cancelled successfully")
}
