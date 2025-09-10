package repository

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/payment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type PaymentRepository interface {
	Save(ctx context.Context, paymentEntity *payment.Payment) error

	Search(ctx context.Context, pagination page.PageInput, searchCriteria map[string]any) (page.Page[[]payment.Payment], error)
	GetByID(ctx context.Context, id valueobject.PaymentID) (payment.Payment, error)
	GetByIDAndPaidFrom(ctx context.Context, CustomerID valueobject.CustomerID) (payment.Payment, error)

	ListByPaidFrom(ctx context.Context, customerID valueobject.CustomerID, pagination page.PageInput) (page.Page[[]payment.Payment], error)
	ListByStatus(ctx context.Context, status enum.PaymentStatus, pagination page.PageInput) (page.Page[[]payment.Payment], error)
	ListOverduePayments(ctx context.Context, pagination page.PageInput) (page.Page[[]payment.Payment], error)
	ListPaymentsByDateRange(ctx context.Context, startDate, endDate time.Time, pagination page.PageInput) (page.Page[[]payment.Payment], error)

	SoftDelete(ctx context.Context, id valueobject.PaymentID) error

	// GetTotalCount() (int, error)
	GetByTransactionID(ctx context.Context, transactionID string) (payment.Payment, error)
}

type PaymentSummary struct {
	Count  int               `json:"count"`
	Amount valueobject.Money `json:"amount"`
}
