package command

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type RescheduleAppointmentCommand struct {
	AppointmentID valueobject.AppointmentID `json:"id" binding:"required"`
	DateTime      time.Time                 `json:"date_time" binding:"required"`
	Reason        *string                   `json:"reason,omitempty"`
}

type RescheduleAppointmentHandler interface {
	Handle(ctx context.Context, command RescheduleAppointmentCommand) shared.CommandResult
}

type rescheduleAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
	service         *service.AppointmentService
}

func NewRescheduleAppointmentHandler(appointmentRepo repository.AppointmentRepository) RescheduleAppointmentHandler {
	return &rescheduleAppointmentHandler{
		appointmentRepo: appointmentRepo,
		service:         &service.AppointmentService{},
	}
}

func (h *rescheduleAppointmentHandler) Handle(ctx context.Context, command RescheduleAppointmentCommand) shared.CommandResult {
	appointment, err := h.appointmentRepo.GetByID(ctx, command.AppointmentID)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	if err := h.service.RescheduleAppointment(&appointment, command.DateTime); err != nil {
		return shared.FailureResult("failed to reschedule appointment", err)
	}

	if err := h.appointmentRepo.Save(ctx, &appointment); err != nil {
		return shared.FailureResult("failed to save rescheduled appointment", err)
	}

	return shared.SuccessResult(appointment.GetID().String(), "appointment rescheduled successfully")
}
