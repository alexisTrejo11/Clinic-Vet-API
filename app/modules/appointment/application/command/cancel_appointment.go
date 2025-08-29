package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type CancelAppointmentCommand struct {
	AppointmentID valueobject.AppointmentID `json:"id" binding:"required"`
	Reason        string                    `json:"reason" binding:"required"`
}

func NewCancelAppointmentCommand(appointmentID valueobject.AppointmentID, reason string) *CancelAppointmentCommand {
	return &CancelAppointmentCommand{
		AppointmentID: appointmentID,
		Reason:        reason,
	}
}

type CancelAppointmentHandler interface {
	Handle(ctx context.Context, command CancelAppointmentCommand) shared.CommandResult
}

type cancelAppointmentHandler struct {
	appointmentRepo   repository.AppointmentRepository
	appointmenService *service.AppointmentService
}

func NewCancelAppointmentHandler(appointmentRepo repository.AppointmentRepository) CancelAppointmentHandler {
	return &cancelAppointmentHandler{
		appointmentRepo:   appointmentRepo,
		appointmenService: &service.AppointmentService{},
	}
}

func (h *cancelAppointmentHandler) Handle(ctx context.Context, command CancelAppointmentCommand) shared.CommandResult {
	// Get existing appointment
	appointment, err := h.appointmentRepo.GetByID(ctx, command.AppointmentID)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	if err := h.appointmenService.Cancel(&appointment); err != nil {
		return shared.FailureResult("failed to cancel appointment", err)
	}

	if err := h.appointmentRepo.Save(ctx, &appointment); err != nil {
		return shared.FailureResult("failed to save cancelled appointment", err)
	}

	return shared.SuccessResult(appointment.GetID().String(), "appointment cancelled successfully")
}
