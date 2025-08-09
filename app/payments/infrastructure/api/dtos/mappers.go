package paymentDTOs

import (
	paymentCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/application/advanced/command"
	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

func (req CreatePaymentRequest) ToCreatePaymentCommand() paymentCmd.CreatePaymentCommand {
	return paymentCmd.CreatePaymentCommand{
		AppointmentId: req.AppointmentId,
		OwnerId:       req.OwnerId,
		Amount:        req.Amount,
		Currency:      req.Currency,
		PaymentMethod: req.PaymentMethod,
		Description:   req.Description,
		DueDate:       req.DueDate,
		TransactionId: req.TransactionId,
	}
}

func (req UpdatePaymentRequest) ToUpdatePaymentCommand(paymentId int) paymentCmd.UpdatePaymentCommand {
	command := paymentCmd.UpdatePaymentCommand{
		PaymentId:     paymentId,
		PaymentMethod: req.PaymentMethod,
		Description:   req.Description,
		DueDate:       req.DueDate,
	}

	if req.Amount != nil && req.Currency != nil {
		money := paymentDomain.NewMoney(*req.Amount, *req.Currency)
		command.Amount = &money
	}

	return command
}

func (req ProcessPaymentRequest) ToProcessPaymentCommand(paymentId int) paymentCmd.ProcessPaymentCommand {
	return paymentCmd.ProcessPaymentCommand{
		PaymentId:     paymentId,
		TransactionId: req.TransactionId,
	}
}

func (req RefundPaymentRequest) ToRefundPaymentCommand(paymentId int) paymentCmd.RefundPaymentCommand {
	return paymentCmd.RefundPaymentCommand{
		PaymentId: paymentId,
		Reason:    req.Reason,
	}
}

func (req CancelPaymentRequest) ToCancelPaymentCommand(paymentId int) paymentCmd.CancelPaymentCommand {
	return paymentCmd.CancelPaymentCommand{
		PaymentId: paymentId,
		Reason:    req.Reason,
	}
}

func ToPaymentResponse(payment *paymentDomain.Payment) PaymentResponse {
	return PaymentResponse{
		Id:            payment.Id,
		AppointmentId: payment.AppointmentId,
		OwnerId:       payment.OwnerId,
		Amount:        payment.Amount.ToFloat(),
		Currency:      payment.Currency,
		PaymentMethod: payment.PaymentMethod,
		Status:        payment.Status,
		TransactionId: payment.TransactionId,
		Description:   payment.Description,
		DueDate:       payment.DueDate,
		PaidAt:        payment.PaidAt,
		RefundedAt:    payment.RefundedAt,
		IsActive:      payment.IsActive,
		CreatedAt:     payment.CreatedAt,
		UpdatedAt:     payment.UpdatedAt,
	}
}

func ToPaymentListResponse(paymentsPage *page.Page[[]paymentDomain.Payment]) PaymentListResponse {
	responses := make([]PaymentResponse, len(paymentsPage.Data))
	for i, payment := range paymentsPage.Data {
		responses[i] = ToPaymentResponse(&payment)
	}

	return PaymentListResponse{
		Data:     responses,
		Metadata: paymentsPage.Metadata,
	}
}

func ToPaymentReportResponse(report paymentDomain.PaymentReport) PaymentReportResponse {
	paymentsByMethod := make(map[paymentDomain.PaymentMethod]PaymentSummary)
	for method, summary := range report.PaymentsByMethod {
		paymentsByMethod[method] = PaymentSummary{
			Count:  summary.Count,
			Amount: summary.Amount.ToFloat(),
		}
	}

	paymentsByStatus := make(map[paymentDomain.PaymentStatus]PaymentSummary)
	for status, summary := range report.PaymentsByStatus {
		paymentsByStatus[status] = PaymentSummary{
			Count:  summary.Count,
			Amount: summary.Amount.ToFloat(),
		}
	}

	return PaymentReportResponse{
		StartDate:        report.StartDate,
		EndDate:          report.EndDate,
		TotalPayments:    report.TotalPayments,
		TotalAmount:      report.TotalAmount.ToFloat(),
		TotalCurrency:    report.TotalAmount.Currency,
		PaidAmount:       report.PaidAmount.ToFloat(),
		PendingAmount:    report.PendingAmount.ToFloat(),
		RefundedAmount:   report.RefundedAmount.ToFloat(),
		OverdueAmount:    report.OverdueAmount.ToFloat(),
		PaymentsByMethod: paymentsByMethod,
		PaymentsByStatus: paymentsByStatus,
	}
}

func (req PaymentSearchRequest) ToSearchCriteria() map[string]interface{} {
	criteria := make(map[string]interface{})

	if req.OwnerId != nil {
		criteria["owner_id"] = *req.OwnerId
	}
	if req.AppointmentId != nil {
		criteria["appointment_id"] = *req.AppointmentId
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
