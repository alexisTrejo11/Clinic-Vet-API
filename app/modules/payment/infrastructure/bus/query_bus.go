// Package bus contains all the setup for command and query bus related to payment
package bus

import (
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/payment/application/query"
	"clinic-vet-api/app/shared/page"
	"context"
)

type PaymentQueryBus struct {
	handler query.PaymentQueryHandler
}

func NewPaymentQueryBus(paymentRepo repository.PaymentRepository) *PaymentQueryBus {
	handler := query.NewPaymentQueryHandler(paymentRepo)
	return &PaymentQueryBus{handler: handler}
}

func (b *PaymentQueryBus) FindByID(ctx context.Context, qry query.FindPaymentByIDQuery) (query.PaymentResult, error) {
	return b.handler.FindByID(ctx, qry)
}

func (b *PaymentQueryBus) FindOverdues(ctx context.Context, qry query.FindOverduePaymentsQuery) (page.Page[query.PaymentResult], error) {
	return b.handler.FindOverdues(ctx, qry)
}

func (b *PaymentQueryBus) FindByStatus(ctx context.Context, qry query.FindPaymentsByStatusQuery) (page.Page[query.PaymentResult], error) {
	return b.handler.FindByStatus(ctx, qry)
}

func (b *PaymentQueryBus) FindByCustomer(ctx context.Context, qry query.FindPaymentsByCustomerQuery) (page.Page[query.PaymentResult], error) {
	return b.handler.FindByCustomer(ctx, qry)
}

func (b *PaymentQueryBus) FindByDateRange(ctx context.Context, qry query.FindPaymentsByDateRangeQuery) (page.Page[query.PaymentResult], error) {
	return b.handler.FindByDateRange(ctx, qry)
}

func (b *PaymentQueryBus) FindBySpecification(ctx context.Context, qry query.FindPaymentsBySpecification) (page.Page[query.PaymentResult], error) {
	return b.handler.FindBySpecification(ctx, qry)
}
