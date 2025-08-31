package dto

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

// CreatePaymentRequest represents the request to create a payment
type CreatePaymentRequest struct {
	AppointmentID int                `json:"appointment_id" validate:"required,min=1"`
	UserID        int                `json:"owner_id" validate:"required,min=1"`
	Amount        float64            `json:"amount" validate:"required,min=0.01"`
	Currency      string             `json:"currency" validate:"required,len=3"`
	PaymentMethod enum.PaymentMethod `json:"payment_method" validate:"required"`
	Description   *string            `json:"description,omitempty"`
	DueDate       *time.Time         `json:"due_date,omitempty"`
	TransactionID *string            `json:"transaction_id,omitempty"`
}

// UpdatePaymentRequest represents the request to update a payment
type UpdatePaymentRequest struct {
	Amount        *float64            `json:"amount,omitempty" validate:"omitempty,min=0.01"`
	Currency      *string             `json:"currency,omitempty" validate:"omitempty,len=3"`
	PaymentMethod *enum.PaymentMethod `json:"payment_method,omitempty"`
	Description   *string             `json:"description,omitempty"`
	DueDate       *time.Time          `json:"due_date,omitempty"`
}

// ProcessPaymentRequest represents the request to process a payment
type ProcessPaymentRequest struct {
	TransactionID string `json:"transaction_id" validate:"required"`
}

// RefundPaymentRequest represents the request to refund a payment
type RefundPaymentRequest struct {
	Reason string `json:"reason" validate:"required,min=1,max=500"`
}

// CancelPaymentRequest represents the request to cancel a payment
type CancelPaymentRequest struct {
	Reason string `json:"reason" validate:"required,min=1,max=500"`
}

// PaymentSearchRequest represents the request to search payments
type PaymentSearchRequest struct {
	UserID        *int                `json:"owner_id,omitempty"`
	AppointmentID *int                `json:"appointment_id,omitempty"`
	Status        *enum.PaymentStatus `json:"status,omitempty"`
	PaymentMethod *enum.PaymentMethod `json:"payment_method,omitempty"`
	MinAmount     *float64            `json:"min_amount,omitempty"`
	MaxAmount     *float64            `json:"max_amount,omitempty"`
	Currency      *string             `json:"currency,omitempty"`
	StartDate     *time.Time          `json:"start_date,omitempty"`
	EndDate       *time.Time          `json:"end_date,omitempty"`
	Page          page.PageData       `json:"page"`
}

// PaymentResponse represents the payment response
type PaymentResponse struct {
	ID            int                `json:"id"`
	AppointmentID int                `json:"appointment_id"`
	UserID        int                `json:"owner_id"`
	Amount        float64            `json:"amount"`
	Currency      string             `json:"currency"`
	PaymentMethod enum.PaymentMethod `json:"payment_method"`
	Status        enum.PaymentStatus `json:"status"`
	TransactionID *string            `json:"transaction_id,omitempty"`
	Description   *string            `json:"description,omitempty"`
	DueDate       *time.Time         `json:"due_date,omitempty"`
	PaidAt        *time.Time         `json:"paid_at,omitempty"`
	RefundedAt    *time.Time         `json:"refunded_at,omitempty"`
	IsActive      bool               `json:"is_active"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
}

// PaymentListResponse represents a paginated list of payments
type PaymentListResponse struct {
	Data     []PaymentResponse `json:"data"`
	Metadata page.PageMetadata `json:"metadata"`
}

// PaymentReportResponse represents a payment report response
type PaymentReportResponse struct {
	StartDate        time.Time                             `json:"start_date"`
	EndDate          time.Time                             `json:"end_date"`
	TotalPayments    int                                   `json:"total_payments"`
	TotalAmount      float64                               `json:"total_amount"`
	TotalCurrency    string                                `json:"total_currency"`
	PaidAmount       float64                               `json:"paid_amount"`
	PendingAmount    float64                               `json:"pending_amount"`
	RefundedAmount   float64                               `json:"refunded_amount"`
	OverdueAmount    float64                               `json:"overdue_amount"`
	PaymentsByMethod map[enum.PaymentMethod]PaymentSummary `json:"payments_by_method"`
	PaymentsByStatus map[enum.PaymentStatus]PaymentSummary `json:"payments_by_status"`
}

// PaymentSummary represents a summary of payments
type PaymentSummary struct {
	Count  int     `json:"count"`
	Amount float64 `json:"amount"`
}

// MarkOverdueResponse represents the response when marking overdue payments
type MarkOverdueResponse struct {
	UpdatedCount int    `json:"updated_count"`
	Message      string `json:"message"`
}

type PaymentReport struct {
	StartDate        time.Time                             `json:"start_date"`
	EndDate          time.Time                             `json:"end_date"`
	TotalPayments    int                                   `json:"total_payments"`
	TotalAmount      float64                               `json:"total_amount"`
	TotalCurrency    string                                `json:"total_currency"`
	PaidAmount       float64                               `json:"paid_amount"`
	PendingAmount    float64                               `json:"pending_amount"`
	RefundedAmount   float64                               `json:"refunded_amount"`
	OverdueAmount    float64                               `json:"overdue_amount"`
	PaymentsByMethod map[enum.PaymentMethod]PaymentSummary `json:"payments_by_method"`
	PaymentsByStatus map[enum.PaymentStatus]PaymentSummary `json:"payments_by_status"`
}
