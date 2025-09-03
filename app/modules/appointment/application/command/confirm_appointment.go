package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type ConfirmAppointmentCommand struct {
	id    valueobject.AppointmentID
	vetID *valueobject.VetID
	ctx   context.Context
}

type ConfirmAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewConfirmAppointmentHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &ConfirmAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *ConfirmAppointmentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(ConfirmAppointmentCommand)

	appointment, err := h.appointmentRepo.GetByID(command.ctx, command.id)
	if err != nil {
		return cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.Confirm(command.vetID); err != nil {
		return cqrs.FailureResult("failed to confirm appointment", err)
	}

	if err := h.appointmentRepo.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save confirmed appointment", err)
	}

	return cqrs.SuccessResult(appointment.GetID().String(), "appointment confirmed successfully")
}
