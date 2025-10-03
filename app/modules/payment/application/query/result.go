// Package query contains all the implementation for retriving data for payments
package query

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type PaymentResult struct {
	ID             uint
	Amount         valueobject.Money
	Status         enum.PaymentStatus
	Method         enum.PaymentMethod
	Description    string
	TransactionID  *string
	DueDate        *time.Time
	PaidAt         *time.Time
	RefundedAt     *time.Time
	PaidByCustomer *valueobject.CustomerID
	MedSessionID   *valueobject.MedSessionID
	InvoiceID      *string
	RefundAmount   *valueobject.Money
	FailureReason  *string
	IsOverdue      bool
	IsRefundable   bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	IsActive       bool
}

func (r *paymentQueryHandler) entityToResult(payment payment.Payment) PaymentResult {
	return PaymentResult{
		ID:            payment.ID().Value(),
		MedSessionID:  payment.MedSessionID(),
		Amount:        payment.Amount(),
		Method:        payment.Method(),
		Status:        payment.Status(),
		TransactionID: payment.TransactionID(),
		Description:   payment.Description(),
		DueDate:       payment.DueDate(),
		PaidAt:        payment.PaidAt(),
		IsActive:      payment.IsActive(),
		RefundedAt:    payment.RefundedAt(),
		IsOverdue:     payment.IsOverdue(),
		IsRefundable:  payment.IsRefundable(),

		PaidByCustomer: payment.PaidByCustomer(),
		InvoiceID:      payment.InvoiceID(),
		RefundAmount:   payment.RefundAmount(),
		FailureReason:  payment.FailureReason(),
		CreatedAt:      payment.CreatedAt(),
		UpdatedAt:      payment.UpdatedAt(),
	}
}
