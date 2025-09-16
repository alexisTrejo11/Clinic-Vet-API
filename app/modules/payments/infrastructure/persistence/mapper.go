package persistence

import (
	"fmt"
	"math/big"

	"clinic-vet-api/app/core/domain/entity/payment"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func sqlcToEntity(sqlcPayment sqlc.Payment) (payment.Payment, error) {
	status, err := enum.ParsePaymentStatus(string(sqlcPayment.Status))
	if err != nil {
		return payment.Payment{}, fmt.Errorf("invalid payment status: %w", err)
	}

	method, err := enum.ParsePaymentMethod(string(sqlcPayment.Method))
	if err != nil {
		return payment.Payment{}, fmt.Errorf("invalid payment method: %w", err)
	}

	// Mapeo de value objects
	paymentID := valueobject.NewPaymentID(uint(sqlcPayment.ID))

	// TODO: Fix
	appointmentID := valueobject.NewAppointmentID(1)

	paidFromCustomer := valueobject.NewUserID(uint(sqlcPayment.PaidFromCustomer.Int32))

	var amount valueobject.Money
	if sqlcPayment.Amount.Valid {
		amount = valueobject.NewMoney(float64(sqlcPayment.Amount.Int.Int64()), sqlcPayment.Currency)
	} else {
		// Valor por defecto o error según tu lógica de negocio
		return payment.Payment{}, fmt.Errorf("payment amount is required")
	}

	// Crear el payment usando functional options
	paymentEntity, err := payment.NewPayment(
		paymentID,
		appointmentID,
		userID,
		payment.WithAmount(amount),
		payment.WithCurrency(sqlcPayment.Currency),
		payment.WithPaymentMethod(method),
		payment.WithStatus(status),
		// TODO: Fix (nullable fields)
		/*)
		payment.WithTransactionID(utils.StringPtrFromNullString(sqlcPayment.TransactionID.String)),
		payment.WithDescription(utils.StringPtrFromNullString(sqlcPayment.Description)),
		payment.WithDueDate(utils.TimePtrFromNullTime(sqlcPayment.DueDate)),
		payment.WithPaidAt(utils.TimePtrFromNullTime(sqlcPayment.PaidAt)),
		payment.WithRefundedAt(utils.TimePtrFromNullTime(sqlcPayment.RefundedAt)),
		*/
	)
	if err != nil {
		return payment.Payment{}, fmt.Errorf("failed to create payment entity: %w", err)
	}

	return *paymentEntity, nil
}

func sqlcToEntityList(sqlcPayments []sqlc.Payment) ([]payment.Payment, error) {
	var payments []payment.Payment
	for _, sqlcPayment := range sqlcPayments {
		payment, err := sqlcToEntity(sqlcPayment)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func entityToCreateParams(payment *payment.Payment) sqlc.CreatePaymentParams {
	return sqlc.CreatePaymentParams{
		Amount:        pgtype.Numeric{Int: big.NewInt(payment.Amount().Amount()), Valid: true},
		Currency:      payment.Amount().Currency(),
		TransactionID: pgtype.Text{String: *payment.TransactionID(), Valid: payment.TransactionID() != nil},
		Status:        models.PaymentStatus(string(payment.Status())),
		Method:        models.PaymentMethod(string(payment.PaymentMethod())),
		Description:   pgtype.Text{String: *payment.Description(), Valid: payment.Description() != nil},
		Duedate:       pgtype.Timestamptz{Time: *payment.DueDate(), Valid: payment.DueDate() != nil},
		PaidAt:        pgtype.Timestamptz{Time: *payment.PaidAt(), Valid: payment.PaidAt() != nil},
		RefundedAt:    pgtype.Timestamptz{Time: *payment.RefundedAt(), Valid: payment.RefundedAt() != nil},
	}
}

func mapDomainToUpdateParams(payment *payment.Payment) sqlc.UpdatePaymentParams {
	return sqlc.UpdatePaymentParams{
		ID:            int32(payment.ID().Value()),
		Amount:        pgtype.Numeric{Int: big.NewInt(payment.Amount().Amount()), Valid: true},
		Currency:      payment.Amount().Currency(),
		TransactionID: pgtype.Text{String: *payment.TransactionID(), Valid: payment.TransactionID() != nil},
		Status:        models.PaymentStatus(string(payment.Status())),
		Method:        models.PaymentMethod(string(payment.PaymentMethod())),
		UserID:        int32(payment.UserID().Value()),
		Description:   pgtype.Text{String: *payment.Description(), Valid: payment.Description() != nil},
		Duedate:       pgtype.Timestamptz{Time: *payment.DueDate(), Valid: payment.DueDate() != nil},
		PaidAt:        pgtype.Timestamptz{Time: *payment.PaidAt(), Valid: payment.PaidAt() != nil},
		RefundedAt:    pgtype.Timestamptz{Time: *payment.RefundedAt(), Valid: payment.RefundedAt() != nil},
	}
}
