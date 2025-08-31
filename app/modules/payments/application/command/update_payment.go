package command

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type UpdatePaymentCommand struct {
	PaymentID     int                 `json:"payment_id"`
	Amount        *valueobject.Money  `json:"amount,omitempty"`
	PaymentMethod *enum.PaymentMethod `json:"payment_method,omitempty"`
	Description   *string             `json:"description,omitempty"`
	DueDate       *time.Time          `json:"due_date,omitempty"`
}

type UpdatePaymentHandler interface {
	Handle(ctx context.Context, command UpdatePaymentCommand) shared.CommandResult
}

type updatePaymentHandler struct {
	paymentRepo repository.PaymentRepository
}

func NewUpdatePaymentHandler(paymentRepo repository.PaymentRepository) UpdatePaymentHandler {
	return &updatePaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *updatePaymentHandler) Handle(ctx context.Context, cmd UpdatePaymentCommand) shared.CommandResult {
	payment, err := h.paymentRepo.GetByID(ctx, cmd.PaymentID)
	if err != nil {
		return shared.FailureResult("error fetching payment", err)
	}

	err = payment.Update(cmd.Amount, cmd.PaymentMethod, cmd.Description, cmd.DueDate)
	if err != nil {
		return shared.FailureResult("error updating payment", err)
	}

	err = h.paymentRepo.Save(ctx, &payment)
	if err != nil {
		return shared.FailureResult("error saving payment", err)
	}

	return shared.SuccessResult(
		string(rune(cmd.PaymentID)),
		"payment updated successfully",
	)
}
