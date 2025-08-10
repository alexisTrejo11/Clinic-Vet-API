package paymentQuery

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListOverduePaymentsHandler interface {
	Handle(ctx context.Context, query ListOverduePaymentsQuery) (*page.Page[[]PaymentResponse], error)
}

type listOverduePaymentsHandlerImpl struct {
	paymentRepository paymentDomain.PaymentRepository
}

type ListOverduePaymentsQuery struct {
	pagination page.PageData
}

func NewListOverduePaymentsQuery(pagination page.PageData) ListOverduePaymentsQuery {
	return ListOverduePaymentsQuery{
		pagination: pagination,
	}
}

func NewListOverduePaymentsHandler(paymentRepository paymentDomain.PaymentRepository) ListOverduePaymentsHandler {
	return &listOverduePaymentsHandlerImpl{
		paymentRepository: paymentRepository,
	}
}

func (h *listOverduePaymentsHandlerImpl) Handle(ctx context.Context, query ListOverduePaymentsQuery) (*page.Page[[]PaymentResponse], error) {
	paymentsPage, err := h.paymentRepository.ListOverduePayments(ctx, query.pagination)
	if err != nil {
		return nil, err
	}

	response := mapPaymentsToResponses(paymentsPage.Data)
	return page.NewPage(response, paymentsPage.Metadata), nil
}
