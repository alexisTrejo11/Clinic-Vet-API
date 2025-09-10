// Package command contains all the implementation to handle all the operations that change the state of the appoinment entity
package command

import (
	"context"

	appt "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type CancelApptCommand struct {
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.EmployeeID
	reason        string
	ctx           context.Context
}

func NewCancelApptCommand(ctx context.Context, id uint, vetID *uint, reason string) *CancelApptCommand {
	var vetIDObj *valueobject.EmployeeID
	if vetID != nil {
		vetIDVal := valueobject.NewEmployeeID(*vetID)
		vetIDObj = &vetIDVal
	}

	return &CancelApptCommand{
		appointmentID: valueobject.NewAppointmentID(id),
		vetID:         vetIDObj,
		reason:        reason,
		ctx:           ctx,
	}
}

type CancelApptHandler struct {
	apptRepository repository.AppointmentRepository
}

func NewCancelApptHandler(apptRepository repository.AppointmentRepository) cqrs.CommandHandler {
	return &CancelApptHandler{
		apptRepository: apptRepository,
	}
}

func (h *CancelApptHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, ok := cmd.(CancelApptCommand)
	if !ok {
		return cqrs.FailureResult(ErrInvalidCommandType, nil)
	}

	appointment, err := h.getAppointment(command)
	if err != nil {
		return cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.Cancel(); err != nil {
		return cqrs.FailureResult(ErrFailedToCancel, err)
	}

	if err := h.apptRepository.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult(ErrUpdateApptFailed, err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), SuccessApptUpdated)
}

func (h *CancelApptHandler) getAppointment(cmd CancelApptCommand) (appt.Appointment, error) {
	if cmd.vetID != nil {
		return h.apptRepository.GetByIDAndEmployeeID(cmd.ctx, cmd.appointmentID, *cmd.vetID)
	}
	return h.apptRepository.GetByID(cmd.ctx, cmd.appointmentID)
}
