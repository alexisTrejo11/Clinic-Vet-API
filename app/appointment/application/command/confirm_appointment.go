package appointmentCmd

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

type ConfirmAppointmentCommand struct {
	Id    int  `json:"id" binding:"required"`
	VetId *int `json:"vet_id,omitempty"`
}

type ConfirmAppointmentHandler interface {
	Handle(ctx context.Context, command ConfirmAppointmentCommand) shared.CommandResult
}

type confirmAppointmentHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewConfirmAppointmentHandler(appointmentRepo appointmentDomain.AppointmentRepository) ConfirmAppointmentHandler {
	return &confirmAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *confirmAppointmentHandler) Handle(ctx context.Context, command ConfirmAppointmentCommand) shared.CommandResult {
	// Get existing appointment
	appointment, err := h.appointmentRepo.GetById(ctx, command.Id)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	// Validate appointment can be confirmed
	if appointment.GetStatus() != appointmentDomain.StatusPending {
		return shared.FailureResult("only pending appointments can be confirmed", nil)
	}

	// Assign vet if provided
	if command.VetId != nil {
		vetId, err := vetDomain.NewVeterinarianId(*command.VetId)
		if err != nil {
			return shared.FailureResult("invalid vet ID", err)
		}
		appointment.SetVetId(&vetId)
	}

	// Confirm appointment by setting status
	appointment.SetStatus(appointmentDomain.StatusPending) // There's no confirmed status in the enum

	// Save updated appointment
	if err := h.appointmentRepo.Save(ctx, appointment); err != nil {
		return shared.FailureResult("failed to save confirmed appointment", err)
	}

	return shared.SuccesResult(appointment.GetId().String(), "appointment confirmed successfully")
}
