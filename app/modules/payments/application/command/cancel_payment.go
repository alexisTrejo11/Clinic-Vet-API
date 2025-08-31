package command

import (
	"context"
	"errors"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type CancelPaymentCommand struct {
	paymentID int
	reason    string
}

func NewCancelPaymentCommand(paymentID int, reason string) CancelPaymentCommand {
	return CancelPaymentCommand{
		paymentID: paymentID,
		reason:    reason,
	}
}

type CancelPaymentHandler interface {
	Handle(ctx context.Context, command CancelPaymentCommand) shared.CommandResult
}

type cancelPaymentHandler struct {
	paymentRepo      repository.PaymentRepository
	paymentProccesor service.PaymentProccesorService
}

func NewCancelPaymentHandler(paymentRepo repository.PaymentRepository) CancelPaymentHandler {
	return &cancelPaymentHandler{
		paymentRepo:      paymentRepo,
		paymentProccesor: service.PaymentProccesorService{},
	}
}

func (h *cancelPaymentHandler) Handle(ctx context.Context, command CancelPaymentCommand) shared.CommandResult {
	if command.paymentID == 0 {
		return shared.FailureResult(
			"payment_id is required",
			errors.New("payment id can't be 0"),
		)
	}

	payment, err := h.paymentRepo.GetByID(ctx, command.paymentID)
	if err != nil {
		return shared.FailureResult("failed to retrieve payment", err)
	}

	if err := h.paymentProccesor.Cancel(&payment); err != nil {
		return shared.FailureResult("failed to cancel payment", err)
	}

	if err := h.paymentRepo.Save(ctx, &payment); err != nil {
		return shared.FailureResult("failed to save canceled payment", err)
	}

	return shared.SuccessResult(string(rune(payment.GetID())), "payment canceled successfully")
}
