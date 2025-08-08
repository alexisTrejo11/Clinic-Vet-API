package paymentDomain

import (
	"time"
)

type Payment struct {
	Id            int           `json:"id" db:"id"`
	AppointmentId int           `json:"appointment_id" db:"appointment_id"`
	OwnerId       int           `json:"owner_id" db:"owner_id"`
	Amount        Money         `json:"amount" db:"amount"`
	Currency      string        `json:"currency" db:"currency"`
	PaymentMethod PaymentMethod `json:"payment_method" db:"payment_method"`
	Status        PaymentStatus `json:"status" db:"status"`
	TransactionId *string       `json:"transaction_id,omitempty" db:"transaction_id"`
	Description   *string       `json:"description,omitempty" db:"description"`
	DueDate       *time.Time    `json:"due_date,omitempty" db:"due_date"`
	PaidAt        *time.Time    `json:"paid_at,omitempty" db:"paid_at"`
	RefundedAt    *time.Time    `json:"refunded_at,omitempty" db:"refunded_at"`
	IsActive      bool          `json:"is_active" db:"is_active"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at" db:"updated_at"`
}

func (p *Payment) MarkAsPaid() {
	now := time.Now()
	p.Status = PAID
	p.PaidAt = &now
	p.UpdatedAt = now
}

func (p *Payment) MarkAsOverdue() {
	p.Status = OVERDUE
	p.UpdatedAt = time.Now()
}

func (p *Payment) markAsRefunded() {
	now := time.Now()
	p.Status = REFUNDED
	p.RefundedAt = &now
	p.UpdatedAt = now
}

func (p *Payment) MarkAsFailed() {
	p.Status = FAILED
	p.UpdatedAt = time.Now()
}

func (p *Payment) IsOverdue() bool {
	if p.DueDate == nil || p.Status == PAID || p.Status == REFUNDED {
		return false
	}
	return time.Now().After(*p.DueDate)
}

func (p *Payment) CanBeRefunded() bool {
	return p.Status == PAID && p.RefundedAt == nil
}

func (p *Payment) GetFormattedAmount() string {
	return p.Amount.FormatWithCurrency(p.Currency)
}

func (p *Payment) Cancel(reason string) error {
	if p.Status == PAID {
		return ErrPaymentAlreadyPaid
	}

	if p.Status == CANCELLED {
		return PaymentStatusConflict(p.Id, ErrPaymentAlreadyCancelled)
	}

	p.Status = CANCELLED
	p.UpdatedAt = time.Now()

	if reason != "" {
		cancelReason := "Cancelled: " + reason
		p.Description = &cancelReason
	}

	return nil
}

func (p *Payment) ValidateDelete() error {
	if p.Status == PAID {
		return NewPaymentError("CANNOT_DELETE", "cannot delete paid payments", p.Id, "")
	}

	if p.Status == REFUNDED {
		return NewPaymentError("CANNOT_DELETE", "cannot delete refunded payments", p.Id, "")
	}

	return nil
}

func (p *Payment) Proccess(transactionId string) error {
	if p.Status == PAID {
		return ErrPaymentAlreadyPaid
	}

	if transactionId != "" {
		p.TransactionId = &transactionId
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
		p.Description = &refundReason
	}

	return nil
}

func (p *Payment) Update(amount *Money, paymentMethod *PaymentMethod, description *string, dueDate *time.Time) error {
	if p.Status != PENDING {
		return NewPaymentError("CANNOT_UPDATE", "can only update pending payments", p.Id, "")
	}

	if amount != nil && (amount.IsZero() || amount.IsNegative()) {
		return ErrInvalidAmount
	}

	if paymentMethod != nil && !paymentMethod.IsValid() {
		return ErrInvalidPaymentMethod
	}

	if description != nil && len(*description) > 255 {
		return NewPaymentError("INVALID_DESCRIPTION", "description cannot exceed 255 characters", p.Id, "")
	}

	if dueDate != nil && dueDate.Before(time.Now()) {
		return NewPaymentError("INVALID_DUE_DATE", "due date cannot be in the past", p.Id, "")
	}

	if paymentMethod != nil {
		p.PaymentMethod = *paymentMethod
	}

	if dueDate != nil {
		p.DueDate = dueDate
	}

	p.UpdatedAt = time.Now()

	return nil
}
