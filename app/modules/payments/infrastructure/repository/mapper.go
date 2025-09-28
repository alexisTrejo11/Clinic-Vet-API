package repositoryimpl

import (
	"clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"
)

func (r *SqlcPaymentRepository) ToCreateParams(p payment.Payment) sqlc.CreatePaymentParams {
	params := &sqlc.CreatePaymentParams{
		Amount:        r.pgMap.PgNumeric.FromMoney(p.Amount()),
		Currency:      p.Currency(),
		Status:        models.PaymentStatus(p.Status().String()),
		Method:        models.PaymentMethod(p.Method().String()),
		MedSessionID:  r.pgMap.PgInt4.FromUint(p.MedSessionID().Value()),
		TransactionID: r.pgMap.PgText.FromStringPtr(p.TransactionID()),
		Description:   r.pgMap.PgText.FromStringPtr(p.Description()),
		DueDate:       r.pgMap.PgTimestamptz.FromTimePtr(p.DueDate()),
		PaidAt:        r.pgMap.PgTimestamptz.FromTimePtr(p.PaidAt()),
		RefundedAt:    r.pgMap.PgTimestamptz.FromTimePtr(p.RefundedAt()),
		InvoiceID:     r.pgMap.PgText.FromStringPtr(p.InvoiceID()),
		RefundAmount:  r.pgMap.PgNumeric.FromMoneyPtr(p.RefundAmount()),
		FailureReason: r.pgMap.PgText.FromStringPtr(p.FailureReason()),
	}
	return *params
}

func (r *SqlcPaymentRepository) ToUpdateParams(p *payment.Payment) sqlc.UpdatePaymentParams {
	params := &sqlc.UpdatePaymentParams{
		ID:               int32(p.ID().Value()),
		Amount:           r.pgMap.PgNumeric.FromMoney(p.Amount()),
		Currency:         p.Currency(),
		MedSessionID:     r.pgMap.PgInt4.FromUint(p.MedSessionID().Value()),
		Status:           models.PaymentStatus(p.Status().String()),
		Method:           models.PaymentMethod(p.Method().String()),
		TransactionID:    r.pgMap.PgText.FromStringPtr(p.TransactionID()),
		Description:      r.pgMap.PgText.FromStringPtr(p.Description()),
		DueDate:          r.pgMap.PgTimestamptz.FromTimePtr(p.DueDate()),
		PaidAt:           r.pgMap.PgTimestamptz.FromTimePtr(p.PaidAt()),
		RefundedAt:       r.pgMap.PgTimestamptz.FromTimePtr(p.RefundedAt()),
		PaidByCustomerID: r.pgMap.PgInt4.FromUint(p.PaidByCustomer().Value()),
		InvoiceID:        r.pgMap.PgText.FromStringPtr(p.InvoiceID()),
		RefundAmount:     r.pgMap.PgNumeric.FromMoneyPtr(p.RefundAmount()),
	}
	return *params
}

func (r *SqlcPaymentRepository) ToEntity(row sqlc.Payment) payment.Payment {
	amount, _ := r.pgMap.PgNumeric.ToMoney(row.Amount, row.Currency)
	customerID := valueobject.NewCustomerID(uint(row.PaidByCustomerID.Int32))
	dueDate := r.pgMap.PgTimestamptz.ToTimePtr(row.DueDate)
	medSession := valueobject.NewMedSessionID(uint(row.MedSessionID.Int32))
	description := r.pgMap.PgText.ToStringPtr(row.Description)
	refoundAmount, _ := r.pgMap.PgNumeric.ToMoney(row.RefundAmount, row.Currency)
	paymentID := valueobject.NewPaymentID(uint(row.ID))

	return *payment.NewPaymentBuilder().
		WithID(paymentID).
		WithInvoiceID(row.InvoiceID.String).
		WithPaidByCustomer(customerID).
		WithAmount(amount).
		WithMedSessionID(medSession).
		WithPaymentMethod(enum.PaymentMethod(row.Method)).
		WithStatus(enum.PaymentStatus(row.Status)).
		WithDueDate(dueDate).
		WithIsActive(row.IsActive).
		WithTransactionID(row.TransactionID.String).
		WithTimeStamps(row.CreatedAt.Time, row.UpdatedAt.Time).
		WithDescription(description).
		WithPaidAt(row.PaidAt.Time).
		WithRefundedAt(row.RefundedAt.Time).
		WithFailureReason(row.FailureReason.String).
		WithInvoiceID(row.InvoiceID.String).
		WithRefundAmount(refoundAmount).
		Build()
}

func (r *SqlcPaymentRepository) sqlRowsToEntities(sqlRows []sqlc.Payment) []payment.Payment {
	payments := make([]payment.Payment, len(sqlRows))
	for i, row := range sqlRows {
		payments[i] = r.ToEntity(row)
	}
	return payments
}
