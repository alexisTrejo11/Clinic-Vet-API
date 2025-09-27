package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type DeletePaymentCommand struct {
	paymentID valueobject.PaymentID
}

func NewDeletePaymentCommand(idInt uint) *DeletePaymentCommand {
	return &DeletePaymentCommand{
		paymentID: valueobject.NewPaymentID(idInt),
	}
}

func (h *paymentCommandHandler) DeletePayment(ctx context.Context, cmd DeletePaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(ctx, cmd.paymentID)
	if err != nil {
		return *cqrs.FailureResult(ErrFetchingPayment, err)
	}

	if err := h.paymentRepository.SoftDelete(ctx, payment.ID()); err != nil {
		return *cqrs.FailureResult(ErrDeletingPayment, err)
	}

	return *cqrs.SuccessResult(
		cmd.paymentID.String(),
		MsgPaymentDeleted,
	)
}
