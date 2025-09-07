package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListPaymentsByOwnerQuery struct {
	ownerID    valueobject.OwnerID
	pagination page.PageInput
	ctx        context.Context
}

func NewListPaymentsByUserQuery(ownerIDInt int, pagination page.PageInput) ListPaymentsByOwnerQuery {
	ownerID, _ := valueobject.NewOwnerID(ownerIDInt)

	return ListPaymentsByOwnerQuery{
		ownerID:    ownerID,
		pagination: pagination,
		ctx:        context.Background(),
	}
}

type ListByOwnerHandler struct {
	repository repository.PaymentRepository
}

func NewListByOwnerHandler(repository repository.PaymentRepository) cqrs.QueryHandler[page.Page[[]PaymentResponse]] {
	return &ListByOwnerHandler{repository: repository}
}

func (h *ListByOwnerHandler) Handle(q cqrs.Query) (page.Page[[]PaymentResponse], error) {
	query := q.(ListPaymentsByOwnerQuery)
	paymentsPage, err := h.repository.ListByPaidFrom(query.ctx, query.ownerID, query.pagination)
	if err != nil {
		return page.Page[[]PaymentResponse]{}, err
	}
	responses := mapPaymentsToResponses(paymentsPage.Data)
	return page.NewPage(responses, paymentsPage.Metadata), nil
}
