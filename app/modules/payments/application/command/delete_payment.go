package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type DeletePaymentCommand struct {
	paymentID valueobject.PaymentID
	ctx       context.Context
}

func NewDeletePaymentCommand(idInt int) (DeletePaymentCommand, error) {
	paymentID, err := valueobject.NewPaymentID(idInt)
	if err != nil {
		return DeletePaymentCommand{}, apperror.MappingError([]string{err.Error()}, "constructor", "command", "payment")
	}

	cmd := &DeletePaymentCommand{
		paymentID: paymentID,
		ctx:       context.Background(),
	}

	return *cmd, nil
}

type DeletePaymentHandler struct {
	paymentRepo repository.PaymentRepository
}

func NewDeletePaymentHandler(paymentRepo repository.PaymentRepository) cqrs.CommandHandler {
	return &DeletePaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *DeletePaymentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(DeletePaymentCommand)
	payment, err := h.paymentRepo.GetByID(command.ctx, command.paymentID)
	if err != nil {
		return cqrs.FailureResult("error fetching payment", err)
	}

	if err := h.paymentRepo.SoftDelete(command.ctx, payment.ID()); err != nil {
		return cqrs.FailureResult("error deleting payment", err)
	}

	return cqrs.SuccessResult(
		command.paymentID.String(),
		"payment deleted successfully",
	)
}
