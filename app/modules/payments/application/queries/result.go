// Package query contains all the implementation for retriving data for payments
package query

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/payment"
)

type PaymentResult struct {
	ID               uint
	Amount           float64
	Currency         string
	Status           string
	Method           string
	TransactionID    *string
	Description      *string
	DueDate          *time.Time
	PaidAt           *time.Time
	RefundedAt       *time.Time
	IsActive         bool
	PaidFromCustomer uint
	PaidToEmployee   uint
	AppointmentID    uint
	InvoiceID        *string
	RefundAmount     *float64
	FailureReason    *string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	IsOverdue        bool
	IsRefundable     bool
}

func entityToResult(payment payment.Payment) PaymentResult {
	result := &PaymentResult{
		ID:               payment.ID().Value(),
		AppointmentID:    payment.AppointmentID().Value(),
		Amount:           payment.Amount().Amount().Float64(),
		Currency:         payment.Currency(),
		Method:           payment.Method().DisplayName(),
		Status:           string(payment.Status()),
		TransactionID:    payment.TransactionID(),
		Description:      payment.Description(),
		DueDate:          payment.DueDate(),
		PaidAt:           payment.PaidAt(),
		IsActive:         payment.IsActive(),
		RefundedAt:       payment.RefundedAt(),
		IsOverdue:        payment.IsOverdue(),
		IsRefundable:     payment.IsRefundable(),
		PaidFromCustomer: payment.PaidFromCustomer().Value(),
		PaidToEmployee:   payment.PaidToEmployee().Value(),
		InvoiceID: func() *string {
			if payment.InvoiceID() != nil {
				invoiceID := payment.InvoiceID()
				return invoiceID
			}
			return nil
		}(),
		RefundAmount: func() *float64 {
			if payment.RefundAmount() != nil {
				refundAmount := payment.RefundAmount().Amount().Float64()
				return &refundAmount
			}
			return nil
		}(),
		FailureReason: payment.FailureReason(),
		CreatedAt:     payment.CreatedAt(),
		UpdatedAt:     payment.UpdatedAt(),
	}

	return *result
}
