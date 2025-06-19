package repository

/*
import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type paymentRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewPaymentRepository(queries *sqlc.Queries) PaymentRepository {
	return &paymentRepositoryImpl{
		queries: queries,
	}
}

func (ar paymentRepositoryImpl) CreatePayment(params sqlc.CreatePaymentParams) error {
	if _, err := ar.queries.CreatePayment(context.Background(), params); err != nil {
		return err
	}

	return nil
}

func (ar paymentRepositoryImpl) GetPaymentByID(PaymentId int32) (*sqlc.Payment, error) {
	Payment, err := ar.queries.GetPaymentByID(context.Background(), PaymentId)
	if err != nil {
		return nil, err
	}

	return &Payment, nil
}

func (ar paymentRepositoryImpl) UpdatePayment(updateParams sqlc.UpdatePaymentParams) error {
	err := ar.queries.UpdatePayment(context.Background(), updateParams)
	if err != nil {
		return err
	}

	return nil
}

func (ar paymentRepositoryImpl) DeletePayment(PaymentId int32) error {
	if err := ar.queries.DeletePayment(context.Background(), PaymentId); err != nil {
		return err
	}

	return nil
}
*/
