package command

import (
	"context"

	appt "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type AppointementCommandHandler interface {
	CancelAppointment(command CancelApptCommand) cqrs.CommandResult
	CreateAppointment(command CreateApptCommand) cqrs.CommandResult
	DeleteAppointment(command DeleteApptCommand) cqrs.CommandResult
	ConfirmAppointment(command ConfirmApptCommand) cqrs.CommandResult
	MarkAppointmentAsNotAttend(command NotAttendApptCommand) cqrs.CommandResult
	RescheduleAppointment(command RescheduleApptCommand) cqrs.CommandResult
	CompleteAppointment(command CompleteApptCommand) cqrs.CommandResult
	UpdateAppointment(command UpdateApptCommand) cqrs.CommandResult

	RequestAppointmentByCustomer(command RequestApptByCustomerCommand) cqrs.CommandResult
}

type apptCommandHandler struct {
	apptRepository repository.AppointmentRepository
}

func NewAppointmentCommandHandler(apptRepository repository.AppointmentRepository) AppointementCommandHandler {
	return &apptCommandHandler{
		apptRepository: apptRepository,
	}
}

func (h *apptCommandHandler) CreateAppointment(command CreateApptCommand) cqrs.CommandResult {
	appointment, err := createCommandToDomain(command)
	if err != nil {
		return cqrs.FailureResult("failed to create appointment domain", err)
	}

	if err := h.apptRepository.Save(command.ctx, appointment); err != nil {
		return cqrs.FailureResult("failed to save appointment", err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), "appointment created successfully")
}

func (h *apptCommandHandler) CancelAppointment(command CancelApptCommand) cqrs.CommandResult {
	appointment, err := h.getAppByIDAndEmployeeID(command.ctx, command.appointmentID, command.vetID)
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

func (h *apptCommandHandler) ConfirmAppointment(command ConfirmApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(command.ctx, command.id)
	if err != nil {
		return cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.Confirm(command.employeeID); err != nil {
		return cqrs.FailureResult("failed to confirm appointment", err)
	}

	if err := h.apptRepository.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save confirmed appointment", err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), "appointment confirmed successfully")
}

func (h *apptCommandHandler) DeleteAppointment(command DeleteApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult(ErrFindingAppt, err)
	}

	if err := h.apptRepository.Delete(command.ctx, command.appointmentID); err != nil {
		return cqrs.FailureResult(ErrFailedToDelete, err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), SuccessApptDeleted)
}

func (h *apptCommandHandler) getAppByIDAndEmployeeID(ctx context.Context, appointID valueobject.AppointmentID, employeeID *valueobject.EmployeeID) (appt.Appointment, error) {
	if employeeID == nil {
		return h.apptRepository.FindByID(ctx, appointID)
	}
	return h.apptRepository.FindByIDAndEmployeeID(ctx, appointID, *employeeID)
}

func (h *apptCommandHandler) RescheduleAppointment(command RescheduleApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.Reschedule(command.datetime); err != nil {
		return cqrs.FailureResult("failed to reschedule appointment", err)
	}

	if err := h.apptRepository.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save rescheduled appointment", err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), "appointment rescheduled successfully")
}

func (h *apptCommandHandler) MarkAppointmentAsNotAttend(command NotAttendApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.MarkAsNotPresented(); err != nil {
		return cqrs.FailureResult(ErrMarkAsNotPresentedFailed, err)
	}

	if err := h.apptRepository.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult(ErrSaveApptFailed, err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), SuccessMarkedAsNotPresented)
}

func (h *apptCommandHandler) CompleteAppointment(command CompleteApptCommand) cqrs.CommandResult {
	appointment, err := h.getAppByIDAndEmployeeID(command.ctx, command.id, command.employeeID)
	if err != nil {
		return cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.Complete(); err != nil {
		return cqrs.FailureResult("failed to complete appointment", err)
	}

	if err := h.apptRepository.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save completed appointment", err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), "appointment completed successfully")
}

func (h *apptCommandHandler) UpdateAppointment(command UpdateApptCommand) cqrs.CommandResult {
	appointment, err := h.apptRepository.FindByID(command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.Update(command.notes, command.vetID, command.service, command.reason); err != nil {
		return cqrs.FailureResult(ErrUpdateApptFailed, err)
	}

	if err := h.apptRepository.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult(ErrSaveApptFailed, err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), SuccessApptUpdated)
}

func (h *apptCommandHandler) RequestAppointmentByCustomer(command RequestApptByCustomerCommand) cqrs.CommandResult {
	appointment, err := requestByCustomerCommandToDomain(command)
	if err != nil {
		return cqrs.FailureResult("failed to create appointment domain", err)
	}

	if err := h.apptRepository.Save(command.ctx, appointment); err != nil {
		return cqrs.FailureResult("failed to save appointment", err)
	}

	// Event

	return cqrs.SuccessResult(appointment.ID().String(), "appointment created successfully")
}
