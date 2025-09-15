package dto

import (
	"time"

	"clinic-vet-api/app/core/domain/enum"
)

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
