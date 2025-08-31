package command

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type RefundPaymentCommand struct {
	paymentID int
	reason    string
}

func NewRefundPaymentCommand(paymentID int, reason string) RefundPaymentCommand {
	return RefundPaymentCommand{
		paymentID: paymentID,
		reason:    reason,
	}
}

type RefundPaymentHandler interface {
	Handle(ctx context.Context, command RefundPaymentCommand) shared.CommandResult
}

type refundPaymentHandler struct {
	paymentRepo      repository.PaymentRepository
	paymentProccesor service.PaymentProccesorService
}

func NewRefundPaymentHandler(paymentRepo repository.PaymentRepository) RefundPaymentHandler {
	return &refundPaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *refundPaymentHandler) Handle(ctx context.Context, command RefundPaymentCommand) shared.CommandResult {
	payment, err := h.paymentRepo.GetByID(ctx, command.paymentID)
	if err != nil {
		return shared.FailureResult("failed to retrieve payment", err)
	}

	if err := h.paymentProccesor.Refund(&payment); err != nil {
		return shared.FailureResult("failed to refund payment", err)
	}

	if err := h.paymentRepo.Save(ctx, &payment); err != nil {
		return shared.FailureResult("failed to save refunded payment", err)
	}

	return shared.SuccessResult(string(payment.GetID()), "payment refunded successfully")
}
