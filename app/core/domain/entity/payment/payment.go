// Package payment defines the Payment entity and its behaviors.
package payment

import (
	"time"

	"clinic-vet-api/app/core/domain/entity/base"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	domainerr "clinic-vet-api/app/core/error"
)

type Payment struct {
	base.Entity[valueobject.PaymentID]
	appointmentID valueobject.AppointmentID
	userID        valueobject.UserID
	amount        valueobject.Money
	currency      string
	paymentMethod enum.PaymentMethod
	status        enum.PaymentStatus
	transactionID *string
	description   *string
	dueDate       *time.Time
	paidAt        *time.Time
	refundedAt    *time.Time
}

func (p *Payment) ID() valueobject.PaymentID {
	return p.Entity.ID()
}

func (p *Payment) AppointmentID() valueobject.AppointmentID {
	return p.appointmentID
}

func (p *Payment) UserID() valueobject.UserID {
	return p.userID
}

func (p *Payment) Amount() valueobject.Money {
	return p.amount
}

func (p *Payment) Currency() string {
	return p.currency
}

func (p *Payment) PaymentMethod() enum.PaymentMethod {
	return p.paymentMethod
}

func (p *Payment) Status() enum.PaymentStatus {
	return p.status
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
		p.paymentMethod = *paymentMethod
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
