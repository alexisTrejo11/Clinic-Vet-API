// Package command contains all the implementation to handle all the operations that change the state of the appoinment entity
package command

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type CancelAppointmentCommand struct {
	appointmentID valueobject.AppointmentID
	reason        string
	ctx           *context.Context
}

func NewCancelAppointmentCommand(ctx *context.Context, id int, reason string) (*CancelAppointmentCommand, error) {
	if ctx == nil {
		return nil, apperror.FieldValidationError("context", "nil", "context is nil")
	}

	appointmentID, err := valueobject.NewAppointmentID(id)
	if err != nil {
		return nil, apperror.FieldValidationError("appointmentID", strconv.Itoa(id), err.Error())
	}

	return &CancelAppointmentCommand{
		appointmentID: appointmentID,
		reason:        reason,
		ctx:           ctx,
	}, nil
}

type CancelAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewCancelAppointmentHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &CancelAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *CancelAppointmentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(CancelAppointmentCommand)

	appointment, err := h.appointmentRepo.GetByID(*command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult("failed finding appointent", err)
	}

	if err := appointment.Cancel(); err != nil {
		return cqrs.FailureResult("failed to cancel appointment", err)
	}

	if err := h.appointmentRepo.Save(*command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save cancelled appointment", err)
	}

	return cqrs.SuccessResult(appointment.GetID().String(), "appointment cancelled successfully")
}
