package command

import (
	"context"

	appt "clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	apperror "clinic-vet-api/app/shared/error/application"
)

type AppointementCommandHandler interface {
	CancelAppointment(ctx context.Context, command CancelApptCommand) cqrs.CommandResult
	CreateAppointment(ctx context.Context, command CreateApptCommand) cqrs.CommandResult
	DeleteAppointment(ctx context.Context, command DeleteApptCommand) cqrs.CommandResult
	ConfirmAppointment(ctx context.Context, command ConfirmApptCommand) cqrs.CommandResult
	MarkAppointmentAsNotAttend(ctx context.Context, command NotAttendApptCommand) cqrs.CommandResult
	RescheduleAppointment(ctx context.Context, command RescheduleApptCommand) cqrs.CommandResult
	CompleteAppointment(ctx context.Context, command CompleteApptCommand) cqrs.CommandResult
	UpdateAppointment(ctx context.Context, command UpdateApptCommand) cqrs.CommandResult

	RequestAppointmentByCustomer(ctx context.Context, command RequestApptByCustomerCommand) cqrs.CommandResult
}

type apptCommandHandler struct {
	apptRepository repository.AppointmentRepository
}

func NewAppointmentCommandHandler(apptRepository repository.AppointmentRepository) AppointementCommandHandler {
	return &apptCommandHandler{
		apptRepository: apptRepository,
	}
}

func (h *apptCommandHandler) CreateAppointment(ctx context.Context, command CreateApptCommand) cqrs.CommandResult {
	appointment, err := createCommandToDomain(ctx, command)
	if err != nil {
		return *cqrs.FailureResult("failed to create appointment domain", err)
	}

	if err := h.apptRepository.Save(ctx, appointment); err != nil {
		return *cqrs.FailureResult("failed to save appointment", err)
	}

	return *cqrs.SuccessResult(appointment.ID().String(), "appointment created successfully")
}

func (h *apptCommandHandler) CancelAppointment(ctx context.Context, command CancelApptCommand) cqrs.CommandResult {
	appointment, err := h.getAppByIDAndEmployeeID(ctx, command.appointmentID, command.vetID)
	if err != nil {
		return *cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.Cancel(ctx); err != nil {
		return *cqrs.FailureResult(ErrFailedToCancel, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult(ErrUpdateApptFailed, err)
	}

	return *cqrs.SuccessResult(appointment.ID().String(), SuccessApptUpdated)
}

func (h *apptCommandHandler) ConfirmAppointment(ctx context.Context, command ConfirmApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(ctx, command.id)
	if err != nil {
		return *cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.Confirm(ctx, command.employeeID); err != nil {
		return *cqrs.FailureResult("failed to confirm appointment", err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult("failed to save confirmed appointment", err)
	}

	return *cqrs.SuccessResult(appointment.ID().String(), "appointment confirmed successfully")
}

func (h *apptCommandHandler) DeleteAppointment(ctx context.Context, command DeleteApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(ctx, command.appointmentID)
	if err != nil {
		return *cqrs.FailureResult(ErrFindingAppt, err)
	}

	if err := h.apptRepository.Delete(ctx, command.appointmentID, false); err != nil {
		return *cqrs.FailureResult(ErrFailedToDelete, err)
	}

	return *cqrs.SuccessResult(appointment.ID().String(), SuccessApptDeleted)
}

func (h *apptCommandHandler) getAppByIDAndEmployeeID(ctx context.Context, appointID valueobject.AppointmentID, employeeID *valueobject.EmployeeID) (appt.Appointment, error) {
	if employeeID == nil {
		return h.apptRepository.FindByID(ctx, appointID)
	}

	spec := specification.ApptByID(appointID).And(specification.ApptByEmployee(*employeeID))
	appoint, err := h.apptRepository.Find(ctx, spec)
	if err != nil {
		return appt.Appointment{}, err
	}

	if len(appoint.Items) == 0 {
		return appt.Appointment{}, apperror.EntityNotFoundValidationError("Appointment", "id", appointID.String())
	}
	return appoint.Items[0], nil
}

func (h *apptCommandHandler) RescheduleAppointment(ctx context.Context, command RescheduleApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(ctx, command.appointmentID)
	if err != nil {
		return *cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.Reschedule(ctx, command.datetime); err != nil {
		return *cqrs.FailureResult("failed to reschedule appointment", err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult("failed to save rescheduled appointment", err)
	}

	return *cqrs.SuccessResult(appointment.ID().String(), "appointment rescheduled successfully")
}

func (h *apptCommandHandler) MarkAppointmentAsNotAttend(ctx context.Context, command NotAttendApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(ctx, command.appointmentID)
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

func (h *apptCommandHandler) CompleteAppointment(ctx context.Context, command CompleteApptCommand) cqrs.CommandResult {
	appointment, err := h.getAppByIDAndEmployeeID(ctx, command.id, command.employeeID)
	if err != nil {
		return *cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.Complete(ctx); err != nil {
		return *cqrs.FailureResult("failed to complete appointment", err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult("failed to save completed appointment", err)
	}

	return *cqrs.SuccessResult(appointment.ID().String(), "appointment completed successfully")
}

func (h *apptCommandHandler) UpdateAppointment(ctx context.Context, command UpdateApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(ctx, command.appointmentID)
	if err != nil {
		return *cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.Update(ctx, command.notes, command.vetID, command.service); err != nil {
		return *cqrs.FailureResult(ErrUpdateApptFailed, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult(ErrSaveApptFailed, err)
	}

	return *cqrs.SuccessResult(appointment.ID().String(), SuccessApptUpdated)
}

func (h *apptCommandHandler) RequestAppointmentByCustomer(ctx context.Context, command RequestApptByCustomerCommand) cqrs.CommandResult {
	appointment, err := requestByCustomerCommandToDomain(ctx, command)
	if err != nil {
		return *cqrs.FailureResult("failed to create appointment domain", err)
	}

	if err := h.apptRepository.Save(ctx, appointment); err != nil {
		return *cqrs.FailureResult("failed to save appointment", err)
	}

	// Event

	return *cqrs.SuccessResult(appointment.ID().String(), "appointment created successfully")
}
