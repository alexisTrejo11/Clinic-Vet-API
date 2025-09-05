package query

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListPaymentsByUserQuery struct {
	userID     int
	pagination page.PageInput
	ctx        context.Context
}

func NewListPaymentsByUserQuery(userID int, pagination page.PageInput) ListPaymentsByUserQuery {
	return ListPaymentsByUserQuery{
		userID:     userID,
		pagination: pagination,
		ctx:        context.Background(),
	}
}

type ListByUserHandler struct {
	repository repository.PaymentRepository
}

func NewListByUserHandler(repository repository.PaymentRepository) cqrs.QueryHandler[page.Page[[]PaymentResponse]] {
	return &ListByUserHandler{repository: repository}
}

func (h *ListByUserHandler) Handle(q cqrs.Query) (page.Page[[]PaymentResponse], error) {
	query := q.(ListPaymentsByUserQuery)
	paymentsPage, err := h.repository.ListByUserID(query.ctx, query.userID, query.pagination)
	if err != nil {
		return page.Page[[]PaymentResponse]{}, err
	}
	responses := mapPaymentsToResponses(paymentsPage.Data)
	return page.NewPage(responses, paymentsPage.Metadata), nil
}
