package paymentQuery

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListByOwnerQuery struct {
	UserId        int `json:"user_id"`
	page.PageData `json:"pagination"`
}

type ListByOwnerQueryHandler interface {
	Handle(query ListByOwnerQuery) (*page.Page[[]PaymentResponse], error)
}

type ListByOwnerHandler struct {
	repository paymentDomain.PaymentRepository
}

func NewListByOwnerHandler(repository paymentDomain.PaymentRepository) *ListByOwnerHandler {
	return &ListByOwnerHandler{repository: repository}
}

func (h *ListByOwnerHandler) Handle(query ListByOwnerQuery) (*page.Page[[]PaymentResponse], error) {
	paymentsPage, err := h.repository.ListByUserId(context.Background(), query.UserId, query.PageData)
	if err != nil {
		return nil, err
	}
	responses := mapPaymentsToResponses(paymentsPage.Data)
	return page.NewPage(responses, paymentsPage.Metadata), nil
}
