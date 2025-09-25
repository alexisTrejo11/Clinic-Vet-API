package payment

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type Payment struct {
	base.Entity[valueobject.PaymentID]
	amount         valueobject.Money
	status         enum.PaymentStatus
	method         enum.PaymentMethod
	transactionID  *string
	description    *string
	dueDate        *time.Time
	paidAt         *time.Time
	refundedAt     *time.Time
	isActive       bool
	paidByCustomer valueobject.CustomerID
	paidToEmployee valueobject.EmployeeID
	medSessionID   *valueobject.MedSessionID
	invoiceID      *string
	refundAmount   *valueobject.Money
	failureReason  *string
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      *time.Time
}

func (p *Payment) ID() valueobject.PaymentID {
	return p.Entity.ID()
}

func (p *Payment) Amount() valueobject.Money {
	return p.amount
}

func (p *Payment) Currency() string {
	return p.amount.Currency()
}

func (p *Payment) Status() enum.PaymentStatus {
	return p.status
}

func (p *Payment) Method() enum.PaymentMethod {
	return p.method
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

func (p *Payment) IsActive() bool {
	return p.isActive
}

func (p *Payment) PaidByCustomer() valueobject.CustomerID {
	return p.paidByCustomer
}

func (p *Payment) PaidToEmployee() valueobject.EmployeeID {
	return p.paidToEmployee
}

func (p *Payment) MedSessionID() *valueobject.MedSessionID {
	return p.medSessionID
}

func (p *Payment) InvoiceID() *string {
	return p.invoiceID
}

func (p *Payment) RefundAmount() *valueobject.Money {
	return p.refundAmount
}

func (p *Payment) FailureReason() *string {
	return p.failureReason
}

func (p *Payment) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Payment) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Payment) DeletedAt() *time.Time {
	return p.deletedAt
}

func (p *Payment) IsOverdue() bool {
	if p.status == enum.PaymentStatusOverdue {
		return true
	}

	if p.paidAt != nil || p.status.IsFinal() {
		return false
	}

	now := time.Now()
	return p.dueDate.Before(now) &&
		p.status != enum.PaymentStatusPaid &&
		p.status != enum.PaymentStatusRefunded &&
		p.status != enum.PaymentStatusCancelled
}

func (p *Payment) IsRefundable() bool {
	return p.status == enum.PaymentStatusPaid &&
		p.paidAt != nil &&
		!p.isRefundPeriodExpired()
}

func (p *Payment) IsPaid() bool {
	return p.status == enum.PaymentStatusPaid
}

func (p *Payment) IsPending() bool {
	return p.status == enum.PaymentStatusPending
}

func (p *Payment) IsFailed() bool {
	return p.status == enum.PaymentStatusFailed
}

func (p *Payment) IsRefunded() bool {
	return p.status == enum.PaymentStatusRefunded
}

func (p *Payment) IsCancelled() bool {
	return p.status == enum.PaymentStatusCancelled
}

func (p *Payment) SetTransactionID(transactionID string) {
	p.transactionID = &transactionID
}

func (p *Payment) SetDescription(description string) {
	p.description = &description
}

func (p *Payment) SetPaidAt(paidAt *time.Time) {
	p.paidAt = paidAt
}

func (p *Payment) SetRefundedAt(refundedAt *time.Time) {
	p.refundedAt = refundedAt
}

func (p *Payment) SetMedSessionID(medSessionID *valueobject.MedSessionID) {
	p.medSessionID = medSessionID
}

func (p *Payment) SetInvoiceID(invoiceID string) {
	p.invoiceID = &invoiceID
}

func (p *Payment) SetRefundAmount(refundAmount *valueobject.Money) {
	p.refundAmount = refundAmount
}

func (p *Payment) SetFailureReason(failureReason string) {
	p.failureReason = &failureReason
}

func (p *Payment) SetDeletedAt(deletedAt *time.Time) {
	p.deletedAt = deletedAt
}
