package command

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type CreatePaymentCommand struct {
	ctx           context.Context
	appointmentID valueobject.AppointmentID
	userID        valueobject.UserID
	amount        valueobject.Money
	paymentMethod enum.PaymentMethod
	description   *string
	dueDate       *time.Time
	transactionID *string
}

func NewCreatePaymentCommand(
	ctx context.Context,
	appointmentIDInt int,
	userIDInt int,
	amountValue float64,
	amountCurrency string,
	paymentMethodStr string,
	description *string,
	dueDate *time.Time,
	transactionID *string,
) (CreatePaymentCommand, error) {
	var errors []string

	appointmentID, err := valueobject.NewAppointmentID(appointmentIDInt)
	if err != nil {
		errors = append(errors, err.Error())
	}

	userID, err := valueobject.NewUserID(userIDInt)
	if err != nil {
		errors = append(errors, err.Error())
	}

	amount, err := valueobject.NewMoney(amountValue, amountCurrency)
	if err != nil {
		errors = append(errors, err.Error())
	}

	paymentMethod, err := enum.NewPaymentMethod(paymentMethodStr)
	if err != nil {
		errors = append(errors, err.Error())
	}

	if description != nil {
		if *description == "" {
			errors = append(errors, "description cannot be empty if provided")
		} else if len(*description) > 500 {
			errors = append(errors, "description must be less than 500 characters")
		}
	}

	if dueDate != nil {
		if dueDate.Before(time.Now()) {
			errors = append(errors, "due date must be in the future")
		}
	}

	if transactionID != nil {
		if *transactionID == "" {
			errors = append(errors, "transaction ID cannot be empty if provided")
		} else if len(*transactionID) > 100 {
			errors = append(errors, "transaction ID must be less than 100 characters")
		}
	}

	if len(errors) > 0 {
		return CreatePaymentCommand{}, apperror.MappingError(errors, "constructor", "command", "payment")
	}

	return CreatePaymentCommand{
		ctx:           ctx,
		appointmentID: appointmentID,
		userID:        userID,
		amount:        amount,
		paymentMethod: paymentMethod,
		description:   description,
		dueDate:       dueDate,
		transactionID: transactionID,
	}, nil
}

type CreatePaymentHandler struct {
	paymentRepo repository.PaymentRepository
}

func NewCreatePaymentHandler(paymentRepo repository.PaymentRepository) cqrs.CommandHandler {
	return &CreatePaymentHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *CreatePaymentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(CreatePaymentCommand)

	payment := h.createCommandToDomain(command)

	if err := h.paymentRepo.Save(command.ctx, payment); err != nil {
		return cqrs.FailureResult("failed to create payment", err)
	}

	return cqrs.SuccessResult(payment.GetID().String(), "payment created successfully")
}

func (h *CreatePaymentHandler) createCommandToDomain(command CreatePaymentCommand) *entity.Payment {
	return entity.NewPaymentBuilder().
		WithAppointmentID(command.appointmentID).
		WithUserID(command.userID).
		WithCurrency(command.amount.Currency).
		WithAmount(command.amount).
		WithPaymentMethod(command.paymentMethod).
		WithDescription(command.description).
		WithDueDate(command.dueDate).
		WithTransactionID(command.transactionID).
		WithStatus(enum.PENDING).
		WithCreatedAt(time.Now()).
		WithUpdatedAt(time.Now()).
		Build()
}
