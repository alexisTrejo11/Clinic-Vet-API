package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type CancelPaymentCommand struct {
	paymentID valueobject.PaymentID
	reason    string
}

func NewCancelPaymentCommand(idInt uint, reason string) CancelPaymentCommand {
	return CancelPaymentCommand{
		paymentID: valueobject.NewPaymentID(idInt),
		reason:    reason,
	}
}

func (h *paymentCommandHandler) CancelPayment(ctx context.Context, cmd CancelPaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(ctx, cmd.paymentID)
	if err != nil {
		return cqrs.FailureResult(ErrFailedRetrievePayment, err)
	}

	if err := payment.Cancel(ctx, cmd.reason); err != nil {
		return cqrs.FailureResult(ErrFailedCancelPayment, err)
	}

	if err := h.paymentRepository.Save(ctx, &payment); err != nil {
		return cqrs.FailureResult(ErrFailedSaveCanceledPayment, err)
	}

	return cqrs.SuccessResult(MsgPaymentCanceled)
}
