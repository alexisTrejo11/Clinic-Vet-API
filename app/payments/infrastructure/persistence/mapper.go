package paymentRepo

import (
	"math/big"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func mapSQLCPaymentToDomain(sqlcPayment sqlc.Payment) (*paymentDomain.Payment, error) {
	status, err := paymentDomain.NewPaymentStatus(sqlcPayment.Status.(string))
	if err != nil {
		return nil, err
	}

	method, err := paymentDomain.NewPaymentMethod(sqlcPayment.Method.(string))
	if err != nil {
		return nil, err
	}

	payment := &paymentDomain.Payment{
		Id: int(sqlcPayment.ID),
		Amount: paymentDomain.Money{
			Amount:   sqlcPayment.Amount.Int.Int64(),
			Currency: sqlcPayment.Currency,
		},
		Status:        status,
		PaymentMethod: method,
		CreatedAt:     sqlcPayment.CreatedAt.Time,
		UpdatedAt:     sqlcPayment.UpdatedAt.Time,
	}

	if sqlcPayment.TransactionID.Valid {
		payment.TransactionId = &sqlcPayment.TransactionID.String
	}

	if sqlcPayment.Description.Valid {
		payment.Description = &sqlcPayment.Description.String
	}

	if sqlcPayment.Duedate.Valid {
		payment.DueDate = &sqlcPayment.Duedate.Time
	}

	if sqlcPayment.PaidAt.Valid {
		payment.PaidAt = &sqlcPayment.PaidAt.Time
	}

	if sqlcPayment.RefundedAt.Valid {
		payment.RefundedAt = &sqlcPayment.RefundedAt.Time
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
		Amount:        pgtype.Numeric{Int: big.NewInt(payment.Amount.Amount), Valid: true},
		Currency:      payment.Amount.Currency,
		TransactionID: pgtype.Text{String: *payment.TransactionId, Valid: payment.TransactionId != nil},
		Status:        string(payment.Status),
		Method:        string(payment.PaymentMethod),
		Description:   pgtype.Text{String: *payment.Description, Valid: payment.Description != nil},
		Duedate:       pgtype.Timestamptz{Time: *payment.DueDate, Valid: payment.DueDate != nil},
		PaidAt:        pgtype.Timestamptz{Time: *payment.PaidAt, Valid: payment.PaidAt != nil},
		RefundedAt:    pgtype.Timestamptz{Time: *payment.RefundedAt, Valid: payment.RefundedAt != nil},
	}
}

func mapDomainToUpdateParams(payment *paymentDomain.Payment) sqlc.UpdatePaymentParams {
	return sqlc.UpdatePaymentParams{
		ID:            int32(payment.Id),
		Amount:        pgtype.Numeric{Int: big.NewInt(payment.Amount.Amount), Valid: true},
		Currency:      payment.Amount.Currency,
		TransactionID: pgtype.Text{String: *payment.TransactionId, Valid: payment.TransactionId != nil},
		Status:        string(payment.Status),
		Method:        string(payment.PaymentMethod),
		UserID:        int32(payment.UserId),
		Description:   pgtype.Text{String: *payment.Description, Valid: payment.Description != nil},
		Duedate:       pgtype.Timestamptz{Time: *payment.DueDate, Valid: payment.DueDate != nil},
		PaidAt:        pgtype.Timestamptz{Time: *payment.PaidAt, Valid: payment.PaidAt != nil},
		RefundedAt:    pgtype.Timestamptz{Time: *payment.RefundedAt, Valid: payment.RefundedAt != nil},
	}
}
