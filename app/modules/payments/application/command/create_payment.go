package command

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type CreatePaymentCommand struct {
	AppointmentID int                `json:"appointment_id"`
	UserID        int                `json:"owner_id"`
	Amount        float64            `json:"amount"`
	Currency      string             `json:"currency"`
	PaymentMethod enum.PaymentMethod `json:"payment_method"`
	Description   *string            `json:"description,omitempty"`
	DueDate       *time.Time         `json:"due_date,omitempty"`
	TransactionID *string            `json:"transaction_id,omitempty"`
}

type CreatePaymentHander interface {
	Handle(ctx context.Context, command CreatePaymentCommand) shared.CommandResult
}
type createPaymentHandler struct {
	paymentRepo repository.PaymentRepository
}

func NewCreatePaymentHandler(paymentRepo repository.PaymentRepository) CreatePaymentHander {
	return &createPaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *createPaymentHandler) Handle(ctx context.Context, command CreatePaymentCommand) shared.CommandResult {
	payment := h.createCommandToDomain(command)

	if err := h.paymentRepo.Save(ctx, payment); err != nil {
		return shared.FailureResult("failed to create payment", err)
	}

	return shared.SuccessResult(string(payment.GetID()), "payment created successfully")
}

func (req *createPaymentHandler) createCommandToDomain(command CreatePaymentCommand) *entity.Payment {
	return entity.NewPaymentBuilder().
		WithAppointmentID(command.AppointmentID).
		WithUserID(command.UserID).
		WithAmount(valueobject.NewMoney(command.Amount, command.Currency)).
		WithCurrency(command.Currency).
		WithPaymentMethod(command.PaymentMethod).
		WithDescription(command.Description).
		WithDueDate(command.DueDate).
		WithTransactionID(command.TransactionID).
		WithStatus(enum.PENDING).
		WithIsActive(true).
		WithCreatedAt(time.Now()).
		WithUpdatedAt(time.Now()).
		Build()
}
