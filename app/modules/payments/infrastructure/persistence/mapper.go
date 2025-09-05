package persistence

import (
	"fmt"
	"math/big"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func sqlcToEntity(sqlcPayment sqlc.Payment) (entity.Payment, error) {
	status, err := enum.NewPaymentStatus(string(sqlcPayment.Status))
	if err != nil {
		return entity.Payment{}, fmt.Errorf("invalid payment status: %w", err)
	}

	method, err := enum.NewPaymentMethod(string(sqlcPayment.Method))
	if err != nil {
		return entity.Payment{}, fmt.Errorf("invalid payment method: %w", err)
	}

	money, err := valueobject.NewMoney(float64(sqlcPayment.Amount.Int.Int64()), sqlcPayment.Currency)
	if err != nil {
		return entity.Payment{}, fmt.Errorf("invalid money amount: %w", err)
	}

	paymentID, err := valueobject.NewPaymentID(int(sqlcPayment.ID))
	if err != nil {
		return entity.Payment{}, err
	}

	builder := entity.NewPaymentBuilder().
		WithID(paymentID).
		WithAmount(money).
		WithStatus(status).
		WithPaymentMethod(method).
		WithCreatedAt(sqlcPayment.CreatedAt.Time).
		WithUpdatedAt(sqlcPayment.UpdatedAt.Time)

	if sqlcPayment.TransactionID.Valid {
		builder = builder.WithTransactionID(&sqlcPayment.TransactionID.String)
	}

	return *builder.Build(), nil
}

func sqlcToEntityList(sqlcPayments []sqlc.Payment) ([]entity.Payment, error) {
	var payments []entity.Payment
	for _, sqlcPayment := range sqlcPayments {
		payment, err := sqlcToEntity(sqlcPayment)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func entityToCreateParams(payment *entity.Payment) sqlc.CreatePaymentParams {
	return sqlc.CreatePaymentParams{
		Amount:        pgtype.Numeric{Int: big.NewInt(payment.GetAmount().Amount), Valid: true},
		Currency:      payment.GetAmount().Currency,
		TransactionID: pgtype.Text{String: *payment.GetTransactionID(), Valid: payment.GetTransactionID() != nil},
		Status:        models.PaymentStatus(string(payment.GetStatus())),
		Method:        models.PaymentMethod(string(payment.GetPaymentMethod())),
		Description:   pgtype.Text{String: *payment.GetDescription(), Valid: payment.GetDescription() != nil},
		Duedate:       pgtype.Timestamptz{Time: *payment.GetDueDate(), Valid: payment.GetDueDate() != nil},
		PaidAt:        pgtype.Timestamptz{Time: *payment.GetPaidAt(), Valid: payment.GetPaidAt() != nil},
		RefundedAt:    pgtype.Timestamptz{Time: *payment.GetRefundedAt(), Valid: payment.GetRefundedAt() != nil},
	}
}

func mapDomainToUpdateParams(payment *entity.Payment) sqlc.UpdatePaymentParams {
	return sqlc.UpdatePaymentParams{
		ID:            int32(payment.GetID().GetValue()),
		Amount:        pgtype.Numeric{Int: big.NewInt(payment.GetAmount().Amount), Valid: true},
		Currency:      payment.GetAmount().Currency,
		TransactionID: pgtype.Text{String: *payment.GetTransactionID(), Valid: payment.GetTransactionID() != nil},
		Status:        models.PaymentStatus(string(payment.GetStatus())),
		Method:        models.PaymentMethod(string(payment.GetPaymentMethod())),
		UserID:        int32(payment.GetUserID().GetValue()),
		Description:   pgtype.Text{String: *payment.GetDescription(), Valid: payment.GetDescription() != nil},
		Duedate:       pgtype.Timestamptz{Time: *payment.GetDueDate(), Valid: payment.GetDueDate() != nil},
		PaidAt:        pgtype.Timestamptz{Time: *payment.GetPaidAt(), Valid: payment.GetPaidAt() != nil},
		RefundedAt:    pgtype.Timestamptz{Time: *payment.GetRefundedAt(), Valid: payment.GetRefundedAt() != nil},
	}
}
