package paymentCmd

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
)

type CancelPaymentCommand struct {
	PaymentId int             `json:"payment_id"`
	Reason    string          `json:"reason"`
	CTX       context.Context `json:"-"`
}

type CancelPaymentHandler interface {
	Handle(command CancelPaymentCommand) error
}

type cancelPaymentHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewCancelPaymentHandler(paymentRepo paymentDomain.PaymentRepository) CancelPaymentHandler {
	return &cancelPaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *cancelPaymentHandler) Handle(command CancelPaymentCommand) error {
	if command.PaymentId == 0 {
		return paymentDomain.InvalidPaymentIdErr(command.PaymentId)
	}

	payment, err := h.paymentRepo.GetById(command.CTX, command.PaymentId)
	if err != nil {
		return err
	}

	if err := payment.Cancel(command.Reason); err != nil {
		return err
	}

	return h.paymentRepo.Save(payment)
}
