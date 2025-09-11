package repository

import (
	"context"
	"time"

	p "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/payment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	vo "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type PaymentRepository interface {
	FindBySpecification(ctx context.Context, spec any) (page.Page[p.Payment], error)
	FindByID(ctx context.Context, id vo.PaymentID) (p.Payment, error)
	FindByTransactionID(ctx context.Context, transactionID string) (p.Payment, error)
	FindByIDAndCustomerID(ctx context.Context, id vo.PaymentID, customerID vo.CustomerID) (p.Payment, error)

	FindAll(ctx context.Context, pageInput page.PageInput) (page.Page[p.Payment], error)
	FindByCustomerID(ctx context.Context, customerID vo.CustomerID, pageInput page.PageInput) (page.Page[p.Payment], error)
	FindByStatus(ctx context.Context, status enum.PaymentStatus, pageInput page.PageInput) (page.Page[p.Payment], error)
	FindOverdue(ctx context.Context, pageInput page.PageInput) (page.Page[p.Payment], error)
	FindByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageInput) (page.Page[p.Payment], error)

	FindRecentByCustomerID(ctx context.Context, customerID vo.CustomerID, limit int) (p.Payment, error)
	FindPendingPayments(ctx context.Context, pageInput page.PageInput) (page.Page[p.Payment], error)
	FindSuccessfulPayments(ctx context.Context, pageInput page.PageInput) (page.Page[p.Payment], error)

	ExistsByID(ctx context.Context, id vo.PaymentID) (bool, error)
	ExistsByTransactionID(ctx context.Context, transactionID string) (bool, error)
	ExistsPendingByCustomerID(ctx context.Context, customerID vo.CustomerID) (bool, error)

	Save(ctx context.Context, paymentEntity *p.Payment) error
	Update(ctx context.Context, paymentEntity *p.Payment) error
	SoftDelete(ctx context.Context, id vo.PaymentID) error
	HardDelete(ctx context.Context, id vo.PaymentID) error

	CountByStatus(ctx context.Context, status enum.PaymentStatus) (int64, error)
	CountByCustomerID(ctx context.Context, customerID vo.CustomerID) (int64, error)
	CountOverdue(ctx context.Context) (int64, error)
	TotalRevenueByDateRange(ctx context.Context, startDate, endDate time.Time) (float64, error)
}

type PaymentSummary struct {
	Count  int      `json:"count"`
	Amount vo.Money `json:"amount"`
}
