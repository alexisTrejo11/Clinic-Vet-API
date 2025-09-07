package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListPaymentsByStatusQuery struct {
	status     enum.PaymentStatus
	pagination page.PageInput
	ctx        context.Context
}

func NewListPaymentsByStatusQuery(status enum.PaymentStatus, pagination page.PageInput) ListPaymentsByStatusQuery {
	return ListPaymentsByStatusQuery{
		status:     status,
		pagination: pagination,
	}
}

type ListPaymentsByStatusHandler struct {
	repo repository.PaymentRepository
}

func NewListPaymentsByStatusHandler(repo repository.PaymentRepository) cqrs.QueryHandler[page.Page[[]PaymentResponse]] {
	return &ListPaymentsByStatusHandler{repo: repo}
}

func (h *ListPaymentsByStatusHandler) Handle(q cqrs.Query) (page.Page[[]PaymentResponse], error) {
	query := q.(ListPaymentsByStatusQuery)

	paymentPage, err := h.repo.ListByStatus(query.ctx, query.status, query.pagination)
	if err != nil {
		return page.Page[[]PaymentResponse]{}, err
	}

	paymentsResponse := mapPaymentsToResponses(paymentPage.Data)
	return page.NewPage(paymentsResponse, paymentPage.Metadata), nil
}
