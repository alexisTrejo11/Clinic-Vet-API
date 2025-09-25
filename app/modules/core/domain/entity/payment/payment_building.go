// payment.go
package payment

import (
	"context"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type PaymentOption func(*Payment)

func WithAmount(amount valueobject.Money) PaymentOption {
	return func(p *Payment) {
		p.amount = amount
	}
}

func WithPaymentMethod(method enum.PaymentMethod) PaymentOption {
	return func(p *Payment) {
		p.method = method
	}
}

func WithStatus(status enum.PaymentStatus) PaymentOption {
	return func(p *Payment) {
		p.status = status
	}
}

func WithTransactionID(transactionID string) PaymentOption {
	return func(p *Payment) {
		p.transactionID = &transactionID
	}
}

func WithDescription(description string) PaymentOption {
	return func(p *Payment) {
		p.description = &description
	}
}

func WithDueDate(dueDate time.Time) PaymentOption {
	return func(p *Payment) {
		p.dueDate = &dueDate
	}
}

func WithPaidAt(paidAt time.Time) PaymentOption {
	return func(p *Payment) {
		p.paidAt = &paidAt
	}
}

func WithRefundedAt(refundedAt time.Time) PaymentOption {
	return func(p *Payment) {
		p.refundedAt = &refundedAt
	}
}

func WithPaidFromCustomer(customerID valueobject.CustomerID) PaymentOption {
	return func(p *Payment) {
		p.paidFromCustomer = customerID
	}
}

func WithPaidToEmployee(employeeID valueobject.EmployeeID) PaymentOption {
	return func(p *Payment) {
		p.paidToEmployee = employeeID
	}
}

func WithAppointmentID(appointmentID valueobject.AppointmentID) PaymentOption {
	return func(p *Payment) {
		p.appointmentID = &appointmentID
	}
}

func WithInvoiceID(invoiceID string) PaymentOption {
	return func(p *Payment) {
		p.invoiceID = &invoiceID
	}
}

func WithRefundAmount(refundAmount valueobject.Money) PaymentOption {
	return func(p *Payment) {
		p.refundAmount = &refundAmount
	}
}

func WithFailureReason(failureReason string) PaymentOption {
	return func(p *Payment) {
		p.failureReason = &failureReason
	}
}

func WithIsActive(isActive bool) PaymentOption {
	return func(p *Payment) {
		p.isActive = isActive
	}
}

func WithTimeStamps(createdAt, updatedAt time.Time, version int) PaymentOption {
	return func(p *Payment) {
		p.Entity = base.NewEntity(p.ID(), &createdAt, &updatedAt, version)
	}
}

func NewPayment(paymentID valueobject.PaymentID, opts ...PaymentOption) *Payment {
	payment := &Payment{
		Entity: base.NewEntity(paymentID, nil, nil, 0),
	}

	for _, opt := range opts {
		opt(payment)
	}

	return payment
}

func CreatePayment(ctx context.Context, paidFromCustomer valueobject.CustomerID, opts ...PaymentOption) (*Payment, error) {
	operation := "CreatePayment"

	if paidFromCustomer.IsZero() {
		return nil, CustomerIDRequiredError(ctx, operation)
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
		opt(payment)
	}

	if err := payment.validate(ctx, operation); err != nil {
		return nil, err
	}

	return payment, nil
}

func (p *Payment) validate(ctx context.Context, operation string) error {
	if p.amount.Amount().IsZero() || p.amount.Amount().IsNegative() {
		return AmountInvalidError(ctx, p.amount, operation)
	}
	if !p.method.IsValid() {
		return MethodInvalidError(ctx, p.method, operation)
	}
	if !p.status.IsValid() {
		return StatusInvalidError(ctx, p.status, operation)
	}
	if p.dueDate != nil && p.dueDate.Before(time.Now()) {
		return DueDatePastError(ctx, operation)
	}
	return nil
}
