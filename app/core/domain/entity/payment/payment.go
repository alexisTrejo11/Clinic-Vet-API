// Package payment defines the Payment entity and its behaviors.
package payment

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

type Payment struct {
	base.Entity[valueobject.PaymentID]
	appointmentID valueobject.AppointmentID
	userID        valueobject.UserID
	amount        valueobject.Money
	currency      string
	paymentMethod enum.PaymentMethod
	status        enum.PaymentStatus
	transactionID *string
	description   *string
	dueDate       *time.Time
	paidAt        *time.Time
	refundedAt    *time.Time
}

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

// NewPayment creates a new Payment with functional options
func NewPayment(
	id valueobject.PaymentID,
	appointmentID valueobject.AppointmentID,
	userID valueobject.UserID,
	opts ...PaymentOption,
) (*Payment, error) {
	if id.Value() == 0 {
		return nil, domainerr.NewValidationError("payment", "id", "payment ID is required")
	}
	if appointmentID.Value() == 0 {
		return nil, domainerr.NewValidationError("payment", "appointmentID", "appointment ID is required")
	}
	if userID.Value() == 0 {
		return nil, domainerr.NewValidationError("payment", "userID", "user ID is required")
	}

	payment := &Payment{
		Entity:        base.NewEntity(id),
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

func (p *Payment) ID() valueobject.PaymentID {
	return p.Entity.ID()
}

func (p *Payment) AppointmentID() valueobject.AppointmentID {
	return p.appointmentID
}

func (p *Payment) UserID() valueobject.UserID {
	return p.userID
}

func (p *Payment) Amount() valueobject.Money {
	return p.amount
}

func (p *Payment) Currency() string {
	return p.currency
}

func (p *Payment) PaymentMethod() enum.PaymentMethod {
	return p.paymentMethod
}

func (p *Payment) Status() enum.PaymentStatus {
	return p.status
}

func (p *Payment) TransactionID() *string {
	return p.transactionID
}

func (p *Payment) Description() *string {
	return p.description
}

func (p *Payment) DueDate() *time.Time {
	return p.dueDate
}

func (p *Payment) PaidAt() *time.Time {
	return p.paidAt
}

func (p *Payment) RefundedAt() *time.Time {
	return p.refundedAt
}

func (p *Payment) Update(amount *valueobject.Money, paymentMethod *enum.PaymentMethod, description *string, dueDate *time.Time) error {
	if amount != nil {
		if amount.Amount() <= 0 {
			return domainerr.NewValidationError("payment", "amount", "amount must be positive")
		}
		p.amount = *amount
	}

	if paymentMethod != nil {
		if !paymentMethod.IsValid() {
			return domainerr.NewValidationError("payment", "payment method", "invalid payment method")
		}
		p.paymentMethod = *paymentMethod
	}

	if description != nil {
		if len(*description) > 500 {
			return domainerr.NewValidationError("payment", "description", "description too long")
		}
		p.description = description
	}

	if dueDate != nil {
		if dueDate.Before(time.Now()) {
			return domainerr.NewValidationError("payment", "due date", "due date cannot be in the past")
		}
		p.dueDate = dueDate
	}

	p.IncrementVersion()
	return nil
}
