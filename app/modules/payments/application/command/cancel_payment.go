// Package command contains all the implemenation of the operations that persist the state of the payments module
package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type CancelPaymentCommand struct {
	paymentID valueobject.PaymentID
	ctx       context.Context
	reason    string
}

func NewCancelPaymentCommand(idInt uint, reason string) *CancelPaymentCommand {
	return &CancelPaymentCommand{
		paymentID: valueobject.NewPaymentID(idInt),
		reason:    reason,
	}
}

type CancelPaymentHandler struct {
	paymentRepo repository.PaymentRepository
}

func NewCancelPaymentHandler(paymentRepo repository.PaymentRepository) cqrs.CommandHandler {
	return &CancelPaymentHandler{paymentRepo: paymentRepo}
}

func (h *CancelPaymentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(CancelPaymentCommand)

	payment, err := h.paymentRepo.FindByID(command.ctx, command.paymentID)
	if err != nil {
		return cqrs.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Cancel(command.reason); err != nil {
		return cqrs.FailureResult("failed to cancel payment", err)
	}

	if err := h.paymentRepo.Save(command.ctx, &payment); err != nil {
		return cqrs.FailureResult("failed to save canceled payment", err)
	}

	return cqrs.SuccessResult(payment.ID().String(), "payment canceled successfully")
}
