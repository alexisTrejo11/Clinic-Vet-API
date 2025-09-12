// Package dto contains data transfer objects for payment operations
package dto

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

// CreatePaymentRequest represents the request to create a payment
// @Description Request body for creating a new payment
type CreatePaymentRequest struct {
	// Appointment ID associated with the payment
	// Required: true
	// Minimum: 1
	AppointmentID int `json:"appointment_id" validate:"required,min=1" example:"123"`

	// User ID (owner) making the payment
	// Required: true
	// Minimum: 1
	UserID int `json:"owner_id" validate:"required,min=1" example:"456"`

	// Payment amount
	// Required: true
	// Minimum: 0.01
	Amount float64 `json:"amount" validate:"required,min=0.01" example:"150.75"`

	// Currency code (3 letters)
	// Required: true
	// Length: 3
	Currency string `json:"currency" validate:"required,len=3" example:"USD"`

	// Payment method used
	// Required: true
	PaymentMethod string `json:"payment_method" validate:"required" example:"credit_card"`

	// Optional payment description
	Description *string `json:"description,omitempty" example:"Veterinary consultation payment"`

	// Optional due date for payment
	DueDate *time.Time `json:"due_date,omitempty" example:"2024-12-31T23:59:59Z"`

	// Optional transaction ID from payment gateway
	TransactionID *string `json:"transaction_id,omitempty" example:"txn_123456789"`
}

// UpdatePaymentRequest represents the request to update a payment
// @Description Request body for updating an existing payment
type UpdatePaymentRequest struct {
	// New payment amount
	// Minimum: 0.01
	Amount *float64 `json:"amount,omitempty" validate:"omitempty,min=0.01" example:"200.50"`

	// New currency code
	// Length: 3
	Currency *string `json:"currency,omitempty" validate:"omitempty,len=3" example:"EUR"`

	// New payment method
	PaymentMethod *string `json:"payment_method,omitempty" example:"paypal"`

	// New payment description
	Description *string `json:"description,omitempty" example:"Updated payment description"`

	// New due date
	DueDate *time.Time `json:"due_date,omitempty" example:"2025-01-15T23:59:59Z"`
}

// ProcessPaymentRequest represents the request to process a payment
// @Description Request body for processing a payment transaction
type ProcessPaymentRequest struct {
	// Transaction ID from payment gateway
	// Required: true
	TransactionID string `json:"transaction_id" validate:"required" example:"txn_987654321"`
}

// RefundPaymentRequest represents the request to refund a payment
// @Description Request body for refunding a processed payment
type RefundPaymentRequest struct {
	// Reason for the refund
	// Required: true
	// Min length: 1, Max length: 500
	Reason string `json:"reason" validate:"required,min=1,max=500" example:"Service not provided"`
}

// CancelPaymentRequest represents the request to cancel a payment
// @Description Request body for canceling a pending payment
type CancelPaymentRequest struct {
	// Reason for cancellation
	// Required: true
	// Min length: 1, Max length: 500
	Reason string `json:"reason" validate:"required,min=1,max=500" example:"Appointment canceled by client"`
}

// PaymentSearchRequest represents the request to search payments
// @Description Request body for searching and filtering payments
type PaymentSearchRequest struct {
	// Filter by user ID (owner)
	UserID *int `json:"owner_id,omitempty" example:"456"`

	// Filter by appointment ID
	AppointmentID *int `json:"appointment_id,omitempty" example:"123"`

	// Filter by payment status
	Status *enum.PaymentStatus `json:"status,omitempty" example:"paid"`

	// Filter by payment method
	PaymentMethod *enum.PaymentMethod `json:"payment_method,omitempty" example:"credit_card"`

	// Minimum amount filter
	MinAmount *float64 `json:"min_amount,omitempty" example:"50.00"`

	// Maximum amount filter
	MaxAmount *float64 `json:"max_amount,omitempty" example:"500.00"`

	// Filter by currency
	Currency *string `json:"currency,omitempty" example:"USD"`

	// Start date for date range filter
	StartDate *time.Time `json:"start_date,omitempty" example:"2024-01-01T00:00:00Z"`

	// End date for date range filter
	EndDate *time.Time `json:"end_date,omitempty" example:"2024-12-31T23:59:59Z"`

	// Pagination parameters
	Page page.PageInput `json:"page"`
}

// PaymentResponse represents the payment response
// @Description Payment information response
type PaymentResponse struct {
	// Payment ID
	ID uint `json:"id" example:"1"`

	// Appointment ID
	AppointmentID uint `json:"appointment_id" example:"123"`

	// User ID (owner)
	UserID uint `json:"owner_id" example:"456"`

	// Payment amount
	Amount float64 `json:"amount" example:"150.75"`

	// Currency code
	Currency string `json:"currency" example:"USD"`

	// Payment method used
	PaymentMethod enum.PaymentMethod `json:"payment_method" example:"credit_card"`

	// Payment status
	Status enum.PaymentStatus `json:"status" example:"paid"`

	// Transaction ID from payment gateway
	TransactionID *string `json:"transaction_id,omitempty" example:"txn_123456789"`

	// Payment description
	Description *string `json:"description,omitempty" example:"Veterinary consultation payment"`

	// Payment due date
	DueDate *time.Time `json:"due_date,omitempty" example:"2024-12-31T23:59:59Z"`

	// Date when payment was made
	PaidAt *time.Time `json:"paid_at,omitempty" example:"2024-01-15T14:30:00Z"`

	// Date when payment was refunded
	RefundedAt *time.Time `json:"refunded_at,omitempty" example:"2024-01-16T10:15:00Z"`

	// Whether the payment is active
	IsActive bool `json:"is_active" example:"true"`

	// Creation timestamp
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T12:00:00Z"`

	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-15T14:30:00Z"`
}

// PaymentListResponse represents a paginated list of payments
// @Description Paginated list of payments with metadata
type PaymentListResponse struct {
	// List of payments
	Items []PaymentResponse `json:"data"`

	// Pagination metadata
	Metadata page.PageMetadata `json:"metadata"`
}

// PaymentReportResponse represents a payment report response
// @Description Comprehensive payment report with statistics and summaries
type PaymentReportResponse struct {
	// Report start date
	StartDate time.Time `json:"start_date" example:"2024-01-01T00:00:00Z"`

	// Report end date
	EndDate time.Time `json:"end_date" example:"2024-12-31T23:59:59Z"`

	// Total number of payments
	TotalPayments int `json:"total_payments" example:"150"`

	// Total amount of all payments
	TotalAmount float64 `json:"total_amount" example:"22500.75"`

	// Base currency for the report
	TotalCurrency string `json:"total_currency" example:"USD"`

	// Total amount of paid payments
	PaidAmount float64 `json:"paid_amount" example:"18000.50"`

	// Total amount of pending payments
	PendingAmount float64 `json:"pending_amount" example:"3000.25"`

	// Total amount of refunded payments
	RefundedAmount float64 `json:"refunded_amount" example:"1500.00"`

	// Total amount of overdue payments
	OverdueAmount float64 `json:"overdue_amount" example:"1000.00"`

	// Payments summarized by method
	PaymentsByMethod map[enum.PaymentMethod]PaymentSummary `json:"payments_by_method"`

	// Payments summarized by status
	PaymentsByStatus map[enum.PaymentStatus]PaymentSummary `json:"payments_by_status"`
}

// PaymentSummary represents a summary of payments
// @Description Summary statistics for payments by method or status
type PaymentSummary struct {
	// Number of payments
	Count int `json:"count" example:"75"`

	// Total amount of payments
	Amount float64 `json:"amount" example:"11250.25"`
}

// MarkOverdueResponse represents the response when marking overdue payments
// @Description Response after marking overdue payments
type MarkOverdueResponse struct {
	// Number of payments updated
	UpdatedCount int `json:"updated_count" example:"5"`

	// Result message
	Message string `json:"message" example:"5 payments marked as overdue"`
}
