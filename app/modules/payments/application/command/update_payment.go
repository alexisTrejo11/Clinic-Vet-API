package command

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

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
	idInt int,
	amountValue *float64,
	amountCurrency *string,
	paymentMethodStr *string,
	description *string,
	dueDate *time.Time,
) (UpdatePaymentCommand, error) {
	var errors []string

	paymentID, err := valueobject.NewPaymentID(idInt)
	if err != nil {
		errors = append(errors, err.Error())
	}

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

type UpdatePaymentHandler struct {
	paymentRepo repository.PaymentRepository
}

func NewUpdatePaymentHandler(paymentRepo repository.PaymentRepository) cqrs.CommandHandler {
	return &UpdatePaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *UpdatePaymentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(UpdatePaymentCommand)
	payment, err := h.paymentRepo.GetByID(command.ctx, command.paymentID)
	if err != nil {
		return cqrs.FailureResult("error fetching payment", err)
	}

	err = payment.Update(command.amount, command.paymentMethod, command.description, command.dueDate)
	if err != nil {
		return cqrs.FailureResult("error updating payment", err)
	}

	err = h.paymentRepo.Save(command.ctx, &payment)
	if err != nil {
		return cqrs.FailureResult("error saving payment", err)
	}

	return cqrs.SuccessResult(
		payment.ID().String(),
		"payment updated successfully",
	)
}
