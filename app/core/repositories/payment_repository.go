package repository

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *entity.Payment) error

	Search(ctx context.Context, pagination page.PageData, searchCriteria interface{}) (page.Page[[]entity.Payment], error)
	GetById(ctx context.Context, id int) (*entity.Payment, error)
	ListByUserId(ctx context.Context, userId int, pagination page.PageData) (page.Page[[]entity.Payment], error)
	ListByStatus(ctx context.Context, status enum.PaymentStatus, pagination page.PageData) (page.Page[[]entity.Payment], error)
	ListOverduePayments(ctx context.Context, pagination page.PageData) (page.Page[[]entity.Payment], error)
	ListPaymentsByDateRange(ctx context.Context, startDate, endDate time.Time, pagination page.PageData) (page.Page[[]entity.Payment], error)

	SoftDelete(ctx context.Context, id int) error

	// GetTotalCount() (int, error)
	GetByTransactionId(ctx context.Context, transactionId string) (entity.Payment, error)
}

type PaymentService interface {
	ProcessPayment(payment *entity.Payment) error
	RefundPayment(paymentId int, reason string) error
	ValidatePayment(payment *entity.Payment) error
	CalculateTotal(appointmentId int) (valueobject.Money, error)

	GetPaymentHistory(ownerId int) (page.Page[[]entity.Payment], error)
	MarkOverduePayments() error
	GeneratePaymentReport(startDate, endDate time.Time) (PaymentReport, error)
}

type PaymentReport struct {
	StartDate        time.Time                             `json:"start_date"`
	EndDate          time.Time                             `json:"end_date"`
	TotalPayments    int                                   `json:"total_payments"`
	TotalAmount      valueobject.Money                     `json:"total_amount"`
	PaidAmount       valueobject.Money                     `json:"paid_amount"`
	PendingAmount    valueobject.Money                     `json:"pending_amount"`
	RefundedAmount   valueobject.Money                     `json:"refunded_amount"`
	OverdueAmount    valueobject.Money                     `json:"overdue_amount"`
	PaymentsByMethod map[enum.PaymentMethod]PaymentSummary `json:"payments_by_method"`
	PaymentsByStatus map[enum.PaymentStatus]PaymentSummary `json:"payments_by_status"`
}

type PaymentSummary struct {
	Count  int               `json:"count"`
	Amount valueobject.Money `json:"amount"`
}
