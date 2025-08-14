package paymentCmd

import (
	"context"
	"fmt"
	"reflect"
	"time"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type Command interface{}

type CommandHandler interface{}

type CommandBus interface {
	Register(commandType reflect.Type, handler CommandHandler) error
	Execute(ctx context.Context, command Command) shared.CommandResult
}

type paymentCommandBus struct {
	handlers map[reflect.Type]CommandHandler
}

func NewPaymentCommandBus(paymentRepo paymentDomain.PaymentRepository) CommandBus {
	bus := &paymentCommandBus{
		handlers: make(map[reflect.Type]CommandHandler),
	}

	bus.registerHandlers(paymentRepo)
	return bus
}

func (bus *paymentCommandBus) registerHandlers(paymentRepo paymentDomain.PaymentRepository) {
	bus.Register(reflect.TypeOf(CreatePaymentCommand{}), NewCreatePaymentHandler(paymentRepo))
	bus.Register(reflect.TypeOf(ProcessPaymentCommand{}), NewProcessPaymentHandler(paymentRepo))
	bus.Register(reflect.TypeOf(RefundPaymentCommand{}), NewRefundPaymentHandler(paymentRepo))
	bus.Register(reflect.TypeOf(CancelPaymentCommand{}), NewCancelPaymentHandler(paymentRepo))
	bus.Register(reflect.TypeOf(UpdatePaymentCommand{}), NewUpdatePaymentHandler(paymentRepo))
	bus.Register(reflect.TypeOf(DeletePaymentCommand{}), NewDeletePaymentHandler(paymentRepo))
	bus.Register(reflect.TypeOf(MarkOverduePaymentsCommand{}), NewMarkOverduePaymentsHandler(paymentRepo))
}

func (bus *paymentCommandBus) Register(commandType reflect.Type, handler CommandHandler) error {
	if handler == nil {
		return fmt.Errorf("handler cannot be nil for command type %s", commandType.Name())
	}

	bus.handlers[commandType] = handler
	return nil
}

func (bus *paymentCommandBus) Execute(ctx context.Context, command Command) shared.CommandResult {
	commandType := reflect.TypeOf(command)
	handler, exists := bus.handlers[commandType]

	if !exists {
		return shared.FailureResult(
			"unhandled command type",
			fmt.Errorf("no handler registered for command type %s", commandType.Name()),
		)
	}

	switch cmd := command.(type) {
	case CreatePaymentCommand:
		h := handler.(CreatePaymentHander)
		return h.Handle(ctx, cmd)

	case ProcessPaymentCommand:
		h := handler.(ProcessPaymentHandler)
		return h.Handle(ctx, cmd)

	case RefundPaymentCommand:
		h := handler.(RefundPaymentHandler)
		return h.Handle(ctx, cmd)

	case CancelPaymentCommand:
		h := handler.(CancelPaymentHandler)
		return h.Handle(ctx, cmd)

	case UpdatePaymentCommand:
		h := handler.(UpdatePaymentHandler)
		return h.Handle(ctx, cmd)

	case DeletePaymentCommand:
		h := handler.(DeletePaymentHandler)
		return h.Handle(ctx, cmd)

	case MarkOverduePaymentsCommand:
		h := handler.(MarkOverduePaymentsHandler)
		return h.Handle(ctx, cmd)

	default:
		return shared.FailureResult("not registred command", fmt.Errorf("unknown command type: %s", commandType.Name()))
	}
}

type PaymentCommandService struct {
	commandBus  CommandBus
	paymentRepo paymentDomain.PaymentRepository
}

func NewPaymentCommandService(ctx context.Context, commandBus CommandBus, paymentRepo paymentDomain.PaymentRepository) *PaymentCommandService {
	return &PaymentCommandService{
		commandBus:  commandBus,
		paymentRepo: paymentRepo,
	}
}

func (s *PaymentCommandService) ProcessPayment(payment *paymentDomain.Payment) shared.CommandResult {
	cmd := NewProcessPaymentCommand(payment.GetId(), *payment.GetTransactionId())
	return s.commandBus.Execute(context.Background(), cmd)

}

func (s *PaymentCommandService) RefundPayment(paymentId int, reason string) shared.CommandResult {
	cmd := NewRefundPaymentCommand(paymentId, reason)
	return s.commandBus.Execute(context.Background(), cmd)
}

func (s *PaymentCommandService) ValidatePayment(payment *paymentDomain.Payment) error {
	if payment == nil {
		return paymentDomain.NewPaymentError("INVALID_PAYMENT", "payment cannot be nil", 0, "")
	}

	if payment.GetAmount().IsZero() || payment.GetAmount().IsNegative() {
		return paymentDomain.ErrInvalidAmount
	}

	if !payment.GetPaymentMethod().IsValid() {
		return paymentDomain.ErrInvalidPaymentMethod
	}

	if !payment.GetStatus().IsValid() {
		return paymentDomain.ErrInvalidPaymentStatus
	}

	return nil
}

func (s *PaymentCommandService) CalculateTotal(appointmentId int) (paymentDomain.Money, error) {
	return paymentDomain.NewMoney(0, "USD"), fmt.Errorf("not implemented")
}

func (s *PaymentCommandService) GetPaymentHistory(ownerId int) (page.Page[[]paymentDomain.Payment], error) {
	return page.Page[[]paymentDomain.Payment]{}, fmt.Errorf("use query handler for read operations")
}

func (s *PaymentCommandService) MarkOverduePayments() shared.CommandResult {
	cmd := MarkOverduePaymentsCommand{}

	return s.commandBus.Execute(context.Background(), cmd)

}

func (s *PaymentCommandService) GeneratePaymentReport(startDate, endDate time.Time) (paymentDomain.PaymentReport, error) {
	return paymentDomain.PaymentReport{}, fmt.Errorf("use query handler for read operations")
}
