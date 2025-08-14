package paymentDomain

import "time"

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

func (pb *PaymentBuilder) WithAmount(amount Money) *PaymentBuilder {
	pb.payment.amount = amount
	return pb
}

func (pb *PaymentBuilder) WithCurrency(currency string) *PaymentBuilder {
	pb.payment.currency = currency
	return pb
}

func (pb *PaymentBuilder) WithPaymentMethod(paymentMethod PaymentMethod) *PaymentBuilder {
	pb.payment.paymentMethod = paymentMethod
	return pb
}

func (pb *PaymentBuilder) WithStatus(status PaymentStatus) *PaymentBuilder {
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
