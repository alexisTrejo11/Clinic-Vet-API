package query

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListPaymentsByUserHandler interface {
	Handle(ctx context.Context, query ListPaymentsByUserQuery) (page.Page[[]PaymentResponse], error)
}

type ListPaymentsByUserQuery struct {
	userID     int
	pagination page.PageData
}

func NewListPaymentsByUserQuery(userID int, pagination page.PageData) ListPaymentsByUserQuery {
	return ListPaymentsByUserQuery{
		userID:     userID,
		pagination: pagination,
	}
}

type listByUserHandlerImpl struct {
	repository repository.PaymentRepository
}

func NewListByUserHandler(repository repository.PaymentRepository) ListPaymentsByUserHandler {
	return &listByUserHandlerImpl{repository: repository}
}

func (h *listByUserHandlerImpl) Handle(ctx context.Context, query ListPaymentsByUserQuery) (page.Page[[]PaymentResponse], error) {
	paymentsPage, err := h.repository.ListByUserID(context.Background(), query.userID, query.pagination)
	if err != nil {
		return page.Page[[]PaymentResponse]{}, err
	}
	responses := mapPaymentsToResponses(paymentsPage.Data)
	return page.NewPage(responses, paymentsPage.Metadata), nil
}
