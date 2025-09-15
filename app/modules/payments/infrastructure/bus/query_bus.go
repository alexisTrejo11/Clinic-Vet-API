// Package bus contains all the setup for command and query bus related to payment
package bus

import (
	query "clinic-vet-api/app/modules/payments/application/queries"
	"clinic-vet-api/app/shared/page"
)

type PaymentQueryBus struct {
	handler query.PaymentQueryHandler
}

func NewPaymentQueryBus(handler query.PaymentQueryHandler) *PaymentQueryBus {
	return &PaymentQueryBus{handler: handler}
}

func (b *PaymentQueryBus) FindByID(qry query.FindPaymentByIDQuery) (query.PaymentResult, error) {
	return b.handler.FindByID(qry)
}

func (b *PaymentQueryBus) FindOverdues(qry query.FindOverduePaymentsQuery) (page.Page[query.PaymentResult], error) {
	return b.handler.FindOverdues(qry)
}

func (b *PaymentQueryBus) FindByStatus(qry query.FindPaymentsByStatusQuery) (page.Page[query.PaymentResult], error) {
	return b.handler.FindByStatus(qry)
}

func (b *PaymentQueryBus) FindByCustomer(qry query.FindPaymentsByCustomerQuery) (page.Page[query.PaymentResult], error) {
	return b.handler.FindByCustomer(qry)
}

func (b *PaymentQueryBus) FindByDateRange(qry query.FindPaymentsByDateRangeQuery) (page.Page[query.PaymentResult], error) {
	return b.handler.FindByDateRange(qry)
}

func (b *PaymentQueryBus) FindBySpecification(qry query.FindPaymentsBySpecification) (page.Page[query.PaymentResult], error) {
	return b.handler.FindBySpecification(qry)
}
