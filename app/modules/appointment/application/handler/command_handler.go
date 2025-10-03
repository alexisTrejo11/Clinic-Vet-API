package handler

import (
	"context"

	c "clinic-vet-api/app/modules/appointment/application/command"
	"clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	apperror "clinic-vet-api/app/shared/error/application"
)

type ApptCommandHandler struct {
	apptRepository repository.AppointmentRepository
}

func NewAppointmentCommandHandler(apptRepository repository.AppointmentRepository) *ApptCommandHandler {
	return &ApptCommandHandler{apptRepository: apptRepository}
}

func (h *ApptCommandHandler) HandleRequestByCustomer(ctx context.Context, cmd c.RequestApptByCustomerCommand) cqrs.CommandResult {
	appointment := appointment.CreateCustomerRequest(
		cmd.PetID(), cmd.CustomerID(), cmd.Service(), cmd.RequestedDate(), cmd.Notes(),
	)

	if err := appointment.ValidatePersistence(ctx); err != nil {
		return cqrs.FailureResult(BusinessRuleFailed, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return cqrs.FailureResult(SaveApptFailed, err)
	}

	return cqrs.SuccessResult(SuccessApptCreated)
}

func (h *ApptCommandHandler) HandleCancel(ctx context.Context, cmd c.CancelApptCommand) cqrs.CommandResult {
	appointment, err := h.getAppByIDAndEmployeeID(ctx, cmd.AppointmentID(), cmd.EmployeeID())
	if err != nil {
		return cqrs.FailureResult(ApptNotFound, err)
	}

	if err := appointment.Cancel(ctx); err != nil {
		return cqrs.FailureResult(FailedToCancel, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return cqrs.FailureResult(UpdateApptFailed, err)
	}

	return cqrs.SuccessResult(SuccessApptUpdated)
}

func (h *ApptCommandHandler) HandleMarkAsNotAttend(ctx context.Context, cmd c.NotAttendApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(ctx, cmd.AppointmentID())
	if err != nil {
		return cqrs.FailureResult(ApptNotFound, err)
	}

	if err := appointment.MarkAsNotPresented(ctx); err != nil {
		return cqrs.FailureResult(MarkAsNotPresentedFailed, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return cqrs.FailureResult(SaveApptFailed, err)
	}

	return cqrs.SuccessResult(SuccessMarkedAsNotPresented)
}

func (h *ApptCommandHandler) HandleReschedule(ctx context.Context, cmd c.RescheduleApptCommand) cqrs.CommandResult {
	appointment, err := h.getAppByIDAndEmployeeID(ctx, cmd.AppointmentID(), cmd.EmployeeID())
	if err != nil {
		return cqrs.FailureResult(FailedToCheckExistence, err)
	}

	if err := appointment.Reschedule(ctx, cmd.DateTime()); err != nil {
		return cqrs.FailureResult(UpdateApptFailed, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return cqrs.FailureResult(SaveApptFailed, err)
	}

	return cqrs.SuccessResult(SuccessApptUpdated)
}

func (h *ApptCommandHandler) HandleDelete(ctx context.Context, cmd c.DeleteApptCommand) cqrs.CommandResult {
	if exists, err := h.apptRepository.ExistsByID(ctx, cmd.AppointmentID()); err != nil {
		return cqrs.FailureResult(FailedToCheckExistence, err)
	} else if !exists {
		return cqrs.FailureResult(AppointmentNotFound, ErrAppointmentNotFound(cmd.AppointmentID()))
	}

	if err := h.apptRepository.Delete(ctx, cmd.AppointmentID(), false); err != nil {
		return cqrs.FailureResult(FailedToDelete, err)
	}

	return cqrs.SuccessResult(SuccessApptDeleted)
}

func (h *ApptCommandHandler) HandleUpdate(ctx context.Context, cmd c.UpdateApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(ctx, cmd.AppointmentID())
	if err != nil {
		return cqrs.FailureResult(ApptNotFound, err)
	}

	if err := appointment.Update(ctx, cmd.Notes(), cmd.Service()); err != nil {
		return cqrs.FailureResult(UpdateApptFailed, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return cqrs.FailureResult(SaveApptFailed, err)
	}

	return cqrs.SuccessResult(SuccessApptUpdated)
}

func (h *ApptCommandHandler) HandleConfirm(ctx context.Context, cmd c.ConfirmApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(ctx, cmd.ID())
	if err != nil {
		return cqrs.FailureResult(FailedToCheckExistence, err)
	}

	if err := appointment.Confirm(ctx, cmd.EmployeeID()); err != nil {
		return cqrs.FailureResult(ConfirmApptFailed, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return cqrs.FailureResult(SaveApptFailed, err)
	}

	return cqrs.SuccessResult(SuccessConfirmedAppt)
}

func (h *ApptCommandHandler) HandleCreate(ctx context.Context, cmd c.CreateApptCommand) cqrs.CommandResult {
	appointment := cmd.ToEntity()

	if err := appointment.ValidatePersistence(ctx); err != nil {
		return cqrs.FailureResult(BusinessRuleFailed, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return cqrs.FailureResult(SaveApptFailed, err)
	}

	return cqrs.SuccessCreateResult(appointment.ID().String(), SuccessApptCreated)
}

func (h *ApptCommandHandler) HandleComplete(ctx context.Context, cmd c.CompleteApptCommand) cqrs.CommandResult {
	appointment, err := h.getAppByIDAndEmployeeID(ctx, cmd.ID(), cmd.EmployeeID())
	if err != nil {
		return cqrs.FailureResult(FailedToCheckExistence, err)
	}

	if err := appointment.Complete(ctx); err != nil {
		return cqrs.FailureResult(CompleteApptFailed, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return cqrs.FailureResult(UpdateApptFailed, err)
	}

	return cqrs.SuccessResult(SuccessApptUpdated)
}

func (h *ApptCommandHandler) getAppByIDAndEmployeeID(ctx context.Context, appointID valueobject.AppointmentID, employeeID *valueobject.EmployeeID) (appointment.Appointment, error) {
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
