package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type MarkAsNotPresentedCommand struct {
	ID valueobject.AppointmentID `json:"id" binding:"required"`
}

type MarkAsNotPresentedHandler interface {
	Handle(ctx context.Context, command MarkAsNotPresentedCommand) shared.CommandResult
}

type markAsNotPresentedHandler struct {
	appointmentRepo repository.AppointmentRepository
	service         *service.AppointmentService
}

func NewMarkAsNotPresentedHandler(appointmentRepo repository.AppointmentRepository) MarkAsNotPresentedHandler {
	return &markAsNotPresentedHandler{
		appointmentRepo: appointmentRepo,
		service:         &service.AppointmentService{},
	}
}

func (h *markAsNotPresentedHandler) Handle(ctx context.Context, command MarkAsNotPresentedCommand) shared.CommandResult {
	appointment, err := h.appointmentRepo.GetByID(ctx, command.ID)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	if err := h.service.NotPresented(&appointment); err != nil {
		return shared.FailureResult("failed to mark appointment as not presented", err)
	}

	if err := h.appointmentRepo.Save(ctx, &appointment); err != nil {
		return shared.FailureResult("failed to save appointment", err)
	}

	return shared.SuccessResult(appointment.GetID().String(), "appointment marked as not presented successfully")
}
