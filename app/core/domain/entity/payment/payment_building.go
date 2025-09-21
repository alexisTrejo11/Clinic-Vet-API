// payment.go
package payment

import (
	"context"
	"time"

	"clinic-vet-api/app/core/domain/entity/base"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
)

type PaymentOption func(context.Context, *Payment) error

func WithAmount(amount valueobject.Money) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if amount.Amount().IsZero() || amount.Amount().IsNegative() {
			return AmountInvalidError(ctx, amount, "WithAmount")
		}
		p.amount = amount
		return nil
	}
}

func WithPaymentMethod(method enum.PaymentMethod) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if !method.IsValid() {
			return MethodInvalidError(ctx, method, "WithPaymentMethod")
		}
		p.method = method
		return nil
	}
}

func WithStatus(status enum.PaymentStatus) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if !status.IsValid() {
			return StatusInvalidError(ctx, status, "WithStatus")
		}
		p.status = status
		return nil
	}
}

func WithTransactionID(transactionID string) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if transactionID == "" {
			return TransactionIDEmptyError(ctx, "WithTransactionID")
		}
		p.transactionID = &transactionID
		return nil
	}
}

func WithDescription(description string) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if len(description) > 500 {
			return DescriptionTooLongError(ctx, len(description), "WithDescription")
		}
		p.description = &description
		return nil
	}
}

func WithDueDate(dueDate time.Time) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if dueDate.Before(time.Now()) {
			return DueDatePastError(ctx, "WithDueDate")
		}
		p.dueDate = &dueDate
		return nil
	}
}

func WithPaidAt(paidAt time.Time) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if paidAt.After(time.Now()) {
			return PaidDateFutureError(ctx, "WithPaidAt")
		}
		p.paidAt = &paidAt
		return nil
	}
}

func WithRefundedAt(refundedAt time.Time) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if refundedAt.After(time.Now()) {
			return RefundDateFutureError(ctx, "WithRefundedAt")
		}
		p.refundedAt = &refundedAt
		return nil
	}
}

func WithPaidFromCustomer(customerID valueobject.CustomerID) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if customerID.IsZero() {
			return CustomerIDRequiredError(ctx, "WithPaidFromCustomer")
		}
		p.paidFromCustomer = customerID
		return nil
	}
}

func WithPaidToEmployee(employeeID valueobject.EmployeeID) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		p.paidToEmployee = employeeID
		return nil
	}
}

func WithAppointmentID(appointmentID valueobject.AppointmentID) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if appointmentID.IsZero() {
			return AppointmentIDRequiredError(ctx, "WithAppointmentID")
		}
		p.appointmentID = &appointmentID
		return nil
	}
}

func WithInvoiceID(invoiceID string) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if invoiceID == "" {
			return InvoiceIDEmptyError(ctx, "WithInvoiceID")
		}
		p.invoiceID = &invoiceID
		return nil
	}
}

func WithRefundAmount(refundAmount valueobject.Money) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if refundAmount.Amount().IsNegative() {
			return RefundAmountNegativeError(ctx, refundAmount, "WithRefundAmount")
		}
		p.refundAmount = &refundAmount
		return nil
	}
}

func WithFailureReason(failureReason string) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		if failureReason == "" {
			return FailureReasonEmptyError(ctx, "WithFailureReason")
		}
		p.failureReason = &failureReason
		return nil
	}
}

func WithIsActive(isActive bool) PaymentOption {
	return func(ctx context.Context, p *Payment) error {
		p.isActive = isActive
		return nil
	}
}

func NewPayment(
	paymentID valueobject.PaymentID,
	createAt time.Time,
	updatedAt time.Time,
	opts ...PaymentOption,
) (*Payment, error) {
	ctx := context.Background()
	return NewPaymentWithContext(ctx, paymentID, createAt, updatedAt, opts...)
}

func NewPaymentWithContext(
	ctx context.Context,
	paymentID valueobject.PaymentID,
	createAt time.Time,
	updatedAt time.Time,
	opts ...PaymentOption,
) (*Payment, error) {
	operation := "NewPaymentWithContext"

	payment := &Payment{
		Entity: base.NewEntity(paymentID, createAt, updatedAt, 1),
	}

	for _, opt := range opts {
		if err := opt(ctx, payment); err != nil {
			return nil, err
		}
	}

	if err := payment.validate(ctx, operation); err != nil {
		return nil, err
	}

	return payment, nil
}

func CreatePayment(
	ctx context.Context,
	paidFromCustomer valueobject.CustomerID,
	opts ...PaymentOption,
) (*Payment, error) {
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
		if err := opt(ctx, payment); err != nil {
			return nil, err
		}
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
