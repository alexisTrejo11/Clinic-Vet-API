// Package dto contains data transfer objects for payment operations
package dto

import (
	"context"
	"time"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/specification"
	query "clinic-vet-api/app/modules/payments/application/queries"
	"clinic-vet-api/app/shared/page"
)

// CreatePaymentRequest represents the request to create a payment
// @Description Request body for creating a new payment
type CreatePaymentRequest struct {
	// Appointment ID associated with the payment
	// Required: true
	// Minimum: 1
	AppointmentID int `json:"appointment_id" validate:"required,min=1" example:"123"`

	// Customer ID making the payment
	// Required: true
	// Minimum: 1
	CustomerID int `json:"customer_id" validate:"required,min=1" example:"456"`

	// Payment amount
	// Required: true
	// Minimum: 0.01
	Amount float64 `json:"amount" validate:"required,min=0.01" example:"150.75"`

	// Currency code (3 letters)
	// Required: true
	// Length: 3
	Currency string `json:"currency" validate:"required,len=3" example:"MXN"`

	// Payment method used
	// Required: true
	// Enum: cash,credit_card,debit_card,bank_transfer,digital_wallet
	PaymentMethod string `json:"payment_method" validate:"required,oneof=cash credit_card debit_card bank_transfer digital_wallet" example:"credit_card"`

	// Payment status
	// Enum: pending,paid,failed,refunded,cancelled
	Status string `json:"status,omitempty" validate:"omitempty,oneof=pending paid failed refunded cancelled" example:"pending"`

	// Optional payment description
	Description *string `json:"description,omitempty" example:"Veterinary consultation payment"`

	// Optional due date for payment
	DueDate time.Time `json:"due_date,omitempty" example:"2024-12-31T23:59:59Z"`

	// Optional transaction ID from payment gateway
	TransactionID *string `json:"transaction_id,omitempty" example:"txn_123456789"`

	// Optional invoice ID
	InvoiceID *string `json:"invoice_id,omitempty" example:"INV-001"`
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

// GetPaymentsByDateRange represents the request to get payments within a date range
// @Description Request body for retrieving payments filtered by a date range with pagination
type PaymentsByDateRangeRequest struct {
	StartDate time.Time `form:"start_date" json:"start_date" validate:"required" example:"2024-01-01T00:00:00Z"`
	EndDate   time.Time `form:"end_date" json:"end_date" validate:"required" example:"2024-12-31T23:59:59Z"`
	page.PageInput
}

func (r *PaymentsByDateRangeRequest) SetDefaultsFieldsIfEmpty() {
	r.PageInput.SetDefaultsFieldsIfEmpty()

	if r.StartDate.IsZero() {
		now := time.Now()
		startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		r.StartDate = startOfToday
	}

	if r.EndDate.IsZero() {
		now := time.Now()
		endOfToday := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
		r.EndDate = endOfToday
	}
}

func (r *PaymentsByDateRangeRequest) ToQuery(ctx context.Context) *query.FindPaymentsByDateRangeQuery {
	return query.NewFindPaymentsByDateRangeQuery(r.StartDate, r.EndDate, r.PageInput)
}

// PaymentSearchRequest represents the request to search payments
// @Description Request body for searching and filtering payments
type PaymentSearchRequest struct {
	// Filter by user ID (customer)
	UserID *int `json:"customer_id,omitempty" example:"456"`

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

// TODO: IMPLEMENT
func (r *PaymentSearchRequest) ToQuery(ctx context.Context) (*query.FindPaymentsBySpecification, error) {
	return query.NewFindPaymentsBySpecification(ctx, specification.PaymentSpecification{}), nil
}
