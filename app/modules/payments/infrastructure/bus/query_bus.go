package bus

import (
	"fmt"
	"reflect"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	query "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/payments/application/queries"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type PaymentQueryBus struct {
	handlers map[reflect.Type]any
}

func NewQueryBus(paymentRepo repository.PaymentRepository) *PaymentQueryBus {
	bus := &PaymentQueryBus{
		handlers: make(map[reflect.Type]any),
	}
	bus.registerHandlers(paymentRepo)
	return bus
}

func (bus *PaymentQueryBus) registerHandlers(paymentRepo repository.PaymentRepository) {
	bus.Register(reflect.TypeOf(query.GetPaymentByIDQuery{}), query.NewGetPaymentByIDHandler(paymentRepo))
	bus.Register(reflect.TypeOf(query.ListPaymentsByOwnerQuery{}), query.NewListByOwnerHandler(paymentRepo))
	bus.Register(reflect.TypeOf(query.ListPaymentsByStatusQuery{}), query.NewListPaymentsByStatusHandler(paymentRepo))
	bus.Register(reflect.TypeOf(query.ListOverduePaymentsQuery{}), query.NewListOverduePaymentsHandler(paymentRepo))
	bus.Register(reflect.TypeOf(query.SearchPaymentsQuery{}), query.NewSearchPaymentsHandler(paymentRepo))
	// bus.Register(reflect.TypeOf(query.GetPaymentHistoryQuery{}), query.NewGetPaymentHistoryHandler(paymentRepo))
}

func (bus *PaymentQueryBus) Execute(qry cqrs.Query) (any, error) {
	queryType := reflect.TypeOf(qry)
	handler, exists := bus.handlers[queryType]

	if !exists {
		return nil, fmt.Errorf("no handler registered for query type %s", queryType.Name())
	}

	switch q := qry.(type) {
	case query.GetPaymentByIDQuery:
		h := handler.(query.GetPaymentByIDHandler)
		return h.Handle(q)
	case query.ListPaymentsByOwnerQuery:
		h := handler.(query.ListPaymentsByStatusHandler)
		return h.Handle(q)

	case query.ListPaymentsByStatusQuery:
		h := handler.(query.ListPaymentsByStatusHandler)
		return h.Handle(q)

	case query.ListOverduePaymentsQuery:
		h := handler.(query.ListOverduePaymentsHandler)
		return h.Handle(q)

	case query.SearchPaymentsQuery:
		h := handler.(query.SearchPaymentsHandler)
		return h.Handle(q)

	default:
		return nil, fmt.Errorf("unknown query type: %s", queryType.Name())
	}
}

func (bus *PaymentQueryBus) Register(queryType reflect.Type, handler any) error {
	if handler == nil {
		return fmt.Errorf("handler cannot be nil for query type %s", queryType.Name())
	}
	bus.handlers[queryType] = handler
	return nil
}
