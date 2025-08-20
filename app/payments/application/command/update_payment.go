package paymentCmd

import (
	"context"
	"time"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type UpdatePaymentCommand struct {
	PaymentId     int                          `json:"payment_id"`
	Amount        *paymentDomain.Money         `json:"amount,omitempty"`
	PaymentMethod *paymentDomain.PaymentMethod `json:"payment_method,omitempty"`
	Description   *string                      `json:"description,omitempty"`
	DueDate       *time.Time                   `json:"due_date,omitempty"`
}

type UpdatePaymentHandler interface {
	Handle(ctx context.Context, command UpdatePaymentCommand) shared.CommandResult
}

type updatePaymentHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewUpdatePaymentHandler(paymentRepo paymentDomain.PaymentRepository) UpdatePaymentHandler {
	return &updatePaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *updatePaymentHandler) Handle(ctx context.Context, cmd UpdatePaymentCommand) shared.CommandResult {
	if cmd.PaymentId == 0 {
		return shared.FailureResult("payment ID is zero", paymentDomain.InvalidPaymentIdErr(cmd.PaymentId))
	}
	payment, err := h.paymentRepo.GetById(ctx, cmd.PaymentId)
	if err != nil {
		return shared.FailureResult("error fetching payment", err)
	}

	err = payment.Update(cmd.Amount, cmd.PaymentMethod, cmd.Description, cmd.DueDate)
	if err != nil {
		return shared.FailureResult("error updating payment", err)
	}

	err = h.paymentRepo.Save(ctx, payment)
	if err != nil {
		return shared.FailureResult("error saving payment", err)
	}

	return shared.SuccessResult(
		string(rune(cmd.PaymentId)),
		"payment updated successfully",
	)
}
