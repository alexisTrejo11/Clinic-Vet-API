package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type SearchPaymentsQuery struct {
	ctx  context.Context
	spec any
}

type SearchPaymentsHandler struct {
	repository repository.PaymentRepository
}

func NewSearchPaymentsHandler(repository repository.PaymentRepository) cqrs.QueryHandler[page.Page[PaymentResponse]] {
	return &SearchPaymentsHandler{repository: repository}
}

func (h *SearchPaymentsHandler) Handle(q cqrs.Query) (page.Page[PaymentResponse], error) {
	query, ok := q.(SearchPaymentsQuery)
	if !ok {
		return page.Page[PaymentResponse]{}, nil
	}

	paymentsPage, err := h.repository.FindBySpecification(query.ctx, query.spec)
	if err != nil {
		return page.Page[PaymentResponse]{}, err
	}

	var responses []PaymentResponse
	for _, payment := range paymentsPage.Items {
		responses = append(responses, NewPaymentResponse(&payment))
	}

	return page.NewPage(responses, paymentsPage.Metadata), nil
}
