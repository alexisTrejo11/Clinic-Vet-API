package query

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListOverduePaymentsHandler struct {
	paymentRepository repository.PaymentRepository
}

type ListOverduePaymentsQuery struct {
	pagination page.PageInput
	ctx        context.Context
}

func NewListOverduePaymentsQuery(pagination page.PageInput) ListOverduePaymentsQuery {
	return ListOverduePaymentsQuery{
		pagination: pagination,
	}
}

func NewListOverduePaymentsHandler(paymentRepository repository.PaymentRepository) cqrs.QueryHandler[page.Page[[]PaymentResponse]] {
	return &ListOverduePaymentsHandler{
		paymentRepository: paymentRepository,
	}
}

func (h *ListOverduePaymentsHandler) Handle(q cqrs.Query) (page.Page[[]PaymentResponse], error) {
	query := q.(ListOverduePaymentsQuery)

	paymentsPage, err := h.paymentRepository.ListOverduePayments(query.ctx, query.pagination)
	if err != nil {
		return page.Page[[]PaymentResponse]{}, err
	}

	response := mapPaymentsToResponses(paymentsPage.Data)
	return page.NewPage(response, paymentsPage.Metadata), nil
}
