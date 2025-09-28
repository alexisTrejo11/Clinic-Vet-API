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
	MedSessionID  valueobject.MedSessionID
	TransactionID *string
	Description   *string
	DueDate       *time.Time
	PaidAt        *time.Time
	PaidBy        *valueobject.CustomerID
	InvoiceID     *string
}

func (h *paymentCommandHandler) CreatePayment(ctx context.Context, cmd CreatePaymentCommand) cqrs.CommandResult {
	payment := cmd.ToEntity()
	if err := payment.ValidatePersistence(ctx); err != nil {
		return *cqrs.FailureResult(ErrInvalidPaymentData, err)
	}

	if err := h.paymentRepository.Save(ctx, &payment); err != nil {
		return *cqrs.FailureResult(ErrFailedCreatePayment, err)
	}

	return *cqrs.SuccessCreateResult(payment.ID().String(), MsgPaymentCreated)
}

func (cmd *CreatePaymentCommand) ToEntity() payment.Payment {
	builder := payment.NewPaymentBuilder()
	if cmd.PaidBy != nil {
		builder = builder.WithPaidByCustomer(*cmd.PaidBy)
	}

	if cmd.TransactionID != nil {
		builder = builder.WithTransactionID(*cmd.TransactionID)
	}

	if cmd.Description != nil {
		builder = builder.WithDescription(cmd.Description)
	}

	if !cmd.DueDate.IsZero() {
		builder = builder.WithDueDate(cmd.DueDate)
	}

	if cmd.PaidAt != nil {
		builder = builder.WithPaidAt(*cmd.PaidAt)
	}

	if cmd.InvoiceID != nil {
		builder = builder.WithInvoiceID(*cmd.InvoiceID)
	}

	return *builder.Build()
}
