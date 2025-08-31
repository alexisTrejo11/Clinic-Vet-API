package command

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type DeletePaymentCommand struct {
	paymentID int
}

func NewDeletePaymentCommand(paymentID int) DeletePaymentCommand {
	return DeletePaymentCommand{
		paymentID: paymentID,
	}
}

type DeletePaymentHandler interface {
	Handle(context context.Context, command DeletePaymentCommand) shared.CommandResult
}

type deletePaymentHandler struct {
	paymentRepo      repository.PaymentRepository
	paymentProccesor service.PaymentProccesorService
}

func NewDeletePaymentHandler(paymentRepo repository.PaymentRepository) DeletePaymentHandler {
	return &deletePaymentHandler{
		paymentRepo:      paymentRepo,
		paymentProccesor: service.PaymentProccesorService{},
	}
}

func (h *deletePaymentHandler) Handle(context context.Context, command DeletePaymentCommand) shared.CommandResult {
	payment, err := h.paymentRepo.GetByID(context, command.paymentID)
	if err != nil {
		return shared.FailureResult("error fetching payment", err)
	}

	if err := h.paymentProccesor.ValidateDelete(&payment); err != nil {
		return shared.FailureResult("payment cannot be deleted", err)
	}

	if err := h.paymentRepo.SoftDelete(context, command.paymentID); err != nil {
		return shared.FailureResult("error deleting payment", err)
	}

	return shared.SuccessResult(
		string(rune(command.paymentID)),
		"payment deleted successfully",
	)
}
