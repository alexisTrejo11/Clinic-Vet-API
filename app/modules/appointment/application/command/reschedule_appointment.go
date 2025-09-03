package command

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type RescheduleAppointmentCommand struct {
	ctx           context.Context
	appointmentID valueobject.AppointmentID
	datetime      time.Time
	reason        *string
}

type RescheduleAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewRescheduleAppointmentHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &RescheduleAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *RescheduleAppointmentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(RescheduleAppointmentCommand)

	appointment, err := h.appointmentRepo.GetByID(command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.RescheduleAppointment(command.datetime); err != nil {
		return cqrs.FailureResult("failed to reschedule appointment", err)
	}

	if err := h.appointmentRepo.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save rescheduled appointment", err)
	}

	return cqrs.SuccessResult(appointment.GetID().String(), "appointment rescheduled successfully")
}
