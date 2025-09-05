// Package bus contains all the setup for command and query bus related to payment
package bus

import (
	"fmt"
	"reflect"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/payments/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type PaymentCommandBus struct {
	handlers map[reflect.Type]cqrs.CommandHandler
}

func NewPaymentCommandBus(paymentRepo repository.PaymentRepository) *PaymentCommandBus {
	bus := &PaymentCommandBus{
		handlers: make(map[reflect.Type]cqrs.CommandHandler),
	}

	bus.registerHandlers(paymentRepo)
	return bus
}

func (bus *PaymentCommandBus) registerHandlers(paymentRepo repository.PaymentRepository) {
	bus.Register(reflect.TypeOf(command.CreatePaymentCommand{}), command.NewCreatePaymentHandler(paymentRepo))
	bus.Register(reflect.TypeOf(command.ProcessPaymentCommand{}), command.NewProcessPaymentHandler(paymentRepo))
	bus.Register(reflect.TypeOf(command.RefundPaymentCommand{}), command.NewRefundPaymentHandler(paymentRepo))
	bus.Register(reflect.TypeOf(command.CancelPaymentCommand{}), command.NewCancelPaymentHandler(paymentRepo))
	bus.Register(reflect.TypeOf(command.UpdatePaymentCommand{}), command.NewUpdatePaymentHandler(paymentRepo))
	bus.Register(reflect.TypeOf(command.DeletePaymentCommand{}), command.NewDeletePaymentHandler(paymentRepo))
	bus.Register(reflect.TypeOf(command.MarkOverduePaymentsCommand{}), command.NewMarkOverduePaymentsHandler(paymentRepo))
}

func (bus *PaymentCommandBus) Register(commandType reflect.Type, handler cqrs.CommandHandler) error {
	if handler == nil {
		return fmt.Errorf("handler cannot be nil for command type %s", commandType.Name())
	}

	bus.handlers[commandType] = handler
	return nil
}

func (bus *PaymentCommandBus) Execute(c cqrs.Command) cqrs.CommandResult {
	commandType := reflect.TypeOf(c)
	handler, exists := bus.handlers[commandType]

	if !exists {
		return cqrs.FailureResult(
			"unhandled command type",
			fmt.Errorf("no handler registered for command type %s", commandType.Name()),
		)
	}

	switch cmd := c.(type) {
	case command.CreatePaymentCommand:
		h := handler.(*command.CreatePaymentHandler)
		return h.Handle(cmd)

	case command.ProcessPaymentCommand:
		h := handler.(*command.ProcessPaymentHandler)
		return h.Handle(cmd)

	case command.RefundPaymentCommand:
		h := handler.(*command.RefundPaymentHandler)
		return h.Handle(cmd)

	case command.CancelPaymentCommand:
		h := handler.(*command.CancelPaymentHandler)
		return h.Handle(cmd)

	case command.UpdatePaymentCommand:
		h := handler.(*command.UpdatePaymentHandler)
		return h.Handle(cmd)

	case command.DeletePaymentCommand:
		h := handler.(*command.DeletePaymentHandler)
		return h.Handle(cmd)

	case command.MarkOverduePaymentsCommand:
		h := handler.(*command.MarkOverduePaymentsHandler)
		return h.Handle(cmd)

	default:
		return cqrs.FailureResult("not registred command", fmt.Errorf("unknown command type: %s", commandType.Name()))
	}
}
