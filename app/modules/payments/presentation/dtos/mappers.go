package dto

import (
	"context"

	"clinic-vet-api/app/core/domain/entity/payment"
	cmd "clinic-vet-api/app/modules/payments/application/command"
	query "clinic-vet-api/app/modules/payments/application/queries"
)

func (req CreatePaymentRequest) ToCreatePaymentCommand() (cmd.CreatePaymentCommand, error) {
	c, err := cmd.NewCreatePaymentCommand(
		context.Background(),
		uint(req.AppointmentID),
		uint(req.UserID),
		req.Amount,
		req.Currency,
		req.PaymentMethod,
		req.Description,
		req.DueDate,
		req.TransactionID,
	)
	if err != nil {
		return cmd.CreatePaymentCommand{}, err
	}

	return c, nil
}

func (req UpdatePaymentRequest) ToUpdatePaymentCommand(ctx context.Context, paymentID uint) (cmd.UpdatePaymentCommand, error) {
	command, errors := cmd.NewUpdatePaymentCommand(
		ctx,
		paymentID,
		req.Amount,
		req.Currency,
		req.PaymentMethod,
		req.Description,
		req.DueDate,
	)

	if errors != nil {
		return cmd.UpdatePaymentCommand{}, errors
	}

	return command, nil
}

func (req ProcessPaymentRequest) ToProcessPaymentCommand(paymentID uint) (*cmd.ProcessPaymentCommand, error) {
	return cmd.NewProcessPaymentCommand(paymentID, req.TransactionID)
}

func (req RefundPaymentRequest) ToRefundPaymentCommand(ctx context.Context, paymentID uint) *cmd.RefundPaymentCommand {
	return cmd.NewRefundPaymentCommand(ctx, paymentID, req.Reason)
}

func (req CancelPaymentRequest) ToCancelPaymentCommand(paymentID uint) *cmd.CancelPaymentCommand {
	return cmd.NewCancelPaymentCommand(paymentID, req.Reason)
}

func ToPaymentResponse(pay any) PaymentResponse {
	payment := pay.(payment.Payment)

	return PaymentResponse{
		ID:            payment.ID().Value(),
		AppointmentID: payment.AppointmentID().Value(),
		UserID:        payment.UserID().Value(),
		Amount:        payment.Amount().ToFloat(),
		Currency:      payment.Currency(),
		PaymentMethod: payment.PaymentMethod(),
		Status:        payment.Status(),
		TransactionID: payment.TransactionID(),
		Description:   payment.Description(),
		DueDate:       payment.DueDate(),
		PaidAt:        payment.PaidAt(),
		RefundedAt:    payment.RefundedAt(),
		CreatedAt:     payment.CreatedAt(),
		UpdatedAt:     payment.UpdatedAt(),
	}
}

func ToPaymentListResponse(payments []query.PaymentResult) []PaymentResponse {
	responses := make([]PaymentResponse, len(payments))
	for i, payment := range payments {
		responses[i] = ToPaymentResponse(&payment)
	}

	return responses
}

/*
func ToPaymentReportResponse(report PaymentReport) PaymentReportResponse {
	paymentsByMethod := make(map[enum.PaymentMethod]PaymentSummary)
	for method, summary := range report.PaymentsByMethod {
		paymentsByMethod[method] = PaymentSummary{
			Count:  summary.Count,
			Amount: summary.Amount,
		}
	}

	paymentsByStatus := make(map[enum.PaymentStatus]PaymentSummary)
	for status, summary := range report.PaymentsByStatus {
		paymentsByStatus[status] = PaymentSummary{
			Count:  summary.Count,
			Amount: summary.Amount,
		}
	}

	return PaymentReportResponse{
		StartDate:        report.StartDate,
		EndDate:          report.EndDate,
		TotalPayments:    report.TotalPayments,
		TotalAmount:      report.TotalAmount,
		TotalCurrency:    report.TotalCurrency,
		PaidAmount:       report.PaidAmount,
		PendingAmount:    report.PendingAmount,
		RefundedAmount:   report.RefundedAmount,
		OverdueAmount:    report.OverdueAmount,
		PaymentsByMethod: paymentsByMethod,
		PaymentsByStatus: paymentsByStatus,
	}
}
*/
