package command

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type RescheduleAppointmentCommand struct {
	ctx            context.Context
	appointmentID  valueobject.AppointmentID
	veterinarianID *valueobject.VetID
	datetime       time.Time
	reason         *string
}

func NewRescheduleAppointmentCommand(
	ctx context.Context,
	appointIDInt int,
	vetID *int,
	dateTime time.Time,
	reason *string,
) (RescheduleAppointmentCommand, error) {
	errorMessages := make([]string, 0)

	appointmentID, err := valueobject.NewAppointmentID(appointIDInt)
	if err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	var veterinarianID *valueobject.VetID
	if vetID != nil {
		vetIDVal, err := valueobject.NewVetID(*vetID)
		if err != nil {
			errorMessages = append(errorMessages, err.Error())
		}
		veterinarianID = &vetIDVal
	}

	if dateTime.IsZero() {
		errorMessages = append(errorMessages, "invalid date time")
	}

	if len(errorMessages) > 0 {
		return RescheduleAppointmentCommand{}, apperror.MappingError(errorMessages, "http_data_body", "command", "appointment")
	}

	cmd := &RescheduleAppointmentCommand{
		ctx:            ctx,
		appointmentID:  appointmentID,
		veterinarianID: veterinarianID,
		reason:         reason,
		datetime:       dateTime,
	}

	return *cmd, nil
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

	if err := appointment.Reschedule(command.datetime); err != nil {
		return cqrs.FailureResult("failed to reschedule appointment", err)
	}

	if err := h.appointmentRepo.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save rescheduled appointment", err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), "appointment rescheduled successfully")
}
