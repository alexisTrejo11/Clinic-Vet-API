// Package query contains all the implementation for retriving data for payments
package query

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/payment"
)

type PaymentResult struct {
	ID             uint
	Amount         float64
	Currency       string
	Status         string
	Method         string
	TransactionID  *string
	Description    *string
	DueDate        *time.Time
	PaidAt         *time.Time
	RefundedAt     *time.Time
	IsActive       bool
	PaidByCustomer uint
	PaidToEmployee uint
	MedSessionID   uint
	InvoiceID      *string
	RefundAmount   *float64
	FailureReason  *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	IsOverdue      bool
	IsRefundable   bool
}

func (r *paymentQueryHandler) entityToResult(payment payment.Payment) PaymentResult {

	result := &PaymentResult{
		ID:             payment.ID().Value(),
		MedSessionID:   payment.MedSessionID().Value(),
		Amount:         payment.Amount().Amount().Float64(),
		Currency:       payment.Currency(),
		Method:         payment.Method().DisplayName(),
		Status:         payment.Status().DisplayName(),
		TransactionID:  payment.TransactionID(),
		Description:    payment.Description(),
		DueDate:        payment.DueDate(),
		PaidAt:         payment.PaidAt(),
		IsActive:       payment.IsActive(),
		RefundedAt:     payment.RefundedAt(),
		IsOverdue:      payment.IsOverdue(),
		IsRefundable:   payment.IsRefundable(),
		PaidByCustomer: payment.PaidByCustomer().Value(),
		PaidToEmployee: payment.PaidToEmployee().Value(),
		InvoiceID:      payment.InvoiceID(),
		RefundAmount:   r.vaueObjMap.MoneyPtrToFloat64Ptr(payment.RefundAmount()),
		FailureReason:  payment.FailureReason(),
		CreatedAt:      payment.CreatedAt(),
		UpdatedAt:      payment.UpdatedAt(),
	}

	return *result
}
