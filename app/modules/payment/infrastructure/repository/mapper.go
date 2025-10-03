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
		Amount:           r.pgMap.PgNumeric.FromMoney(p.Amount()),
		Currency:         p.Currency(),
		Status:           models.PaymentStatus(p.Status().String()),
		Method:           models.PaymentMethod(p.Method().String()),
		Description:      r.pgMap.PgText.FromString(p.Description()),
		PaidByCustomerID: r.pgMap.PgInt4.FromCustomerIDPtr(p.PaidByCustomer()),
		MedSessionID:     r.pgMap.PgInt4.FromMedSessionIDPtr(p.MedSessionID()),
		TransactionID:    r.pgMap.PgText.FromStringPtr(p.TransactionID()),
		DueDate:          r.pgMap.PgTimestamptz.FromTimePtr(p.DueDate()),
		PaidAt:           r.pgMap.PgTimestamptz.FromTimePtr(p.PaidAt()),
		RefundedAt:       r.pgMap.PgTimestamptz.FromTimePtr(p.RefundedAt()),
		InvoiceID:        r.pgMap.PgText.FromStringPtr(p.InvoiceID()),
		RefundAmount:     r.pgMap.PgNumeric.FromMoneyPtr(p.RefundAmount()),
		FailureReason:    r.pgMap.PgText.FromStringPtr(p.FailureReason()),
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
		Description:      r.pgMap.PgText.FromString(p.Description()),
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
	amount := r.pgMap.PgNumeric.ToMoney(row.Amount, row.Currency)
	customerID := r.pgMap.PgInt4.ToCustomerIDPtr(row.PaidByCustomerID)
	dueDate := r.pgMap.PgTimestamptz.ToTimePtr(row.DueDate)
	medSession := r.pgMap.PgInt4.ToMedSessionIDPtr(row.MedSessionID)
	description := r.pgMap.PgText.ToString(row.Description)
	refoundAmount, _ := r.pgMap.PgNumeric.ToMoneyPtr(row.RefundAmount, row.Currency)
	paymentID := valueobject.NewPaymentID(uint(row.ID))

	return *payment.NewPaymentBuilder().
		// Mandatory fields
		WithID(paymentID).
		WithAmount(amount).
		WithPaymentMethod(enum.PaymentMethod(row.Method)).
		WithStatus(enum.PaymentStatus(row.Status)).
		WithDescription(description).

		// Optional fields
		WithPaidByCustomer(customerID).
		WithMedSessionID(medSession).
		WithDueDate(dueDate).
		WithInvoiceID(r.pgMap.PgText.ToStringPtr(row.InvoiceID)).
		WithTransactionID(r.pgMap.PgText.ToStringPtr(row.TransactionID)).
		WithPaidAt(r.pgMap.PgTimestamptz.ToTimePtr(row.PaidAt)).
		WithRefundedAt(r.pgMap.PgTimestamptz.ToTimePtr(row.RefundedAt)).
		WithFailureReason(r.pgMap.PgText.ToStringPtr(row.FailureReason)).
		WithInvoiceID(r.pgMap.PgText.ToStringPtr(row.InvoiceID)).
		WithRefundAmount(refoundAmount).

		// Audit fields
		WithIsActive(row.IsActive).
		WithTimeStamps(row.CreatedAt.Time, row.UpdatedAt.Time).
		Build()
}

func (r *SqlcPaymentRepository) sqlRowsToEntities(sqlRows []sqlc.Payment) []payment.Payment {
	payments := make([]payment.Payment, len(sqlRows))
	for i, row := range sqlRows {
		payments[i] = r.ToEntity(row)
	}
	return payments
}
