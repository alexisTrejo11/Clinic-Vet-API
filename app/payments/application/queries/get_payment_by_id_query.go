package paymentQuery

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
)

type GetPaymentByIdQuery struct {
	Id int `json:"id"`
}

type GetPaymentByIdQueryHandler interface {
	Handle(query GetPaymentByIdQuery) (*PaymentResponse, error)
}

type GetPaymentByIdHandler struct {
	repository paymentDomain.PaymentRepository
}

func NewGetPaymentByIdHandler(repository paymentDomain.PaymentRepository) *GetPaymentByIdHandler {
	return &GetPaymentByIdHandler{repository: repository}
}

func (h *GetPaymentByIdHandler) Handle(query GetPaymentByIdQuery) (*PaymentResponse, error) {
	payment, err := h.repository.GetById(context.Background(), query.Id)
	if err != nil {
		return nil, err
	}

	if payment == nil {
		return nil, paymentDomain.PaymentNotFoundErr(query.Id)
	}

	response := NewPaymentResponse(payment)
	return &response, nil
}
