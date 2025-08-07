package paymentDomain

import (
	"context"
	"time"
)

type PaymentRepository interface {
	Save(payment *Payment) error

	Search(ctx context.Context, limit, offset int, searchCriteria map[string]interface{}) ([]Payment, error)
	GetById(ctx context.Context, id int) (*Payment, error)
	ListByAppointmentId(appointmentId int) ([]Payment, error)
	ListByOwnerId(ctx context.Context, ownerId int) ([]Payment, error)
	ListByStatus(ctx context.Context, status PaymentStatus) ([]Payment, error)
	ListOverduePayments(ctx context.Context) ([]Payment, error)
	ListPaymentsByDateRange(startDate, endDate time.Time) ([]Payment, error)

	SoftDelete(id int) error

	//GetTotalCount() (int, error)
	GetByTransactionId(transactionId string) (*Payment, error)
}

type PaymentService interface {
	ProcessPayment(payment *Payment) error
	RefundPayment(paymentId int, reason string) error
	ValidatePayment(payment *Payment) error
	CalculateTotal(appointmentId int) (Money, error)

	GetPaymentHistory(ownerId int) ([]Payment, error)
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
