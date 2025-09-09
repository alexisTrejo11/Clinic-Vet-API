package command

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

	SuccessApptUpdated          = "appointment updated successfully"
	SuccessApptDeleted          = "appointment deleted successfully"
	SuccessApptCanceled         = "appointment canceled successfully"
	SuccessMarkedAsNotPresented = "appointment marked as not presented successfully"
)
