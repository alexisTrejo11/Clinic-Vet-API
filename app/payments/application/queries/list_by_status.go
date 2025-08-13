package paymentQuery

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListPaymentsByStatusQuery struct {
	status     paymentDomain.PaymentStatus
	pagination page.PageData
}

func NewListPaymentsByStatusQuery(status paymentDomain.PaymentStatus, pagination page.PageData) ListPaymentsByStatusQuery {
	return ListPaymentsByStatusQuery{
		status:     status,
		pagination: pagination,
	}
}

type ListPaymentsByStatusHandler interface {
	Handle(ctx context.Context, query ListPaymentsByStatusQuery) (page.Page[[]PaymentResponse], error)
}

type listPaymentsByStatusHandler struct {
	repo paymentDomain.PaymentRepository
}

func NewListPaymentsByStatusHandler(repo paymentDomain.PaymentRepository) ListPaymentsByStatusHandler {
	return &listPaymentsByStatusHandler{repo: repo}
}

func (h *listPaymentsByStatusHandler) Handle(ctx context.Context, query ListPaymentsByStatusQuery) (page.Page[[]PaymentResponse], error) {
	if !query.status.IsValid() {
		return page.Page[[]PaymentResponse]{}, paymentDomain.InvalidPaymentStatusErr(query.status)
	}

	paymentPage, err := h.repo.ListByStatus(ctx, query.status, query.pagination)
	if err != nil {
		return page.Page[[]PaymentResponse]{}, err
	}

	response := mapPaymentsToResponses(paymentPage.Data)
	return *page.NewPage(response, paymentPage.Metadata), nil
}
