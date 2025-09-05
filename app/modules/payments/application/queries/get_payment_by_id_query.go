package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type GetPaymentByIDQuery struct {
	id  valueobject.PaymentID
	ctx context.Context
}

func NewGetPaymentByIDQuery(idInt int) (GetPaymentByIDQuery, error) {
	paymentID, err := valueobject.NewPaymentID(idInt)
	if err != nil {
		return GetPaymentByIDQuery{}, apperror.MappingError([]string{err.Error()}, "constructor", "command", "payment")
	}

	q := &GetPaymentByIDQuery{id: paymentID}

	return *q, nil
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
	payment, err := h.repository.GetByID(query.ctx, query.id.GetValue())
	if err != nil {
		return PaymentResponse{}, err
	}

	return NewPaymentResponse(&payment), nil
}
