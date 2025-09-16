package payment

import (
	"fmt"
	"slices"
	"time"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	domainerr "clinic-vet-api/app/core/error"
)

func (p *Payment) MarkAsPaid(transactionID string, paidAt time.Time) error {
	if p.status != enum.PaymentStatusPending {
		return fmt.Errorf("payment is not in pending status")
	}

	p.status = enum.PaymentStatusPaid
	p.transactionID = &transactionID
	p.paidAt = &paidAt
	p.updatedAt = time.Now()

	return nil
}

func (p *Payment) RequestRefund(refundAmount valueobject.Money, reason string) error {
	if !p.IsRefundable() {
		return fmt.Errorf("payment is not refundable")
	}

	if refundAmount.Amount() > p.amount.Amount() {
		return fmt.Errorf("refund amount cannot exceed original payment amount")
	}

	p.status = enum.PaymentStatusRefunded
	p.refundAmount = &refundAmount
	now := time.Now()
	p.refundedAt = &now
	p.updatedAt = now

	return nil
}

func (p *Payment) SoftDelete() {
	p.isActive = false
	deletedAt := time.Now()
	p.deletedAt = &deletedAt
	p.updatedAt = time.Now()
}

func (p *Payment) Restore() {
	p.isActive = true
	p.deletedAt = nil
	p.updatedAt = time.Now()
}

func (p *Payment) isRefundPeriodExpired() bool {
	if p.paidAt == nil {
		return true
	}

	// Refund period: 30 days from payment date
	refundDeadline := p.paidAt.Add(30 * 24 * time.Hour)
	return time.Now().After(refundDeadline)
}

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

func (p *Payment) CanRetry() bool {
	return p.status == enum.PaymentStatusFailed || p.status == enum.PaymentStatusPending
}

func (p *Payment) AmountDue() valueobject.Money {
	if p.IsPaid() || p.IsRefunded() || p.IsCancelled() {
		return valueobject.NewMoney(0, p.amount.Currency())
	}
	return p.amount
}

// canTransitionStatus checks if the current status allows transition checking if it is in the allowedStatuses slice
func (p *Payment) canTransitionStatus(allowedStatuses []enum.PaymentStatus) bool {
	return slices.Contains(allowedStatuses, p.status)
}

func (p *Payment) RequiresPaymentProcessing() bool {
	return p.status == enum.PaymentStatusPending &&
		p.method.RequiresOnlineProcessing()
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

func (p *Payment) Update(amount *valueobject.Money, paymentMethod *enum.PaymentMethod, description *string, dueDate *time.Time) error {
	if amount != nil {
		if amount.Amount() <= 0 {
			return domainerr.NewValidationError("payment", "amount", "amount must be positive")
		}
		p.amount = *amount
	}

	if paymentMethod != nil {
		if !paymentMethod.IsValid() {
			return domainerr.NewValidationError("payment", "payment method", "invalid payment method")
		}
		p.method = *paymentMethod
	}

	if description != nil {
		if len(*description) > 500 {
			return domainerr.NewValidationError("payment", "description", "description too long")
		}
		p.description = description
	}

	if dueDate != nil {
		if dueDate.Before(time.Now()) {
			return domainerr.NewValidationError("payment", "due date", "due date cannot be in the past")
		}
		p.dueDate = dueDate
	}

	p.IncrementVersion()
	return nil
}
