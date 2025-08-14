package paymentDomain

import (
	"time"
)

type Payment struct {
	id            int
	appointmentId int
	userId        int
	amount        Money
	currency      string
	paymentMethod PaymentMethod
	status        PaymentStatus
	transactionId *string
	description   *string
	dueDate       *time.Time
	paidAt        *time.Time
	refundedAt    *time.Time
	isActive      bool
	createdAt     time.Time
	updatedAt     time.Time
}

func (p *Payment) GetId() int {
	return p.id
}

func (p *Payment) SetId(id int) {
	p.id = id
}

func (p *Payment) GetAppointmentId() int {
	return p.appointmentId
}

func (p *Payment) SetAppointmentId(appointmentId int) {
	p.appointmentId = appointmentId
}

func (p *Payment) GetUserId() int {
	return p.userId
}

func (p *Payment) SetUserId(userId int) {
	p.userId = userId
}

func (p *Payment) GetAmount() Money {
	return p.amount
}

func (p *Payment) SetAmount(amount Money) {
	p.amount = amount
}

func (p *Payment) GetCurrency() string {
	return p.currency
}

func (p *Payment) SetCurrency(currency string) {
	p.currency = currency
}

func (p *Payment) GetPaymentMethod() PaymentMethod {
	return p.paymentMethod
}

func (p *Payment) SetPaymentMethod(paymentMethod PaymentMethod) {
	p.paymentMethod = paymentMethod
}

func (p *Payment) GetStatus() PaymentStatus {
	return p.status
}

func (p *Payment) SetStatus(status PaymentStatus) {
	p.status = status
}

func (p *Payment) GetTransactionId() *string {
	return p.transactionId
}

func (p *Payment) SetTransactionId(transactionId *string) {
	p.transactionId = transactionId
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
