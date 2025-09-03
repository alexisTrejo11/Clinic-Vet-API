package command

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type CompleteAppointmentCommand struct {
	id    valueobject.AppointmentID
	notes *string
	ctx   *context.Context
}

func NewCompleteAppointmenCommand(ctx context.Context, id int, notes *string) (*CompleteAppointmentCommand, error) {
	appointmentID, err := valueobject.NewAppointmentID(id)
	if err != nil {
		return nil, apperror.FieldValidationError("appointmentID", strconv.Itoa(id), err.Error())
	}
	return &CompleteAppointmentCommand{
		id:    appointmentID,
		notes: notes,
	}, nil
}

type CompleteAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewCompleteAppointmentHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &CompleteAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *CompleteAppointmentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(CompleteAppointmentCommand)

	appointment, err := h.appointmentRepo.GetByID(*command.ctx, command.id)
	if err != nil {
		return cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.Complete(); err != nil {
		return cqrs.FailureResult("failed to complete appointment", err)
	}

	if err := h.appointmentRepo.Save(*command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save completed appointment", err)
	}

	return cqrs.SuccessResult(appointment.GetID().String(), "appointment completed successfully")
}
