package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type FindPaymentsByCustomerQuery struct {
	customerID valueobject.CustomerID
	pagination page.PageInput
	ctx        context.Context
}

func NewFindPaymentsByUserQuery(customerIDInt uint, pagination page.PageInput) FindPaymentsByCustomerQuery {
	customerID := valueobject.NewCustomerID(customerIDInt)

	return FindPaymentsByCustomerQuery{
		customerID: customerID,
		pagination: pagination,
		ctx:        context.Background(),
	}
}

type FindByCustomerHandler struct {
	repository repository.PaymentRepository
}

func NewFindByCustomerHandler(repository repository.PaymentRepository) cqrs.QueryHandler[page.Page[PaymentResponse]] {
	return &FindByCustomerHandler{repository: repository}
}

func (h *FindByCustomerHandler) Handle(q cqrs.Query) (page.Page[PaymentResponse], error) {
	query := q.(FindPaymentsByCustomerQuery)
	paymentsPage, err := h.repository.FindByCustomerID(query.ctx, query.customerID, query.pagination)
	if err != nil {
		return page.Page[PaymentResponse]{}, err
	}
	responses := mapPaymentsToResponses(paymentsPage.Items)
	return page.NewPage(responses, paymentsPage.Metadata), nil
}
