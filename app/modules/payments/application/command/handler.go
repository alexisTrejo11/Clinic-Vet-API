// Package command contains the command handlers for payment-related operations.
package command

import (
	"context"
	"fmt"

	"clinic-vet-api/app/core/domain/entity/payment"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/page"
)

type PaymentCommandHandler interface {
	CreatePayment(command CreatePaymentCommand) cqrs.CommandResult
	ProcessPayment(command ProcessPaymentCommand) cqrs.CommandResult
	MarkOverudePayments(command MarkOverduePaymentsCommand) cqrs.CommandResult
	UpdatePayment(command UpdatePaymentCommand) cqrs.CommandResult
	RefundPayment(command RefundPaymentCommand) cqrs.CommandResult
	CancelPayment(command CancelPaymentCommand) cqrs.CommandResult
	DeletePayment(command DeletePaymentCommand) cqrs.CommandResult
}

type paymentCommandHandler struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentCommandHandler(paymentRepository repository.PaymentRepository) PaymentCommandHandler {
	return &paymentCommandHandler{paymentRepository: paymentRepository}
}

func (h *paymentCommandHandler) ProcessPayment(command ProcessPaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(command.ctx, command.paymentID)
	if err != nil {
		return *cqrs.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Pay(command.ctx, command.transactionID); err != nil {
		return *cqrs.FailureResult("failed to process payment", err)
	}

	return *cqrs.SuccessResult(payment.ID().String(), "payment processed successfully")
}

func (h *paymentCommandHandler) CreatePayment(command CreatePaymentCommand) cqrs.CommandResult {
	payment, err := command.ToDomain()
	if err != nil {
		return *cqrs.FailureResult("failed to create payment domain", err)
	}

	if err := h.paymentRepository.Save(command.Ctx, payment); err != nil {
		return *cqrs.FailureResult("failed to create payment", err)
	}

	return *cqrs.SuccessResult(payment.ID().String(), "payment created successfully")
}

func (h *paymentCommandHandler) RefundPayment(command RefundPaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(command.ctx, command.paymentID)
	if err != nil {
		return *cqrs.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Refund(command.ctx); err != nil {
		return *cqrs.FailureResult("failed to refund payment", err)
	}

	if err := h.paymentRepository.Save(command.ctx, &payment); err != nil {
		return *cqrs.FailureResult("failed to save refunded payment", err)
	}

	return *cqrs.SuccessResult(payment.ID().String(), "payment refunded successfully")
}

func (h *paymentCommandHandler) CancelPayment(command CancelPaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(command.ctx, command.paymentID)
	if err != nil {
		return *cqrs.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Cancel(command.ctx, command.reason); err != nil {
		return *cqrs.FailureResult("failed to cancel payment", err)
	}

	if err := h.paymentRepository.Save(command.ctx, &payment); err != nil {
		return *cqrs.FailureResult("failed to save canceled payment", err)
	}

	return *cqrs.SuccessResult(payment.ID().String(), "payment canceled successfully")
}

func (h *paymentCommandHandler) DeletePayment(command DeletePaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(command.ctx, command.paymentID)
	if err != nil {
		return *cqrs.FailureResult("error fetching payment", err)
	}

	if err := h.paymentRepository.SoftDelete(command.ctx, payment.ID()); err != nil {
		return *cqrs.FailureResult("error deleting payment", err)
	}

	return *cqrs.SuccessResult(
		command.paymentID.String(),
		"payment deleted successfully",
	)
}

func (h *paymentCommandHandler) UpdatePayment(command UpdatePaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(command.ctx, command.paymentID)
	if err != nil {
		return *cqrs.FailureResult("error fetching payment", err)
	}

	err = payment.Update(command.ctx, command.amount, command.paymentMethod, command.description, command.dueDate)
	if err != nil {
		return *cqrs.FailureResult("error updating payment", err)
	}

	err = h.paymentRepository.Save(command.ctx, &payment)
	if err != nil {
		return *cqrs.FailureResult("error saving payment", err)
	}

	return *cqrs.SuccessResult(
		payment.ID().String(),
		"payment updated successfully",
	)
}

func (h *paymentCommandHandler) MarkOverudePayments(command MarkOverduePaymentsCommand) cqrs.CommandResult {
	pagination := page.PageInput{
		PageSize: 100,
		Page:     1,
	}

	var updatedCount int
	for {
		paymentsPage, err := h.paymentRepository.FindByStatus(command.context, enum.PaymentStatusOverdue, pagination)
		if err != nil {
			return *cqrs.FailureResult("failed to search payments", err)
		}

		payments := paymentsPage.Items
		if h.IsEmptyList(payments) {
			break
		}

		for _, payment := range payments {
			if err := h.UpdatePaymentOverdued(command.context, &payment); err != nil {
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

func (h *paymentCommandHandler) IsLastPage(pagination page.PageInput, totalPages int) bool {
	return pagination.Page >= totalPages
}

func (h *paymentCommandHandler) IsEmptyList(payments []payment.Payment) bool {
	return len(payments) == 0
}
