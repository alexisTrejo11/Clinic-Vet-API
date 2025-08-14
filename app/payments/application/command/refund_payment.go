package paymentCmd

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type RefundPaymentCommand struct {
	paymentId int
	reason    string
}

func NewRefundPaymentCommand(paymentId int, reason string) RefundPaymentCommand {
	return RefundPaymentCommand{
		paymentId: paymentId,
		reason:    reason,
	}
}

type RefundPaymentHandler interface {
	Handle(ctx context.Context, command RefundPaymentCommand) shared.CommandResult
}

type refundPaymentHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewRefundPaymentHandler(paymentRepo paymentDomain.PaymentRepository) RefundPaymentHandler {
	return &refundPaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *refundPaymentHandler) Handle(ctx context.Context, command RefundPaymentCommand) shared.CommandResult {
	if command.paymentId == 0 {
		return shared.FailureResult("payment_id is required", paymentDomain.InvalidPaymentIdErr(command.paymentId))
	}

	payment, err := h.paymentRepo.GetById(ctx, command.paymentId)
	if err != nil {
		return shared.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Refund(command.reason); err != nil {
		return shared.FailureResult("failed to refund payment", err)
	}

	if err := h.paymentRepo.Save(ctx, payment); err != nil {
		return shared.FailureResult("failed to save refunded payment", err)
	}

	return shared.SuccesResult(string(payment.GetId()), "payment refunded successfully")
}
