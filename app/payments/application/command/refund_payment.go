package paymentCmd

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
)

type RefundPaymentCommand struct {
	PaymentId int             `json:"payment_id"`
	Reason    string          `json:"reason"`
	CTX       context.Context `json:"-"`
}

type RefundPaymentHandler interface {
	Handle(command RefundPaymentCommand) error
}

type refundPaymentHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewRefundPaymentHandler(paymentRepo paymentDomain.PaymentRepository) RefundPaymentHandler {
	return &refundPaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *refundPaymentHandler) Handle(command RefundPaymentCommand) error {
	if command.PaymentId == 0 {
		return paymentDomain.PaymentNotFoundErr(command.PaymentId)
	}

	payment, err := h.paymentRepo.GetById(command.CTX, command.PaymentId)
	if err != nil {
		return err
	}

	if err := payment.Refund(command.Reason); err != nil {
		return err
	}

	return h.paymentRepo.Save(payment)
}
