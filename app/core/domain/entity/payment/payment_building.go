package payment

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

type PaymentOption func(*Payment) error

func WithAmount(amount valueobject.Money) PaymentOption {
	return func(p *Payment) error {
		if amount.Amount() <= 0 {
			return domainerr.NewValidationError("payment", "amount", "amount must be positive")
		}
		p.amount = amount
		return nil
	}
}

func WithCurrency(currency string) PaymentOption {
	return func(p *Payment) error {
		if currency == "" {
			return domainerr.NewValidationError("payment", "currency", "currency is required")
		}
		if len(currency) != 3 {
			return domainerr.NewValidationError("payment", "currency", "currency must be 3-letter code")
		}
		p.currency = currency
		return nil
	}
}

func WithPaymentMethod(method enum.PaymentMethod) PaymentOption {
	return func(p *Payment) error {
		if !method.IsValid() {
			return domainerr.NewValidationError("payment", "paymentMethod", "invalid payment method")
		}
		p.paymentMethod = method
		return nil
	}
}

func WithStatus(status enum.PaymentStatus) PaymentOption {
	return func(p *Payment) error {
		if !status.IsValid() {
			return domainerr.NewValidationError("payment", "status", "invalid payment status")
		}
		p.status = status
		return nil
	}
}

func WithTransactionID(transactionID *string) PaymentOption {
	return func(p *Payment) error {
		if transactionID != nil && *transactionID == "" {
			return domainerr.NewValidationError("payment", "transactionID", "transaction ID cannot be empty")
		}
		p.transactionID = transactionID
		return nil
	}
}

func WithDescription(description *string) PaymentOption {
	return func(p *Payment) error {
		if description != nil && len(*description) > 500 {
			return domainerr.NewValidationError("payment", "description", "description too long")
		}
		p.description = description
		return nil
	}
}

func WithDueDate(dueDate *time.Time) PaymentOption {
	return func(p *Payment) error {
		if dueDate != nil && dueDate.Before(time.Now()) {
			return domainerr.NewValidationError("payment", "dueDate", "due date cannot be in the past")
		}
		p.dueDate = dueDate
		return nil
	}
}

func WithPaidAt(paidAt *time.Time) PaymentOption {
	return func(p *Payment) error {
		if paidAt != nil && paidAt.After(time.Now()) {
			return domainerr.NewValidationError("payment", "paidAt", "paid date cannot be in the future")
		}
		p.paidAt = paidAt
		return nil
	}
}

func WithRefundedAt(refundedAt *time.Time) PaymentOption {
	return func(p *Payment) error {
		if refundedAt != nil && refundedAt.After(time.Now()) {
			return domainerr.NewValidationError("payment", "refundedAt", "refunded date cannot be in the future")
		}
		p.refundedAt = refundedAt
		return nil
	}
}

func WithTimeStamps(createdAt, updatedAt time.Time) PaymentOption {
	return func(p *Payment) error {
		if createdAt.IsZero() || updatedAt.IsZero() {
			return domainerr.NewValidationError("payment", "timestamps", "createdAt and updatedAt are required")
		}
		p.SetTimeStamps(createdAt, updatedAt)
		return nil
	}
}

// NewPayment creates a new Payment with functional options
func NewPayment(
	id valueobject.PaymentID,
	appointmentID valueobject.AppointmentID,
	userID valueobject.UserID,
	opts ...PaymentOption,
) (*Payment, error) {
	payment := &Payment{
		Entity:        base.NewEntity(id, time.Now(), time.Now(), 1),
		appointmentID: appointmentID,
		userID:        userID,
	}

	for _, opt := range opts {
		if err := opt(payment); err != nil {
			return nil, err
		}
	}

	return payment, nil
}

func CreatePayment(
	appointmentID valueobject.AppointmentID,
	userID valueobject.UserID,
	opts ...PaymentOption,
) (*Payment, error) {
	if appointmentID.Value() == 0 {
		return nil, domainerr.NewValidationError("payment", "appointmentID", "appointment ID is required")
	}
	if userID.Value() == 0 {
		return nil, domainerr.NewValidationError("payment", "userID", "user ID is required")
	}

	payment := &Payment{
		Entity:        base.CreateEntity(valueobject.PaymentID{}),
		appointmentID: appointmentID,
		userID:        userID,
		status:        enum.PaymentStatusPending, // Default status
		currency:      "USD",                     // Default currency
	}

	for _, opt := range opts {
		if err := opt(payment); err != nil {
			return nil, err
		}
	}

	if err := payment.validate(); err != nil {
		return nil, err
	}

	return payment, nil
}

// Validation
func (p *Payment) validate() error {
	if p.amount.Amount() <= 0 {
		return domainerr.NewValidationError("payment", "amount", "amount must be positive")
	}
	if p.currency == "" {
		return domainerr.NewValidationError("payment", "currency", "currency is required")
	}
	if !p.paymentMethod.IsValid() {
		return domainerr.NewValidationError("payment", "payment Method", "payment method is required")
	}
	if !p.status.IsValid() {
		return domainerr.NewValidationError("payment", "status", "status is required")
	}
	return nil
}
