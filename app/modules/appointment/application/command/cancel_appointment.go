// Package command contains all the implementation to handle all the operations that change the state of the appoinment entity
package command

import (
	"context"
	"strconv"

	appt "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type CancelAppointmentCommand struct {
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.VetID
	reason        string
	ctx           context.Context
}

func NewCancelAppointmentCommand(ctx context.Context, id int, vetID *int, reason string) (*CancelAppointmentCommand, error) {
	appointmentID, err := valueobject.NewAppointmentID(id)
	if err != nil {
		return nil, apperror.FieldValidationError("appointment-ID", strconv.Itoa(id), err.Error())
	}

	var vetIDObj *valueobject.VetID
	if vetID != nil {
		vetIDVal, err := valueobject.NewVetID(*vetID)
		vetIDObj = &vetIDVal
		if err != nil {
			return nil, apperror.FieldValidationError("vet-ID", strconv.Itoa(*vetID), err.Error())
		}
	}

	return &CancelAppointmentCommand{
		appointmentID: appointmentID,
		vetID:         vetIDObj,
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
	command, ok := cmd.(CancelAppointmentCommand)
	if !ok {
		return cqrs.FailureResult(ErrInvalidCommandType, nil)
	}

	appointment, err := h.getAppointment(command)
	if err != nil {
		return cqrs.FailureResult(ErrAppointmentNotFound, err)
	}

	if err := appointment.Cancel(); err != nil {
		return cqrs.FailureResult(ErrFailedToCancel, err)
	}

	if err := h.appointmentRepo.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult(ErrUpdateAppointmentFailed, err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), SuccessAppointmentUpdated)
}

func (h *CancelAppointmentHandler) getAppointment(cmd CancelAppointmentCommand) (appt.Appointment, error) {
	if cmd.vetID != nil {
		return h.appointmentRepo.GetByIDAndVetID(cmd.ctx, cmd.appointmentID, *cmd.vetID)
	}
	return h.appointmentRepo.GetByID(cmd.ctx, cmd.appointmentID)
}
