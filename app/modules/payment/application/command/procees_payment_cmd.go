package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type ProcessPaymentCommand struct {
	paymentID     valueobject.PaymentID
	transactionID string
}

func NewProcessPaymentCommand(idInt uint, transactionID string) ProcessPaymentCommand {
	return ProcessPaymentCommand{
		paymentID:     valueobject.NewPaymentID(idInt),
		transactionID: transactionID,
	}
}

func (h *paymentCommandHandler) ProcessPayment(ctx context.Context, cmd ProcessPaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(ctx, cmd.paymentID)
	if err != nil {
		return cqrs.FailureResult(ErrFailedRetrievePayment, err)
	}

	if err := payment.Pay(ctx, cmd.transactionID); err != nil {
		return cqrs.FailureResult(ErrFailedProcessPayment, err)
	}

	return cqrs.SuccessResult(MsgPaymentProcessed)
}
