package entity

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

type Payment struct {
	id            int
	appointmentId int
	userId        int
	amount        valueobject.Money
	currency      string
	paymentMethod enum.PaymentMethod
	status        enum.PaymentStatus
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

type PaymentBuilder struct {
	payment *Payment
}

func NewPaymentBuilder() *PaymentBuilder {
	return &PaymentBuilder{payment: &Payment{}}
}

func (pb *PaymentBuilder) WithId(id int) *PaymentBuilder {
	pb.payment.id = id
	return pb
}

func (pb *PaymentBuilder) WithAppointmentId(appointmentId int) *PaymentBuilder {
	pb.payment.appointmentId = appointmentId
	return pb
}

func (pb *PaymentBuilder) WithUserId(userId int) *PaymentBuilder {
	pb.payment.userId = userId
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

func (pb *PaymentBuilder) WithTransactionId(transactionId *string) *PaymentBuilder {
	pb.payment.transactionId = transactionId
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
