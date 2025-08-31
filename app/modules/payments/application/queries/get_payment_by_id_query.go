package query

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type GetPaymentByIDQuery struct {
	id int
}

func NewGetPaymentByIDQuery(id int) GetPaymentByIDQuery {
	return GetPaymentByIDQuery{id: id}
}

type GetPaymentByIDHandler interface {
	Handle(ctx context.Context, query GetPaymentByIDQuery) (PaymentResponse, error)
}

type getPaymentByIDHandlerImpl struct {
	repository repository.PaymentRepository
}

func NewGetPaymentByIDHandler(repository repository.PaymentRepository) GetPaymentByIDHandler {
	return &getPaymentByIDHandlerImpl{
		repository: repository,
	}
}

func (h *getPaymentByIDHandlerImpl) Handle(ctx context.Context, query GetPaymentByIDQuery) (PaymentResponse, error) {
	payment, err := h.repository.GetByID(ctx, query.id)
	if err != nil {
		return PaymentResponse{}, err
	}

	return NewPaymentResponse(&payment), nil
}
