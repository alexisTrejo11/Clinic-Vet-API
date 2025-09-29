package handler

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

const (
	ApptNotFound             = "appointment not found"
	MarkAsNotPresentedFailed = "failed to mark appointment as not presented"
	SaveApptFailed           = "failed to save appointment"
	FindingAppt              = "error finding appointment"
	FailedToDelete           = "failed to delete appointment"
	FailedToCancel           = "failed to cancel appointment"
	UpdateApptFailed         = "failed to update appointment"
	InvalidCommand           = "invalid command data"
	CompleteApptFailed       = "failed to complete appointment"
	ConfirmApptFailed        = "failed to confirm appointment"
	FailedToCheckExistence   = "failed to check appointment existence"
	BusinessRuleFailed       = "business rule validation failed"
	AppointmentNotFound      = "appointment not found"

	SuccessApptCreated          = "appointment created successfully"
	SuccessApptUpdated          = "appointment updated successfully"
	SuccessApptDeleted          = "appointment deleted successfully"
	SuccessApptCanceled         = "appointment canceled successfully"
	SuccessMarkedAsNotPresented = "appointment marked as not presented successfully"
	SuccessConfirmedAppt        = "appointment confirmed successfully"
)

func ErrAppointmentNotFound(id valueobject.AppointmentID) error {
	return apperror.EntityNotFoundValidationError("Appointment", "id", id.String())
}
