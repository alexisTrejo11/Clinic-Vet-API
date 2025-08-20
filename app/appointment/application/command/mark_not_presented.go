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
	appointment, err := h.appointmentRepo.GetById(ctx, command.Id)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	if err := appointment.NotPresented(); err != nil {
		return shared.FailureResult("failed to mark appointment as not presented", err)
	}

	if err := h.appointmentRepo.Save(ctx, &appointment); err != nil {
		return shared.FailureResult("failed to save appointment", err)
	}

	return shared.SuccessResult(appointment.GetId().String(), "appointment marked as not presented successfully")
}
