package command

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type CancelPaymentCommand struct {
	paymentID valueobject.PaymentID
	reason    string
}

func NewCancelPaymentCommand(idInt uint, reason string) *CancelPaymentCommand {
	return &CancelPaymentCommand{
		paymentID: valueobject.NewPaymentID(idInt),
		reason:    reason,
	}
}

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

type MarkOverduePaymentsCommand struct {
}

func NewMarkOverduePaymentsCommand() *MarkOverduePaymentsCommand {
	return &MarkOverduePaymentsCommand{}
}

type DeletePaymentCommand struct {
	paymentID valueobject.PaymentID
}

func NewDeletePaymentCommand(idInt uint) *DeletePaymentCommand {
	return &DeletePaymentCommand{
		paymentID: valueobject.NewPaymentID(idInt),
	}
}

type RefundPaymentCommand struct {
	paymentID valueobject.PaymentID
	reason    string
}

func NewRefundPaymentCommand(paymentID uint, reason string) *RefundPaymentCommand {
	return &RefundPaymentCommand{
		paymentID: valueobject.NewPaymentID(paymentID),
		reason:    reason,
	}
}

type UpdatePaymentCommand struct {
	paymentID     valueobject.PaymentID
	amount        *valueobject.Money
	paymentMethod *enum.PaymentMethod
	description   *string
	dueDate       *time.Time
}

func NewUpdatePaymentCommand(id uint, amount *valueobject.Money, paymentMethod *string, description *string, dueDate *time.Time) UpdatePaymentCommand {
	paymentID := valueobject.NewPaymentID(id)

	var paymentMethodObj *enum.PaymentMethod
	if paymentMethod != nil {
		pm := enum.PaymentMethod(*paymentMethod)
		paymentMethodObj = &pm

	}

	return UpdatePaymentCommand{
		paymentID:     paymentID,
		amount:        amount,
		paymentMethod: paymentMethodObj,
		description:   description,
		dueDate:       dueDate,
	}
}

type ProcessPaymentCommand struct {
	paymentID     valueobject.PaymentID
	transactionID string
}

func NewProcessPaymentCommand(idInt uint, transactionID string) *ProcessPaymentCommand {
	return &ProcessPaymentCommand{
		paymentID:     valueobject.NewPaymentID(idInt),
		transactionID: transactionID,
	}
}
