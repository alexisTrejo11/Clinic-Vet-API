package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type ConfirmAppointmentCommand struct {
	id    valueobject.AppointmentID
	vetID valueobject.VetID
	ctx   context.Context
}

func NewConfirmAppointmentCommand(ctx context.Context, appointIDInt, vetIDInt int) (ConfirmAppointmentCommand, error) {
	errorMessage := make([]string, 0)

	appointmentID, err := valueobject.NewAppointmentID(appointIDInt)
	if err != nil {
		errorMessage = append(errorMessage, err.Error())
	}

	vetID, err := valueobject.NewVetID(vetIDInt)
	if err != nil {
		errorMessage = append(errorMessage, err.Error())
	}

	if len(errorMessage) > 0 {
		return ConfirmAppointmentCommand{}, apperror.MappingError(errorMessage, "contructor", "command", "confirmAppointmentCommand")
	}

	cmd := ConfirmAppointmentCommand{
		ctx:   ctx,
		vetID: vetID,
		id:    appointmentID,
	}
	return cmd, nil
}

type ConfirmAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewConfirmAppointmentHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &ConfirmAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *ConfirmAppointmentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(ConfirmAppointmentCommand)

	appointment, err := h.appointmentRepo.GetByID(command.ctx, command.id)
	if err != nil {
		return cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.Confirm(command.vetID); err != nil {
		return cqrs.FailureResult("failed to confirm appointment", err)
	}

	if err := h.appointmentRepo.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save confirmed appointment", err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), "appointment confirmed successfully")
}
