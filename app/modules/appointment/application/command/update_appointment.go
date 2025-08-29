package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type UpdateAppointmentCommand struct {
	AppointmentID int                     `json:"appoinment_id" binding:"required"`
	VetID         *int                    `json:"vet_id,omitempty"`
	Service       *enum.ClinicService     `json:"service,omitempty"`
	Status        *enum.AppointmentStatus `json:"status,omitempty"`
	Reason        *string                 `json:"reason,omitempty"`
	Notes         *string                 `json:"notes,omitempty"`
	IsEmergency   *bool                   `json:"is_emergency,omitempty"`
}

type UpdateAppointmentHandler interface {
	Handle(ctx context.Context, command UpdateAppointmentCommand) shared.CommandResult
}

type updateAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
	service         *service.AppointmentService
}

func NewUpdateAppointmentHandler(appointmentRepo repository.AppointmentRepository) UpdateAppointmentHandler {
	return &updateAppointmentHandler{
		appointmentRepo: appointmentRepo,
		service:         &service.AppointmentService{},
	}
}

func (h *updateAppointmentHandler) Handle(ctx context.Context, command UpdateAppointmentCommand) shared.CommandResult {
	appointmentID, err := valueobject.NewAppointmentID(command.AppointmentID)
	if err != nil {
		return shared.FailureResult("invalid appointment ID", err)
	}

	appointment, err := h.appointmentRepo.GetByID(ctx, appointmentID)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	if err := h.updateFields(&appointment, command); err != nil {
		return shared.FailureResult("failed to update appointment fields", err)
	}

	if err := h.service.ValidateFields(&appointment); err != nil {
		return shared.FailureResult("appointment validation failed", err)
	}

	if err := h.appointmentRepo.Save(ctx, &appointment); err != nil {
		return shared.FailureResult("failed to update appointment", err)
	}

	return shared.SuccessResult(appointment.GetID().String(), "appointment updated successfully")
}

func (h *updateAppointmentHandler) updateFields(appointment *entity.Appointment, command UpdateAppointmentCommand) error {
	if command.VetID != nil {
		vetID, err := valueobject.NewVetID(*command.VetID)
		if err != nil {
			return err
		}

		appointment.SetVetID(&vetID)
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

	return nil
}
