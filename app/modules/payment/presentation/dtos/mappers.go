package dto

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	cmd "clinic-vet-api/app/modules/payment/application/command"
	"clinic-vet-api/app/modules/payment/application/query"
	"clinic-vet-api/app/shared/mapper"
)

func (req CreatePaymentRequest) ToCommand() cmd.CreatePaymentCommand {
	amount := valueobject.NewMoney(valueobject.NewDecimalFromFloat(req.Amount), req.Currency)
	return cmd.CreatePaymentCommand{
		Amount:        amount,
		Status:        enum.PaymentStatus(req.Status),
		Method:        enum.PaymentMethod(req.PaymentMethod),
		TransactionID: req.TransactionID,
		Description:   req.Description,
		DueDate:       req.DueDate,
		PaidBy:        valueobject.NewOptCustomerID(req.CustomerID),
		MedSessionID:  valueobject.NewOptMedSessionID(req.MedSessionID),
		InvoiceID:     req.InvoiceID,
	}
}

func (req UpdatePaymentRequest) ToCommand(paymentID uint) cmd.UpdatePaymentCommand {
	var amount *valueobject.Money
	if req.Amount != nil && req.Currency != nil {
		amountoObj := valueobject.NewMoney(valueobject.NewDecimalFromFloat(*req.Amount), *req.Currency)
		amount = &amountoObj
	}

	return cmd.NewUpdatePaymentCommand(paymentID, amount, req.PaymentMethod, req.Description, req.DueDate)
}

func (req ProcessPaymentRequest) ToCommand(paymentID uint) cmd.ProcessPaymentCommand {
	return cmd.NewProcessPaymentCommand(paymentID, req.TransactionID)
}

func (req RefundPaymentRequest) ToCommand(paymentID uint) cmd.RefundPaymentCommand {
	return cmd.NewRefundPaymentCommand(paymentID, req.Reason)
}

func (req CancelPaymentRequest) ToCommand(paymentID uint) cmd.CancelPaymentCommand {
	return cmd.NewCancelPaymentCommand(paymentID, req.Reason)
}

func ToPaymentResponse(payment query.PaymentResult) PaymentResponse {
	return PaymentResponse{
		ID:            payment.ID,
		Amount:        payment.Amount.Amount().Float64(),
		Currency:      payment.Amount.Currency(),
		Method:        payment.Method.DisplayName(),
		Status:        payment.Status.DisplayName(),
		TransactionID: payment.TransactionID,
		Description:   payment.Description,
		DueDate:       payment.DueDate,
		PaidAt:        payment.PaidAt,
		MedSessionID:  mapper.MedSessionIDPtrToPtr(payment.MedSessionID),
		CustomerID:    mapper.PaidByCustomerPtrToPtr(payment.PaidByCustomer),
		InvoiceID:     payment.InvoiceID,
		RefundAmount:  mapper.MoneyPtrToFloat64Ptr(payment.RefundAmount),
		FailureReason: payment.FailureReason,
		IsActive:      payment.IsActive,
		RefundedAt:    payment.RefundedAt,
		CreatedAt:     payment.CreatedAt,
		UpdatedAt:     payment.UpdatedAt,
	}
}

func ToPaymentListResponse(payments []query.PaymentResult) []PaymentResponse {
	responses := make([]PaymentResponse, len(payments))
	for i, payment := range payments {
		responses[i] = ToPaymentResponse(payment)
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
