package paymentDomain

import "time"

func (p *Payment) MarkAsPaid() {
	now := time.Now()
	p.status = PAID
	p.paidAt = &now
	p.updatedAt = now
}

func (p *Payment) MarkAsOverdue() {
	p.status = OVERDUE
	p.updatedAt = time.Now()
}

func (p *Payment) markAsRefunded() {
	now := time.Now()
	p.status = REFUNDED
	p.refundedAt = &now
	p.updatedAt = now
}

func (p *Payment) MarkAsFailed() {
	p.status = FAILED
	p.updatedAt = time.Now()
}

func (p *Payment) IsOverdue() bool {
	if p.dueDate == nil || p.status == PAID || p.status == REFUNDED {
		return false
	}
	return time.Now().After(*p.dueDate)
}

func (p *Payment) CanBeRefunded() bool {
	return p.status == PAID && p.refundedAt == nil
}

func (p *Payment) GetFormattedAmount() string {
	return p.amount.FormatWithCurrency(p.currency)
}

func (p *Payment) Cancel(reason string) error {
	if p.status == PAID {
		return ErrPaymentAlreadyPaid
	}

	if p.status == CANCELLED {
		return PaymentStatusConflict(p.id, ErrPaymentAlreadyCancelled)
	}

	p.status = CANCELLED
	p.updatedAt = time.Now()

	if reason != "" {
		cancelReason := "Cancelled: " + reason
		p.description = &cancelReason
	}

	return nil
}

func (p *Payment) ValidateDelete() error {
	if p.status == PAID {
		return NewPaymentError("CANNOT_DELETE", "cannot delete paid payments", p.id, "")
	}

	if p.status == REFUNDED {
		return NewPaymentError("CANNOT_DELETE", "cannot delete refunded payments", p.id, "")
	}

	return nil
}

func (p *Payment) Process(transactionId string) error {
	if p.status == PAID {
		return ErrPaymentAlreadyPaid
	}

	if transactionId != "" {
		p.transactionId = &transactionId
	}

	p.MarkAsPaid()

	return nil
}

func (p *Payment) Refund(reason string) error {
	if !p.CanBeRefunded() {
		return ErrPaymentCannotBeRefunded
	}

	p.markAsRefunded()

	if reason != "" {
		refundReason := "Refunded: " + reason
		p.description = &refundReason
	}

	return nil
}

func (p *Payment) Update(amount *Money, paymentMethod *PaymentMethod, description *string, dueDate *time.Time) error {
	if p.status != PENDING {
		return NewPaymentError("CANNOT_UPDATE", "can only update pending payments", p.id, "")
	}

	if amount != nil && (amount.IsZero() || amount.IsNegative()) {
		return ErrInvalidAmount
	}

	if paymentMethod != nil && !paymentMethod.IsValid() {
		return ErrInvalidPaymentMethod
	}

	if description != nil && len(*description) > 255 {
		return NewPaymentError("INVALID_DESCRIPTION", "description cannot exceed 255 characters", p.id, "")
	}

	if dueDate != nil && dueDate.Before(time.Now()) {
		return NewPaymentError("INVALID_DUE_DATE", "due date cannot be in the past", p.id, "")
	}

	if paymentMethod != nil {
		p.paymentMethod = *paymentMethod
	}

	if dueDate != nil {
		p.dueDate = dueDate
	}

	p.updatedAt = time.Now()

	return nil
}
