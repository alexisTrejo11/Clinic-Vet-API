package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type NotAttendApptCommand struct {
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.EmployeeID
}

func NewNotAttendApptCommand(id uint, vetIDUint *uint) *NotAttendApptCommand {
	var vetID *valueobject.EmployeeID
	if vetIDUint != nil {
		vetIDVal := valueobject.NewEmployeeID(*vetIDUint)
		vetID = &vetIDVal
	}

	return &NotAttendApptCommand{

		appointmentID: valueobject.NewAppointmentID(id),
		vetID:         vetID,
	}
}

func (h *apptCommandHandler) MarkAppointmentAsNotAttend(ctx context.Context, cmd NotAttendApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(ctx, cmd.appointmentID)
	if err != nil {
		return *cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.MarkAsNotPresented(ctx); err != nil {
		return *cqrs.FailureResult(ErrMarkAsNotPresentedFailed, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult(ErrSaveApptFailed, err)
	}

	return *cqrs.SuccessResult(appointment.ID().String(), SuccessMarkedAsNotPresented)
}
