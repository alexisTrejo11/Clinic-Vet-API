package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type ProcessPaymentCommand struct {
	paymentID     valueobject.PaymentID
	transactionID string
	ctx           context.Context
}

func NewProcessPaymentCommand(idInt uint, transactionID string) (*ProcessPaymentCommand, error) {
	if idInt == 0 {
		return nil, apperror.FieldValidationError("id", "", "Payment ID cannot be zero")
	}

	if transactionID == "" {
		return nil, apperror.FieldValidationError("transaction_id", "", "Transaction ID cannot be empty")
	}

	return &ProcessPaymentCommand{
		paymentID:     valueobject.NewPaymentID(idInt),
		transactionID: transactionID,
		ctx:           context.Background(),
	}, nil
}

type ProcessPaymentHandler struct {
	paymentRepo repository.PaymentRepository
}

func NewProcessPaymentHandler(paymentRepo repository.PaymentRepository) cqrs.CommandHandler {
	return &ProcessPaymentHandler{paymentRepo: paymentRepo}
}

func (h *ProcessPaymentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(ProcessPaymentCommand)
	payment, err := h.paymentRepo.FindByID(command.ctx, command.paymentID)
	if err != nil {
		return cqrs.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Pay(command.transactionID); err != nil {
		return cqrs.FailureResult("failed to process payment", err)
	}

	if err := h.paymentRepo.Save(command.ctx, &payment); err != nil {
		return cqrs.FailureResult("failed to save processed payment", err)
	}

	return cqrs.SuccessResult(payment.ID().String(), "payment processed successfully")
}
