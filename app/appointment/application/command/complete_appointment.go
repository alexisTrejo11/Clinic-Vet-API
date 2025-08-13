package appointmentCmd

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type CompleteAppointmentCommand struct {
	Id int `json:"id" binding:"required"`

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
	appointment, err := h.appointmentRepo.GetById(ctx, command.Id)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	if err := appointment.Complete(); err != nil {
		return shared.FailureResult("failed to complete appointment", err)
	}

	if err := h.appointmentRepo.Save(ctx, &appointment); err != nil {
		return shared.FailureResult("failed to save completed appointment", err)
	}

	return shared.SuccesResult(appointment.GetId().String(), "appointment completed successfully")
}
