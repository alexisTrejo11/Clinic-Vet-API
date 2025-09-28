// Package repositoryimpl implements the PaymentRepository interface using SQLC.
package repositoryimpl

import (
	p "clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/mapper"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcPaymentRepository struct {
	queries *sqlc.Queries
	pgMap   mapper.SqlcFieldMapper
}

func NewSqlcPaymentRepository(queries *sqlc.Queries) repository.PaymentRepository {
	return &SqlcPaymentRepository{
		queries: queries,
		pgMap:   mapper.NewSqlcFieldMapper(),
	}
}

func (r *SqlcPaymentRepository) Save(ctx context.Context, paymentEntity *p.Payment) error {
	if paymentEntity.ID().IsZero() {
		return r.create(ctx, paymentEntity)
	}
	return r.update(ctx, paymentEntity)
}

func (r *SqlcPaymentRepository) CountByStatus(ctx context.Context, status enum.PaymentStatus) (int64, error) {
	count, err := r.queries.CountPaymentsByStatus(ctx, models.PaymentStatus(status.String()))
	if err != nil {
		return 0, r.dbError(OpCount, ErrMsgListPaymentsByStatus, err)
	}
	return count, nil
}

func (r *SqlcPaymentRepository) CountByCustomerID(ctx context.Context, customerID vo.CustomerID) (int64, error) {
	count, err := r.queries.CountPaymentsByCustomerID(ctx, pgtype.Int4{Int32: int32(customerID.Int32()), Valid: true})
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

func (r *SqlcPaymentRepository) TotalRevenueByDateRange(
	ctx context.Context, startDate, endDate time.Time,
) (float64, error) {
	params := sqlc.TotalRevenueByDateRangeParams{
		PaidAt:   r.pgMap.PgTimestamptz.FromTime(startDate),
		PaidAt_2: r.pgMap.PgTimestamptz.FromTime(endDate),
	}

	total, err := r.queries.TotalRevenueByDateRange(ctx, params)
	if err != nil {
		return 0, r.dbError(OpSelect, ErrMsgListPaymentsByDateRange, err)
	}

	totalFloat, ok := total.(float64)
	if !ok {
		return 0, r.wrapConversionError(fmt.Errorf("failed to convert total revenue to float64"))
	}
	return totalFloat, nil
}

func (r *SqlcPaymentRepository) Delete(ctx context.Context, id vo.PaymentID, isHard bool) error {
	if isHard {
		return r.hardDelete(ctx, id)
	}
	return r.softDelete(ctx, id)
}

func (r *SqlcPaymentRepository) create(ctx context.Context, paymentEntity *p.Payment) error {
	createParams := r.ToCreateParams(*paymentEntity)

	_, err := r.queries.CreatePayment(ctx, createParams)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreatePayment, err)
	}

	return nil
}

func (r *SqlcPaymentRepository) update(ctx context.Context, payment *p.Payment) error {
	updateParams := r.ToUpdateParams(payment)

	_, err := r.queries.UpdatePayment(ctx, updateParams)
	if err != nil {
		return r.dbError(OpUpdate, ErrMsgUpdatePayment, err)
	}

	return nil
}

func (r *SqlcPaymentRepository) softDelete(ctx context.Context, id vo.PaymentID) error {
	err := r.queries.SoftDeletePayment(ctx, id.Int32())
	if err != nil {
		return r.dbError(OpDelete, ErrMsgSoftDeletePayment, err)
	}
	return nil
}

func (r *SqlcPaymentRepository) hardDelete(ctx context.Context, id vo.PaymentID) error {
	err := r.queries.HardDeletePayment(ctx, id.Int32())
	if err != nil {
		return r.dbError(OpDelete, ErrMsgSoftDeletePayment, err)
	}
	return nil
}
