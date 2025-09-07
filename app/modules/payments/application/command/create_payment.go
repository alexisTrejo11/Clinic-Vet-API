package command

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/payment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
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

	amount := valueobject.NewMoney(amountValue, amountCurrency)

	paymentMethod, err := enum.ParsePaymentMethod(paymentMethodStr)
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

	payment, err := h.createCommandToDomain(command)
	if err != nil {
		return cqrs.FailureResult("failed to create payment domain", err)
	}

	if err := h.paymentRepo.Save(command.ctx, payment); err != nil {
		return cqrs.FailureResult("failed to create payment", err)
	}

	return cqrs.SuccessResult(payment.ID().String(), "payment created successfully")
}

func (h *CreatePaymentHandler) createCommandToDomain(command CreatePaymentCommand) (*payment.Payment, error) {
	paymentEntity, err := payment.NewPayment(
		valueobject.PaymentID{},
		command.appointmentID,
		command.userID,
		payment.WithCurrency(command.amount.Currency()),
		payment.WithAmount(command.amount),
		payment.WithPaymentMethod(command.paymentMethod),
		payment.WithDescription(command.description),
		payment.WithDueDate(command.dueDate),
		payment.WithTransactionID(command.transactionID),
		payment.WithStatus(enum.PaymentStatusPending))
	if err != nil {
		return nil, apperror.MappingError([]string{err.Error()}, "constructor", "domain", "Payment")
	}

	return paymentEntity, nil
}
