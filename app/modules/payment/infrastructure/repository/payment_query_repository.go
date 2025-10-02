package repositoryimpl

import (
	p "clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
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

func (r *SqlcPaymentRepository) FindBySpecification(ctx context.Context, specification specification.PaymentSpecification) (page.Page[p.Payment], error) {
	return page.Page[p.Payment]{}, r.dbError(OpSelect, ErrMsgSearchPayments, fmt.Errorf("method not implemented"))
}

func (r *SqlcPaymentRepository) FindByID(ctx context.Context, id vo.PaymentID) (p.Payment, error) {
	sqlRow, err := r.queries.FindPaymentByID(ctx, id.Int32())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return p.Payment{}, r.notFoundError("id", fmt.Sprintf("%d", id.Value()))
		}
		return p.Payment{}, r.dbError(OpSelect, ErrMsgGetPayment, err)
	}

	return r.ToEntity(sqlRow), nil
}

func (r *SqlcPaymentRepository) FindByTransactionID(ctx context.Context, transactionID string) (p.Payment, error) {
	sqlRow, err := r.queries.FindPaymentByTransactionID(ctx, pgtype.Text{String: transactionID, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return p.Payment{}, r.notFoundError("transaction_id", transactionID)
		}
		return p.Payment{}, r.dbError(OpSelect, ErrMsgGetPaymentByTransaction, err)
	}

	return r.ToEntity(sqlRow), nil
}

func (r *SqlcPaymentRepository) FindByIDAndCustomerID(ctx context.Context, id vo.PaymentID, customerID vo.CustomerID) (p.Payment, error) {
	sqlRow, err := r.queries.FindPaymentByIDAndCustomerID(ctx, sqlc.FindPaymentByIDAndCustomerIDParams{
		ID:               int32(id.Value()),
		PaidByCustomerID: pgtype.Int4{Int32: int32(customerID.Int32()), Valid: true},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return p.Payment{}, fmt.Errorf("payment with ID %d for customer %d not found", id.Value(), customerID.Value())
		}
		return p.Payment{}, fmt.Errorf("failed to find payment by ID and customer ID: %w", err)
	}

	return r.ToEntity(sqlRow), nil
}

func (r *SqlcPaymentRepository) FindByCustomerID(ctx context.Context, customerID vo.CustomerID, pagination page.PaginationRequest) (page.Page[p.Payment], error) {
	sqlRows, err := r.queries.FindPaymentsByCustomerID(ctx, sqlc.FindPaymentsByCustomerIDParams{
		PaidByCustomerID: r.pgMap.UintToPgInt4(customerID.Value()),
		Limit:            int32(pagination.Limit()),
		Offset:           int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[p.Payment]{}, r.dbError(OpSelect, ErrMsgListPaymentsByUser, err)
	}

	total, err := r.queries.CountPaymentsByCustomerID(ctx, r.pgMap.UintToPgInt4(customerID.Value()))
	if err != nil {
		return page.Page[p.Payment]{}, r.dbError(OpCount, ErrMsgCountPayments, err)
	}

	payments := r.sqlRowsToEntities(sqlRows)
	return page.NewPage(payments, total, pagination), nil
}

func (r *SqlcPaymentRepository) FindByStatus(ctx context.Context, status enum.PaymentStatus, pagination page.PaginationRequest) (page.Page[p.Payment], error) {
	sqlRows, err := r.queries.FindPaymentsByStatus(ctx, sqlc.FindPaymentsByStatusParams{
		Status: models.PaymentStatus(status.String()),
		Limit:  int32(pagination.Limit()),
		Offset: int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[p.Payment]{}, fmt.Errorf("failed to find payments by status: %w", err)
	}

	total, err := r.queries.CountPaymentsByStatus(ctx, models.PaymentStatus(status.String()))
	if err != nil {
		return page.Page[p.Payment]{}, fmt.Errorf("failed to count payments by status: %w", err)
	}

	payments := r.sqlRowsToEntities(sqlRows)
	return page.NewPage(payments, total, pagination), nil
}

func (r *SqlcPaymentRepository) FindOverdue(ctx context.Context, pagination page.PaginationRequest) (page.Page[p.Payment], error) {
	sqlRows, err := r.queries.FindOverduePayments(ctx, sqlc.FindOverduePaymentsParams{
		Limit:  int32(pagination.Limit()),
		Offset: int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[p.Payment]{}, fmt.Errorf("failed to find overdue payments: %w", err)
	}

	total, err := r.queries.CountOverduePayments(ctx)
	if err != nil {
		return page.Page[p.Payment]{}, fmt.Errorf("failed to count overdue payments: %w", err)
	}

	payments := r.sqlRowsToEntities(sqlRows)
	return page.NewPage(payments, total, pagination), nil
}

func (r *SqlcPaymentRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time, pagination page.PaginationRequest) (page.Page[p.Payment], error) {
	sqlRows, err := r.queries.FindPaymentsByDateRange(ctx, sqlc.FindPaymentsByDateRangeParams{
		CreatedAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		CreatedAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Limit:       int32(pagination.Limit()),
		Offset:      int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[p.Payment]{}, fmt.Errorf("failed to find payments by date range: %w", err)
	}

	total, err := r.queries.CountPaymentsByDateRange(ctx, sqlc.CountPaymentsByDateRangeParams{
		CreatedAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		CreatedAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return page.Page[p.Payment]{}, fmt.Errorf("failed to count payments by date range: %w", err)
	}

	payments := r.sqlRowsToEntities(sqlRows)
	return page.NewPage(payments, total, pagination), nil
}

func (r *SqlcPaymentRepository) FindRecentByCustomerID(ctx context.Context, customerID vo.CustomerID, limit int) ([]p.Payment, error) {
	sqlRows, err := r.queries.FindRecentPaymentsByCustomerID(ctx, sqlc.FindRecentPaymentsByCustomerIDParams{
		PaidByCustomerID: pgtype.Int4{Int32: int32(customerID.Int32()), Valid: true},
		Limit:            int32(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find recent payments by customer ID: %w", err)
	}

	payments := r.sqlRowsToEntities(sqlRows)
	return payments, nil
}

func (r *SqlcPaymentRepository) FindPendingPayments(ctx context.Context, pagination page.PaginationRequest) (page.Page[p.Payment], error) {
	return r.FindByStatus(ctx, enum.PaymentStatusPending, pagination)
}

func (r *SqlcPaymentRepository) FindSuccessfulPayments(ctx context.Context, pagination page.PaginationRequest) (page.Page[p.Payment], error) {
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
