package bus

import "clinic-vet-api/app/modules/core/repository"

type PaymentBus struct {
	CommandBus PaymentCommandBus
	QueryBus   PaymentQueryBus
}

func NewPaymentBus(paymentRepository repository.PaymentRepository) *PaymentBus {
	return &PaymentBus{CommandBus: *NewPaymentCommandBus(paymentRepository), QueryBus: *NewPaymentQueryBus(paymentRepository)}
}
