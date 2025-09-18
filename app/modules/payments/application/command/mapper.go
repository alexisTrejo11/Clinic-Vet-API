package command

import (
	"clinic-vet-api/app/core/domain/entity/payment"
	apperror "clinic-vet-api/app/shared/error/application"
)

func (command *CreatePaymentCommand) ToDomain() (*payment.Payment, error) {
	opts := []payment.PaymentOption{
		payment.WithAmount(command.Amount),
		payment.WithStatus(command.Status),
		payment.WithPaymentMethod(command.Method),
		payment.WithPaidFromCustomer(command.PaidFromCustomer),
	}

	if command.TransactionID != nil {
		opts = append(opts, payment.WithTransactionID(*command.TransactionID))
	}

	if command.Description != nil {
		opts = append(opts, payment.WithDescription(*command.Description))
	}

	if !command.DueDate.IsZero() {
		opts = append(opts, payment.WithDueDate(command.DueDate))
	}

	if command.PaidAt != nil {
		opts = append(opts, payment.WithPaidAt(*command.PaidAt))
	}

	if command.RefundedAt != nil {
		opts = append(opts, payment.WithRefundedAt(*command.RefundedAt))
	}

	if command.AppointmentID != nil {
		opts = append(opts, payment.WithAppointmentID(*command.AppointmentID))
	}

	if command.InvoiceID != nil {
		opts = append(opts, payment.WithInvoiceID(*command.InvoiceID))
	}

	paymentEntity, err := payment.CreatePayment(
		command.Ctx,
		command.PaidFromCustomer,
		opts...,
	)

	if err != nil {
		return nil, apperror.MappingError([]string{err.Error()}, "constructor", "domain", "Payment")
	}

	return paymentEntity, nil
}
