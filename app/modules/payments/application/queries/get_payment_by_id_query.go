package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type GetPaymentByIDQuery struct {
	id  valueobject.PaymentID
	ctx context.Context
}

func NewGetPaymentByIDQuery(idInt uint) (*GetPaymentByIDQuery, error) {
	paymentID := valueobject.NewPaymentID(idInt)
	if paymentID.IsZero() {
		return nil, apperror.FieldValidationError("id", "0", "invalid payment id")
	}

	return &GetPaymentByIDQuery{id: paymentID}, nil
}

type GetPaymentByIDHandler struct {
	repository repository.PaymentRepository
}

func NewGetPaymentByIDHandler(repository repository.PaymentRepository) cqrs.QueryHandler[PaymentResponse] {
	return &GetPaymentByIDHandler{
		repository: repository,
	}
}

func (h *GetPaymentByIDHandler) Handle(q cqrs.Query) (PaymentResponse, error) {
	query := q.(GetPaymentByIDQuery)
	payment, err := h.repository.FindByID(query.ctx, query.id)
	if err != nil {
		return PaymentResponse{}, err
	}

	return NewPaymentResponse(&payment), nil
}
