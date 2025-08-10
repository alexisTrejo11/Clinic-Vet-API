package paymentCmd

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type CancelPaymentCommand struct {
	paymentId int
	reason    string
}

func NewCancelPaymentCommand(paymentId int, reason string) CancelPaymentCommand {
	return CancelPaymentCommand{
		paymentId: paymentId,
		reason:    reason,
	}
}

type CancelPaymentHandler interface {
	Handle(ctx context.Context, command CancelPaymentCommand) shared.CommandResult
}

type cancelPaymentHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewCancelPaymentHandler(paymentRepo paymentDomain.PaymentRepository) CancelPaymentHandler {
	return &cancelPaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *cancelPaymentHandler) Handle(ctx context.Context, command CancelPaymentCommand) shared.CommandResult {
	if command.paymentId == 0 {
		return shared.FailureResult(
			"payment_id is required",
			paymentDomain.InvalidPaymentIdErr(command.paymentId),
		)
	}

	payment, err := h.paymentRepo.GetById(ctx, command.paymentId)
	if err != nil {
		return shared.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Cancel(command.reason); err != nil {
		return shared.FailureResult("failed to cancel payment", err)
	}

	if err := h.paymentRepo.Save(ctx, payment); err != nil {
		return shared.FailureResult("failed to save canceled payment", err)
	}

	return shared.SuccesResult(string(rune(payment.Id)), "payment canceled successfully")
}
