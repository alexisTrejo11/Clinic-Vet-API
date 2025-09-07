// Package query contains all the implementation for retriving data for payments
package query

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/payment"
)

type PaymentResponse struct {
	ID            int     `json:"id"`
	AppointmentID int     `json:"appointment_id"`
	UserID        int     `json:"owner_id"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`
	TransactionID *string `json:"transaction_id,omitempty"`
	Description   *string `json:"description,omitempty"`
	DueDate       *string `json:"due_date,omitempty"`
	PaidAt        *string `json:"paid_at,omitempty"`
	RefundedAt    *string `json:"refunded_at,omitempty"`
	IsActive      bool    `json:"is_active"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func NewPaymentResponse(payment *payment.Payment) PaymentResponse {
	return PaymentResponse{
		ID:            payment.ID().Value(),
		AppointmentID: payment.AppointmentID().Value(),
		UserID:        payment.UserID().Value(),
		Amount:        payment.Amount().ToFloat(),
		Currency:      payment.Currency(),
		PaymentMethod: string(payment.PaymentMethod()),
		Status:        string(payment.Status()),
		TransactionID: payment.TransactionID(),
		Description:   payment.Description(),
		DueDate:       formatTime(payment.DueDate()),
		PaidAt:        formatTime(payment.PaidAt()),
		RefundedAt:    formatTime(payment.RefundedAt()),
		CreatedAt:     payment.CreatedAt().Format(time.RFC3339),
		UpdatedAt:     payment.UpdatedAt().Format(time.RFC3339),
	}
}

func formatTime(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format(time.RFC3339)
	return &formatted
}

func mapPaymentsToResponses(payments []payment.Payment) []PaymentResponse {
	var responses []PaymentResponse
	for _, payment := range payments {
		responses = append(responses, NewPaymentResponse(&payment))
	}
	return responses
}
