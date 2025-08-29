package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type CompleteAppointmentCommand struct {
	ID valueobject.AppointmentID `json:"id" binding:"required"`

	Notes *string `json:"notes,omitempty"`
}

type CompleteAppointmentHandler interface {
	Handle(ctx context.Context, command CompleteAppointmentCommand) shared.CommandResult
}

type completeAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
	service         *service.AppointmentService
}

func NewCompleteAppointmentHandler(appointmentRepo repository.AppointmentRepository) CompleteAppointmentHandler {
	return &completeAppointmentHandler{
		appointmentRepo: appointmentRepo,
		service:         &service.AppointmentService{},
	}
}

func (h *completeAppointmentHandler) Handle(ctx context.Context, command CompleteAppointmentCommand) shared.CommandResult {
	appointment, err := h.appointmentRepo.GetByID(ctx, command.ID)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	if err := h.service.Complete(&appointment); err != nil {
		return shared.FailureResult("failed to complete appointment", err)
	}

	if err := h.appointmentRepo.Save(ctx, &appointment); err != nil {
		return shared.FailureResult("failed to save completed appointment", err)
	}

	return shared.SuccessResult(appointment.GetID().String(), "appointment completed successfully")
}
