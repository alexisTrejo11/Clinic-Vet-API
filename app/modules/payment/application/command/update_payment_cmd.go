package command

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
	"time"
)

type UpdatePaymentCommand struct {
	paymentID     valueobject.PaymentID
	amount        *valueobject.Money
	paymentMethod *enum.PaymentMethod
	description   *string
	dueDate       *time.Time
}

func NewUpdatePaymentCommand(id uint, amount *valueobject.Money, paymentMethod *string, description *string, dueDate *time.Time) UpdatePaymentCommand {
	paymentID := valueobject.NewPaymentID(id)

	var paymentMethodObj *enum.PaymentMethod
	if paymentMethod != nil {
		pm := enum.PaymentMethod(*paymentMethod)
		paymentMethodObj = &pm

	}

	return UpdatePaymentCommand{
		paymentID:     paymentID,
		amount:        amount,
		paymentMethod: paymentMethodObj,
		description:   description,
		dueDate:       dueDate,
	}
}

func (h *paymentCommandHandler) UpdatePayment(ctx context.Context, cmd UpdatePaymentCommand) cqrs.CommandResult {
	payment, err := h.paymentRepository.FindByID(ctx, cmd.paymentID)
	if err != nil {
		return *cqrs.FailureResult(ErrFetchingPayment, err)
	}

	err = payment.Update(ctx, cmd.amount, cmd.paymentMethod, cmd.description, cmd.dueDate)
	if err != nil {
		return *cqrs.FailureResult(ErrUpdatingPayment, err)
	}

	err = h.paymentRepository.Save(ctx, &payment)
	if err != nil {
		return *cqrs.FailureResult(ErrSavingPayment, err)
	}

	return *cqrs.SuccessResult(MsgPaymentUpdated)
}
