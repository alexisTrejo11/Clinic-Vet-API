package dto

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	cmd "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/payments/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

func (req CreatePaymentRequest) ToCreatePaymentCommand() (cmd.CreatePaymentCommand, error) {
	c, err := cmd.NewCreatePaymentCommand(
		context.Background(),
		req.AppointmentID,
		req.UserID,
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

func (req UpdatePaymentRequest) ToUpdatePaymentCommand(ctx context.Context, paymentID int) (cmd.UpdatePaymentCommand, error) {
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

func (req ProcessPaymentRequest) ToProcessPaymentCommand(paymentID int) (cmd.ProcessPaymentCommand, error) {
	return cmd.NewProcessPaymentCommand(paymentID, req.TransactionID)
}

func (req RefundPaymentRequest) ToRefundPaymentCommand(paymentID int) (cmd.RefundPaymentCommand, error) {
	return cmd.NewRefundPaymentCommand(paymentID, req.Reason)
}

func (req CancelPaymentRequest) ToCancelPaymentCommand(paymentID int) (cmd.CancelPaymentCommand, error) {
	return cmd.NewCancelPaymentCommand(paymentID, req.Reason)
}

func ToPaymentResponse(pay any) PaymentResponse {
	payment := pay.(entity.Payment)

	return PaymentResponse{
		ID:            payment.GetID().GetValue(),
		AppointmentID: payment.GetAppointmentID().GetValue(),
		UserID:        payment.GetUserID().GetValue(),
		Amount:        payment.GetAmount().ToFloat(),
		Currency:      payment.GetCurrency(),
		PaymentMethod: payment.GetPaymentMethod(),
		Status:        payment.GetStatus(),
		TransactionID: payment.GetTransactionID(),
		Description:   payment.GetDescription(),
		DueDate:       payment.GetDueDate(),
		PaidAt:        payment.GetPaidAt(),
		RefundedAt:    payment.GetRefundedAt(),
		CreatedAt:     payment.GetCreatedAt(),
		UpdatedAt:     payment.GetUpdatedAt(),
	}
}

func ToPaymentListResponse(data interface{}) PaymentListResponse {
	paymentsPage := data.(page.Page[[]entity.Payment])

	responses := make([]PaymentResponse, len(paymentsPage.Data))
	for i, payment := range paymentsPage.Data {
		responses[i] = ToPaymentResponse(&payment)
	}

	return PaymentListResponse{
		Data:     responses,
		Metadata: paymentsPage.Metadata,
	}
}

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

func (req PaymentSearchRequest) ToSearchCriteria() map[string]interface{} {
	criteria := make(map[string]interface{})

	if req.UserID != nil {
		criteria["owner_id"] = *req.UserID
	}
	if req.AppointmentID != nil {
		criteria["appointment_id"] = *req.AppointmentID
	}
	if req.Status != nil {
		criteria["status"] = *req.Status
	}
	if req.PaymentMethod != nil {
		criteria["payment_method"] = *req.PaymentMethod
	}
	if req.MinAmount != nil {
		criteria["min_amount"] = *req.MinAmount
	}
	if req.MaxAmount != nil {
		criteria["max_amount"] = *req.MaxAmount
	}
	if req.Currency != nil {
		criteria["currency"] = *req.Currency
	}
	if req.StartDate != nil {
		criteria["start_date"] = *req.StartDate
	}
	if req.EndDate != nil {
		criteria["end_date"] = *req.EndDate
	}

	return criteria
}
