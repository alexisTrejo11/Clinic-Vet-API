package command

const (
	ErrAppointmentNotFound      = "appointment not found"
	ErrMarkAsNotPresentedFailed = "failed to mark appointment as not presented"
	ErrSaveAppointmentFailed    = "failed to save appointment"
	ErrInvalidCommandType       = "invalid command type"
	ErrFindingAppointment       = "error finding appointment"
	ErrCannotDeleteCompleted    = "cannot delete completed appointment"
	ErrFailedToDelete           = "failed to delete appointment"
	ErrFailedToCancel           = "failed to cancel appointment"
	ErrUpdateAppointmentFailed  = "failed to update appointment"
	ErrDeleteAppointmentFailed  = "failed to delete appointment"
	ErrMarkNotPresentedFailed   = "failed to mark as not presented"

	SuccessAppointmentUpdated   = "appointment updated successfully"
	SuccessAppointmentDeleted   = "appointment deleted successfully"
	SuccessAppointmentCanceled  = "appointment canceled successfully"
	SuccessMarkedAsNotPresented = "appointment marked as not presented successfully"
)
