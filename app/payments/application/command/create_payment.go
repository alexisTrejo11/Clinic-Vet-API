package paymentCmd

import (
	"context"
	"time"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type CreatePaymentCommand struct {
	AppointmentId int                         `json:"appointment_id"`
	UserId        int                         `json:"owner_id"`
	Amount        float64                     `json:"amount"`
	Currency      string                      `json:"currency"`
	PaymentMethod paymentDomain.PaymentMethod `json:"payment_method"`
	Description   *string                     `json:"description,omitempty"`
	DueDate       *time.Time                  `json:"due_date,omitempty"`
	TransactionId *string                     `json:"transaction_id,omitempty"`
}

type CreatePaymentHander interface {
	Handle(ctx context.Context, command CreatePaymentCommand) shared.CommandResult
}
type createPaymentHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewCreatePaymentHandler(paymentRepo paymentDomain.PaymentRepository) CreatePaymentHander {
	return &createPaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *createPaymentHandler) Handle(ctx context.Context, command CreatePaymentCommand) shared.CommandResult {
	payment := h.createCommandToDomain(command)

	if err := h.paymentRepo.Save(ctx, payment); err != nil {
		return shared.FailureResult("failed to create payment", err)
	}

	return shared.SuccessResult(string(payment.GetId()), "payment created successfully")
}

func (req *createPaymentHandler) createCommandToDomain(command CreatePaymentCommand) *paymentDomain.Payment {
	return paymentDomain.NewPaymentBuilder().
		WithAppointmentId(command.AppointmentId).
		WithUserId(command.UserId).
		WithAmount(paymentDomain.NewMoney(command.Amount, command.Currency)).
		WithCurrency(command.Currency).
		WithPaymentMethod(command.PaymentMethod).
		WithDescription(command.Description).
		WithDueDate(command.DueDate).
		WithTransactionId(command.TransactionId).
		WithStatus(paymentDomain.PENDING).
		WithIsActive(true).
		WithCreatedAt(time.Now()).
		WithUpdatedAt(time.Now()).
		Build()
}
