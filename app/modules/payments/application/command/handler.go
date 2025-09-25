// Package command contains the command handlers for payment-related operations.
package command

import (
	"context"
	"fmt"

	"clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/page"
)

type PaymentCommandHandler interface {
	CreatePayment(ctx context.Context, cmd CreatePaymentCommand) cqrs.CommandResult
	ProcessPayment(ctx context.Context, cmd ProcessPaymentCommand) cqrs.CommandResult
	MarkOverduePayments(ctx context.Context, cmd MarkOverduePaymentsCommand) cqrs.CommandResult
	UpdatePayment(ctx context.Context, cmd UpdatePaymentCommand) cqrs.CommandResult
	RefundPayment(ctx context.Context, cmd RefundPaymentCommand) cqrs.CommandResult
	CancelPayment(ctx context.Context, cmd CancelPaymentCommand) cqrs.CommandResult
	DeletePayment(ctx context.Context, cmd DeletePaymentCommand) cqrs.CommandResult
}

type paymentCommandHandler struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentCommandHandler(paymentRepository repository.PaymentRepository) PaymentCommandHandler {
	return &paymentCommandHandler{paymentRepository: paymentRepository}
}

func (h *paymentCommandHandler) ProcessPayment(ctx context.Context, cmd ProcessPaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(ctx, cmd.paymentID)
	if err != nil {
		return *cqrs.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Pay(ctx, cmd.transactionID); err != nil {
		return *cqrs.FailureResult("failed to process payment", err)
	}

	return *cqrs.SuccessResult(payment.ID().String(), "payment processed successfully")
}

func (h *paymentCommandHandler) CreatePayment(ctx context.Context, cmd CreatePaymentCommand) cqrs.CommandResult {
	payment, err := cmd.ToDomain(ctx)
	if err != nil {
		return *cqrs.FailureResult("failed to create payment domain", err)
	}

	if err := h.paymentRepository.Save(ctx, payment); err != nil {
		return *cqrs.FailureResult("failed to create payment", err)
	}

	return *cqrs.SuccessResult(payment.ID().String(), "payment created successfully")
}

func (h *paymentCommandHandler) RefundPayment(ctx context.Context, cmd RefundPaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(ctx, cmd.paymentID)
	if err != nil {
		return *cqrs.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Refund(ctx); err != nil {
		return *cqrs.FailureResult("failed to refund payment", err)
	}

	if err := h.paymentRepository.Save(ctx, &payment); err != nil {
		return *cqrs.FailureResult("failed to save refunded payment", err)
	}

	return *cqrs.SuccessResult(payment.ID().String(), "payment refunded successfully")
}

func (h *paymentCommandHandler) CancelPayment(ctx context.Context, cmd CancelPaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(ctx, cmd.paymentID)
	if err != nil {
		return *cqrs.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Cancel(ctx, cmd.reason); err != nil {
		return *cqrs.FailureResult("failed to cancel payment", err)
	}

	if err := h.paymentRepository.Save(ctx, &payment); err != nil {
		return *cqrs.FailureResult("failed to save canceled payment", err)
	}

	return *cqrs.SuccessResult(payment.ID().String(), "payment canceled successfully")
}

func (h *paymentCommandHandler) DeletePayment(ctx context.Context, cmd DeletePaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(ctx, cmd.paymentID)
	if err != nil {
		return *cqrs.FailureResult("error fetching payment", err)
	}

	if err := h.paymentRepository.SoftDelete(ctx, payment.ID()); err != nil {
		return *cqrs.FailureResult("error deleting payment", err)
	}

	return *cqrs.SuccessResult(
		cmd.paymentID.String(),
		"payment deleted successfully",
	)
}

func (h *paymentCommandHandler) UpdatePayment(ctx context.Context, cmd UpdatePaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(ctx, cmd.paymentID)
	if err != nil {
		return *cqrs.FailureResult("error fetching payment", err)
	}

	err = payment.Update(ctx, cmd.amount, cmd.paymentMethod, cmd.description, cmd.dueDate)
	if err != nil {
		return *cqrs.FailureResult("error updating payment", err)
	}

	err = h.paymentRepository.Save(ctx, &payment)
	if err != nil {
		return *cqrs.FailureResult("error saving payment", err)
	}

	return *cqrs.SuccessResult(payment.ID().String(), "payment updated successfully")
}

func (h *paymentCommandHandler) MarkOverduePayments(ctx context.Context, cmd MarkOverduePaymentsCommand) cqrs.CommandResult {
	pagination := page.PaginationRequest{
		PageSize: 100,
		Page:     1,
	}

	var updatedCount int
	for {
		paymentsPage, err := h.paymentRepository.FindByStatus(ctx, enum.PaymentStatusOverdue, pagination)
		if err != nil {
			return *cqrs.FailureResult("failed to search payments", err)
		}

		payments := paymentsPage.Items
		if h.IsEmptyList(payments) {
			break
		}

		for _, payment := range payments {
			if err := h.UpdatePaymentOverdued(ctx, &payment); err != nil {
				fmt.Printf("Error updating payment %d: %v\n", payment.ID(), err)
				continue
			}
			updatedCount++
		}

		pagination.Page++

		if h.IsLastPage(pagination, paymentsPage.Metadata.TotalPages) {
			break
		}
	}

	return *cqrs.SuccessResult("", fmt.Sprintf("Updated %d overdue payments", updatedCount))
}

func (h *paymentCommandHandler) UpdatePaymentOverdued(ctx context.Context, payment *payment.Payment) error {
	if err := payment.MarkAsOverdue(ctx); err != nil {
		return err
	}

	if err := h.paymentRepository.Save(ctx, payment); err != nil {
		return err
	}

	return nil
}

func (h *paymentCommandHandler) IsLastPage(pagination page.PaginationRequest, totalPages int) bool {
	return pagination.Page >= totalPages
}

func (h *paymentCommandHandler) IsEmptyList(payments []payment.Payment) bool {
	return len(payments) == 0
}
