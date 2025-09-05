package command

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type RefundPaymentCommand struct {
	paymentID int
	reason    string
	ctx       context.Context
}

func NewRefundPaymentCommand(paymentID int, reason string) (RefundPaymentCommand, error) {
	cmd := &RefundPaymentCommand{
		paymentID: paymentID,
		reason:    reason,
		ctx:       context.Background(),
	}

	return *cmd, nil
}

type RefundPaymentHandler struct {
	paymentRepo repository.PaymentRepository
}

func NewRefundPaymentHandler(paymentRepo repository.PaymentRepository) cqrs.CommandHandler {
	return &RefundPaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *RefundPaymentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(RefundPaymentCommand)

	payment, err := h.paymentRepo.GetByID(command.ctx, command.paymentID)
	if err != nil {
		return cqrs.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Refund(); err != nil {
		return cqrs.FailureResult("failed to refund payment", err)
	}

	if err := h.paymentRepo.Save(command.ctx, &payment); err != nil {
		return cqrs.FailureResult("failed to save refunded payment", err)
	}

	return cqrs.SuccessResult(payment.GetID().String(), "payment refunded successfully")
}
