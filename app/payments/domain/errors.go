package paymentDomain

import (
	"errors"
	"strconv"
	"time"

	domainErr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/domain"
)

var (
	ErrInvalidPaymentStatus    = errors.New("invalid payment status")
	ErrInvalidPaymentMethod    = errors.New("invalid payment method")
	ErrInvalidAmount           = errors.New("invalid payment amount")
	ErrInvalidCurrency         = errors.New("invalid currency")
	ErrPaymentAlreadyPaid      = errors.New("payment is already paid")
	ErrPaymentAlreadyCancelled = errors.New("payment is already cancelled")
	ErrPaymentCannotBeRefunded = errors.New("payment cannot be refunded")
	ErrPaymentProcessingFailed = errors.New("payment processing failed")
	ErrInsufficientFunds       = errors.New("insufficient funds")
	ErrPaymentExpired          = errors.New("payment has expired")
	ErrDuplicatePayment        = errors.New("duplicate payment")
	ErrCurrencyMismatch        = errors.New("currency mismatch")
)

type PaymentError struct {
	PaymentID int       `json:"payment_id,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	domainErr.BaseDomainError
}

func (e PaymentError) Error() string {
	return e.Message
}

func NewPaymentError(code, message string, paymentID int, details string) *PaymentError {
	return &PaymentError{
		BaseDomainError: domainErr.BaseDomainError{
			Code:    code,
			Type:    "Domain Error",
			Message: message,
		},
		PaymentID: paymentID,
		Timestamp: time.Now(),
	}
}

func PaymentNotFoundErr(paymentID int) *PaymentError {
	return &PaymentError{
		BaseDomainError: domainErr.BaseDomainError{
			Code:    "ERR_PAYMENT_NOT_FOUND",
			Type:    "Domain Error",
			Message: "Payment not found",
		},
		PaymentID: paymentID,
		Timestamp: time.Now(),
	}
}

func InvalidPaymentIdErr(paymentId int) *PaymentError {
	var message string

	if paymentId <= 0 {
		message = "Payment ID must provided and greater than zero"
	} else {
		message = "Invalid payment ID: " + strconv.Itoa(paymentId)
	}

	return &PaymentError{
		BaseDomainError: domainErr.BaseDomainError{
			Code:    "ERR_INVALID_PAYMENT_ID",
			Type:    "Domain Error",
			Message: message,
		},
		PaymentID: paymentId,
		Timestamp: time.Now(),
	}
}

func InvalidPaymentStatusErr(status PaymentStatus) *PaymentError {
	return &PaymentError{
		BaseDomainError: domainErr.BaseDomainError{
			Code:    "ERR_INVALID_PAYMENT_STATUS",
			Type:    "Domain Error",
			Message: "Invalid payment status: " + string(status),
		},
		Timestamp: time.Now(),
	}
}

func PaymentStatusConflict(paymentId int, err error) error {
	if errors.Is(err, ErrPaymentAlreadyPaid) {
		return NewPaymentError("ALREADY_PAID", "payment is already paid", paymentId, "")
	}
	if errors.Is(err, ErrPaymentAlreadyCancelled) {
		return NewPaymentError("ALREADY_CANCELLED", "payment is already cancelled", paymentId, "")
	}

	message := "invalid payment status"
	if errors.Is(err, ErrInvalidPaymentStatus) {
		message += " for payment ID " + strconv.Itoa(paymentId)
	} else {
		message += " for unknown payment"

		return NewPaymentError("INVALID_STATUS", message, paymentId, "")
	}
	return NewPaymentError("INVALID_STATUS", "an error ocurred with handling payment status", paymentId, "")
}
