package appointmentCmd

import (
	"context"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

type UpdateAppointmentCommand struct {
	AppointmentId int                                  `json:"appoinment_id" binding:"required"`
	VetId         *int                                 `json:"vet_id,omitempty"`
	Service       *appointmentDomain.ClinicService     `json:"service,omitempty"`
	Status        *appointmentDomain.AppointmentStatus `json:"status,omitempty"`
	Reason        *string                              `json:"reason,omitempty"`
	Notes         *string                              `json:"notes,omitempty"`
	IsEmergency   *bool                                `json:"is_emergency,omitempty"`
}

type UpdateAppointmentHandler interface {
	Handle(ctx context.Context, command UpdateAppointmentCommand) shared.CommandResult
}

type updateAppointmentHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewUpdateAppointmentHandler(appointmentRepo appointmentDomain.AppointmentRepository) UpdateAppointmentHandler {
	return &updateAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *updateAppointmentHandler) Handle(ctx context.Context, command UpdateAppointmentCommand) shared.CommandResult {
	// Get existing appointment
	appointmentId, err := appointmentDomain.NewAppointmentId(command.AppointmentId)
	if err != nil {
		return shared.FailureResult("invalid appointment ID", err)
	}

	appointment, err := h.appointmentRepo.GetById(ctx, appointmentId.GetValue())
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	// Apply updates
	if command.VetId != nil {
		vetId, err := vetDomain.NewVeterinarianId(*command.VetId)
		if err != nil {
			return shared.FailureResult("invalid vet ID", err)
		}

		appointment.SetVetId(&vetId)
	}

	if command.Service != nil {
		appointment.SetService(*command.Service)
	}

	if command.Status != nil {
		if command.IsEmergency != nil {
			appointment.SetStatus(*command.Status)
		} else {
			appointment.SetStatus(*command.Status)
		}
	}

	if command.Reason != nil {
		appointment.SetReason(*command.Reason)
	}

	if command.Notes != nil {
		appointment.SetNotes(command.Notes)
	}

	// Validate and save
	if err := appointment.ValidateFields(); err != nil {
		return shared.FailureResult("appointment validation failed", err)
	}

	if err := h.appointmentRepo.Save(ctx, appointment); err != nil {
		return shared.FailureResult("failed to update appointment", err)
	}

	return shared.SuccesResult(appointment.GetId().String(), "appointment updated successfully")
}
