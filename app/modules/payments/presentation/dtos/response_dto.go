package dto

import (
	"time"

	"clinic-vet-api/app/core/domain/enum"
)

// PaymentResponse represents a payment API response
// @Description Payment information response
type PaymentResponse struct {
	// Unique identifier of the payment
	// Required: true
	// Example: "123"
	ID uint `json:"id"`

	// Payment amount
	// Required: true
	// Minimum: 0.01
	// Example: 150.75
	Amount float64 `json:"amount"`

	// Currency code (ISO 4217)
	// Required: true
	// Example: "MXN"
	Currency string `json:"currency"`

	// Current status of the payment
	// Required: true
	// Enum: pending,paid,failed,refunded,cancelled
	// Example: "paid"
	Status string `json:"status"`

	// Payment method used
	// Required: true
	// Enum: cash,credit_card,debit_card,bank_transfer,digital_wallet
	// Example: "credit_card"
	Method string `json:"method"`

	// Transaction ID from payment gateway
	// Example: "txn_123456789"
	TransactionID *string `json:"transaction_id,omitempty"`

	// Payment description
	// Example: "Veterinary consultation payment"
	Description *string `json:"description,omitempty"`

	// Due date for the payment
	// Required: true
	// Example: "2024-12-31T23:59:59Z"
	DueDate *time.Time `json:"due_date"`

	// Date when payment was completed
	// Example: "2024-01-15T10:30:00Z"
	PaidAt *time.Time `json:"paid_at,omitempty"`

	// Date when payment was refunded
	// Example: "2024-01-16T14:20:00Z"
	RefundedAt *time.Time `json:"refunded_at,omitempty"`

	// Customer ID who made the payment
	// Required: true
	// Example: 456
	CustomerID uint `json:"customer_id"`

	// Appointment ID associated with the payment
	// Example: 123
	AppointmentID uint `json:"appointment_id,omitempty"`

	// Invoice number
	// Example: "INV-001"
	InvoiceID *string `json:"invoice_id,omitempty"`

	// Refund amount if applicable
	// Example: 150.75
	RefundAmount *float64 `json:"refund_amount,omitempty"`

	// Reason for payment failure
	// Example: "Insufficient funds"
	FailureReason *string `json:"failure_reason,omitempty"`

	// Indicates if payment is active
	// Required: true
	// Example: true
	IsActive bool `json:"is_active"`

	// Date when payment was created
	// Required: true
	// Example: "2024-01-15T10:00:00Z"
	CreatedAt time.Time `json:"created_at"`

	// Date when payment was last updated
	// Required: true
	// Example: "2024-01-15T10:30:00Z"
	UpdatedAt time.Time `json:"updated_at"`

	// Date when payment was deleted (soft delete)
	// Example: "2024-01-20T12:00:00Z"
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
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
