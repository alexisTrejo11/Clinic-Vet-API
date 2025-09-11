package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type RefundPaymentCommand struct {
	paymentID valueobject.PaymentID
	reason    string
	ctx       context.Context
}

func NewRefundPaymentCommand(paymentID uint, reason string) (*RefundPaymentCommand, error) {
	if paymentID == 0 {
		return nil, apperror.FieldValidationError("id", "0", "Payment ID cannot be zero")
	}

	paymentIDVO := valueobject.NewPaymentID(paymentID)

	return &RefundPaymentCommand{
		paymentID: paymentIDVO,
		reason:    reason,
		ctx:       context.Background(),
	}, nil
}

type RefundPaymentHandler struct {
	paymentRepo repository.PaymentRepository
}

func NewRefundPaymentHandler(paymentRepo repository.PaymentRepository) cqrs.CommandHandler {
	return &RefundPaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *RefundPaymentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(RefundPaymentCommand)

	payment, err := h.paymentRepo.FindByID(command.ctx, command.paymentID)
	if err != nil {
		return cqrs.FailureResult("failed to retrieve payment", err)
	}

	if err := payment.Refund(); err != nil {
		return cqrs.FailureResult("failed to refund payment", err)
	}

	if err := h.paymentRepo.Save(command.ctx, &payment); err != nil {
		return cqrs.FailureResult("failed to save refunded payment", err)
	}

	return cqrs.SuccessResult(payment.ID().String(), "payment refunded successfully")
}
