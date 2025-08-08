package paymentQuery

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListByStatusQuery struct {
	Status     paymentDomain.PaymentStatus `json:"status"`
	Pagination page.PageData               `json:"pagination"`
}

type ListByStatusQueryHandler interface {
	Handle(query ListByStatusQuery) (*page.Page[[]PaymentResponse], error)
}

type listByStatusQueryHandler struct {
	repo paymentDomain.PaymentRepository
}

func NewListByStatusQueryHandler(repo paymentDomain.PaymentRepository) ListByStatusQueryHandler {
	return &listByStatusQueryHandler{repo: repo}
}

func (h *listByStatusQueryHandler) Handle(query ListByStatusQuery) (*page.Page[[]PaymentResponse], error) {
	if !query.Status.IsValid() {
		return nil, paymentDomain.InvalidPaymentStatusErr(query.Status)
	}

	paymentPage, err := h.repo.ListByStatus(context.Background(), query.Status, query.Pagination)
	if err != nil {
		return nil, err
	}

	response := mapPaymentsToResponses(paymentPage.Data)
	return page.NewPage(response, paymentPage.Metadata), nil
}
