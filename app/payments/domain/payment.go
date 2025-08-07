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

func (p *Payment) MarkAsRefunded() {
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
