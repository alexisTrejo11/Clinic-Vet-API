// Package repositoryimpl implements the PaymentRepository interface using SQLC.
package repositoryimpl

import (
	"clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcPaymentRepository struct {
	queries *sqlc.Queries
}

func NewSqlcPaymentRepository(queries *sqlc.Queries) repository.PaymentRepository {
	return &SqlcPaymentRepository{
		queries: queries,
	}
}

func (r *SqlcPaymentRepository) FindBySpecification(ctx context.Context, specification specification.PaymentSpecification) (page.Page[payment.Payment], error) {
	return page.Page[payment.Payment]{}, r.dbError(OpSelect, ErrMsgSearchPayments, fmt.Errorf("method not implemented"))
}

func (r *SqlcPaymentRepository) FindByID(ctx context.Context, id valueobject.PaymentID) (payment.Payment, error) {
	sqlRow, err := r.queries.FindPaymentByID(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return payment.Payment{}, r.notFoundError("id", fmt.Sprintf("%d", id.Value()))
		}
		return payment.Payment{}, r.dbError(OpSelect, ErrMsgGetPayment, err)
	}

	return *sqlcRowToEntity(sqlRow), nil
}

func (r *SqlcPaymentRepository) FindByTransactionID(ctx context.Context, transactionID string) (payment.Payment, error) {
	sqlRow, err := r.queries.FindPaymentByTransactionID(ctx, pgtype.Text{String: transactionID, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return payment.Payment{}, r.notFoundError("transaction_id", transactionID)
		}
		return payment.Payment{}, r.dbError(OpSelect, ErrMsgGetPaymentByTransaction, err)
	}

	return *sqlcRowToEntity(sqlRow), nil
}

func (r *SqlcPaymentRepository) FindByIDAndCustomerID(ctx context.Context, id valueobject.PaymentID, customerID valueobject.CustomerID) (payment.Payment, error) {
	sqlRow, err := r.queries.FindPaymentByIDAndCustomerID(ctx, sqlc.FindPaymentByIDAndCustomerIDParams{
		ID:               int32(id.Value()),
		PaidByCustomerID: pgtype.Int4{Int32: int32(customerID.Value()), Valid: true},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return payment.Payment{}, fmt.Errorf("payment with ID %d for customer %d not found", id.Value(), customerID.Value())
		}
		return payment.Payment{}, fmt.Errorf("failed to find payment by ID and customer ID: %w", err)
	}

	return *sqlcRowToEntity(sqlRow), nil
}

func (r *SqlcPaymentRepository) FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID, pagination page.PaginationRequest) (page.Page[payment.Payment], error) {
	sqlRows, err := r.queries.FindPaymentsByCustomerID(ctx, sqlc.FindPaymentsByCustomerIDParams{
		PaidByCustomerID: pgtype.Int4{Int32: int32(customerID.Value()), Valid: true},
		Limit:            int32(pagination.Limit()),
		Offset:           int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[payment.Payment]{}, r.dbError(OpSelect, ErrMsgListPaymentsByUser, err)
	}

	total, err := r.queries.CountPaymentsByCustomerID(ctx, pgtype.Int4{Int32: int32(customerID.Value()), Valid: true})
	if err != nil {
		return page.Page[payment.Payment]{}, r.dbError(OpCount, ErrMsgCountPayments, err)
	}

	payments, err := r.sqlRowsToEntities(sqlRows)
	if err != nil {
		return page.Page[payment.Payment]{}, err
	}

	return page.NewPage(payments, int(total), pagination), nil
}

func (r *SqlcPaymentRepository) FindByStatus(ctx context.Context, status enum.PaymentStatus, pagination page.PaginationRequest) (page.Page[payment.Payment], error) {
	sqlRows, err := r.queries.FindPaymentsByStatus(ctx, sqlc.FindPaymentsByStatusParams{
		Status: models.PaymentStatus(status.String()),
		Limit:  int32(pagination.Limit()),
		Offset: int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[payment.Payment]{}, fmt.Errorf("failed to find payments by status: %w", err)
	}

	total, err := r.queries.CountPaymentsByStatus(ctx, models.PaymentStatus(status.String()))
	if err != nil {
		return page.Page[payment.Payment]{}, fmt.Errorf("failed to count payments by status: %w", err)
	}

	payments, err := r.sqlRowsToEntities(sqlRows)
	if err != nil {
		return page.Page[payment.Payment]{}, err
	}

	return page.NewPage(payments, int(total), pagination), nil
}

func (r *SqlcPaymentRepository) FindOverdue(ctx context.Context, pagination page.PaginationRequest) (page.Page[payment.Payment], error) {
	sqlRows, err := r.queries.FindOverduePayments(ctx, sqlc.FindOverduePaymentsParams{
		Limit:  int32(pagination.Limit()),
		Offset: int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[payment.Payment]{}, fmt.Errorf("failed to find overdue payments: %w", err)
	}

	total, err := r.queries.CountOverduePayments(ctx)
	if err != nil {
		return page.Page[payment.Payment]{}, fmt.Errorf("failed to count overdue payments: %w", err)
	}

	payments, err := r.sqlRowsToEntities(sqlRows)
	if err != nil {
		return page.Page[payment.Payment]{}, err
	}

	return page.NewPage(payments, int(total), pagination), nil
}

func (r *SqlcPaymentRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time, pagination page.PaginationRequest) (page.Page[payment.Payment], error) {
	sqlRows, err := r.queries.FindPaymentsByDateRange(ctx, sqlc.FindPaymentsByDateRangeParams{
		CreatedAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		CreatedAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Limit:       int32(pagination.Limit()),
		Offset:      int32(pagination.Offset()),
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

	payments, err := r.sqlRowsToEntities(sqlRows)
	if err != nil {
		return page.Page[payment.Payment]{}, err
	}

	return page.NewPage(payments, int(total), pagination), nil
}

func (r *SqlcPaymentRepository) FindRecentByCustomerID(ctx context.Context, customerID valueobject.CustomerID, limit int) ([]payment.Payment, error) {
	sqlRows, err := r.queries.FindRecentPaymentsByCustomerID(ctx, sqlc.FindRecentPaymentsByCustomerIDParams{
		PaidByCustomerID: pgtype.Int4{Int32: int32(customerID.Value()), Valid: true},
		Limit:            int32(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find recent payments by customer ID: %w", err)
	}

	return r.sqlRowsToEntities(sqlRows)
}

func (r *SqlcPaymentRepository) FindPendingPayments(ctx context.Context, pagination page.PaginationRequest) (page.Page[payment.Payment], error) {
	return r.FindByStatus(ctx, enum.PaymentStatusPending, pagination)
}

func (r *SqlcPaymentRepository) FindSuccessfulPayments(ctx context.Context, pagination page.PaginationRequest) (page.Page[payment.Payment], error) {
	return r.FindByStatus(ctx, enum.PaymentStatusPaid, pagination)
}

func (r *SqlcPaymentRepository) ExistsByID(ctx context.Context, id valueobject.PaymentID) (bool, error) {
	exists, err := r.queries.ExistsPaymentByID(ctx, int32(id.Value()))
	if err != nil {
		return false, r.dbError(OpSelect, ErrMsgGetPayment, err)
	}
	return exists, nil
}

func (r *SqlcPaymentRepository) ExistsByTransactionID(ctx context.Context, transactionID string) (bool, error) {
	exists, err := r.queries.ExistsPaymentByTransactionID(ctx, pgtype.Text{String: transactionID, Valid: true})
	if err != nil {
		return false, r.dbError(OpSelect, ErrMsgGetPaymentByTransaction, err)
	}
	return exists, nil
}

func (r *SqlcPaymentRepository) ExistsPendingByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (bool, error) {
	exists, err := r.queries.ExistsPendingPaymentByCustomerID(ctx, pgtype.Int4{Int32: int32(customerID.Value()), Valid: true})
	if err != nil {
		return false, fmt.Errorf("failed to check pending payment existence by customer ID: %w", err)
	}
	return exists, nil
}

func (r *SqlcPaymentRepository) Save(ctx context.Context, paymentEntity *payment.Payment) error {
	if paymentEntity.ID().IsZero() {
		return r.create(ctx, paymentEntity)
	}
	return r.update(ctx, paymentEntity)
}

func (r *SqlcPaymentRepository) SoftDelete(ctx context.Context, id valueobject.PaymentID) error {
	err := r.queries.SoftDeletePayment(ctx, int32(id.Value()))
	if err != nil {
		return r.dbError(OpDelete, ErrMsgSoftDeletePayment, err)
	}
	return nil
}

func (r *SqlcPaymentRepository) HardDelete(ctx context.Context, id valueobject.PaymentID) error {
	err := r.queries.HardDeletePayment(ctx, int32(id.Value()))
	if err != nil {
		return r.dbError(OpDelete, ErrMsgSoftDeletePayment, err)
	}
	return nil
}

func (r *SqlcPaymentRepository) CountByStatus(ctx context.Context, status enum.PaymentStatus) (int64, error) {
	count, err := r.queries.CountPaymentsByStatus(ctx, models.PaymentStatus(status.String()))
	if err != nil {
		return 0, r.dbError(OpCount, ErrMsgListPaymentsByStatus, err)
	}
	return count, nil
}

func (r *SqlcPaymentRepository) CountByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (int64, error) {
	count, err := r.queries.CountPaymentsByCustomerID(ctx, pgtype.Int4{Int32: int32(customerID.Value()), Valid: true})
	if err != nil {
		return 0, r.dbError(OpCount, ErrMsgCountPayments, err)
	}
	return count, nil
}

func (r *SqlcPaymentRepository) CountOverdue(ctx context.Context) (int64, error) {
	count, err := r.queries.CountOverduePayments(ctx)
	if err != nil {
		return 0, r.dbError(OpCount, ErrMsgListOverduePayments, err)
	}
	return count, nil
}

func (r *SqlcPaymentRepository) TotalRevenueByDateRange(ctx context.Context, startDate, endDate time.Time) (float64, error) {
	total, err := r.queries.TotalRevenueByDateRange(ctx, sqlc.TotalRevenueByDateRangeParams{
		PaidAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		PaidAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return 0, r.dbError(OpSelect, ErrMsgListPaymentsByDateRange, err)
	}

	totalFloat, ok := total.(float64)
	if !ok {
		return 0, r.wrapConversionError(fmt.Errorf("failed to convert total revenue to float64"))
	}
	return totalFloat, nil
}

func (r *SqlcPaymentRepository) create(ctx context.Context, paymentEntity *payment.Payment) error {
	params := entityToCreateParams(*paymentEntity)

	_, err := r.queries.CreatePayment(ctx, *params)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreatePayment, err)
	}

	return nil
}

func (r *SqlcPaymentRepository) update(ctx context.Context, paymentEntity *payment.Payment) error {
	params := entityToUpdateParams(paymentEntity)

	_, err := r.queries.UpdatePayment(ctx, *params)
	if err != nil {
		return r.dbError(OpUpdate, ErrMsgUpdatePayment, err)
	}

	return nil
}
