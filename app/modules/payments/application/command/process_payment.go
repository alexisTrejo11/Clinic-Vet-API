package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type ProcessPaymentCommand struct {
	paymentID     valueobject.PaymentID
	transactionID string
	ctx           context.Context
}

func NewProcessPaymentCommand(idInt int, transactionID string) (ProcessPaymentCommand, error) {
	paymentID, err := valueobject.NewPaymentID(idInt)
	if err != nil {
		return ProcessPaymentCommand{}, apperror.MappingError([]string{err.Error()}, "constructor", "command", "payment")
	}

	cmd := &ProcessPaymentCommand{
		paymentID:     paymentID,
		transactionID: transactionID,
		ctx:           context.Background(),
	}

	return *cmd, nil
}

type ProcessPaymentHandler struct {
	paymentRepo repository.PaymentRepository
}

func NewProcessPaymentHandler(paymentRepo repository.PaymentRepository) cqrs.CommandHandler {
	return &ProcessPaymentHandler{paymentRepo: paymentRepo}
}

func (h *ProcessPaymentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(ProcessPaymentCommand)
	payment, err := h.paymentRepo.GetByID(command.ctx, command.paymentID)
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
