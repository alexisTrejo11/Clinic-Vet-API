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
	// Get existing appointment
	appointment, err := h.appointmentRepo.GetById(ctx, command.AppointmentId)
	if err != nil {
		return shared.FailureResult("appointment not found", err)
	}

	// Validate appointment can be rescheduled
	if appointment.GetStatus() == appointmentDomain.StatusCompleted {
		return shared.FailureResult("cannot reschedule completed appointment", nil)
	}

	if appointment.GetStatus() == appointmentDomain.StatusCancelled {
		return shared.FailureResult("cannot reschedule cancelled appointment", nil)
	}

	// Validate new date is in the future
	if command.DateTime.Before(time.Now()) {
		return shared.FailureResult("appointment date must be in the future", nil)
	}

	// Reschedule appointment
	appointment.RescheduleAppointment(command.DateTime)

	// Save updated appointment
	if err := h.appointmentRepo.Save(ctx, appointment); err != nil {
		return shared.FailureResult("failed to save rescheduled appointment", err)
	}

	return shared.SuccesResult(appointment.GetId().String(), "appointment rescheduled successfully")
}
