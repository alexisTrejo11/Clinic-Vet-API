package paymentQuery

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListOverduePaymentsQuery struct {
	pagination page.PageData `json:"page_data"`
}

type ListOverduePaymentsQueryHandler interface {
	Handle(query ListOverduePaymentsQuery) (*page.Page[[]PaymentResponse], error)
}

type listOverduePaymentsHandler struct {
	paymentRepository paymentDomain.PaymentRepository
}

func NewListOverduePaymentsHandler(paymentRepository paymentDomain.PaymentRepository) ListOverduePaymentsQueryHandler {
	return &listOverduePaymentsHandler{
		paymentRepository: paymentRepository,
	}
}

func (h *listOverduePaymentsHandler) Handle(query ListOverduePaymentsQuery) (*page.Page[[]PaymentResponse], error) {
	paymentsPage, err := h.paymentRepository.ListOverduePayments(context.Background(), query.pagination)
	if err != nil {
		return nil, err
	}

	response := mapPaymentsToResponses(paymentsPage.Data)
	return page.NewPage(response, paymentsPage.Metadata), nil
}
