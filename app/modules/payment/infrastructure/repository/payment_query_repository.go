package repositoryimpl

import (
	pay "clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	p "clinic-vet-api/app/shared/page"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *SqlcPaymentRepository) FindBySpecification(ctx context.Context, specification specification.PaymentSpecification) (p.Page[pay.Payment], error) {
	return p.Page[pay.Payment]{}, r.dbError(OpSelect, ErrMsgSearchPayments, fmt.Errorf("method not implemented"))
}

func (r *SqlcPaymentRepository) FindByID(ctx context.Context, id vo.PaymentID) (pay.Payment, error) {
	sqlRow, err := r.queries.FindPaymentByID(ctx, id.Int32())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pay.Payment{}, r.notFoundError("id", fmt.Sprintf("%d", id.Value()))
		}
		return pay.Payment{}, r.dbError(OpSelect, ErrMsgGetPayment, err)
	}

	return r.ToEntity(sqlRow), nil
}

func (r *SqlcPaymentRepository) FindByTransactionID(ctx context.Context, transactionID string) (pay.Payment, error) {
	sqlRow, err := r.queries.FindPaymentByTransactionID(ctx, pgtype.Text{String: transactionID, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pay.Payment{}, r.notFoundError("transaction_id", transactionID)
		}
		return pay.Payment{}, r.dbError(OpSelect, ErrMsgGetPaymentByTransaction, err)
	}

	return r.ToEntity(sqlRow), nil
}

func (r *SqlcPaymentRepository) FindByIDAndCustomerID(ctx context.Context, id vo.PaymentID, customerID vo.CustomerID) (pay.Payment, error) {
	sqlRow, err := r.queries.FindPaymentByIDAndCustomerID(ctx, sqlc.FindPaymentByIDAndCustomerIDParams{
		ID:               id.Int32(),
		PaidByCustomerID: r.pgMap.PgInt4.FromInt32(customerID.Int32()),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pay.Payment{}, fmt.Errorf("payment with ID %d for customer %d not found", id.Value(), customerID.Value())
		}
		return pay.Payment{}, fmt.Errorf("failed to find payment by ID and customer ID: %w", err)
	}

	return r.ToEntity(sqlRow), nil
}

func (r *SqlcPaymentRepository) FindByCustomerID(ctx context.Context, customerID vo.CustomerID, pagination p.PaginationRequest) (p.Page[pay.Payment], error) {
	sqlRows, err := r.queries.FindPaymentsByCustomerID(ctx, sqlc.FindPaymentsByCustomerIDParams{
		PaidByCustomerID: r.pgMap.UintToPgInt4(customerID.Value()),
		Limit:            pagination.Limit(),
		Offset:           pagination.Offset(),
	})
	if err != nil {
		return p.Page[pay.Payment]{}, r.dbError(OpSelect, ErrMsgListPaymentsByUser, err)
	}

	total, err := r.queries.CountPaymentsByCustomerID(ctx, r.pgMap.UintToPgInt4(customerID.Value()))
	if err != nil {
		return p.Page[pay.Payment]{}, r.dbError(OpCount, ErrMsgCountPayments, err)
	}

	payments := r.sqlRowsToEntities(sqlRows)
	return p.NewPage(payments, total, pagination), nil
}

func (r *SqlcPaymentRepository) FindByStatus(ctx context.Context, status enum.PaymentStatus, pagination p.PaginationRequest) (p.Page[pay.Payment], error) {
	sqlRows, err := r.queries.FindPaymentsByStatus(ctx, sqlc.FindPaymentsByStatusParams{
		Status: models.PaymentStatus(status.String()),
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	})
	if err != nil {
		return p.Page[pay.Payment]{}, fmt.Errorf("failed to find payments by status: %w", err)
	}

	total, err := r.queries.CountPaymentsByStatus(ctx, models.PaymentStatus(status.String()))
	if err != nil {
		return p.Page[pay.Payment]{}, fmt.Errorf("failed to count payments by status: %w", err)
	}

	payments := r.sqlRowsToEntities(sqlRows)
	return p.NewPage(payments, total, pagination), nil
}

func (r *SqlcPaymentRepository) FindOverdue(ctx context.Context, pagination p.PaginationRequest) (p.Page[pay.Payment], error) {
	sqlRows, err := r.queries.FindOverduePayments(ctx, sqlc.FindOverduePaymentsParams{
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	})
	if err != nil {
		return p.Page[pay.Payment]{}, fmt.Errorf("failed to find overdue payments: %w", err)
	}

	total, err := r.queries.CountOverduePayments(ctx)
	if err != nil {
		return p.Page[pay.Payment]{}, fmt.Errorf("failed to count overdue payments: %w", err)
	}

	payments := r.sqlRowsToEntities(sqlRows)
	return p.NewPage(payments, total, pagination), nil
}

func (r *SqlcPaymentRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time, pagination p.PaginationRequest) (p.Page[pay.Payment], error) {
	sqlRows, err := r.queries.FindPaymentsByDateRange(ctx, sqlc.FindPaymentsByDateRangeParams{
		CreatedAt:   r.pgMap.PgTimestamptz.FromTime(startDate),
		CreatedAt_2: r.pgMap.PgTimestamptz.FromTime(endDate),
		Limit:       pagination.Limit(),
		Offset:      pagination.Offset(),
	})
	if err != nil {
		return p.Page[pay.Payment]{}, fmt.Errorf("failed to find payments by date range: %w", err)
	}

	total, err := r.queries.CountPaymentsByDateRange(ctx, sqlc.CountPaymentsByDateRangeParams{
		CreatedAt:   r.pgMap.PgTimestamptz.FromTime(startDate),
		CreatedAt_2: r.pgMap.PgTimestamptz.FromTime(endDate),
	})
	if err != nil {
		return p.Page[pay.Payment]{}, fmt.Errorf("failed to count payments by date range: %w", err)
	}

	payments := r.sqlRowsToEntities(sqlRows)
	return p.NewPage(payments, total, pagination), nil
}

func (r *SqlcPaymentRepository) FindRecentByCustomerID(ctx context.Context, customerID vo.CustomerID, limit int) ([]pay.Payment, error) {
	sqlRows, err := r.queries.FindRecentPaymentsByCustomerID(ctx, sqlc.FindRecentPaymentsByCustomerIDParams{
		PaidByCustomerID: r.pgMap.PgInt4.FromInt32(customerID.Int32()),
		Limit:            int32(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find recent payments by customer ID: %w", err)
	}

	payments := r.sqlRowsToEntities(sqlRows)
	return payments, nil
}

func (r *SqlcPaymentRepository) FindPendingPayments(ctx context.Context, pagination p.PaginationRequest) (p.Page[pay.Payment], error) {
	return r.FindByStatus(ctx, enum.PaymentStatusPending, pagination)
}

func (r *SqlcPaymentRepository) FindSuccessfulPayments(ctx context.Context, pagination p.PaginationRequest) (p.Page[pay.Payment], error) {
	return r.FindByStatus(ctx, enum.PaymentStatusPaid, pagination)
}

func (r *SqlcPaymentRepository) ExistsByID(ctx context.Context, id vo.PaymentID) (bool, error) {
	exists, err := r.queries.ExistsPaymentByID(ctx, id.Int32())
	if err != nil {
		return false, r.dbError(OpSelect, ErrMsgGetPayment, err)
	}
	return exists, nil
}

func (r *SqlcPaymentRepository) ExistsByTransactionID(ctx context.Context, transactionID string) (bool, error) {
	exists, err := r.queries.ExistsPaymentByTransactionID(
		ctx,
		r.pgMap.StringPtrToPgText(&transactionID),
	)
	if err != nil {
		return false, r.dbError(OpSelect, ErrMsgGetPaymentByTransaction, err)
	}
	return exists, nil
}

func (r *SqlcPaymentRepository) ExistsPendingByCustomerID(ctx context.Context, customerID vo.CustomerID) (bool, error) {
	exists, err := r.queries.ExistsPendingPaymentByCustomerID(
		ctx,
		r.pgMap.UintToPgInt4(customerID.Value()),
	)

	if err != nil {
		return false, fmt.Errorf("failed to check pending payment existence by customer ID: %w", err)
	}
	return exists, nil
}
