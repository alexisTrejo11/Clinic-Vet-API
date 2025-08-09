package paymentCmd

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
)

type ProcessPaymentCommand struct {
	PaymentId     int             `json:"payment_id"`
	TransactionId string          `json:"transaction_id"`
	CTX           context.Context `json:"-"`
}

type ProcessPaymentHandler interface {
	Handle(command ProcessPaymentCommand) error
}

type processPaymentHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewProcessPaymentHandler(paymentRepo paymentDomain.PaymentRepository) ProcessPaymentHandler {
	return &processPaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *processPaymentHandler) Handle(command ProcessPaymentCommand) error {
	if command.PaymentId == 0 {
		return paymentDomain.InvalidPaymentIdErr(command.PaymentId)
	}

	payment, err := h.paymentRepo.GetById(command.CTX, command.PaymentId)
	if err != nil {
		return err
	}

	if err := payment.Proccess(command.TransactionId); err != nil {
		return err
	}

	return h.paymentRepo.Save(payment)
}
