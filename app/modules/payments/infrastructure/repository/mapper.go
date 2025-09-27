package repositoryimpl

import (
	"clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"
	"fmt"
	"math/big"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func entityToCreateParams(p payment.Payment) *sqlc.CreatePaymentParams {
	params := &sqlc.CreatePaymentParams{
		Amount:           moneyToNumeric(p.Amount()),
		Currency:         p.Currency(),
		Status:           models.PaymentStatus(p.Status().String()),
		Method:           models.PaymentMethod(p.Method().String()),
		MedSessionID:     pgtype.Int4{Int32: int32(p.Entity.ID().Value()), Valid: true},
		TransactionID:    stringToPgText(p.TransactionID()),
		Description:      stringToPgText(p.Description()),
		DueDate:          timeToPgTimestamptzPtr(p.DueDate()),
		PaidAt:           timeToPgTimestamptzPtr(p.PaidAt()),
		RefundedAt:       timeToPgTimestamptzPtr(p.RefundedAt()),
		PaidByCustomerID: customerIDToPgInt4(p.PaidByCustomer()),
	}

	if p.MedSessionID() != nil {
		params.MedSessionID = MedSessionIDToPgInt4(*p.MedSessionID())
	}

	if p.InvoiceID() != nil {
		params.InvoiceID = stringToPgText(p.InvoiceID())
	}

	if p.RefundAmount() != nil {
		params.RefundAmount = moneyToNumeric(*p.RefundAmount())
	}

	if p.FailureReason() != nil {
		params.FailureReason = stringToPgText(p.FailureReason())
	}

	return params
}

func entityToUpdateParams(p *payment.Payment) *sqlc.UpdatePaymentParams {
	params := &sqlc.UpdatePaymentParams{
		ID:               int32(p.ID().Value()),
		Amount:           moneyToNumeric(p.Amount()),
		Currency:         p.Currency(),
		MedSessionID:     pgtype.Int4{Int32: int32(p.Entity.ID().Value()), Valid: true},
		Status:           models.PaymentStatus(p.Status().String()),
		Method:           models.PaymentMethod(p.Method().String()),
		TransactionID:    stringToPgText(p.TransactionID()),
		Description:      stringToPgText(p.Description()),
		DueDate:          timeToPgTimestamptz(*p.DueDate()),
		PaidAt:           timeToPgTimestamptzPtr(p.PaidAt()),
		RefundedAt:       timeToPgTimestamptzPtr(p.RefundedAt()),
		PaidByCustomerID: customerIDToPgInt4(p.PaidByCustomer()),
	}

	if p.MedSessionID() != nil {
		params.MedSessionID = MedSessionIDToPgInt4(*p.MedSessionID())
	}

	if p.InvoiceID() != nil {
		params.InvoiceID = stringToPgText(p.InvoiceID())
	}

	if p.RefundAmount() != nil {
		params.RefundAmount = moneyToNumeric(*p.RefundAmount())
	}

	if p.FailureReason() != nil {
		params.FailureReason = stringToPgText(p.FailureReason())
	}

	return params
}

func sqlcRowToEntity(row sqlc.Payment) *payment.Payment {
	amount, _ := numericToMoney(row.Amount, row.Currency)
	builder := payment.NewPaymentBuilder().
		WithID(valueobject.NewPaymentID(uint(row.ID))).
		WithInvoiceID(row.InvoiceID.String).
		WithPaidByCustomer(pgInt4ToCustomerID(row.PaidByCustomerID)).
		WithAmount(amount).
		WithMedSessionID(pgInt4ToMedSessionID(row.MedSessionID)).
		WithPaymentMethod(enum.PaymentMethod(row.Method)).
		WithStatus(enum.PaymentStatus(row.Status)).
		WithDueDate(pgTimestamptzToTime(row.DueDate)).
		WithIsActive(row.IsActive).
		WithTimeStamps(row.CreatedAt.Time, row.UpdatedAt.Time)

	if row.TransactionID.Valid {
		builder.WithTransactionID(row.TransactionID.String)
	}

	if row.Description.Valid {
		builder.WithDescription(&row.Description.String)
	}

	if row.PaidAt.Valid {
		paidAt := row.PaidAt.Time
		builder.WithPaidAt(paidAt)
	}

	if row.RefundedAt.Valid {
		refundedAt := row.RefundedAt.Time
		builder.WithRefundedAt(refundedAt)
	}

	if row.MedSessionID.Valid {
		medSessionID := pgInt4ToMedSessionID(row.MedSessionID)
		builder.WithMedSessionID(medSessionID)
	}

	if row.InvoiceID.Valid {
		builder.WithInvoiceID(row.InvoiceID.String)
	}

	if row.RefundAmount.Valid {
		refundAmount, _ := numericToMoney(row.RefundAmount, row.Currency)
		builder.WithRefundAmount(refundAmount)
	}

	if row.FailureReason.Valid {
		builder.WithFailureReason(row.FailureReason.String)
	}

	return builder.Build()
}

func (r *SqlcPaymentRepository) sqlRowsToEntities(sqlRows []sqlc.Payment) ([]payment.Payment, error) {
	payments := make([]payment.Payment, len(sqlRows))
	for i, row := range sqlRows {
		paymentEntity := sqlcRowToEntity(row)
		if paymentEntity == nil {
			return nil, fmt.Errorf("failed to convert SQL row to payment at index %d", i)
		}
		payments[i] = *paymentEntity
	}
	return payments, nil
}

func moneyToNumeric(money valueobject.Money) pgtype.Numeric {
	amount := money.Amount()
	return pgtype.Numeric{
		Int:   big.NewInt(int64(amount.Float64())),
		Valid: true,
	}
}

func numericToMoney(num pgtype.Numeric, currency string) (valueobject.Money, error) {
	if !num.Valid {
		return valueobject.Money{}, fmt.Errorf("numeric value is null")
	}

	var amount float64
	if err := num.Scan(&amount); err != nil {
		return valueobject.Money{}, fmt.Errorf("failed to convert numeric to amount: %w", err)
	}

	return valueobject.NewMoney(valueobject.NewDecimalFromFloat(amount), currency), nil
}

func stringToPgText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}

func timeToPgTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{Time: t, Valid: true}
}

func timeToPgTimestamptzPtr(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{Valid: false}
	}
	return pgtype.Timestamptz{Time: *t, Valid: true}
}

func pgTimestamptzToTime(t pgtype.Timestamptz) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

func customerIDToPgInt4(id valueobject.CustomerID) pgtype.Int4 {
	return pgtype.Int4{Int32: int32(id.Value()), Valid: true}
}

func MedSessionIDToPgInt4(id valueobject.MedSessionID) pgtype.Int4 {
	return pgtype.Int4{Int32: int32(id.Value()), Valid: true}
}

func pgInt4ToCustomerID(num pgtype.Int4) valueobject.CustomerID {
	if !num.Valid {
		return valueobject.CustomerID{}
	}
	return valueobject.NewCustomerID(uint(num.Int32))
}

func pgInt4ToMedSessionID(num pgtype.Int4) valueobject.MedSessionID {
	if !num.Valid {
		return valueobject.MedSessionID{}
	}
	return valueobject.NewMedSessionID(uint(num.Int32))
}
