package paymentQuery

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
)

type GetPaymentByIdQuery struct {
	id int `json:"id"`
}

func NewGetPaymentByIdQuery(id int) GetPaymentByIdQuery {
	return GetPaymentByIdQuery{id: id}
}

type GetPaymentByIdHandler interface {
	Handle(ctx context.Context, query GetPaymentByIdQuery) (*PaymentResponse, error)
}

type getPaymentByIdHandlerImpl struct {
	repository paymentDomain.PaymentRepository
}

func NewGetPaymentByIdHandler(repository paymentDomain.PaymentRepository) GetPaymentByIdHandler {
	return &getPaymentByIdHandlerImpl{repository: repository}
}

func (h *getPaymentByIdHandlerImpl) Handle(ctx context.Context, query GetPaymentByIdQuery) (*PaymentResponse, error) {
	payment, err := h.repository.GetById(ctx, query.id)
	if err != nil {
		return nil, err
	}

	if payment == nil {
		return nil, paymentDomain.PaymentNotFoundErr(query.id)
	}

	response := NewPaymentResponse(payment)
	return &response, nil
}
