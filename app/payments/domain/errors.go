package paymentDomain

import (
	"errors"
	"time"

	domainErr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/domain"
)

var (
	ErrPaymentNotFound         = errors.New("payment not found")
	ErrInvalidPaymentStatus    = errors.New("invalid payment status")
	ErrInvalidPaymentMethod    = errors.New("invalid payment method")
	ErrInvalidAmount           = errors.New("invalid payment amount")
	ErrInvalidCurrency         = errors.New("invalid currency")
	ErrPaymentAlreadyPaid      = errors.New("payment is already paid")
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
