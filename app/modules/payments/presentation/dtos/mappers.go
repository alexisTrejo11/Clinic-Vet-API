package dto

import (
	"context"
	"fmt"

	"clinic-vet-api/app/core/domain/entity/payment"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	cmd "clinic-vet-api/app/modules/payments/application/command"
	query "clinic-vet-api/app/modules/payments/application/queries"
)

func (req CreatePaymentRequest) ToCreatePaymentCommand(ctx context.Context) (cmd.CreatePaymentCommand, error) {
	paymentMethod, err := enum.ParsePaymentMethod(req.PaymentMethod)
	if err != nil {
		return cmd.CreatePaymentCommand{}, fmt.Errorf("invalid payment method: %w", err)
	}

	var paymentStatus enum.PaymentStatus
	if req.Status != "" {
		paymentStatus, err = enum.ParsePaymentStatus(req.Status)
		if err != nil {
			return cmd.CreatePaymentCommand{}, fmt.Errorf("invalid payment status: %w", err)
		}
	} else {
		paymentStatus = enum.PaymentStatusPending
	}

	amount := valueobject.NewMoney(req.Amount, req.Currency)

	customerID := valueobject.NewCustomerID(uint(req.CustomerID))
	appointmentID := valueobject.NewAppointmentID(uint(req.AppointmentID))

	c := cmd.CreatePaymentCommand{
		Ctx:              ctx,
		Amount:           amount,
		Status:           paymentStatus,
		Method:           paymentMethod,
		TransactionID:    req.TransactionID,
		Description:      req.Description,
		DueDate:          req.DueDate,
		PaidFromCustomer: customerID,
		AppointmentID:    &appointmentID,
		InvoiceID:        req.InvoiceID,
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
		Amount:        payment.Amount().ToFloat(),
		Currency:      payment.Currency(),
		Method:        payment.Method().DisplayName(),
		Status:        payment.Status().DisplayName(),
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

func PaymentFromResult(result *query.PaymentResult) *PaymentResponse {
	if result == nil {
		return nil
	}

	return &PaymentResponse{
		ID:            result.ID,
		Amount:        result.Amount,
		Currency:      result.Currency,
		Status:        result.Status,
		Method:        result.Method,
		TransactionID: result.TransactionID,
		Description:   result.Description,
		DueDate:       result.DueDate,
		PaidAt:        result.PaidAt,
		RefundedAt:    result.RefundedAt,
		CustomerID:    result.PaidFromCustomer,
		AppointmentID: result.AppointmentID,
		InvoiceID:     result.InvoiceID,
		RefundAmount:  result.RefundAmount,
		FailureReason: result.FailureReason,
		IsActive:      result.IsActive,
		CreatedAt:     result.CreatedAt,
		UpdatedAt:     result.UpdatedAt,
	}
}
