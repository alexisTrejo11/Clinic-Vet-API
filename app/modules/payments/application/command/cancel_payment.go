// Package command contains all the implemenation of the operations that persist the state of the payments module
package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type CancelPaymentCommand struct {
	paymentID valueobject.PaymentID
	ctx       context.Context
	reason    string
}

func NewCancelPaymentCommand(idInt int, reason string) (CancelPaymentCommand, error) {
	paymentID, err := valueobject.NewPaymentID(idInt)
	if err != nil {
		return CancelPaymentCommand{}, appError.MappingError([]string{err.Error()}, "constructor", "command", "payment")
	}

	cmd := &CancelPaymentCommand{
		paymentID: paymentID,
		reason:    reason,
	}
	return *cmd, nil
}

type CancelPaymentHandler struct {
	paymentRepo repository.PaymentRepository
}

func NewCancelPaymentHandler(paymentRepo repository.PaymentRepository) cqrs.CommandHandler {
	return &CancelPaymentHandler{paymentRepo: paymentRepo}
}

func (h *CancelPaymentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(CancelPaymentCommand)

	payment, err := h.paymentRepo.GetByID(command.ctx, command.paymentID)
	if err != nil {
		return cqrs.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Cancel(command.reason); err != nil {
		return cqrs.FailureResult("failed to cancel payment", err)
	}

	if err := h.paymentRepo.Save(command.ctx, &payment); err != nil {
		return cqrs.FailureResult("failed to save canceled payment", err)
	}

	return cqrs.SuccessResult(payment.ID().String(), "payment canceled successfully")
}
