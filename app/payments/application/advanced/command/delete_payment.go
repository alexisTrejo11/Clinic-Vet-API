package paymentCmd

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
)

type DeletePaymentCommand struct {
	PaymentId int             `json:"payment_id"`
	CTX       context.Context `json:"-"`
}

type DeletePaymentHandler interface {
	Handle(command DeletePaymentCommand) error
}

type deletePaymentHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewDeletePaymentHandler(paymentRepo paymentDomain.PaymentRepository) DeletePaymentHandler {
	return &deletePaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *deletePaymentHandler) Handle(command DeletePaymentCommand) error {
	if command.PaymentId == 0 {
		return paymentDomain.InvalidPaymentIdErr(command.PaymentId)
	}

	payment, err := h.paymentRepo.GetById(command.CTX, command.PaymentId)
	if err != nil {
		return err
	}

	if err := payment.ValidateDelete(); err != nil {
		return err
	}

	return h.paymentRepo.SoftDelete(command.PaymentId)
}
