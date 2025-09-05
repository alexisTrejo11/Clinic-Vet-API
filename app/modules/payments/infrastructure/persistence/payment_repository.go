package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

// SQLCPaymentRepository implements PaymentRepository using SQLC
type SQLCPaymentRepository struct {
	queries *sqlc.Queries
}

// NewSQLCPaymentRepository creates a new payment repository instance
func NewSQLCPaymentRepository(queries *sqlc.Queries) repository.PaymentRepository {
	return &SQLCPaymentRepository{
		queries: queries,
	}
}

// Save creates or updates a payment based on whether it has an ID
func (r *SQLCPaymentRepository) Save(ctx context.Context, payment *entity.Payment) error {
	if payment.GetID().IsZero() {
		return r.create(ctx, payment)
	}
	return r.update(ctx, payment)
}

// GetByID retrieves a payment by ID
func (r *SQLCPaymentRepository) GetByID(ctx context.Context, id int) (entity.Payment, error) {
	sqlRow, err := r.queries.GetPaymentById(ctx, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Payment{}, r.notFoundError("id", fmt.Sprintf("%d", id))
		}
		return entity.Payment{}, r.dbError(OpSelect, fmt.Sprintf("failed to get payment with ID %d", id), err)
	}

	payment, err := sqlcToEntity(sqlRow)
	if err != nil {
		return entity.Payment{}, r.wrapConversionError(err)
	}

	return payment, nil
}

// GetByTransactionID retrieves a payment by transaction ID
func (r *SQLCPaymentRepository) GetByTransactionID(ctx context.Context, transactionID string) (entity.Payment, error) {
	sqlRow, err := r.queries.GetPaymentByTransactionId(
		ctx,
		pgtype.Text{String: transactionID, Valid: true},
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Payment{}, r.notFoundError("transaction_id", transactionID)
		}
		return entity.Payment{}, r.dbError(OpSelect, fmt.Sprintf("failed to get payment with transaction ID %s", transactionID), err)
	}

	payment, err := sqlcToEntity(sqlRow)
	if err != nil {
		return entity.Payment{}, r.wrapConversionError(err)
	}

	return payment, nil
}

// ListByUserID retrieves payments for a specific user with pagination
func (r *SQLCPaymentRepository) ListByUserID(ctx context.Context, userID int, pageInput page.PageInput) (page.Page[[]entity.Payment], error) {
	params := sqlc.ListPaymentsByUserIdParams{
		UserID: int32(userID),
		Limit:  int32(pageInput.PageSize),
		Offset: r.calculateOffset(pageInput),
	}

	sqlRows, err := r.queries.ListPaymentsByUserId(ctx, params)
	if err != nil {
		return page.Page[[]entity.Payment]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list payments for user ID %d", userID), err)
	}

	payments, err := sqlcToEntityList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Payment]{}, r.wrapConversionError(err)
	}

	// Get total count for user payments
	totalCount, err := r.queries.CountPaymentsByUserId(ctx, int32(userID))
	if err != nil {
		return page.Page[[]entity.Payment]{}, r.dbError(OpCount, fmt.Sprintf("failed to count payments for user ID %d", userID), err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(payments, *pageMetadata), nil
}

// ListByStatus retrieves payments by status with pagination
func (r *SQLCPaymentRepository) ListByStatus(ctx context.Context, status enum.PaymentStatus, pageInput page.PageInput) (page.Page[[]entity.Payment], error) {
	params := sqlc.ListPaymentsByStatusParams{
		Status: models.PaymentStatus(status),
		Limit:  int32(pageInput.PageSize),
		Offset: r.calculateOffset(pageInput),
	}

	sqlRows, err := r.queries.ListPaymentsByStatus(ctx, params)
	if err != nil {
		return page.Page[[]entity.Payment]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list payments with status %s", status), err)
	}

	payments, err := sqlcToEntityList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Payment]{}, r.wrapConversionError(err)
	}

	// Get total count for status
	totalCount, err := r.queries.CountPaymentsByStatus(ctx, models.PaymentStatus(status))
	if err != nil {
		return page.Page[[]entity.Payment]{}, r.dbError(OpCount, fmt.Sprintf("failed to count payments with status %s", status), err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(payments, *pageMetadata), nil
}

// ListOverduePayments retrieves overdue payments with pagination
func (r *SQLCPaymentRepository) ListOverduePayments(ctx context.Context, pageInput page.PageInput) (page.Page[[]entity.Payment], error) {
	params := sqlc.ListOverduePaymentsParams{
		Limit:  int32(pageInput.PageSize),
		Offset: r.calculateOffset(pageInput),
	}

	sqlRows, err := r.queries.ListOverduePayments(ctx, params)
	if err != nil {
		return page.Page[[]entity.Payment]{}, r.dbError(OpSelect, ErrMsgListOverduePayments, err)
	}

	payments, err := sqlcToEntityList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Payment]{}, r.wrapConversionError(err)
	}

	// Get total count for overdue payments
	totalCount, err := r.queries.CountOverduePayments(ctx)
	if err != nil {
		return page.Page[[]entity.Payment]{}, r.dbError(OpCount, "failed to count overdue payments", err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(payments, *pageMetadata), nil
}

// ListPaymentsByDateRange retrieves payments within a date range with pagination
func (r *SQLCPaymentRepository) ListPaymentsByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageInput) (page.Page[[]entity.Payment], error) {
	params := sqlc.ListPaymentsByDateRangeParams{
		CreatedAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		CreatedAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Limit:       int32(pageInput.PageSize),
		Offset:      r.calculateOffset(pageInput),
	}

	sqlRows, err := r.queries.ListPaymentsByDateRange(ctx, params)
	if err != nil {
		return page.Page[[]entity.Payment]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list payments between %s and %s", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")), err)
	}

	payments, err := sqlcToEntityList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Payment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountPaymentsByDateRange(ctx, sqlc.CountPaymentsByDateRangeParams{
		CreatedAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		CreatedAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return page.Page[[]entity.Payment]{}, r.dbError(OpCount, fmt.Sprintf("failed to count payments between %s and %s", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")), err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(payments, *pageMetadata), nil
}

// Search performs a search on payments based on criteria with pagination
func (r *SQLCPaymentRepository) Search(ctx context.Context, pageInput page.PageInput, searchCriteria map[string]any) (page.Page[[]entity.Payment], error) {
	return page.Page[[]entity.Payment]{}, r.dbError(OpSelect, ErrMsgSearchPayments, errors.New("search method not implemented"))
}

// SoftDelete marks a payment as deleted without removing the record
func (r *SQLCPaymentRepository) SoftDelete(ctx context.Context, id int) error {
	if err := r.queries.SoftDeletePayment(ctx, int32(id)); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("failed to soft delete payment with ID %d", id), err)
	}
	return nil
}

// create inserts a new payment into the database
func (r *SQLCPaymentRepository) create(ctx context.Context, payment *entity.Payment) error {
	createParams := entityToCreateParams(payment)

	_, err := r.queries.CreatePayment(ctx, createParams)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreatePayment, err)
	}

	return nil
}

// update modifies an existing payment in the database
func (r *SQLCPaymentRepository) update(ctx context.Context, payment *entity.Payment) error {
	params := mapDomainToUpdateParams(payment)

	_, err := r.queries.UpdatePayment(ctx, params)
	if err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("failed to update payment with ID %d", payment.GetID().GetValue()), err)
	}

	return nil
}
