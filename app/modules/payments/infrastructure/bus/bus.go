package bus

import (
	"clinic-vet-api/app/modules/payments/application/command"
	query "clinic-vet-api/app/modules/payments/application/queries"
)

type PaymentBus struct {
	CommandBus PaymentCommandBus
	QueryBus   PaymentQueryBus
}

func NewPaymentBus(commandHandler command.PaymentCommandHandler, queryHandler query.PaymentQueryHandler) *PaymentBus {
	return &PaymentBus{CommandBus: *NewPaymentCommandBus(commandHandler), QueryBus: *NewPaymentQueryBus(queryHandler)}
}
