package paymentCmd

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type DeletePaymentCommand struct {
	paymentId int
}

func NewDeletePaymentCommand(paymentId int) DeletePaymentCommand {
	return DeletePaymentCommand{
		paymentId: paymentId,
	}
}

type DeletePaymentHandler interface {
	Handle(context context.Context, command DeletePaymentCommand) shared.CommandResult
}

type deletePaymentHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewDeletePaymentHandler(paymentRepo paymentDomain.PaymentRepository) DeletePaymentHandler {
	return &deletePaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *deletePaymentHandler) Handle(context context.Context, command DeletePaymentCommand) shared.CommandResult {
	if command.paymentId == 0 {
		return shared.FailureResult("payment ID is zero", paymentDomain.InvalidPaymentIdErr(command.paymentId))
	}

	payment, err := h.paymentRepo.GetById(context, command.paymentId)
	if err != nil {
		return shared.FailureResult("error fetching payment", err)
	}

	if err := payment.ValidateDelete(); err != nil {
		return shared.FailureResult("payment cannot be deleted", err)
	}

	if err := h.paymentRepo.SoftDelete(context, command.paymentId); err != nil {
		return shared.FailureResult("error deleting payment", err)
	}

	return shared.SuccesResult(
		string(rune(command.paymentId)),
		"payment deleted successfully",
	)
}
