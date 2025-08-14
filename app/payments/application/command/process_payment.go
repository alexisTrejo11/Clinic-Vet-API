package paymentCmd

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type ProcessPaymentCommand struct {
	paymentId     int
	transactionId string
}

func NewProcessPaymentCommand(paymentId int, transactionId string) ProcessPaymentCommand {
	return ProcessPaymentCommand{
		paymentId:     paymentId,
		transactionId: transactionId,
	}
}

type ProcessPaymentHandler interface {
	Handle(ctx context.Context, command ProcessPaymentCommand) shared.CommandResult
}

type processPaymentHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewProcessPaymentHandler(paymentRepo paymentDomain.PaymentRepository) ProcessPaymentHandler {
	return &processPaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *processPaymentHandler) Handle(ctx context.Context, command ProcessPaymentCommand) shared.CommandResult {
	if command.paymentId == 0 {
		return shared.FailureResult("payment_id is required", paymentDomain.InvalidPaymentIdErr(command.paymentId))
	}

	payment, err := h.paymentRepo.GetById(ctx, command.paymentId)
	if err != nil {
		return shared.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Process(command.transactionId); err != nil {
		return shared.FailureResult("failed to process payment", err)
	}

	if err := h.paymentRepo.Save(ctx, payment); err != nil {
		return shared.FailureResult("failed to save processed payment", err)
	}

	return shared.SuccesResult(string(payment.GetId()), "payment processed successfully")
}
