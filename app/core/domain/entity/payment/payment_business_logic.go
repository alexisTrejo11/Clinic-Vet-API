package payment

import (
	"slices"
	"time"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	domainerr "clinic-vet-api/app/core/error"
)

func (p *Payment) Cancel(reason string) error {
	allowedStatuses := []enum.PaymentStatus{
		enum.PaymentStatusPending,
		enum.PaymentStatusPaid,
		enum.PaymentStatusFailed,
	}

	if !p.canTransitionStatus(allowedStatuses) {
		return domainerr.NewBusinessRuleError("payment", "cancel", "payment cannot be cancelled in current status")
	}

	if reason == "" {
		return domainerr.NewValidationError("payment", "reason", "cancellation reason is required")
	}

	p.status = enum.PaymentStatusCancelled
	p.IncrementVersion()
	return nil
}

func (p *Payment) Pay(transactionID string) error {
	allowedStatuses := []enum.PaymentStatus{
		enum.PaymentStatusFailed,
		enum.PaymentStatusPending,
	}

	if !p.canTransitionStatus(allowedStatuses) {
		return domainerr.NewBusinessRuleError("payment", "pay", "payment cannot be paid in current status")
	}

	if transactionID == "" {
		return domainerr.NewValidationError("payment", "transaction ID", "transaction ID is required")
	}

	now := time.Now()
	p.status = enum.PaymentStatusPaid
	p.transactionID = &transactionID
	p.paidAt = &now
	p.IncrementVersion()

	return nil
}

func (p *Payment) Refund() error {
	allowedStatuses := []enum.PaymentStatus{enum.PaymentStatusPaid}

	if !p.canTransitionStatus(allowedStatuses) {
		return domainerr.NewBusinessRuleError("payment", "refund", "payment cannot be refunded in current status")
	}

	now := time.Now()
	p.status = enum.PaymentStatusRefunded
	p.refundedAt = &now
	p.IncrementVersion()

	return nil
}

func (p *Payment) MarkAsOverdue() error {
	allowedStatuses := []enum.PaymentStatus{
		enum.PaymentStatusFailed,
		enum.PaymentStatusPending,
	}

	if !p.canTransitionStatus(allowedStatuses) {
		return domainerr.NewBusinessRuleError("payment", "overdue", "payment cannot be marked as overdue in current status")
	}

	if p.dueDate != nil && time.Now().Before(*p.dueDate) {
		return domainerr.NewBusinessRuleError("payment", "overdue", "cannot mark as overdue before due date")
	}

	p.status = enum.PaymentStatusOverdue
	p.IncrementVersion()
	return nil
}

func (p *Payment) MarkAsFailed() error {
	allowedStatuses := []enum.PaymentStatus{enum.PaymentStatusPending}

	if !p.canTransitionStatus(allowedStatuses) {
		return domainerr.NewBusinessRuleError("payment", "failed", "payment cannot be marked as failed in current status")
	}

	p.status = enum.PaymentStatusFailed
	p.IncrementVersion()
	return nil
}

func (p *Payment) IsOverdue() bool {
	if p.status == enum.PaymentStatusOverdue {
		return true
	}

	if p.paidAt != nil || p.status.IsFinal() {
		return false
	}

	if p.dueDate != nil && time.Now().After(*p.dueDate) {
		return true
	}

	return false
}

func (p *Payment) IsPaid() bool {
	return p.status == enum.PaymentStatusPaid
}

func (p *Payment) IsRefunded() bool {
	return p.status == enum.PaymentStatusRefunded
}

func (p *Payment) IsCancelled() bool {
	return p.status == enum.PaymentStatusCancelled
}

func (p *Payment) CanRetry() bool {
	return p.status == enum.PaymentStatusFailed || p.status == enum.PaymentStatusPending
}

func (p *Payment) AmountDue() valueobject.Money {
	if p.IsPaid() || p.IsRefunded() || p.IsCancelled() {
		return valueobject.NewMoney(0, p.currency)
	}
	return p.amount
}

// canTransitionStatus checks if the current status allows transition checking if it is in the allowedStatuses slice
func (p *Payment) canTransitionStatus(allowedStatuses []enum.PaymentStatus) bool {
	return slices.Contains(allowedStatuses, p.status)
}

func (p *Payment) RequiresPaymentProcessing() bool {
	return p.status == enum.PaymentStatusPending &&
		p.paymentMethod.RequiresOnlineProcessing()
}

func (p *Payment) CanBeRefunded() bool {
	return p.status == enum.PaymentStatusPaid &&
		p.paidAt != nil &&
		p.paidAt.After(time.Now().AddDate(0, -6, 0)) // Within 6 months
}

func (p *Payment) DaysOverdue() int {
	if !p.IsOverdue() || p.dueDate == nil {
		return 0
	}
	return int(time.Since(*p.dueDate).Hours() / 24)
}
