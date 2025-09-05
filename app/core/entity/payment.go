package entity

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

type Payment struct {
	id            valueobject.PaymentID
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
	createdAt     time.Time
	updatedAt     time.Time
}

func (p *Payment) Update(amount *valueobject.Money, paymentMethod *enum.PaymentMethod, description *string, dueDate *time.Time) error {
	if amount != nil {
		p.amount = *amount
	}
	if paymentMethod != nil {
		p.paymentMethod = *paymentMethod
	}

	p.description = description
	p.dueDate = dueDate
	return nil
}

func (p *Payment) Cancel(reason string) error {
	allowedStatuses := []enum.PaymentStatus{enum.PENDING, enum.PAID, enum.FAILED}

	if !slices.Contains(allowedStatuses, p.status) {
		return fmt.Errorf("cannot cancel payment with status '%s'. Allowed statuses: %v", p.status, allowedStatuses)
	}

	if strings.TrimSpace(reason) == "" {
		return errors.New("cancellation reason cannot be empty")
	}

	p.status = enum.CANCELLED
	p.updatedAt = time.Now()

	return nil
}

func (p *Payment) Pay(transactionID string) error {
	allowedStatuses := []enum.PaymentStatus{enum.FAILED, enum.PENDING}
	if !slices.Contains(allowedStatuses, p.status) {
		return fmt.Errorf("cannot cancel payment with status '%s'. Allowed statuses: %v", p.status, allowedStatuses)
	}

	p.updatedAt = time.Now()
	p.status = enum.PAID
	p.transactionID = &transactionID

	return nil
}

func (p *Payment) Refund() error {
	allowedStatuses := []enum.PaymentStatus{enum.PAID}
	if !slices.Contains(allowedStatuses, p.status) {
		return fmt.Errorf("cannot cancel payment with status '%s'. Allowed statuses: %v", p.status, allowedStatuses)
	}

	now := time.Now()

	p.status = enum.CANCELLED
	p.paidAt = nil
	p.refundedAt = &now
	p.updatedAt = now

	return nil
}

func (p *Payment) Overdue() error {
	allowedStatuses := []enum.PaymentStatus{enum.FAILED, enum.PENDING}
	if !slices.Contains(allowedStatuses, p.status) {
		return fmt.Errorf("cannot cancel payment with status '%s'. Allowed statuses: %v", p.status, allowedStatuses)
	}

	if p.dueDate != nil && time.Now().Before(*p.dueDate) {
		return fmt.Errorf("can't set a payment as overdue before his overdue date")
	}
	p.updatedAt = time.Now()
	p.status = enum.OVERDUE
	return nil
}

func (p *Payment) IsOverdue() bool {
	if p.status == enum.OVERDUE {
		return true
	}

	if p.paidAt != nil && p.paidAt.IsZero() {
		return false
	}

	if p.dueDate != nil && time.Now().Before(*p.dueDate) {
		return true
	}

	return false
}

func (p *Payment) GetID() valueobject.PaymentID {
	return p.id
}

func (p *Payment) GetAppointmentID() valueobject.AppointmentID {
	return p.appointmentID
}

func (p *Payment) GetUserID() valueobject.UserID {
	return p.userID
}

func (p *Payment) GetAmount() valueobject.Money {
	return p.amount
}

func (p *Payment) GetCurrency() string {
	return p.currency
}

func (p *Payment) SetCurrency(currency string) {
	p.currency = currency
}

func (p *Payment) GetPaymentMethod() enum.PaymentMethod {
	return p.paymentMethod
}

func (p *Payment) GetStatus() enum.PaymentStatus {
	return p.status
}

func (p *Payment) GetTransactionID() *string {
	return p.transactionID
}

func (p *Payment) GetDescription() *string {
	return p.description
}

func (p *Payment) GetDueDate() *time.Time {
	return p.dueDate
}

func (p *Payment) GetPaidAt() *time.Time {
	return p.paidAt
}

func (p *Payment) GetRefundedAt() *time.Time {
	return p.refundedAt
}

func (p *Payment) GetCreatedAt() time.Time {
	return p.createdAt
}

func (p *Payment) GetUpdatedAt() time.Time {
	return p.updatedAt
}

type PaymentBuilder struct {
	payment *Payment
}

func NewPaymentBuilder() *PaymentBuilder {
	return &PaymentBuilder{payment: &Payment{}}
}

func (pb *PaymentBuilder) WithID(id valueobject.PaymentID) *PaymentBuilder {
	pb.payment.id = id
	return pb
}

func (pb *PaymentBuilder) WithAppointmentID(appointmentID valueobject.AppointmentID) *PaymentBuilder {
	pb.payment.appointmentID = appointmentID
	return pb
}

func (pb *PaymentBuilder) WithUserID(userID valueobject.UserID) *PaymentBuilder {
	pb.payment.userID = userID
	return pb
}

func (pb *PaymentBuilder) WithAmount(amount valueobject.Money) *PaymentBuilder {
	pb.payment.amount = amount
	return pb
}

func (pb *PaymentBuilder) WithCurrency(currency string) *PaymentBuilder {
	pb.payment.currency = currency
	return pb
}

func (pb *PaymentBuilder) WithPaymentMethod(paymentMethod enum.PaymentMethod) *PaymentBuilder {
	pb.payment.paymentMethod = paymentMethod
	return pb
}

func (pb *PaymentBuilder) WithStatus(status enum.PaymentStatus) *PaymentBuilder {
	pb.payment.status = status
	return pb
}

func (pb *PaymentBuilder) WithTransactionID(transactionID *string) *PaymentBuilder {
	pb.payment.transactionID = transactionID
	return pb
}

func (pb *PaymentBuilder) WithDescription(description *string) *PaymentBuilder {
	pb.payment.description = description
	return pb
}

func (pb *PaymentBuilder) WithDueDate(dueDate *time.Time) *PaymentBuilder {
	pb.payment.dueDate = dueDate
	return pb
}

func (pb *PaymentBuilder) WithPaidAt(paidAt *time.Time) *PaymentBuilder {
	pb.payment.paidAt = paidAt
	return pb
}

func (pb *PaymentBuilder) WithRefundedAt(refundedAt *time.Time) *PaymentBuilder {
	pb.payment.refundedAt = refundedAt
	return pb
}

func (pb *PaymentBuilder) WithCreatedAt(createdAt time.Time) *PaymentBuilder {
	pb.payment.createdAt = createdAt
	return pb
}

func (pb *PaymentBuilder) WithUpdatedAt(updatedAt time.Time) *PaymentBuilder {
	pb.payment.updatedAt = updatedAt
	return pb
}

func (pb *PaymentBuilder) Build() *Payment {
	return pb.payment
}

type PaymentReport struct {
	StartDate        time.Time                  `json:"start_date"`
	EndDate          time.Time                  `json:"end_date"`
	TotalPayments    int                        `json:"total_payments"`
	TotalAmount      float64                    `json:"total_amount"`
	TotalCurrency    string                     `json:"total_currency"`
	PaidAmount       float64                    `json:"paid_amount"`
	PendingAmount    float64                    `json:"pending_amount"`
	RefundedAmount   float64                    `json:"refunded_amount"`
	OverdueAmount    float64                    `json:"overdue_amount"`
	PaymentsByMethod map[enum.PaymentMethod]any `json:"payments_by_method"`
	PaymentsByStatus map[enum.PaymentStatus]any `json:"payments_by_status"`
}
