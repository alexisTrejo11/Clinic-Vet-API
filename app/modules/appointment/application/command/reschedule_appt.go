package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
	"time"
)

type RescheduleApptCommand struct {
	appointmentID  valueobject.AppointmentID
	veterinarianID *valueobject.EmployeeID
	datetime       time.Time
}

func NewRescheduleApptCommand(appointIDInt uint, vetID *uint, dateTime time.Time, reason *string) *RescheduleApptCommand {
	var veterinarianID *valueobject.EmployeeID
	if vetID != nil {
		vetIDVal := valueobject.NewEmployeeID(*vetID)
		veterinarianID = &vetIDVal
	}
	return &RescheduleApptCommand{
		appointmentID:  valueobject.NewAppointmentID(appointIDInt),
		veterinarianID: veterinarianID,
		datetime:       dateTime,
	}
}

func (h *apptCommandHandler) RescheduleAppointment(ctx context.Context, cmd RescheduleApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(ctx, cmd.appointmentID)
	if err != nil {
		return *cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.Reschedule(ctx, cmd.datetime); err != nil {
		return *cqrs.FailureResult("failed to reschedule appointment", err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult("failed to save rescheduled appointment", err)
	}

	return *cqrs.SuccessResult(appointment.ID().String(), "appointment rescheduled successfully")
}
