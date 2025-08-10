package paymentQuery

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListPaymentsByUserHandler interface {
	Handle(ctx context.Context, query ListPaymentsByUserQuery) (*page.Page[[]PaymentResponse], error)
}

type ListPaymentsByUserQuery struct {
	userId     int
	pagination page.PageData
}

func NewListPaymentsByUserQuery(userId int, pagination page.PageData) ListPaymentsByUserQuery {
	return ListPaymentsByUserQuery{
		userId:     userId,
		pagination: pagination,
	}
}

type listByUserHandlerImpl struct {
	repository paymentDomain.PaymentRepository
}

func NewListByUserHandler(repository paymentDomain.PaymentRepository) ListPaymentsByUserHandler {
	return &listByUserHandlerImpl{repository: repository}
}

func (h *listByUserHandlerImpl) Handle(ctx context.Context, query ListPaymentsByUserQuery) (*page.Page[[]PaymentResponse], error) {
	paymentsPage, err := h.repository.ListByUserId(context.Background(), query.userId, query.pagination)
	if err != nil {
		return nil, err
	}
	responses := mapPaymentsToResponses(paymentsPage.Data)
	return page.NewPage(responses, paymentsPage.Metadata), nil
}
