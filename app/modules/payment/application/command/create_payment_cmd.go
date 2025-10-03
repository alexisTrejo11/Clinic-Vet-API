package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
	"time"
)

type CreatePaymentCommand struct {
	Amount        valueobject.Money
	Status        enum.PaymentStatus
	Method        enum.PaymentMethod
	Description   string
	MedSessionID  *valueobject.MedSessionID
	PaidBy        *valueobject.CustomerID
	InvoiceID     *string
	TransactionID *string
	DueDate       *time.Time
}

func (h *paymentCommandHandler) CreatePayment(ctx context.Context, cmd CreatePaymentCommand) cqrs.CommandResult {
	payment := cmd.ToEntity()
	if err := payment.ValidatePersistence(ctx); err != nil {
		return cqrs.FailureResult(ErrInvalidPaymentData, err)
	}

	if err := h.paymentRepository.Save(ctx, payment); err != nil {
		return cqrs.FailureResult(ErrFailedCreatePayment, err)
	}

	return cqrs.SuccessCreateResult(payment.ID().String(), MsgPaymentCreated)
}

func (cmd *CreatePaymentCommand) ToEntity() *payment.Payment {
	return payment.NewPaymentBuilder().
		WithAmount(cmd.Amount).
		WithStatus(cmd.Status).
		WithPaymentMethod(cmd.Method).
		WithMedSessionID(cmd.MedSessionID).
		WithPaidByCustomer(cmd.PaidBy).
		WithTransactionID(cmd.TransactionID).
		WithDescription(cmd.Description).
		WithDueDate(cmd.DueDate).
		WithInvoiceID(cmd.InvoiceID).
		Build()
}
