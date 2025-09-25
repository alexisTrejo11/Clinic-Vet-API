package dto

import (
	"clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	cmd "clinic-vet-api/app/modules/payments/application/command"
	query "clinic-vet-api/app/modules/payments/application/queries"
)

func (req CreatePaymentRequest) ToCommand() cmd.CreatePaymentCommand {
	paymentMethod := enum.PaymentMethod(req.PaymentMethod)
	paymentStatus := enum.PaymentStatus(req.Status)
	amount := valueobject.NewMoney(valueobject.NewDecimalFromFloat(req.Amount), req.Currency)
	medSessionID := valueobject.NewMedSessionID(uint(req.MedSessionID))

	var customerID *valueobject.CustomerID
	if req.CustomerID != nil {
		custID := valueobject.NewCustomerID(uint(*req.CustomerID))
		customerID = &custID
	}

	c := &cmd.CreatePaymentCommand{
		Amount:        amount,
		Status:        paymentStatus,
		Method:        paymentMethod,
		TransactionID: req.TransactionID,
		Description:   req.Description,
		DueDate:       req.DueDate,
		PaidBy:        customerID,
		MedSessionID:  medSessionID,
		InvoiceID:     req.InvoiceID,
	}

	return *c
}

func (req UpdatePaymentRequest) ToCommand(paymentID uint) cmd.UpdatePaymentCommand {
	var amount *valueobject.Money
	if req.Amount != nil && req.Currency != nil {
		amountoObj := valueobject.NewMoney(valueobject.NewDecimalFromFloat(*req.Amount), *req.Currency)
		amount = &amountoObj
	}

	return cmd.NewUpdatePaymentCommand(paymentID, amount, req.PaymentMethod, req.Description, req.DueDate)
}

func (req ProcessPaymentRequest) ToCommand(paymentID uint) *cmd.ProcessPaymentCommand {
	return cmd.NewProcessPaymentCommand(paymentID, req.TransactionID)
}

func (req RefundPaymentRequest) ToCommand(paymentID uint) *cmd.RefundPaymentCommand {
	return cmd.NewRefundPaymentCommand(paymentID, req.Reason)
}

func (req CancelPaymentRequest) ToCommand(paymentID uint) *cmd.CancelPaymentCommand {
	return cmd.NewCancelPaymentCommand(paymentID, req.Reason)
}

func ToPaymentResponse(pay any) PaymentResponse {
	payment := pay.(payment.Payment)

	return PaymentResponse{
		ID:            payment.ID().Value(),
		MedSessionID:  payment.MedSessionID().Value(),
		Amount:        payment.Amount().Amount().Float64(),
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
		CustomerID:    result.PaidByCustomer,
		MedSessionID:  result.MedSessionID,
		InvoiceID:     result.InvoiceID,
		RefundAmount:  result.RefundAmount,
		FailureReason: result.FailureReason,
		IsActive:      result.IsActive,
		CreatedAt:     result.CreatedAt,
		UpdatedAt:     result.UpdatedAt,
	}
}
