package bus

import (
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/payments/application/command"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type PaymentCommandBus struct {
	handlers command.PaymentCommandHandler
}

func NewPaymentCommandBus(paymentRepo repository.PaymentRepository) *PaymentCommandBus {
	handlers := command.NewPaymentCommandHandler(paymentRepo)

	return &PaymentCommandBus{handlers: handlers}
}

func (bus *PaymentCommandBus) CreatePayment(ctx context.Context, cmd command.CreatePaymentCommand) cqrs.CommandResult {
	return bus.handlers.CreatePayment(ctx, cmd)
}

func (bus *PaymentCommandBus) ProcessPayment(ctx context.Context, cmd command.ProcessPaymentCommand) cqrs.CommandResult {
	return bus.handlers.ProcessPayment(ctx, cmd)
}

func (bus *PaymentCommandBus) RefundPayment(ctx context.Context, cmd command.RefundPaymentCommand) cqrs.CommandResult {
	return bus.handlers.RefundPayment(ctx, cmd)
}

func (bus *PaymentCommandBus) CancelPayment(ctx context.Context, cmd command.CancelPaymentCommand) cqrs.CommandResult {
	return bus.handlers.CancelPayment(ctx, cmd)
}

func (bus *PaymentCommandBus) UpdatePayment(ctx context.Context, cmd command.UpdatePaymentCommand) cqrs.CommandResult {
	return bus.handlers.UpdatePayment(ctx, cmd)
}

func (bus *PaymentCommandBus) DeletePayment(ctx context.Context, cmd command.DeletePaymentCommand) cqrs.CommandResult {
	return bus.handlers.DeletePayment(ctx, cmd)
}

func (bus *PaymentCommandBus) MarkOverduePayments(ctx context.Context, cmd command.MarkOverduePaymentsCommand) cqrs.CommandResult {
	return bus.handlers.MarkOverduePayments(ctx, cmd)
}
