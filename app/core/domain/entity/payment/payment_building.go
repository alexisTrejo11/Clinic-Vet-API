package payment

import (
	"time"

	"clinic-vet-api/app/core/domain/entity/base"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	domainerr "clinic-vet-api/app/core/error"
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

func WithPaymentMethod(method enum.PaymentMethod) PaymentOption {
	return func(p *Payment) error {
		if !method.IsValid() {
			return domainerr.NewValidationError("payment", "paymentMethod", "invalid payment method")
		}
		p.method = method
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

func WithTransactionID(transactionID string) PaymentOption {
	return func(p *Payment) error {
		if transactionID == "" {
			return domainerr.NewValidationError("payment", "transactionID", "transaction ID cannot be empty")
		}
		p.transactionID = &transactionID
		return nil
	}
}

func WithDescription(description string) PaymentOption {
	return func(p *Payment) error {
		if len(description) > 500 {
			return domainerr.NewValidationError("payment", "description", "description too long")
		}
		p.description = &description
		return nil
	}
}

func WithDueDate(dueDate *time.Time) PaymentOption {
	return func(p *Payment) error {
		if dueDate.Before(time.Now()) {
			return domainerr.NewValidationError("payment", "dueDate", "due date cannot be in the past")
		}
		p.dueDate = dueDate
		return nil
	}
}

func WithPaidAt(paidAt time.Time) PaymentOption {
	return func(p *Payment) error {
		if paidAt.After(time.Now()) {
			return domainerr.NewValidationError("payment", "paidAt", "paid date cannot be in the future")
		}
		p.paidAt = &paidAt
		return nil
	}
}

func WithRefundedAt(refundedAt time.Time) PaymentOption {
	return func(p *Payment) error {
		if refundedAt.After(time.Now()) {
			return domainerr.NewValidationError("payment", "refundedAt", "refunded date cannot be in the future")
		}
		p.refundedAt = &refundedAt
		return nil
	}
}

func WithPaidFromCustomer(customerID valueobject.CustomerID) PaymentOption {
	return func(p *Payment) error {
		if customerID.IsZero() {
			return domainerr.NewValidationError("payment", "paidFromCustomer", "customer ID is required")
		}
		p.paidFromCustomer = customerID
		return nil
	}
}

func WithPaidToEmployee(employeeID valueobject.EmployeeID) PaymentOption {
	return func(p *Payment) error {
		if employeeID.IsZero() {
			return domainerr.NewValidationError("payment", "paidToEmployee", "employee ID is required")
		}
		p.paidToEmployee = employeeID
		return nil
	}
}

func WithAppointmentID(appointmentID valueobject.AppointmentID) PaymentOption {
	return func(p *Payment) error {
		if appointmentID.IsZero() {
			return domainerr.NewValidationError("payment", "appointmentID", "appointment ID is required")
		}
		p.appointmentID = &appointmentID
		return nil
	}
}

func WithInvoiceID(invoiceID string) PaymentOption {
	return func(p *Payment) error {
		if invoiceID == "" {
			return domainerr.NewValidationError("payment", "invoiceID", "invoice ID cannot be empty")
		}
		p.invoiceID = &invoiceID
		return nil
	}
}

func WithRefundAmount(refundAmount valueobject.Money) PaymentOption {
	return func(p *Payment) error {
		if refundAmount.Amount() < 0 {
			return domainerr.NewValidationError("payment", "refundAmount", "refund amount cannot be negative")
		}
		p.refundAmount = &refundAmount
		return nil
	}
}

func WithFailureReason(failureReason string) PaymentOption {
	return func(p *Payment) error {
		if failureReason == "" {
			return domainerr.NewValidationError("payment", "failureReason", "failure reason cannot be empty")
		}
		p.failureReason = &failureReason
		return nil
	}
}

func WithIsActive(isActive bool) PaymentOption {
	return func(p *Payment) error {
		p.isActive = isActive
		return nil
	}
}

func NewPayment(
	paymentID valueobject.PaymentID,
	createAt time.Time,
	updatedAt time.Time,
	opts ...PaymentOption,
) *Payment {
	payment := &Payment{
		Entity: base.NewEntity(paymentID, createAt, updatedAt, 1),
	}
	for _, opt := range opts {
		if err := opt(payment); err != nil {
			return nil
		}
	}

	return payment
}

func CreatePayment(
	paidFromCustomer valueobject.CustomerID,
	opts ...PaymentOption,
) (*Payment, error) {

	if paidFromCustomer.IsZero() {
		return nil, domainerr.NewValidationError("payment", "paidFromCustomer", "customer ID is required")
	}

	payment := &Payment{
		Entity:           base.CreateEntity(valueobject.NewPaymentID(0)),
		paidFromCustomer: paidFromCustomer,
		status:           enum.PaymentStatusPending,
		method:           enum.PaymentMethodCash,
		isActive:         true,
		dueDate:          func() *time.Time { t := time.Now().Add(7 * 24 * time.Hour); return &t }(), // Default: 7 days from now
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

func (p *Payment) validate() error {
	if p.amount.Amount() <= 0 {
		return domainerr.NewValidationError("payment", "amount", "amount must be positive")
	}
	if !p.method.IsValid() {
		return domainerr.NewValidationError("payment", "paymentMethod", "payment method is required")
	}
	if !p.status.IsValid() {
		return domainerr.NewValidationError("payment", "status", "status is required")
	}
	if p.dueDate.Before(time.Now()) {
		return domainerr.NewValidationError("payment", "dueDate", "due date cannot be in the past")
	}
	return nil
}
