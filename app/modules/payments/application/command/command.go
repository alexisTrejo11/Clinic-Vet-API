package command

import (
	"context"
	"time"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type CancelPaymentCommand struct {
	paymentID valueobject.PaymentID
	ctx       context.Context
	reason    string
}

func NewCancelPaymentCommand(idInt uint, reason string) *CancelPaymentCommand {
	return &CancelPaymentCommand{
		paymentID: valueobject.NewPaymentID(idInt),
		reason:    reason,
	}
}

type CreatePaymentCommand struct {
	Ctx              context.Context
	Amount           valueobject.Money
	Status           enum.PaymentStatus
	Method           enum.PaymentMethod
	TransactionID    *string
	Description      *string
	DueDate          time.Time
	PaidAt           *time.Time
	RefundedAt       *time.Time
	PaidFromCustomer valueobject.CustomerID
	AppointmentID    *valueobject.AppointmentID
	InvoiceID        *string
}

type MarkOverduePaymentsCommand struct {
	context context.Context
}

func NewMarkOverduePaymentsCommand(ctx context.Context) *MarkOverduePaymentsCommand {
	return &MarkOverduePaymentsCommand{context: ctx}
}

type DeletePaymentCommand struct {
	paymentID valueobject.PaymentID
	ctx       context.Context
}

func NewDeletePaymentCommand(ctx context.Context, idInt uint) *DeletePaymentCommand {
	return &DeletePaymentCommand{
		paymentID: valueobject.NewPaymentID(idInt),
		ctx:       ctx,
	}
}

type RefundPaymentCommand struct {
	paymentID valueobject.PaymentID
	reason    string
	ctx       context.Context
}

func NewRefundPaymentCommand(ctx context.Context, paymentID uint, reason string) *RefundPaymentCommand {
	return &RefundPaymentCommand{
		paymentID: valueobject.NewPaymentID(paymentID),
		reason:    reason,
		ctx:       context.Background(),
	}
}

type UpdatePaymentCommand struct {
	paymentID     valueobject.PaymentID
	amount        *valueobject.Money
	paymentMethod *enum.PaymentMethod
	description   *string
	dueDate       *time.Time
	ctx           context.Context
}

func NewUpdatePaymentCommand(
	ctx context.Context,
	idInt uint,
	amountValue *float64,
	amountCurrency *string,
	paymentMethodStr *string,
	description *string,
	dueDate *time.Time,
) (UpdatePaymentCommand, error) {
	var errors []string

	if idInt == 0 {
		errors = append(errors, "payment ID is required")
	}
	paymentID := valueobject.NewPaymentID(idInt)

	var amount *valueobject.Money
	if amountValue != nil {
		if amountCurrency == nil {
			errors = append(errors, "currency is required when amount is provided")
		} else {
			money := valueobject.NewMoney(*amountValue, *amountCurrency)
			amount = &money
		}
	} else if amountCurrency != nil {
		errors = append(errors, "amount is required when currency is provided")
	}

	var paymentMethod *enum.PaymentMethod
	if paymentMethodStr != nil {
		pm, err := enum.ParsePaymentMethod(*paymentMethodStr)
		if err != nil {
			errors = append(errors, err.Error())
		} else {
			paymentMethod = &pm
		}
	}

	if description != nil && len(*description) > 500 {
		errors = append(errors, "description must be less than 500 characters")
	}

	if dueDate != nil && dueDate.Before(time.Now()) {
		errors = append(errors, "due date must be in the future")
	}

	if len(errors) > 0 {
		return UpdatePaymentCommand{}, apperror.MappingError(errors, "constructor", "command", "payment")
	}

	return UpdatePaymentCommand{
		paymentID:     paymentID,
		amount:        amount,
		paymentMethod: paymentMethod,
		description:   description,
		dueDate:       dueDate,
		ctx:           ctx,
	}, nil
}

type ProcessPaymentCommand struct {
	paymentID     valueobject.PaymentID
	transactionID string
	ctx           context.Context
}

func NewProcessPaymentCommand(idInt uint, transactionID string) (*ProcessPaymentCommand, error) {
	if idInt == 0 {
		return nil, apperror.FieldValidationError("id", "", "Payment ID cannot be zero")
	}

	if transactionID == "" {
		return nil, apperror.FieldValidationError("transaction_id", "", "Transaction ID cannot be empty")
	}

	return &ProcessPaymentCommand{
		paymentID:     valueobject.NewPaymentID(idInt),
		transactionID: transactionID,
		ctx:           context.Background(),
	}, nil
}
