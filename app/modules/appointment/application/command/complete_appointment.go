package command

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type CompleteAppointmentCommand struct {
	id    valueobject.AppointmentID
	vetID *valueobject.VetID
	notes *string
	ctx   context.Context
}

func NewCompleteAppointmenCommand(ctx context.Context, id int, vetID *int, notes *string) (*CompleteAppointmentCommand, error) {
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

	return &CompleteAppointmentCommand{
		id:    appointmentID,
		notes: notes,
		vetID: vetIDObj,
		ctx:   ctx,
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
	command, valid := cmd.(CompleteAppointmentCommand)
	if !valid {
		return cqrs.FailureResult(ErrInvalidCommandType, nil)
	}

	appointment, err := h.getAppointment(command)
	if err != nil {
		return cqrs.FailureResult(ErrAppointmentNotFound, err)
	}

	if err := appointment.Complete(); err != nil {
		return cqrs.FailureResult("failed to complete appointment", err)
	}

	if err := h.appointmentRepo.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save completed appointment", err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), "appointment completed successfully")
}

func (h *CompleteAppointmentHandler) getAppointment(command CompleteAppointmentCommand) (appointment.Appointment, error) {
	if command.vetID != nil {
		appointment, err := h.appointmentRepo.GetByIDAndVetID(command.ctx, command.id, *command.vetID)
		if err != nil {
			return appointment, err
		}
		return appointment, nil
	}

	appointment, err := h.appointmentRepo.GetByID(command.ctx, command.id)
	if err != nil {
		return appointment, err
	}
	return appointment, nil
}
