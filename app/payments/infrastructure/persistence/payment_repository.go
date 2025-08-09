package paymentRepo

import (
	"context"
	"time"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCPaymentRepository struct {
	queries *sqlc.Queries
}

func NewSQLCPaymentRepository(queries *sqlc.Queries) paymentDomain.PaymentRepository {
	return &SQLCPaymentRepository{
		queries: queries,
	}
}

func (c *SQLCPaymentRepository) Save(payment *paymentDomain.Payment) error {
	if payment.Id == 0 {
		return c.create(context.Background(), payment)
	}
	return c.update(context.Background(), payment)
}

func (c *SQLCPaymentRepository) Search(ctx context.Context, pageInput page.PageData, searchCriteria map[string]interface{}) (*page.Page[[]paymentDomain.Payment], error) {
	return nil, nil
}

func (c *SQLCPaymentRepository) GetById(ctx context.Context, id int) (*paymentDomain.Payment, error) {
	sqlRow, err := c.queries.GetPaymentById(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	payment, err := mapSQLCPaymentToDomain(sqlRow)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (c *SQLCPaymentRepository) ListByUserId(ctx context.Context, userId int, pageInput page.PageData) (*page.Page[[]paymentDomain.Payment], error) {
	sqlRows, err := c.queries.ListPaymentsByUserId(ctx, sqlc.ListPaymentsByUserIdParams{
		UserID: int32(userId),
		Limit:  int32(pageInput.PageSize),
		Offset: int32(pageInput.PageNumber*pageInput.PageSize) - 1,
	})

	if err != nil {
		return nil, err
	}

	return toPageablePayments(sqlRows, pageInput, len(sqlRows))
}

func (c *SQLCPaymentRepository) ListByStatus(ctx context.Context, status paymentDomain.PaymentStatus, pageInput page.PageData) (*page.Page[[]paymentDomain.Payment], error) {
	sqlRows, err := c.queries.ListPaymentsByStatus(ctx, sqlc.ListPaymentsByStatusParams{
		Status: pgtype.Text{String: string(status), Valid: true},
	})

	if err != nil {
		return nil, err
	}

	return toPageablePayments(sqlRows, pageInput, len(sqlRows))
}

func (c *SQLCPaymentRepository) ListOverduePayments(ctx context.Context, pageInput page.PageData) (*page.Page[[]paymentDomain.Payment], error) {
	sqlRows, err := c.queries.ListOverduePayments(ctx, sqlc.ListOverduePaymentsParams{
		Limit:  int32(pageInput.PageSize),
		Offset: int32(pageInput.PageNumber*pageInput.PageSize) - 1,
	})

	if err != nil {
		return nil, err
	}

	return toPageablePayments(sqlRows, pageInput, len(sqlRows))
}

func (c *SQLCPaymentRepository) ListPaymentsByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageData) (*page.Page[[]paymentDomain.Payment], error) {
	sqlRows, err := c.queries.ListPaymentsByDateRange(ctx, sqlc.ListPaymentsByDateRangeParams{
		CreatedAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		CreatedAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Limit:       int32(pageInput.PageSize),
		Offset:      int32(pageInput.PageNumber*pageInput.PageSize) - 1,
	})
	if err != nil {
		return nil, err
	}

	return toPageablePayments(sqlRows, pageInput, len(sqlRows))
}

func (c *SQLCPaymentRepository) SoftDelete(ctx context.Context, id int) error {
	if err := c.queries.SoftDeletePayment(ctx, int32(id)); err != nil {
		return err
	}
	return nil
}

func (c *SQLCPaymentRepository) GetByTransactionId(ctx context.Context, transactionId string) (*paymentDomain.Payment, error) {
	sqlRow, err := c.queries.GetPaymentByTransactionId(
		ctx,
		pgtype.Text{String: transactionId, Valid: true},
	)
	if err != nil {
		return nil, err
	}

	return mapSQLCPaymentToDomain(sqlRow)
}

func (c *SQLCPaymentRepository) create(context context.Context, payment *paymentDomain.Payment) error {
	createdRow, err := c.queries.CreatePayment(context, mapDomainToCreateParams(payment))
	if err != nil {
		return err
	}

	payment.SetId(int(createdRow.ID))
	return nil
}

func (c *SQLCPaymentRepository) update(context context.Context, payment *paymentDomain.Payment) error {
	_, err := c.queries.UpdatePayment(context, mapDomainToUpdateParams(payment))
	if err != nil {
		return err
	}

	return nil
}

func toPageablePayments(sqlRows []sqlc.Payment, pageInput page.PageData, totalItems int) (*page.Page[[]paymentDomain.Payment], error) {
	payments, err := mapSQLCPaymentListToDomain(sqlRows)
	if err != nil {
		return nil, err
	}
	return page.NewPage(payments, *page.GetPageMetadata(totalItems, pageInput)), nil
}
