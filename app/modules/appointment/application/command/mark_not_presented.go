package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type NotAttendAppointmentCommand struct {
	ctx           context.Context
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.VetID
}

func NewNotAttendAppointmentCommand(ctx context.Context, id int, vetID *int) (*NotAttendAppointmentCommand, error) {
	appointmentID, err := valueobject.NewAppointmentID(id)
	if err != nil {
		return nil, err
	}

	var vetIDObj *valueobject.VetID
	if vetID != nil {
		vetIDVal, err := valueobject.NewVetID(*vetID)
		vetIDObj = &vetIDVal
		if err != nil {
			return nil, err
		}
	}

	return &NotAttendAppointmentCommand{
		ctx:           ctx,
		appointmentID: appointmentID,
		vetID:         vetIDObj,
	}, nil
}

type NotAttendAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewNotAttendAppointmentHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &NotAttendAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *NotAttendAppointmentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, ok := cmd.(*NotAttendAppointmentCommand)
	if !ok {
		return cqrs.FailureResult(ErrInvalidCommandType, nil)
	}

	appointment, err := h.getAppointment(command)
	if err != nil {
		return cqrs.FailureResult(ErrAppointmentNotFound, err)
	}

	if err := appointment.MarkAsNotPresented(); err != nil {
		return cqrs.FailureResult(ErrMarkAsNotPresentedFailed, err)
	}

	if err := h.appointmentRepo.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult(ErrSaveAppointmentFailed, err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), SuccessMarkedAsNotPresented)
}

func (h *NotAttendAppointmentHandler) getAppointment(command *NotAttendAppointmentCommand) (appointment.Appointment, error) {
	if command.vetID != nil {
		return h.appointmentRepo.GetByIDAndVetID(command.ctx, command.appointmentID, *command.vetID)
	}

	return h.appointmentRepo.GetByID(command.ctx, command.appointmentID)
}
