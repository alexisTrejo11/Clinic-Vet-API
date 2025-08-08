package paymentRepo

import (
	"context"
	"time"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

type SQLCPaymentRepository struct {
	queries sqlc.Queries
}

func NewSQLCPaymentRepository(queries sqlc.Queries) paymentDomain.PaymentRepository {
	return &SQLCPaymentRepository{
		queries: queries,
	}
}

func (c *SQLCPaymentRepository) Save(payment *paymentDomain.Payment) error {
	return nil
}

func (c *SQLCPaymentRepository) Search(ctx context.Context, pagination page.PageData, searchCriteria map[string]interface{}) (*page.Page[[]paymentDomain.Payment], error) {
	return nil, nil
}

func (c *SQLCPaymentRepository) GetById(ctx context.Context, id int) (*paymentDomain.Payment, error) {
	return nil, nil
}

func (c *SQLCPaymentRepository) ListByUserId(ctx context.Context, userId int, pagination page.PageData) (*page.Page[[]paymentDomain.Payment], error) {
	return nil, nil
}

func (c *SQLCPaymentRepository) ListByStatus(ctx context.Context, status paymentDomain.PaymentStatus, pagination page.PageData) (*page.Page[[]paymentDomain.Payment], error) {
	return nil, nil
}

func (c *SQLCPaymentRepository) ListOverduePayments(ctx context.Context, pagination page.PageData) (*page.Page[[]paymentDomain.Payment], error) {
	return nil, nil
}

func (c *SQLCPaymentRepository) ListPaymentsByDateRange(startDate, endDate time.Time) (*page.Page[[]paymentDomain.Payment], error) {
	return nil, nil
}

func (c *SQLCPaymentRepository) SoftDelete(id int) error {
	return nil
}

func (c *SQLCPaymentRepository) GetByTransactionId(transactionId string) (*paymentDomain.Payment, error) {
	return nil, nil
}
