package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type FindOverduePaymentsHandler struct {
	paymentRepository repository.PaymentRepository
}

type FindOverduePaymentsQuery struct {
	pagination page.PageInput
	ctx        context.Context
}

func NewFindOverduePaymentsQuery(pagination page.PageInput) FindOverduePaymentsQuery {
	return FindOverduePaymentsQuery{
		pagination: pagination,
	}
}

func NewFindOverduePaymentsHandler(paymentRepository repository.PaymentRepository) cqrs.QueryHandler[page.Page[PaymentResponse]] {
	return &FindOverduePaymentsHandler{
		paymentRepository: paymentRepository,
	}
}

func (h *FindOverduePaymentsHandler) Handle(q cqrs.Query) (page.Page[PaymentResponse], error) {
	query := q.(FindOverduePaymentsQuery)

	paymentsPage, err := h.paymentRepository.FindOverdue(query.ctx, query.pagination)
	if err != nil {
		return page.Page[PaymentResponse]{}, err
	}

	response := mapPaymentsToResponses(paymentsPage.Items)
	return page.NewPage(response, paymentsPage.Metadata), nil
}
