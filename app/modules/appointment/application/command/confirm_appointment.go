package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type ConfirmAppointmentCommand struct {
	ID    valueobject.AppointmentID `json:"id" binding:"required"`
	VetID *valueobject.VetID        `json:"vet_id,omitempty"`
}

type ConfirmAppointmentHandler interface {
	Handle(ctx context.Context, command ConfirmAppointmentCommand) shared.CommandResult
}

type confirmAppointmentHandlerImpl struct {
	appointmentRepo repository.AppointmentRepository
	service         *service.AppointmentService
}

func NewConfirmAppointmentHandler(appointmentRepo repository.AppointmentRepository) ConfirmAppointmentHandler {
	return &confirmAppointmentHandlerImpl{
		appointmentRepo: appointmentRepo,
		service:         &service.AppointmentService{},
	}
}

func (h *confirmAppointmentHandlerImpl) Handle(ctx context.Context, command ConfirmAppointmentCommand) shared.CommandResult {
	appointment, err := h.appointmentRepo.GetByID(ctx, command.ID)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	if err := h.service.Confirm(command.VetID, &appointment); err != nil {
		return shared.FailureResult("failed to confirm appointment", err)
	}

	if err := h.appointmentRepo.Save(ctx, &appointment); err != nil {
		return shared.FailureResult("failed to save confirmed appointment", err)
	}

	return shared.SuccessResult(appointment.GetID().String(), "appointment confirmed successfully")
}
