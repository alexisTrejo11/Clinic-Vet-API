package appointmentCmd

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

type ConfirmAppointmentCommand struct {
	Id    int              `json:"id" binding:"required"`
	VetId *vetDomain.VetId `json:"vet_id,omitempty"`
}

type ConfirmAppointmentHandler interface {
	Handle(ctx context.Context, command ConfirmAppointmentCommand) shared.CommandResult
}

type confirmAppointmentHandlerImpl struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewConfirmAppointmentHandler(appointmentRepo appointmentDomain.AppointmentRepository) ConfirmAppointmentHandler {
	return &confirmAppointmentHandlerImpl{
		appointmentRepo: appointmentRepo,
	}
}

func (h *confirmAppointmentHandlerImpl) Handle(ctx context.Context, command ConfirmAppointmentCommand) shared.CommandResult {
	appointment, err := h.appointmentRepo.GetById(ctx, command.Id)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	if err := appointment.Confirm(command.VetId); err != nil {
		return shared.FailureResult("failed to confirm appointment", err)
	}

	if err := h.appointmentRepo.Save(ctx, &appointment); err != nil {
		return shared.FailureResult("failed to save confirmed appointment", err)
	}

	return shared.SuccesResult(appointment.GetId().String(), "appointment confirmed successfully")
}
