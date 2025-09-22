// payment.go (métodos actualizados)
package payment

import (
	"context"
	"slices"
	"time"

	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/log"

	"go.uber.org/zap"
)

func (p *Payment) MarkAsPaid(ctx context.Context, transactionID string, paidAt time.Time) error {
	operation := "MarkAsPaid"

	if p.status != enum.PaymentStatusPending {
		return CannotMarkAsPaidError(ctx, p.status, operation)
	}

	if transactionID == "" {
		return TransactionIDRequiredError(ctx, operation)
	}

	if paidAt.After(time.Now()) {
		return PaidDateFutureError(ctx, operation)
	}

	p.status = enum.PaymentStatusPaid
	p.transactionID = &transactionID
	p.paidAt = &paidAt
	p.updatedAt = time.Now()
	p.IncrementVersion()

	return nil
}

func (p *Payment) RequestRefund(ctx context.Context, refundAmount valueobject.Money, reason string) error {
	operation := "RequestRefund"

	if !p.IsRefundable() {
		return CannotRequestRefundError(ctx, p.status, operation)
	}

	if p.isRefundPeriodExpired() {
		return RefundPeriodExpiredError(ctx, *p.paidAt, operation)
	}

	diff, err := refundAmount.Subtract(p.amount)
	if err != nil {
		return err
	}
	if diff.IsPositive() {
		return RefundAmountExceededError(ctx, refundAmount, p.amount, operation)
	}

	if reason == "" {
		return FailureReasonEmptyError(ctx, operation)
	}

	p.status = enum.PaymentStatusRefunded
	p.refundAmount = &refundAmount
	now := time.Now()
	p.refundedAt = &now
	p.failureReason = &reason
	p.updatedAt = now
	p.IncrementVersion()

	return nil
}

func (p *Payment) SoftDelete(ctx context.Context) {
	operation := "SoftDelete"

	log.App.Info("Soft deleting payment",
		zap.String("payment_id", p.ID().String()),
		zap.String("operation", operation))

	p.isActive = false
	deletedAt := time.Now()
	p.deletedAt = &deletedAt
	p.updatedAt = time.Now()
	p.IncrementVersion()
}

func (p *Payment) Restore(ctx context.Context) {
	operation := "Restore"

	log.App.Info("Restoring payment",
		zap.String("payment_id", p.ID().String()),
		zap.String("operation", operation))

	p.isActive = true
	p.deletedAt = nil
	p.updatedAt = time.Now()
	p.IncrementVersion()
}

func (p *Payment) isRefundPeriodExpired() bool {
	if p.paidAt == nil {
		return true
	}

	// Refund period: 30 days from payment date
	refundDeadline := p.paidAt.Add(30 * 24 * time.Hour)
	return time.Now().After(refundDeadline)
}

func (p *Payment) Cancel(ctx context.Context, reason string) error {
	operation := "CancelPayment"

	allowedStatuses := []enum.PaymentStatus{
		enum.PaymentStatusPending,
		enum.PaymentStatusPaid,
		enum.PaymentStatusFailed,
	}

	if !p.canTransitionStatus(allowedStatuses) {
		return CannotCancelError(ctx, p.status, operation)
	}

	if reason == "" {
		return CancellationReasonRequiredError(ctx, operation)
	}

	p.status = enum.PaymentStatusCancelled
	p.failureReason = &reason
	p.IncrementVersion()
	p.updatedAt = time.Now()

	return nil
}

func (p *Payment) Pay(ctx context.Context, transactionID string) error {
	operation := "PayPayment"

	allowedStatuses := []enum.PaymentStatus{
		enum.PaymentStatusFailed,
		enum.PaymentStatusPending,
	}

	if !p.canTransitionStatus(allowedStatuses) {
		return CannotProcessError(ctx, p.status, operation)
	}

	if transactionID == "" {
		return TransactionIDRequiredError(ctx, operation)
	}

	now := time.Now()
	p.status = enum.PaymentStatusPaid
	p.transactionID = &transactionID
	p.paidAt = &now
	p.IncrementVersion()
	p.updatedAt = now

	return nil
}

func (p *Payment) Refund(ctx context.Context) error {
	operation := "RefundPayment"

	allowedStatuses := []enum.PaymentStatus{enum.PaymentStatusPaid}

	if !p.canTransitionStatus(allowedStatuses) {
		return CannotRefundError(ctx, p.status, operation)
	}

	if p.isRefundPeriodExpired() {
		return RefundPeriodExpiredError(ctx, *p.paidAt, operation)
	}

	now := time.Now()
	p.status = enum.PaymentStatusRefunded
	p.refundedAt = &now
	p.IncrementVersion()
	p.updatedAt = now

	return nil
}

func (p *Payment) MarkAsOverdue(ctx context.Context) error {
	operation := "MarkAsOverdue"

	allowedStatuses := []enum.PaymentStatus{
		enum.PaymentStatusFailed,
		enum.PaymentStatusPending,
	}

	if !p.canTransitionStatus(allowedStatuses) {
		return CannotMarkOverdueError(ctx, p.status, operation)
	}

	if p.dueDate != nil && time.Now().Before(*p.dueDate) {
		return DueDateNotReachedError(ctx, *p.dueDate, operation)
	}

	p.status = enum.PaymentStatusOverdue
	p.IncrementVersion()
	p.updatedAt = time.Now()

	return nil
}

func (p *Payment) MarkAsFailed(ctx context.Context) error {
	operation := "MarkAsFailed"

	allowedStatuses := []enum.PaymentStatus{enum.PaymentStatusPending}

	if !p.canTransitionStatus(allowedStatuses) {
		return CannotMarkFailedError(ctx, p.status, operation)
	}

	p.status = enum.PaymentStatusFailed
	p.IncrementVersion()
	p.updatedAt = time.Now()

	return nil
}

func (p *Payment) Update(ctx context.Context, amount *valueobject.Money, paymentMethod *enum.PaymentMethod, description *string, dueDate *time.Time) error {
	operation := "UpdatePayment"

	if amount != nil {
		if amount.Amount().IsZero() || amount.Amount().IsNegative() {
			return AmountInvalidError(ctx, *amount, operation)
		}
		p.amount = *amount
	}

	if paymentMethod != nil {
		if !paymentMethod.IsValid() {
			return MethodInvalidError(ctx, *paymentMethod, operation)
		}
		p.method = *paymentMethod
	}

	if description != nil {
		if len(*description) > 500 {
			return DescriptionTooLongError(ctx, len(*description), operation)
		}
		p.description = description
	}

	if dueDate != nil {
		if dueDate.Before(time.Now()) {
			return DueDatePastError(ctx, operation)
		}
		p.dueDate = dueDate
	}

	p.IncrementVersion()
	p.updatedAt = time.Now()

	return nil
}

// Métodos auxiliares (se mantienen igual pero con mejor documentación)
func (p *Payment) CanRetry() bool {
	return p.status == enum.PaymentStatusFailed || p.status == enum.PaymentStatusPending
}

func (p *Payment) AmountDue() valueobject.Money {
	if p.IsPaid() || p.IsRefunded() || p.IsCancelled() {
		return valueobject.NewMoney(p.Amount().Amount(), p.amount.Currency())
	}
	return p.amount
}

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
		!p.isRefundPeriodExpired()
}

func (p *Payment) DaysOverdue() int {
	if !p.IsOverdue() || p.dueDate == nil {
		return 0
	}
	return int(time.Since(*p.dueDate).Hours() / 24)
}
