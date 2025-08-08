package paymentCmd

import (
	"time"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
)

type CreatePaymentCommand struct {
	AppointmentId int                         `json:"appointment_id"`
	OwnerId       int                         `json:"owner_id"`
	Amount        float64                     `json:"amount"`
	Currency      string                      `json:"currency"`
	PaymentMethod paymentDomain.PaymentMethod `json:"payment_method"`
	Description   *string                     `json:"description,omitempty"`
	DueDate       *time.Time                  `json:"due_date,omitempty"`
	TransactionId *string                     `json:"transaction_id,omitempty"`
}

type CreatePaymentHander interface {
	Handle(command CreatePaymentCommand) (*paymentDomain.Payment, error)
}

type createPaymentHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewCreatePaymentHandler(paymentRepo paymentDomain.PaymentRepository) CreatePaymentHander {
	return &createPaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *createPaymentHandler) Handle(command CreatePaymentCommand) (*paymentDomain.Payment, error) {
	payment := &paymentDomain.Payment{
		AppointmentId: command.AppointmentId,
		OwnerId:       command.OwnerId,
		Amount:        paymentDomain.NewMoney(command.Amount, command.Currency),
		Currency:      command.Currency,
		PaymentMethod: command.PaymentMethod,
		Description:   command.Description,
		DueDate:       command.DueDate,
		TransactionId: command.TransactionId,
		Status:        paymentDomain.PENDING,
		IsActive:      true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := h.paymentRepo.Save(payment); err != nil {
		return nil, err
	}

	return payment, nil
}
