package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"clinic-vet-api/app/core/domain/entity/payment"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"

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

func (r *SQLCPaymentRepository) FindBySpecification(ctx context.Context, specification specification.PaymentSpecification) (page.Page[payment.Payment], error) {
	return page.Page[payment.Payment]{}, fmt.Errorf("FindBySpecification method is not implemented")
}

func (r *SQLCPaymentRepository) FindByID(ctx context.Context, id valueobject.PaymentID) (payment.Payment, error) {
	sqlRow, err := r.queries.FindPaymentByID(ctx, int32(id.Value()))
	if err != nil {
		if err == sql.ErrNoRows {
			return payment.Payment{}, fmt.Errorf("payment with ID %d not found", id.Value())
		}
		return payment.Payment{}, fmt.Errorf("failed to find payment by ID: %w", err)
	}

	paymentEntity, err := sqlcRowToEntity(&sqlRow)
	if err != nil {
		return payment.Payment{}, fmt.Errorf("failed to convert SQL row to payment: %w", err)
	}

	return *paymentEntity, nil
}

func (r *SQLCPaymentRepository) FindByTransactionID(ctx context.Context, transactionID string) (payment.Payment, error) {
	sqlRow, err := r.queries.FindPaymentByTransactionID(ctx, pgtype.Text{String: transactionID, Valid: true})
	if err != nil {
		if err == sql.ErrNoRows {
			return payment.Payment{}, fmt.Errorf("payment with transaction ID %s not found", transactionID)
		}
		return payment.Payment{}, fmt.Errorf("failed to find payment by transaction ID: %w", err)
	}

	paymentEntity, err := sqlcRowToEntity(&sqlRow)
	if err != nil {
		return payment.Payment{}, fmt.Errorf("failed to convert SQL row to payment: %w", err)
	}

	return *paymentEntity, nil
}

func (r *SQLCPaymentRepository) FindByIDAndCustomerID(ctx context.Context, id valueobject.PaymentID, customerID valueobject.CustomerID) (payment.Payment, error) {
	sqlRow, err := r.queries.FindPaymentByIDAndCustomerID(ctx, sqlc.FindPaymentByIDAndCustomerIDParams{
		ID:               int32(id.Value()),
		PaidFromCustomer: pgtype.Int4{Int32: int32(customerID.Value()), Valid: true},
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return payment.Payment{}, fmt.Errorf("payment with ID %d for customer %d not found", id.Value(), customerID.Value())
		}
		return payment.Payment{}, fmt.Errorf("failed to find payment by ID and customer ID: %w", err)
	}

	paymentEntity, err := sqlcRowToEntity(&sqlRow)
	if err != nil {
		return payment.Payment{}, fmt.Errorf("failed to convert SQL row to payment: %w", err)
	}

	return *paymentEntity, nil
}

func (r *SQLCPaymentRepository) FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID, pageInput page.PageInput) (page.Page[payment.Payment], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize

	sqlRows, err := r.queries.FindPaymentsByCustomerID(ctx, sqlc.FindPaymentsByCustomerIDParams{
		PaidFromCustomer: pgtype.Int4{Int32: int32(customerID.Value()), Valid: true},
		Limit:            int32(pageInput.PageSize),
		Offset:           int32(offset),
	})
	if err != nil {
		return page.Page[payment.Payment]{}, fmt.Errorf("failed to find payments by customer ID: %w", err)
	}

	total, err := r.queries.CountPaymentsByCustomerID(ctx, pgtype.Int4{Int32: int32(customerID.Value()), Valid: true})
	if err != nil {
		return page.Page[payment.Payment]{}, fmt.Errorf("failed to count payments by customer ID: %w", err)
	}

	payments, err := r.convertSQLRowsToPayments(sqlRows)
	if err != nil {
		return page.Page[payment.Payment]{}, err
	}

	return page.NewPage(payments, *page.GetPageMetadata(int(total), pageInput)), nil
}

func (r *SQLCPaymentRepository) FindByStatus(ctx context.Context, status enum.PaymentStatus, pageInput page.PageInput) (page.Page[payment.Payment], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize

	sqlRows, err := r.queries.FindPaymentsByStatus(ctx, sqlc.FindPaymentsByStatusParams{
		Status: models.PaymentStatus(status.String()),
		Limit:  int32(pageInput.PageSize),
		Offset: int32(offset),
	})
	if err != nil {
		return page.Page[payment.Payment]{}, fmt.Errorf("failed to find payments by status: %w", err)
	}

	total, err := r.queries.CountPaymentsByStatus(ctx, models.PaymentStatus(status.String()))
	if err != nil {
		return page.Page[payment.Payment]{}, fmt.Errorf("failed to count payments by status: %w", err)
	}

	payments, err := r.convertSQLRowsToPayments(sqlRows)
	if err != nil {
		return page.Page[payment.Payment]{}, err
	}

	return page.NewPage(payments, *page.GetPageMetadata(int(total), pageInput)), nil
}

func (r *SQLCPaymentRepository) FindOverdue(ctx context.Context, pageInput page.PageInput) (page.Page[payment.Payment], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize

	sqlRows, err := r.queries.FindOverduePayments(ctx, sqlc.FindOverduePaymentsParams{
		Limit:  int32(pageInput.PageSize),
		Offset: int32(offset),
	})
	if err != nil {
		return page.Page[payment.Payment]{}, fmt.Errorf("failed to find overdue payments: %w", err)
	}

	total, err := r.queries.CountOverduePayments(ctx)
	if err != nil {
		return page.Page[payment.Payment]{}, fmt.Errorf("failed to count overdue payments: %w", err)
	}

	payments, err := r.convertSQLRowsToPayments(sqlRows)
	if err != nil {
		return page.Page[payment.Payment]{}, err
	}

	return page.NewPage(payments, *page.GetPageMetadata(int(total), pageInput)), nil
}

func (r *SQLCPaymentRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageInput) (page.Page[payment.Payment], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize

	sqlRows, err := r.queries.FindPaymentsByDateRange(ctx, sqlc.FindPaymentsByDateRangeParams{
		CreatedAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		CreatedAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Limit:       int32(pageInput.PageSize),
		Offset:      int32(offset),
	})
	if err != nil {
		return page.Page[payment.Payment]{}, fmt.Errorf("failed to find payments by date range: %w", err)
	}

	total, err := r.queries.CountPaymentsByDateRange(ctx, sqlc.CountPaymentsByDateRangeParams{
		CreatedAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		CreatedAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return page.Page[payment.Payment]{}, fmt.Errorf("failed to count payments by date range: %w", err)
	}

	payments, err := r.convertSQLRowsToPayments(sqlRows)
	if err != nil {
		return page.Page[payment.Payment]{}, err
	}

	return page.NewPage(payments, *page.GetPageMetadata(int(total), pageInput)), nil
}

func (r *SQLCPaymentRepository) FindRecentByCustomerID(ctx context.Context, customerID valueobject.CustomerID, limit int) ([]payment.Payment, error) {
	sqlRows, err := r.queries.FindRecentPaymentsByCustomerID(ctx, sqlc.FindRecentPaymentsByCustomerIDParams{
		PaidFromCustomer: pgtype.Int4{Int32: int32(customerID.Value()), Valid: true},
		Limit:            int32(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find recent payments by customer ID: %w", err)
	}

	return r.convertSQLRowsToPayments(sqlRows)
}

func (r *SQLCPaymentRepository) FindPendingPayments(ctx context.Context, pageInput page.PageInput) (page.Page[payment.Payment], error) {
	return r.FindByStatus(ctx, enum.PaymentStatusPending, pageInput)
}

func (r *SQLCPaymentRepository) FindSuccessfulPayments(ctx context.Context, pageInput page.PageInput) (page.Page[payment.Payment], error) {
	return r.FindByStatus(ctx, enum.PaymentStatusPaid, pageInput)
}

func (r *SQLCPaymentRepository) ExistsByID(ctx context.Context, id valueobject.PaymentID) (bool, error) {
	exists, err := r.queries.ExistsPaymentByID(ctx, int32(id.Value()))
	if err != nil {
		return false, fmt.Errorf("failed to check payment existence by ID: %w", err)
	}
	return exists, nil
}

func (r *SQLCPaymentRepository) ExistsByTransactionID(ctx context.Context, transactionID string) (bool, error) {
	exists, err := r.queries.ExistsPaymentByTransactionID(ctx, pgtype.Text{String: transactionID, Valid: true})
	if err != nil {
		return false, fmt.Errorf("failed to check payment existence by transaction ID: %w", err)
	}
	return exists, nil
}

func (r *SQLCPaymentRepository) ExistsPendingByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (bool, error) {
	exists, err := r.queries.ExistsPendingPaymentByCustomerID(ctx, pgtype.Int4{Int32: int32(customerID.Value()), Valid: true})
	if err != nil {
		return false, fmt.Errorf("failed to check pending payment existence by customer ID: %w", err)
	}
	return exists, nil
}

func (r *SQLCPaymentRepository) Save(ctx context.Context, paymentEntity *payment.Payment) error {
	if paymentEntity.ID().IsZero() {
		return r.create(ctx, paymentEntity)
	}
	return r.update(ctx, paymentEntity)
}

func (r *SQLCPaymentRepository) SoftDelete(ctx context.Context, id valueobject.PaymentID) error {
	err := r.queries.SoftDeletePayment(ctx, int32(id.Value()))
	if err != nil {
		return fmt.Errorf("failed to soft delete payment: %w", err)
	}
	return nil
}

func (r *SQLCPaymentRepository) HardDelete(ctx context.Context, id valueobject.PaymentID) error {
	err := r.queries.HardDeletePayment(ctx, int32(id.Value()))
	if err != nil {
		return fmt.Errorf("failed to hard delete payment: %w", err)
	}
	return nil
}

func (r *SQLCPaymentRepository) CountByStatus(ctx context.Context, status enum.PaymentStatus) (int64, error) {
	count, err := r.queries.CountPaymentsByStatus(ctx, models.PaymentStatus(status.String()))
	if err != nil {
		return 0, fmt.Errorf("failed to count payments by status: %w", err)
	}
	return count, nil
}

func (r *SQLCPaymentRepository) CountByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (int64, error) {
	count, err := r.queries.CountPaymentsByCustomerID(ctx, pgtype.Int4{Int32: int32(customerID.Value()), Valid: true})
	if err != nil {
		return 0, fmt.Errorf("failed to count payments by customer ID: %w", err)
	}
	return count, nil
}

func (r *SQLCPaymentRepository) CountOverdue(ctx context.Context) (int64, error) {
	count, err := r.queries.CountOverduePayments(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to count overdue payments: %w", err)
	}
	return count, nil
}

func (r *SQLCPaymentRepository) TotalRevenueByDateRange(ctx context.Context, startDate, endDate time.Time) (float64, error) {
	total, err := r.queries.TotalRevenueByDateRange(ctx, sqlc.TotalRevenueByDateRangeParams{
		PaidAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		PaidAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return 0, fmt.Errorf("failed to calculate total revenue: %w", err)
	}

	totalFloat, ok := total.(float64)
	if !ok {
		return 0, fmt.Errorf("failed to convert total revenue to float64")
	}
	return totalFloat, nil
}

func (r *SQLCPaymentRepository) create(ctx context.Context, paymentEntity *payment.Payment) error {
	params, err := entityToCreateParams(paymentEntity)
	if err != nil {
		return fmt.Errorf("failed to convert payment to create params: %w", err)
	}

	_, err = r.queries.CreatePayment(ctx, *params)
	if err != nil {
		return fmt.Errorf("failed to create payment: %w", err)
	}

	return nil
}

func (r *SQLCPaymentRepository) update(ctx context.Context, paymentEntity *payment.Payment) error {
	params, err := entityToUpdateParams(paymentEntity)
	if err != nil {
		return fmt.Errorf("failed to convert payment to update params: %w", err)
	}

	_, err = r.queries.UpdatePayment(ctx, *params)
	if err != nil {
		return fmt.Errorf("failed to update payment: %w", err)
	}

	return nil
}
