package bus

import (
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/payments/application/command"
	"clinic-vet-api/app/shared/cqrs"
)

type PaymentCommandBus struct {
	handlers command.PaymentCommandHandler
}

func NewPaymentCommandBus(paymentRepo repository.PaymentRepository) *PaymentCommandBus {
	handlers := command.NewPaymentCommandHandler(paymentRepo)

	return &PaymentCommandBus{handlers: handlers}
}

func (bus *PaymentCommandBus) CreatePayment(cmd command.CreatePaymentCommand) cqrs.CommandResult {
	return bus.handlers.CreatePayment(cmd)
}

func (bus *PaymentCommandBus) ProcessPayment(cmd command.ProcessPaymentCommand) cqrs.CommandResult {
	return bus.handlers.ProcessPayment(cmd)
}

func (bus *PaymentCommandBus) RefundPayment(cmd command.RefundPaymentCommand) cqrs.CommandResult {
	return bus.handlers.RefundPayment(cmd)
}

func (bus *PaymentCommandBus) CancelPayment(cmd command.CancelPaymentCommand) cqrs.CommandResult {
	return bus.handlers.CancelPayment(cmd)
}

func (bus *PaymentCommandBus) UpdatePayment(cmd command.UpdatePaymentCommand) cqrs.CommandResult {
	return bus.handlers.UpdatePayment(cmd)
}

func (bus *PaymentCommandBus) DeletePayment(cmd command.DeletePaymentCommand) cqrs.CommandResult {
	return bus.handlers.DeletePayment(cmd)
}

func (bus *PaymentCommandBus) MarkOverduePayments(cmd command.MarkOverduePaymentsCommand) cqrs.CommandResult {
	return bus.handlers.MarkOverudePayments(cmd)
}
