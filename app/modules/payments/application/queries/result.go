// Package query contains all the implementation for retriving data for payments
package query

import (
	"time"

	"clinic-vet-api/app/core/domain/entity/payment"
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

func entityToResult(payment *payment.Payment) PaymentResult {
	return PaymentResult{
		ID:            payment.ID().Value(),
		AppointmentID: payment.AppointmentID().Value(),
		Amount:        payment.Amount().ToFloat(),
		Currency:      payment.Currency(),
		Method:        payment.Method().DisplayName(),
		Status:        string(payment.Status()),
		TransactionID: payment.TransactionID(),
		Description:   payment.Description(),
		DueDate:       payment.DueDate(),
		PaidAt:        payment.PaidAt(),
		RefundedAt:    payment.RefundedAt(),
		CreatedAt:     payment.CreatedAt(),
		UpdatedAt:     payment.UpdatedAt(),
	}
}

func entityToResults(payments []payment.Payment) []PaymentResult {
	var responses []PaymentResult
	for _, payment := range payments {
		responses = append(responses, entityToResult(&payment))
	}
	return responses
}
