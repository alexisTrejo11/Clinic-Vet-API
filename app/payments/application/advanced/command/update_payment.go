package paymentCmd

import (
	"context"
	"time"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
)

type UpdatePaymentCommand struct {
	PaymentId     int                          `json:"payment_id"`
	Amount        *paymentDomain.Money         `json:"amount,omitempty"`
	PaymentMethod *paymentDomain.PaymentMethod `json:"payment_method,omitempty"`
	Description   *string                      `json:"description,omitempty"`
	DueDate       *time.Time                   `json:"due_date,omitempty"`
	CTX           context.Context              `json:"-"`
}

type UpdatePaymentHandler interface {
	Handle(command UpdatePaymentCommand) (*paymentDomain.Payment, error)
}

type updatePaymentHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewUpdatePaymentHandler(paymentRepo paymentDomain.PaymentRepository) UpdatePaymentHandler {
	return &updatePaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *updatePaymentHandler) Handle(cmd UpdatePaymentCommand) (*paymentDomain.Payment, error) {
	if cmd.PaymentId == 0 {
		return nil, paymentDomain.PaymentNotFoundErr(cmd.PaymentId)
	}

	payment, err := h.paymentRepo.GetById(cmd.CTX, cmd.PaymentId)
	if err != nil {
		return nil, err
	}

	err = payment.Update(cmd.Amount, cmd.PaymentMethod, cmd.Description, cmd.DueDate)
	if err != nil {
		return nil, err
	}

	err = h.paymentRepo.Save(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
