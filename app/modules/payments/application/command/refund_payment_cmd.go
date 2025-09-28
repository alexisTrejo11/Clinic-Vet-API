package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type RefundPaymentCommand struct {
	paymentID valueobject.PaymentID
	reason    string
}

func NewRefundPaymentCommand(paymentID uint, reason string) *RefundPaymentCommand {
	return &RefundPaymentCommand{
		paymentID: valueobject.NewPaymentID(paymentID),
		reason:    reason,
	}
}

func (h *paymentCommandHandler) RefundPayment(ctx context.Context, cmd RefundPaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(ctx, cmd.paymentID)
	if err != nil {
		return *cqrs.FailureResult(ErrFailedRetrievePayment, err)
	}

	if err := payment.Refund(ctx); err != nil {
		return *cqrs.FailureResult(ErrFailedRefundPayment, err)
	}

	if err := h.paymentRepository.Save(ctx, &payment); err != nil {
		return *cqrs.FailureResult(ErrFailedSaveRefundedPayment, err)
	}

	return *cqrs.SuccessResult(MsgPaymentRefunded)
}
