package appointmentCmd

import (
	"context"
	"time"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type RescheduleAppointmentCommand struct {
	AppointmentId int       `json:"id" binding:"required"`
	DateTime      time.Time `json:"date_time" binding:"required"`
	Reason        *string   `json:"reason,omitempty"`
}

type RescheduleAppointmentHandler interface {
	Handle(ctx context.Context, command RescheduleAppointmentCommand) shared.CommandResult
}

type rescheduleAppointmentHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewRescheduleAppointmentHandler(appointmentRepo appointmentDomain.AppointmentRepository) RescheduleAppointmentHandler {
	return &rescheduleAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *rescheduleAppointmentHandler) Handle(ctx context.Context, command RescheduleAppointmentCommand) shared.CommandResult {
	appointment, err := h.appointmentRepo.GetById(ctx, command.AppointmentId)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	if err := appointment.RescheduleAppointment(command.DateTime); err != nil {
		return shared.FailureResult("failed to reschedule appointment", err)
	}

	if err := h.appointmentRepo.Save(ctx, &appointment); err != nil {
		return shared.FailureResult("failed to save rescheduled appointment", err)
	}

	return shared.SuccessResult(appointment.GetId().String(), "appointment rescheduled successfully")
}
