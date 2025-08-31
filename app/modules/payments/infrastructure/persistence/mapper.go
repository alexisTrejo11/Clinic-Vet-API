package paymentRepo

import (
	"math/big"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func mapSQLCPaymentToDomain(sqlcPayment sqlc.Payment) (entity.Payment, error) {
	status, err := enum.NewPaymentStatus(string(sqlcPayment.Status))
	if err != nil {
		return entity.Payment{}, err
	}

	method, err := enum.NewPaymentMethod(string(sqlcPayment.Method))
	if err != nil {
		return entity.Payment{}, err
	}

	payment := &entity.Payment{}

	payment.SetID(int(sqlcPayment.ID))
	payment.SetAmount(valueobject.Money{
		Amount:   sqlcPayment.Amount.Int.Int64(),
		Currency: sqlcPayment.Currency,
	})
	payment.SetStatus(status)
	payment.SetPaymentMethod(method)
	payment.SetCreatedAt(sqlcPayment.CreatedAt.Time)
	payment.SetUpdatedAt(sqlcPayment.UpdatedAt.Time)

	if sqlcPayment.TransactionID.Valid {
		payment.SetTransactionID(&sqlcPayment.TransactionID.String)
	}

	if sqlcPayment.Description.Valid {
		payment.SetDescription(&sqlcPayment.Description.String)
	}

	if sqlcPayment.Duedate.Valid {
		payment.SetDueDate(&sqlcPayment.Duedate.Time)
	}

	if sqlcPayment.PaidAt.Valid {
		payment.SetPaidAt(&sqlcPayment.PaidAt.Time)
	}

	if sqlcPayment.RefundedAt.Valid {
		payment.SetRefundedAt(&sqlcPayment.RefundedAt.Time)
	}

	return *payment, nil
}

func mapSQLCPaymentListToDomain(sqlcPayments []sqlc.Payment) ([]entity.Payment, error) {
	var payments []entity.Payment
	for _, sqlcPayment := range sqlcPayments {
		payment, err := mapSQLCPaymentToDomain(sqlcPayment)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func mapDomainToCreateParams(payment *entity.Payment) sqlc.CreatePaymentParams {
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
		ID:            int32(payment.GetID()),
		Amount:        pgtype.Numeric{Int: big.NewInt(payment.GetAmount().Amount), Valid: true},
		Currency:      payment.GetAmount().Currency,
		TransactionID: pgtype.Text{String: *payment.GetTransactionID(), Valid: payment.GetTransactionID() != nil},
		Status:        models.PaymentStatus(string(payment.GetStatus())),
		Method:        models.PaymentMethod(string(payment.GetPaymentMethod())),
		UserID:        int32(payment.GetUserID()),
		Description:   pgtype.Text{String: *payment.GetDescription(), Valid: payment.GetDescription() != nil},
		Duedate:       pgtype.Timestamptz{Time: *payment.GetDueDate(), Valid: payment.GetDueDate() != nil},
		PaidAt:        pgtype.Timestamptz{Time: *payment.GetPaidAt(), Valid: payment.GetPaidAt() != nil},
		RefundedAt:    pgtype.Timestamptz{Time: *payment.GetRefundedAt(), Valid: payment.GetRefundedAt() != nil},
	}
}
