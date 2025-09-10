package command

import (
	"context"
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type RescheduleApptCommand struct {
	ctx            context.Context
	appointmentID  valueobject.AppointmentID
	veterinarianID *valueobject.EmployeeID
	datetime       time.Time
	reason         *string
}

func NewRescheduleApptCommand(ctx context.Context, appointIDInt uint, vetID *uint, dateTime time.Time, reason *string) *RescheduleApptCommand {
	var veterinarianID *valueobject.EmployeeID
	if vetID != nil {
		vetIDVal := valueobject.NewEmployeeID(*vetID)
		veterinarianID = &vetIDVal
	}

	return &RescheduleApptCommand{
		ctx:            ctx,
		appointmentID:  valueobject.NewAppointmentID(appointIDInt),
		veterinarianID: veterinarianID,
		reason:         reason,
		datetime:       dateTime,
	}
}

type RescheduleApptHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewRescheduleApptHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &RescheduleApptHandler{appointmentRepo: appointmentRepo}
}

func (h *RescheduleApptHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, valid := cmd.(RescheduleApptCommand)
	if !valid {
		return cqrs.FailureResult("invalid command type", errors.New("expected RescheduleApptCommand"))
	}

	appointment, err := h.appointmentRepo.GetByID(command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.Reschedule(command.datetime); err != nil {
		return cqrs.FailureResult("failed to reschedule appointment", err)
	}

	if err := h.appointmentRepo.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save rescheduled appointment", err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), "appointment rescheduled successfully")
}
