package command

// Error messages
const (
	ErrFailedRetrievePayment     = "failed to retrieve payment"
	ErrFailedProcessPayment      = "failed to process payment"
	ErrInvalidPaymentData        = "invalid payment data"
	ErrFailedCreatePayment       = "failed to create payment"
	ErrFailedRefundPayment       = "failed to refund payment"
	ErrFailedSaveRefundedPayment = "failed to save refunded payment"
	ErrFailedCancelPayment       = "failed to cancel payment"
	ErrFailedSaveCanceledPayment = "failed to save canceled payment"
	ErrFetchingPayment           = "error fetching payment"
	ErrDeletingPayment           = "error deleting payment"
	ErrUpdatingPayment           = "error updating payment"
	ErrSavingPayment             = "error saving payment"
	ErrFailedSearchPayments      = "failed to search payments"
)

// Success messages
const (
	MsgPaymentProcessed = "payment processed successfully"
	MsgPaymentCreated   = "payment created successfully"
	MsgPaymentRefunded  = "payment refunded successfully"
	MsgPaymentCanceled  = "payment canceled successfully"
	MsgPaymentDeleted   = "payment deleted successfully"
	MsgPaymentUpdated   = "payment updated successfully"
	MsgOverduePayments  = "Updated %d overdue payments"
)

// Pagination constants
const (
	DefaultPageSize = 100
	InitialPage     = 1
)
