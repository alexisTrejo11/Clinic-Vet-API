package appointmentCmd

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type MarkAsNotPresentedCommand struct {
	Id int `json:"id" binding:"required"`
}

type MarkAsNotPresentedHandler interface {
	Handle(ctx context.Context, command MarkAsNotPresentedCommand) shared.CommandResult
}

type markAsNotPresentedHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewMarkAsNotPresentedHandler(appointmentRepo appointmentDomain.AppointmentRepository) MarkAsNotPresentedHandler {
	return &markAsNotPresentedHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *markAsNotPresentedHandler) Handle(ctx context.Context, command MarkAsNotPresentedCommand) shared.CommandResult {
	// Get existing appointment
	appointment, err := h.appointmentRepo.GetById(ctx, command.Id)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	// Validate appointment can be marked as not presented
	if appointment.GetStatus() == appointmentDomain.StatusCompleted {
		return shared.FailureResult("cannot mark completed appointment as not presented", nil)
	}

	if appointment.GetStatus() == appointmentDomain.StatusCancelled {
		return shared.FailureResult("cannot mark cancelled appointment as not presented", nil)
	}

	if appointment.GetStatus() == appointmentDomain.StatusNotPresented {
		return shared.FailureResult("appointment is already marked as not presented", nil)
	}

	// Mark as not presented
	appointment.MarkAsNotPresented()

	// Save updated appointment
	if err := h.appointmentRepo.Save(ctx, appointment); err != nil {
		return shared.FailureResult("failed to save appointment", err)
	}

	return shared.SuccesResult(appointment.GetId().String(), "appointment marked as not presented successfully")
}
