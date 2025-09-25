package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/payment"
	"context"
)

func (cmd *CreatePaymentCommand) ToDomain(ctx context.Context) (*payment.Payment, error) {
	opts := []payment.PaymentOption{}

	if cmd.PaidBy != nil {
		opts = append(opts, payment.WithPaidByCustomer(*cmd.PaidBy))
	}

	if cmd.TransactionID != nil {
		opts = append(opts, payment.WithTransactionID(*cmd.TransactionID))
	}

	if cmd.Description != nil {
		opts = append(opts, payment.WithDescription(*cmd.Description))
	}

	if !cmd.DueDate.IsZero() {
		opts = append(opts, payment.WithDueDate(cmd.DueDate))
	}

	if cmd.PaidAt != nil {
		opts = append(opts, payment.WithPaidAt(*cmd.PaidAt))
	}

	if cmd.InvoiceID != nil {
		opts = append(opts, payment.WithInvoiceID(*cmd.InvoiceID))
	}

	return payment.CreatePayment(ctx, cmd.Amount, cmd.Method, cmd.Status, opts...)
}
