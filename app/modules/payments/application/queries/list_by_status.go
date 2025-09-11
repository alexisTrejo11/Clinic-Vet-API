package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type FindPaymentsByStatusQuery struct {
	status     enum.PaymentStatus
	pagination page.PageInput
	ctx        context.Context
}

func NewFindPaymentsByStatusQuery(status string, pagination page.PageInput) (*FindPaymentsByStatusQuery, error) {
	statusEnum, err := enum.ParsePaymentStatus(status)
	if err != nil {
		return nil, err
	}

	return &FindPaymentsByStatusQuery{
		status:     statusEnum,
		pagination: pagination,
	}, nil
}

type FindPaymentsByStatusHandler struct {
	repo repository.PaymentRepository
}

func NewFindPaymentsByStatusHandler(repo repository.PaymentRepository) cqrs.QueryHandler[page.Page[PaymentResponse]] {
	return &FindPaymentsByStatusHandler{repo: repo}
}

func (h *FindPaymentsByStatusHandler) Handle(q cqrs.Query) (page.Page[PaymentResponse], error) {
	query := q.(FindPaymentsByStatusQuery)

	paymentPage, err := h.repo.FindByStatus(query.ctx, query.status, query.pagination)
	if err != nil {
		return page.Page[PaymentResponse]{}, err
	}

	paymentsResponse := mapPaymentsToResponses(paymentPage.Items)
	return page.NewPage(paymentsResponse, paymentPage.Metadata), nil
}
