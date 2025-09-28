package repository

import (
	"context"
	"time"

	p "clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
)

type PaymentRepository interface {
	FindBySpecification(ctx context.Context, specification specification.PaymentSpecification) (page.Page[p.Payment], error)
	FindByID(ctx context.Context, id vo.PaymentID) (p.Payment, error)
	FindByTransactionID(ctx context.Context, transactionID string) (p.Payment, error)
	FindByIDAndCustomerID(ctx context.Context, id vo.PaymentID, customerID vo.CustomerID) (p.Payment, error)

	FindByCustomerID(ctx context.Context, customerID vo.CustomerID, PaginationRequest page.PaginationRequest) (page.Page[p.Payment], error)
	FindByStatus(ctx context.Context, status enum.PaymentStatus, PaginationRequest page.PaginationRequest) (page.Page[p.Payment], error)
	FindOverdue(ctx context.Context, PaginationRequest page.PaginationRequest) (page.Page[p.Payment], error)
	FindByDateRange(ctx context.Context, startDate, endDate time.Time, PaginationRequest page.PaginationRequest) (page.Page[p.Payment], error)

	FindRecentByCustomerID(ctx context.Context, customerID vo.CustomerID, limit int) ([]p.Payment, error)
	FindPendingPayments(ctx context.Context, PaginationRequest page.PaginationRequest) (page.Page[p.Payment], error)
	FindSuccessfulPayments(ctx context.Context, PaginationRequest page.PaginationRequest) (page.Page[p.Payment], error)

	ExistsByID(ctx context.Context, id vo.PaymentID) (bool, error)
	ExistsByTransactionID(ctx context.Context, transactionID string) (bool, error)
	ExistsPendingByCustomerID(ctx context.Context, customerID vo.CustomerID) (bool, error)

	Save(ctx context.Context, paymentEntity *p.Payment) error
	Delete(ctx context.Context, id vo.PaymentID, isHard bool) error

	CountByStatus(ctx context.Context, status enum.PaymentStatus) (int64, error)
	CountByCustomerID(ctx context.Context, customerID vo.CustomerID) (int64, error)
	CountOverdue(ctx context.Context) (int64, error)
	TotalRevenueByDateRange(ctx context.Context, startDate, endDate time.Time) (float64, error)
}

type PaymentSummary struct {
	Count  int      `json:"count"`
	Amount vo.Money `json:"amount"`
}
