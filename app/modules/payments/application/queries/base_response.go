package paymentQuery

import (
	"time"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
)

type PaymentResponse struct {
	Id            int     `json:"id"`
	AppointmentId int     `json:"appointment_id"`
	UserId        int     `json:"owner_id"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`
	TransactionId *string `json:"transaction_id,omitempty"`
	Description   *string `json:"description,omitempty"`
	DueDate       *string `json:"due_date,omitempty"`
	PaidAt        *string `json:"paid_at,omitempty"`
	RefundedAt    *string `json:"refunded_at,omitempty"`
	IsActive      bool    `json:"is_active"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func NewPaymentResponse(payment *paymentDomain.Payment) PaymentResponse {
	return PaymentResponse{
		Id:            payment.GetId(),
		AppointmentId: payment.GetAppointmentId(),
		UserId:        payment.GetUserId(),
		Amount:        payment.GetAmount().ToFloat(),
		Currency:      payment.GetCurrency(),
		PaymentMethod: string(payment.GetPaymentMethod()),
		Status:        string(payment.GetStatus()),
		TransactionId: payment.GetTransactionId(),
		Description:   payment.GetDescription(),
		DueDate:       formatTime(payment.GetDueDate()),
		PaidAt:        formatTime(payment.GetPaidAt()),
		RefundedAt:    formatTime(payment.GetRefundedAt()),
		IsActive:      payment.GetIsActive(),
		CreatedAt:     payment.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt:     payment.GetUpdatedAt().Format(time.RFC3339),
	}
}

func formatTime(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format(time.RFC3339)
	return &formatted
}

func mapPaymentsToResponses(payments []paymentDomain.Payment) []PaymentResponse {
	var responses []PaymentResponse
	for _, payment := range payments {
		responses = append(responses, NewPaymentResponse(&payment))
	}
	return responses
}
