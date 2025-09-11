package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type DeletePaymentCommand struct {
	paymentID valueobject.PaymentID
	ctx       context.Context
}

func NewDeletePaymentCommand(ctx context.Context, idInt uint) *DeletePaymentCommand {
	return &DeletePaymentCommand{
		paymentID: valueobject.NewPaymentID(idInt),
		ctx:       ctx,
	}
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
	payment, err := h.paymentRepo.FindByID(command.ctx, command.paymentID)
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
