package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	apperror "clinic-vet-api/app/shared/error/application"
	"clinic-vet-api/app/shared/mapper"
	"context"
	"errors"
)

type CompleteApptCommand struct {
	id         valueobject.AppointmentID
	employeeID *valueobject.EmployeeID
}

func NewCompleteApptCommand(id uint, employeeID *uint) *CompleteApptCommand {
	return &CompleteApptCommand{
		id:         valueobject.NewAppointmentID(id),
		employeeID: mapper.PtrToEmployeeIDPtr(employeeID),
	}
}

func (h *apptCommandHandler) CompleteAppointment(ctx context.Context, cmd CompleteApptCommand) cqrs.CommandResult {
	if err := cmd.Validate(); err != nil {
		return *cqrs.FailureResult(ErrInvalidCommand, err)
	}

	appointment, err := h.getAppByIDAndEmployeeID(ctx, cmd.id, cmd.employeeID)
	if err != nil {
		return *cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.Complete(ctx); err != nil {
		return *cqrs.FailureResult("failed to complete appointment", err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult("failed to save completed appointment", err)
	}

	return *cqrs.SuccessResult("appointment completed successfully")
}

func (c *CompleteApptCommand) Validate() error {
	if c.id.IsZero() {
		return errors.New("appointment ID required")
	}

	if c.employeeID != nil && c.employeeID.IsZero() {
		return errors.New("employee ID cannot be zero if provided")
	}

	return nil
}

func (h *apptCommandHandler) getAppByIDAndEmployeeID(ctx context.Context, appointID valueobject.AppointmentID, employeeID *valueobject.EmployeeID) (appointment.Appointment, error) {
	if employeeID == nil {
		return h.apptRepository.FindByID(ctx, appointID)
	}

	spec := specification.ApptByID(appointID).And(specification.ApptByEmployee(*employeeID))
	appoint, err := h.apptRepository.Find(ctx, spec)
	if err != nil {
		return appointment.Appointment{}, err
	}

	if len(appoint.Items) == 0 {
		return appointment.Appointment{}, apperror.EntityNotFoundValidationError("Appointment", "id", appointID.String())
	}
	return appoint.Items[0], nil
}
