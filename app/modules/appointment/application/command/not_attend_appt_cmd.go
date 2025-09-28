package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/mapper"
	"context"
)

type NotAttendApptCommand struct {
	appointmentID valueobject.AppointmentID
	employeeID    *valueobject.EmployeeID
}

func NewNotAttendApptCommand(id uint, employeeIDUint *uint) *NotAttendApptCommand {
	return &NotAttendApptCommand{
		appointmentID: valueobject.NewAppointmentID(id),
		employeeID:    mapper.PtrToEmployeeIDPtr(employeeIDUint),
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

	return *cqrs.SuccessResult(SuccessMarkedAsNotPresented)
}
