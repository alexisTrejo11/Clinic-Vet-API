package paymentDomain

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type PaymentRepository interface {
	Save(payment *Payment) error

	Search(ctx context.Context, pagination page.PageData, searchCriteria map[string]interface{}) (page.Page[[]Payment], error)
	GetById(ctx context.Context, id int) (*Payment, error)
	ListByUserId(ctx context.Context, userId int, pagination page.PageData) (page.Page[[]Payment], error)
	ListByStatus(ctx context.Context, status PaymentStatus, pagination page.PageData) (page.Page[[]Payment], error)
	ListOverduePayments(ctx context.Context, pagination page.PageData) (page.Page[[]Payment], error)
	ListPaymentsByDateRange(startDate, endDate time.Time) (page.Page[[]Payment], error)

	SoftDelete(id int) error

	//GetTotalCount() (int, error)
	GetByTransactionId(transactionId string) (*Payment, error)
}

type PaymentService interface {
	ProcessPayment(payment *Payment) error
	RefundPayment(paymentId int, reason string) error
	ValidatePayment(payment *Payment) error
	CalculateTotal(appointmentId int) (Money, error)

	GetPaymentHistory(ownerId int) (page.Page[[]Payment], error)
	MarkOverduePayments() error
	GeneratePaymentReport(startDate, endDate time.Time) (PaymentReport, error)
}

type PaymentReport struct {
	StartDate        time.Time                        `json:"start_date"`
	EndDate          time.Time                        `json:"end_date"`
	TotalPayments    int                              `json:"total_payments"`
	TotalAmount      Money                            `json:"total_amount"`
	PaidAmount       Money                            `json:"paid_amount"`
	PendingAmount    Money                            `json:"pending_amount"`
	RefundedAmount   Money                            `json:"refunded_amount"`
	OverdueAmount    Money                            `json:"overdue_amount"`
	PaymentsByMethod map[PaymentMethod]PaymentSummary `json:"payments_by_method"`
	PaymentsByStatus map[PaymentStatus]PaymentSummary `json:"payments_by_status"`
}

type PaymentSummary struct {
	Count  int   `json:"count"`
	Amount Money `json:"amount"`
}
