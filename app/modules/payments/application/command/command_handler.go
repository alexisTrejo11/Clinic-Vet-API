// Package command contains the command handlers for payment-related operations.
package command

import (
	"context"

	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/cqrs"
)

// PaymentCommandHandler defines the contract for handling payment-related commands
// This interface encapsulates all business operations that modify payment state
type PaymentCommandHandler interface {
	// CreatePayment handles the creation of a new payment record
	// Validates payment data and persists it to the repository
	CreatePayment(ctx context.Context, cmd CreatePaymentCommand) cqrs.CommandResult

	// ProcessPayment executes the payment processing workflow
	// Retrieves the payment, processes it with the provided transaction ID, and updates its status
	ProcessPayment(ctx context.Context, cmd ProcessPaymentCommand) cqrs.CommandResult

	// MarkOverduePayments identifies and marks overdue payments in batches
	// Processes payments in paginated manner to handle large datasets efficiently
	MarkOverduePayments(ctx context.Context, cmd MarkOverduePaymentsCommand) cqrs.CommandResult

	// UpdatePayment modifies existing payment information
	// Allows updating amount, payment method, description, and due date
	UpdatePayment(ctx context.Context, cmd UpdatePaymentCommand) cqrs.CommandResult

	// RefundPayment handles the refund process for a payment
	// Changes payment status to refunded and persists the changes
	RefundPayment(ctx context.Context, cmd RefundPaymentCommand) cqrs.CommandResult

	// CancelPayment cancels a payment with a specified reason
	// Updates payment status and records cancellation reason
	CancelPayment(ctx context.Context, cmd CancelPaymentCommand) cqrs.CommandResult

	// DeletePayment performs a soft delete of a payment record
	// Maintains data integrity while removing the payment from active use
	DeletePayment(ctx context.Context, cmd DeletePaymentCommand) cqrs.CommandResult
}

// paymentCommandHandler implements the PaymentCommandHandler interface
// It encapsulates the business logic for payment command operations
type paymentCommandHandler struct {
	paymentRepository repository.PaymentRepository
}

// NewPaymentCommandHandler creates a new instance of PaymentCommandHandler
// Parameters:
//   - paymentRepository: Repository implementation for data persistence
//
// Returns:
//   - PaymentCommandHandler: Configured command handler instance
func NewPaymentCommandHandler(paymentRepository repository.PaymentRepository) PaymentCommandHandler {
	return &paymentCommandHandler{paymentRepository: paymentRepository}
}
