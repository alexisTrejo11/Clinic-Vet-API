package entity

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

type Payment struct {
	id            int
	appointmentID int
	userID        int
	amount        valueobject.Money
	currency      string
	paymentMethod enum.PaymentMethod
	status        enum.PaymentStatus
	transactionID *string
	description   *string
	dueDate       *time.Time
	paidAt        *time.Time
	refundedAt    *time.Time
	isActive      bool
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

func (p *Payment) GetID() int {
	return p.id
}

func (p *Payment) SetID(id int) {
	p.id = id
}

func (p *Payment) GetAppointmentID() int {
	return p.appointmentID
}

func (p *Payment) SetAppointmentID(appointmentID int) {
	p.appointmentID = appointmentID
}

func (p *Payment) GetUserID() int {
	return p.userID
}

func (p *Payment) SetUserID(userID int) {
	p.userID = userID
}

func (p *Payment) GetAmount() valueobject.Money {
	return p.amount
}

func (p *Payment) SetAmount(amount valueobject.Money) {
	p.amount = amount
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

func (p *Payment) SetPaymentMethod(paymentMethod enum.PaymentMethod) {
	p.paymentMethod = paymentMethod
}

func (p *Payment) GetStatus() enum.PaymentStatus {
	return p.status
}

func (p *Payment) SetStatus(status enum.PaymentStatus) {
	p.status = status
}

func (p *Payment) GetTransactionID() *string {
	return p.transactionID
}

func (p *Payment) SetTransactionID(transactionID *string) {
	p.transactionID = transactionID
}

func (p *Payment) GetDescription() *string {
	return p.description
}

func (p *Payment) SetDescription(description *string) {
	p.description = description
}

func (p *Payment) GetDueDate() *time.Time {
	return p.dueDate
}

func (p *Payment) SetDueDate(dueDate *time.Time) {
	p.dueDate = dueDate
}

func (p *Payment) GetPaidAt() *time.Time {
	return p.paidAt
}

func (p *Payment) SetPaidAt(paidAt *time.Time) {
	p.paidAt = paidAt
}

func (p *Payment) GetRefundedAt() *time.Time {
	return p.refundedAt
}

func (p *Payment) SetRefundedAt(refundedAt *time.Time) {
	p.refundedAt = refundedAt
}

func (p *Payment) GetIsActive() bool {
	return p.isActive
}

func (p *Payment) SetIsActive(isActive bool) {
	p.isActive = isActive
}

func (p *Payment) GetCreatedAt() time.Time {
	return p.createdAt
}

func (p *Payment) SetCreatedAt(createdAt time.Time) {
	p.createdAt = createdAt
}

func (p *Payment) GetUpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Payment) SetUpdatedAt(updatedAt time.Time) {
	p.updatedAt = updatedAt
}

type PaymentBuilder struct {
	payment *Payment
}

func NewPaymentBuilder() *PaymentBuilder {
	return &PaymentBuilder{payment: &Payment{}}
}

func (pb *PaymentBuilder) WithID(id int) *PaymentBuilder {
	pb.payment.id = id
	return pb
}

func (pb *PaymentBuilder) WithAppointmentID(appointmentID int) *PaymentBuilder {
	pb.payment.appointmentID = appointmentID
	return pb
}

func (pb *PaymentBuilder) WithUserID(userID int) *PaymentBuilder {
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

func (pb *PaymentBuilder) WithIsActive(isActive bool) *PaymentBuilder {
	pb.payment.isActive = isActive
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
