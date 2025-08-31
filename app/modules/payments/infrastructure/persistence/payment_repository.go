package paymentRepo

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCPaymentRepository struct {
	queries *sqlc.Queries
}

func NewSQLCPaymentRepository(queries *sqlc.Queries) repository.PaymentRepository {
	return &SQLCPaymentRepository{
		queries: queries,
	}
}

func (c *SQLCPaymentRepository) Save(ctx context.Context, payment *entity.Payment) error {
	if payment.GetID() == 0 {
		return c.create(ctx, payment)
	}
	return c.update(ctx, payment)
}

func (c *SQLCPaymentRepository) Search(ctx context.Context, pageInput page.PageData, searchCriteria map[string]any) (page.Page[[]entity.Payment], error) {
	return page.Page[[]entity.Payment]{}, nil
}

func (c *SQLCPaymentRepository) GetByID(ctx context.Context, id int) (entity.Payment, error) {
	sqlRow, err := c.queries.GetPaymentById(ctx, int32(id))
	if err != nil {
		return entity.Payment{}, err
	}

	payment, err := mapSQLCPaymentToDomain(sqlRow)
	if err != nil {
		return entity.Payment{}, err
	}

	return payment, nil
}

func (c *SQLCPaymentRepository) ListByUserID(ctx context.Context, userID int, pageInput page.PageData) (page.Page[[]entity.Payment], error) {
	sqlRows, err := c.queries.ListPaymentsByUserId(ctx, sqlc.ListPaymentsByUserIdParams{
		UserID: int32(userID),
		Limit:  int32(pageInput.PageSize),
		Offset: int32(pageInput.PageNumber*pageInput.PageSize) - 1,
	})
	if err != nil {
		return page.Page[[]entity.Payment]{}, err
	}

	return toPageablePayments(sqlRows, pageInput, len(sqlRows))
}

func (c *SQLCPaymentRepository) ListByStatus(ctx context.Context, status enum.PaymentStatus, pageInput page.PageData) (page.Page[[]entity.Payment], error) {
	sqlRows, err := c.queries.ListPaymentsByStatus(ctx, sqlc.ListPaymentsByStatusParams{
		Status: models.PaymentStatus(status),
	})
	if err != nil {
		return page.Page[[]entity.Payment]{}, err
	}

	return toPageablePayments(sqlRows, pageInput, len(sqlRows))
}

func (c *SQLCPaymentRepository) ListOverduePayments(ctx context.Context, pageInput page.PageData) (page.Page[[]entity.Payment], error) {
	sqlRows, err := c.queries.ListOverduePayments(ctx, sqlc.ListOverduePaymentsParams{
		Limit:  int32(pageInput.PageSize),
		Offset: int32(pageInput.PageNumber*pageInput.PageSize) - 1,
	})
	if err != nil {
		return page.Page[[]entity.Payment]{}, err
	}

	return toPageablePayments(sqlRows, pageInput, len(sqlRows))
}

func (c *SQLCPaymentRepository) ListPaymentsByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageData) (page.Page[[]entity.Payment], error) {
	sqlRows, err := c.queries.ListPaymentsByDateRange(ctx, sqlc.ListPaymentsByDateRangeParams{
		CreatedAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		CreatedAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Limit:       int32(pageInput.PageSize),
		Offset:      int32(pageInput.PageNumber*pageInput.PageSize) - 1,
	})
	if err != nil {
		return page.Page[[]entity.Payment]{}, err
	}

	return toPageablePayments(sqlRows, pageInput, len(sqlRows))
}

func (c *SQLCPaymentRepository) SoftDelete(ctx context.Context, id int) error {
	if err := c.queries.SoftDeletePayment(ctx, int32(id)); err != nil {
		return err
	}
	return nil
}

func (c *SQLCPaymentRepository) GetByTransactionID(ctx context.Context, transactionID string) (entity.Payment, error) {
	sqlRow, err := c.queries.GetPaymentByTransactionId(
		ctx,
		pgtype.Text{String: transactionID, Valid: true},
	)
	if err != nil {
		return entity.Payment{}, err
	}

	return mapSQLCPaymentToDomain(sqlRow)
}

func (c *SQLCPaymentRepository) create(context context.Context, payment *entity.Payment) error {
	createdRow, err := c.queries.CreatePayment(context, mapDomainToCreateParams(payment))
	if err != nil {
		return err
	}

	payment.SetID(int(createdRow.ID))
	return nil
}

func (c *SQLCPaymentRepository) update(context context.Context, payment *entity.Payment) error {
	_, err := c.queries.UpdatePayment(context, mapDomainToUpdateParams(payment))
	if err != nil {
		return err
	}

	return nil
}

func toPageablePayments(sqlRows []sqlc.Payment, pageInput page.PageData, totalItems int) (page.Page[[]entity.Payment], error) {
	payments, err := mapSQLCPaymentListToDomain(sqlRows)
	if err != nil {
		return page.Page[[]entity.Payment]{}, err
	}
	return page.NewPage(payments, *page.GetPageMetadata(totalItems, pageInput)), nil
}
