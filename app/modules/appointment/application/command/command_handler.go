package command

import (
	"context"

	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/cqrs"
)

const (
	ErrApptNotFound             = "appointment not found"
	ErrMarkAsNotPresentedFailed = "failed to mark appointment as not presented"
	ErrSaveApptFailed           = "failed to save appointment"
	ErrInvalidCommandType       = "invalid command type"
	ErrFindingAppt              = "error finding appointment"
	ErrCannotDeleteCompleted    = "cannot delete completed appointment"
	ErrFailedToDelete           = "failed to delete appointment"
	ErrFailedToCancel           = "failed to cancel appointment"
	ErrUpdateApptFailed         = "failed to update appointment"
	ErrDeleteApptFailed         = "failed to delete appointment"
	ErrMarkNotPresentedFailed   = "failed to mark as not presented"
	ErrInvalidCommand           = "invalid command data"

	SuccessApptUpdated          = "appointment updated successfully"
	SuccessApptDeleted          = "appointment deleted successfully"
	SuccessApptCanceled         = "appointment canceled successfully"
	SuccessMarkedAsNotPresented = "appointment marked as not presented successfully"
	ErrFailedToCheckExistence   = "failed to check appointment existence"
)

type AppointementCommandHandler interface {
	CancelAppointment(ctx context.Context, cmd CancelApptCommand) cqrs.CommandResult
	CreateAppointment(ctx context.Context, cmd CreateApptCommand) cqrs.CommandResult
	DeleteAppointment(ctx context.Context, cmd DeleteApptCommand) cqrs.CommandResult

	ConfirmAppointment(ctx context.Context, cmd ConfirmApptCommand) cqrs.CommandResult

	MarkAppointmentAsNotAttend(ctx context.Context, cmd NotAttendApptCommand) cqrs.CommandResult
	RescheduleAppointment(ctx context.Context, cmd RescheduleApptCommand) cqrs.CommandResult
	CompleteAppointment(ctx context.Context, cmd CompleteApptCommand) cqrs.CommandResult
	UpdateAppointment(ctx context.Context, cmd UpdateApptCommand) cqrs.CommandResult

	RequestAppointmentByCustomer(ctx context.Context, cmd RequestApptByCustomerCommand) cqrs.CommandResult
}

type apptCommandHandler struct {
	apptRepository repository.AppointmentRepository
}

func NewAppointmentCommandHandler(apptRepository repository.AppointmentRepository) AppointementCommandHandler {
	return &apptCommandHandler{
		apptRepository: apptRepository,
	}
}
