package paymentRepo

import (
	"math/big"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func mapSQLCPaymentToDomain(sqlcPayment sqlc.Payment) (*paymentDomain.Payment, error) {
	status, err := paymentDomain.NewPaymentStatus(string(sqlcPayment.Status))
	if err != nil {
		return nil, err
	}

	method, err := paymentDomain.NewPaymentMethod(string(sqlcPayment.Method))
	if err != nil {
		return nil, err
	}

	payment := &paymentDomain.Payment{}

	payment.SetId(int(sqlcPayment.ID))
	payment.SetAmount(paymentDomain.Money{
		Amount:   sqlcPayment.Amount.Int.Int64(),
		Currency: sqlcPayment.Currency,
	})
	payment.SetStatus(status)
	payment.SetPaymentMethod(method)
	payment.SetCreatedAt(sqlcPayment.CreatedAt.Time)
	payment.SetUpdatedAt(sqlcPayment.UpdatedAt.Time)

	if sqlcPayment.TransactionID.Valid {
		payment.SetTransactionId(&sqlcPayment.TransactionID.String)
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

	return payment, nil
}

func mapSQLCPaymentListToDomain(sqlcPayments []sqlc.Payment) ([]paymentDomain.Payment, error) {
	var payments []paymentDomain.Payment
	for _, sqlcPayment := range sqlcPayments {
		payment, err := mapSQLCPaymentToDomain(sqlcPayment)
		if err != nil {
			return nil, err
		}
		payments = append(payments, *payment)
	}
	return payments, nil
}

func mapDomainToCreateParams(payment *paymentDomain.Payment) sqlc.CreatePaymentParams {
	return sqlc.CreatePaymentParams{
		Amount:        pgtype.Numeric{Int: big.NewInt(payment.GetAmount().Amount), Valid: true},
		Currency:      payment.GetAmount().Currency,
		TransactionID: pgtype.Text{String: *payment.GetTransactionId(), Valid: payment.GetTransactionId() != nil},
		Status:        models.PaymentStatus(string(payment.GetStatus())),
		Method:        models.PaymentMethod(string(payment.GetPaymentMethod())),
		Description:   pgtype.Text{String: *payment.GetDescription(), Valid: payment.GetDescription() != nil},
		Duedate:       pgtype.Timestamptz{Time: *payment.GetDueDate(), Valid: payment.GetDueDate() != nil},
		PaidAt:        pgtype.Timestamptz{Time: *payment.GetPaidAt(), Valid: payment.GetPaidAt() != nil},
		RefundedAt:    pgtype.Timestamptz{Time: *payment.GetRefundedAt(), Valid: payment.GetRefundedAt() != nil},
	}
}

func mapDomainToUpdateParams(payment *paymentDomain.Payment) sqlc.UpdatePaymentParams {
	return sqlc.UpdatePaymentParams{
		ID:            int32(payment.GetId()),
		Amount:        pgtype.Numeric{Int: big.NewInt(payment.GetAmount().Amount), Valid: true},
		Currency:      payment.GetAmount().Currency,
		TransactionID: pgtype.Text{String: *payment.GetTransactionId(), Valid: payment.GetTransactionId() != nil},
		Status:        models.PaymentStatus(string(payment.GetStatus())),
		Method:        models.PaymentMethod(string(payment.GetPaymentMethod())),
		UserID:        int32(payment.GetUserId()),
		Description:   pgtype.Text{String: *payment.GetDescription(), Valid: payment.GetDescription() != nil},
		Duedate:       pgtype.Timestamptz{Time: *payment.GetDueDate(), Valid: payment.GetDueDate() != nil},
		PaidAt:        pgtype.Timestamptz{Time: *payment.GetPaidAt(), Valid: payment.GetPaidAt() != nil},
		RefundedAt:    pgtype.Timestamptz{Time: *payment.GetRefundedAt(), Valid: payment.GetRefundedAt() != nil},
	}
}
