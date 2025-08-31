package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListPaymentsByStatusQuery struct {
	status     enum.PaymentStatus
	pagination page.PageData
}

func NewListPaymentsByStatusQuery(status enum.PaymentStatus, pagination page.PageData) ListPaymentsByStatusQuery {
	return ListPaymentsByStatusQuery{
		status:     status,
		pagination: pagination,
	}
}

type ListPaymentsByStatusHandler interface {
	Handle(ctx context.Context, query ListPaymentsByStatusQuery) (page.Page[[]PaymentResponse], error)
}

type listPaymentsByStatusHandler struct {
	repo repository.PaymentRepository
}

func NewListPaymentsByStatusHandler(repo repository.PaymentRepository) ListPaymentsByStatusHandler {
	return &listPaymentsByStatusHandler{repo: repo}
}

func (h *listPaymentsByStatusHandler) Handle(ctx context.Context, query ListPaymentsByStatusQuery) (page.Page[[]PaymentResponse], error) {
	paymentPage, err := h.repo.ListByStatus(ctx, query.status, query.pagination)
	if err != nil {
		return page.Page[[]PaymentResponse]{}, err
	}

	paymentsResponse := mapPaymentsToResponses(paymentPage.Data)
	return page.NewPage(paymentsResponse, paymentPage.Metadata), nil
}
