package command

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type ProcessPaymentCommand struct {
	paymentID     int
	transactionID string
}

func NewProcessPaymentCommand(paymentID int, transactionID string) ProcessPaymentCommand {
	return ProcessPaymentCommand{
		paymentID:     paymentID,
		transactionID: transactionID,
	}
}

type ProcessPaymentHandler interface {
	Handle(ctx context.Context, command ProcessPaymentCommand) shared.CommandResult
}

type processPaymentHandler struct {
	paymentRepo      repository.PaymentRepository
	paymentProccesor service.PaymentProccesorService
}

func NewProcessPaymentHandler(paymentRepo repository.PaymentRepository) ProcessPaymentHandler {
	return &processPaymentHandler{
		paymentRepo:      paymentRepo,
		paymentProccesor: service.PaymentProccesorService{},
	}
}

func (h *processPaymentHandler) Handle(ctx context.Context, command ProcessPaymentCommand) shared.CommandResult {
	payment, err := h.paymentRepo.GetByID(ctx, command.paymentID)
	if err != nil {
		return shared.FailureResult("failed to retrieve payment", err)
	}

	if err := h.paymentProccesor.Process(&payment, command.transactionID); err != nil {
		return shared.FailureResult("failed to process payment", err)
	}

	if err := h.paymentRepo.Save(ctx, &payment); err != nil {
		return shared.FailureResult("failed to save processed payment", err)
	}

	return shared.SuccessResult(string(payment.GetID()), "payment processed successfully")
}
