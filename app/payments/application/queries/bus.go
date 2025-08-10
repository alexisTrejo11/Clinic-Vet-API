package paymentQuery

import (
	"context"
	"fmt"
	"reflect"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
)

type Query interface{}

type QuerHandler[T any] interface {
	Handle(query Query) (T, error)
}

type QueryBus interface {
	Execute(ctx context.Context, query Query) (interface{}, error)
	Register(queryType reflect.Type, handler interface{}) error
}

type paymentQueryBus struct {
	handlers map[reflect.Type]interface{}
}

func NewQueryBus(paymentRepo paymentDomain.PaymentRepository) QueryBus {
	bus := &paymentQueryBus{
		handlers: make(map[reflect.Type]interface{}),
	}
	bus.registerHandlers(paymentRepo)
	return bus
}

func (bus *paymentQueryBus) registerHandlers(paymentRepo paymentDomain.PaymentRepository) {
	bus.Register(reflect.TypeOf(GetPaymentByIdQuery{}), NewGetPaymentByIdHandler(paymentRepo))
	bus.Register(reflect.TypeOf(ListPaymentsByUserQuery{}), NewListByUserHandler(paymentRepo))
	bus.Register(reflect.TypeOf(ListPaymentsByStatusQuery{}), NewListPaymentsByStatusHandler(paymentRepo))
	bus.Register(reflect.TypeOf(ListOverduePaymentsQuery{}), NewListOverduePaymentsHandler(paymentRepo))
	bus.Register(reflect.TypeOf(SearchPaymentsQuery{}), NewSearchPaymentsHandler(paymentRepo))
	//bus.Register(reflect.TypeOf(GetPaymentHistoryQuery{}), NewGetPaymentHistoryHandler(paymentRepo))
}

func (bus *paymentQueryBus) Execute(ctx context.Context, query Query) (interface{}, error) {
	queryType := reflect.TypeOf(query)
	handler, exists := bus.handlers[queryType]

	if !exists {
		return nil, fmt.Errorf("no handler registered for query type %s", queryType.Name())
	}

	// Type switch para manejar diferentes tipos de respuesta
	switch q := query.(type) {
	case GetPaymentByIdQuery:
		h := handler.(GetPaymentByIdHandler)
		return h.Handle(ctx, q)

	case ListPaymentsByUserQuery:
		h := handler.(ListPaymentsByUserHandler)
		return h.Handle(ctx, q)

	case ListPaymentsByStatusQuery:
		h := handler.(ListPaymentsByStatusHandler)
		return h.Handle(ctx, q)

	case ListOverduePaymentsQuery:
		h := handler.(ListOverduePaymentsHandler)
		return h.Handle(ctx, q)

	case SearchPaymentsQuery:
		h := handler.(SearchPaymentsHandler)
		return h.Handle(ctx, q)

	default:
		return nil, fmt.Errorf("unknown query type: %s", queryType.Name())
	}
}

func (bus *paymentQueryBus) Register(queryType reflect.Type, handler interface{}) error {
	if handler == nil {
		return fmt.Errorf("handler cannot be nil for query type %s", queryType.Name())
	}
	bus.handlers[queryType] = handler
	return nil
}
