package repositoryimpl

import (
	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func entityToCreateParams(p *payment.Payment) (*sqlc.CreatePaymentParams, error) {
	if p == nil {
		return nil, fmt.Errorf("payment cannot be nil")
	}

	amountNumeric, err := moneyToNumeric(p.Amount())
	if err != nil {
		return nil, fmt.Errorf("failed to convert amount: %w", err)
	}

	params := &sqlc.CreatePaymentParams{
		Amount:           amountNumeric,
		Currency:         p.Currency(),
		Status:           models.PaymentStatus(p.Status().String()),
		Method:           models.PaymentMethod(p.Method().String()),
		AppointmentID:    pgtype.Int4{Int32: int32(p.Entity.ID().Value()), Valid: true},
		TransactionID:    stringToPgText(p.TransactionID()),
		Description:      stringToPgText(p.Description()),
		DueDate:          timeToPgTimestamptzPtr(p.DueDate()),
		PaidAt:           timeToPgTimestamptzPtr(p.PaidAt()),
		RefundedAt:       timeToPgTimestamptzPtr(p.RefundedAt()),
		PaidFromCustomer: customerIDToPgInt4(p.PaidFromCustomer()),
		PaidToEmployee:   employeeIDToPgInt4(p.PaidToEmployee()),
	}

	// Campos opcionales
	if p.AppointmentID() != nil {
		params.AppointmentID = appointmentIDToPgInt4(*p.AppointmentID())
	}

	if p.InvoiceID() != nil {
		params.InvoiceID = stringToPgText(p.InvoiceID())
	}

	if p.RefundAmount() != nil {
		refundAmountNumeric, err := moneyToNumeric(*p.RefundAmount())
		if err != nil {
			return nil, fmt.Errorf("failed to convert refund amount: %w", err)
		}
		params.RefundAmount = refundAmountNumeric
	}

	if p.FailureReason() != nil {
		params.FailureReason = stringToPgText(p.FailureReason())
	}

	return params, nil
}

func entityToUpdateParams(p *payment.Payment) (*sqlc.UpdatePaymentParams, error) {
	if p == nil {
		return nil, fmt.Errorf("payment cannot be nil")
	}

	if p.ID().IsZero() {
		return nil, fmt.Errorf("payment ID is required for update")
	}

	amountNumeric, err := moneyToNumeric(p.Amount())
	if err != nil {
		return nil, fmt.Errorf("failed to convert amount: %w", err)
	}

	params := &sqlc.UpdatePaymentParams{
		ID:               int32(p.ID().Value()),
		Amount:           amountNumeric,
		Currency:         p.Currency(),
		AppointmentID:    pgtype.Int4{Int32: int32(p.Entity.ID().Value()), Valid: true},
		Status:           models.PaymentStatus(p.Status().String()),
		Method:           models.PaymentMethod(p.Method().String()),
		TransactionID:    stringToPgText(p.TransactionID()),
		Description:      stringToPgText(p.Description()),
		DueDate:          timeToPgTimestamptz(*p.DueDate()),
		PaidAt:           timeToPgTimestamptzPtr(p.PaidAt()),
		RefundedAt:       timeToPgTimestamptzPtr(p.RefundedAt()),
		PaidFromCustomer: customerIDToPgInt4(p.PaidFromCustomer()),
		PaidToEmployee:   employeeIDToPgInt4(p.PaidToEmployee()),
	}

	// Campos opcionales
	if p.AppointmentID() != nil {
		params.AppointmentID = appointmentIDToPgInt4(*p.AppointmentID())
	}

	if p.InvoiceID() != nil {
		params.InvoiceID = stringToPgText(p.InvoiceID())
	}

	if p.RefundAmount() != nil {
		refundAmountNumeric, err := moneyToNumeric(*p.RefundAmount())
		if err != nil {
			return nil, fmt.Errorf("failed to convert refund amount: %w", err)
		}
		params.RefundAmount = refundAmountNumeric
	}

	if p.FailureReason() != nil {
		params.FailureReason = stringToPgText(p.FailureReason())
	}

	return params, nil
}

func sqlcRowToEntity(row *sqlc.Payment) (*payment.Payment, error) {
	if row == nil {
		return nil, fmt.Errorf("row cannot be nil")
	}

	paymentID := valueobject.NewPaymentID(uint(row.ID))

	amount, err := numericToMoney(row.Amount, row.Currency)
	if err != nil {
		return nil, fmt.Errorf("failed to convert amount: %w", err)
	}

	status := enum.PaymentStatus(row.Status)
	if !status.IsValid() {
		return nil, fmt.Errorf("invalid payment status: %s", row.Status)
	}

	method := enum.PaymentMethod(row.Method)
	if !method.IsValid() {
		return nil, fmt.Errorf("invalid payment method: %s", row.Method)
	}

	customerID := pgInt4ToCustomerID(row.PaidFromCustomer)

	p, err := payment.NewPayment(
		valueobject.NewPaymentID(uint(row.ID)),
		row.CreatedAt.Time,
		row.UpdatedAt.Time,
		payment.WithInvoiceID(row.InvoiceID.String),
		payment.WithPaidFromCustomer(customerID),
		payment.WithPaidToEmployee(valueobject.NewEmployeeID(uint(row.PaidToEmployee.Int32))),
		payment.WithAmount(amount),
		payment.WithAppointmentID(pgInt4ToAppointmentID(row.AppointmentID)),
		payment.WithPaymentMethod(method),
		payment.WithStatus(status),
		payment.WithDueDate(pgTimestamptzToTime(row.DueDate)),
		payment.WithIsActive(row.IsActive),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	p.Entity = base.NewEntity(paymentID, pgTimestamptzToTime(row.CreatedAt), pgTimestamptzToTime(row.UpdatedAt), 1)

	if row.TransactionID.Valid {
		p.SetTransactionID(row.TransactionID.String)
	}

	if row.Description.Valid {
		p.SetDescription(row.Description.String)
	}

	if row.PaidAt.Valid {
		paidAt := row.PaidAt.Time
		p.SetPaidAt(&paidAt)
	}

	if row.RefundedAt.Valid {
		refundedAt := row.RefundedAt.Time
		p.SetRefundedAt(&refundedAt)
	}

	if row.AppointmentID.Valid {
		appointmentID := valueobject.NewAppointmentID(uint(row.AppointmentID.Int32))
		p.SetAppointmentID(&appointmentID)
	}

	if row.InvoiceID.Valid {
		p.SetInvoiceID(row.InvoiceID.String)
	}

	if row.RefundAmount.Valid {
		refundAmount, err := numericToMoney(row.RefundAmount, p.Currency())
		if err != nil {
			return nil, fmt.Errorf("failed to convert refund amount: %w", err)
		}
		p.SetRefundAmount(&refundAmount)
	}

	if row.FailureReason.Valid {
		p.SetFailureReason(row.FailureReason.String)
	}

	if row.DeletedAt.Valid {
		deletedAt := row.DeletedAt.Time
		p.SetDeletedAt(&deletedAt)
	}

	return p, nil
}

func (r *SqlcPaymentRepository) sqlRowsToEntities(sqlRows []sqlc.Payment) ([]payment.Payment, error) {
	payments := make([]payment.Payment, len(sqlRows))
	for i, row := range sqlRows {
		paymentEntity, err := sqlcRowToEntity(&row)
		if err != nil {
			return nil, fmt.Errorf("failed to convert SQL row to payment at index %d: %w", i, err)
		}
		payments[i] = *paymentEntity
	}
	return payments, nil
}

func moneyToNumeric(money valueobject.Money) (pgtype.Numeric, error) {
	amount := money.Amount()

	num := pgtype.Numeric{}
	err := num.Scan(amount)
	return num, err
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

func pgTimestamptzToTime(t pgtype.Timestamptz) time.Time {
	if !t.Valid {
		return time.Time{}
	}
	return t.Time
}

func customerIDToPgInt4(id valueobject.CustomerID) pgtype.Int4 {
	return pgtype.Int4{Int32: int32(id.Value()), Valid: true}
}

func employeeIDToPgInt4(id valueobject.EmployeeID) pgtype.Int4 {
	return pgtype.Int4{Int32: int32(id.Value()), Valid: true}
}

func appointmentIDToPgInt4(id valueobject.AppointmentID) pgtype.Int4 {
	return pgtype.Int4{Int32: int32(id.Value()), Valid: true}
}

func pgInt4ToCustomerID(num pgtype.Int4) valueobject.CustomerID {
	if !num.Valid {
		return valueobject.CustomerID{}
	}
	return valueobject.NewCustomerID(uint(num.Int32))
}

func pgInt4ToAppointmentID(num pgtype.Int4) valueobject.AppointmentID {
	if !num.Valid {
		return valueobject.AppointmentID{}
	}
	return valueobject.NewAppointmentID(uint(num.Int32))
}
