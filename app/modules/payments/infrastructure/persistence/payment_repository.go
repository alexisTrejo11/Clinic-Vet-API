package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/payment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
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

func (r *SQLCPaymentRepository) GetByIDAndPaidFrom(ctx context.Context, id valueobject.OwnerID) (payment.Payment, error) {
	return payment.Payment{}, errors.New("method not implemented")
}

// Save creates or updates a payment based on whether it has an ID
func (r *SQLCPaymentRepository) Save(ctx context.Context, payment *payment.Payment) error {
	if payment.ID().IsZero() {
		return r.create(ctx, payment)
	}
	return r.update(ctx, payment)
}

func (r *SQLCPaymentRepository) GetByID(ctx context.Context, id valueobject.PaymentID) (payment.Payment, error) {
	sqlRow, err := r.queries.GetPaymentById(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return payment.Payment{}, r.notFoundError("id", fmt.Sprintf("%d", id))
		}
		return payment.Payment{}, r.dbError(OpSelect, fmt.Sprintf("failed to get payment with ID %d", id), err)
	}

	paymentEntity, err := sqlcToEntity(sqlRow)
	if err != nil {
		return payment.Payment{}, r.wrapConversionError(err)
	}

	return paymentEntity, nil
}

func (r *SQLCPaymentRepository) GetByTransactionID(ctx context.Context, transactionID string) (payment.Payment, error) {
	sqlRow, err := r.queries.GetPaymentByTransactionId(
		ctx,
		pgtype.Text{String: transactionID, Valid: true},
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return payment.Payment{}, r.notFoundError("transaction_id", transactionID)
		}
		return payment.Payment{}, r.dbError(OpSelect, fmt.Sprintf("failed to get payment with transaction ID %s", transactionID), err)
	}

	payment, err := sqlcToEntity(sqlRow)
	if err != nil {
		return payment.Payment{}, r.wrapConversionError(err)
	}

	return payment, nil
}

// TODO: fix
func (r *SQLCPaymentRepository) ListByPaidFrom(ctx context.Context, ownerID valueobject.OwnerID, pageInput page.PageInput) (page.Page[[]payment.Payment], error) {
	params := sqlc.ListPaymentsByUserIdParams{
		UserID: int32(ownerID.Value()),
		Limit:  int32(pageInput.PageSize),
		Offset: r.calculateOffset(pageInput),
	}

	sqlRows, err := r.queries.ListPaymentsByUserId(ctx, params)
	if err != nil {
		return page.Page[[]payment.Payment]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list payments for user ID %d", ownerID.Value()), err)
	}

	payments, err := sqlcToEntityList(sqlRows)
	if err != nil {
		return page.Page[[]payment.Payment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountPaymentsByUserId(ctx, int32(ownerID.Value()))
	if err != nil {
		return page.Page[[]payment.Payment]{}, r.dbError(OpCount, fmt.Sprintf("failed to count payments for user ID %d", ownerID.Value()), err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(payments, *pageMetadata), nil
}

func (r *SQLCPaymentRepository) ListByStatus(ctx context.Context, status enum.PaymentStatus, pageInput page.PageInput) (page.Page[[]payment.Payment], error) {
	params := sqlc.ListPaymentsByStatusParams{
		Status: models.PaymentStatus(status),
		Limit:  int32(pageInput.PageSize),
		Offset: r.calculateOffset(pageInput),
	}

	sqlRows, err := r.queries.ListPaymentsByStatus(ctx, params)
	if err != nil {
		return page.Page[[]payment.Payment]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list payments with status %s", status), err)
	}

	payments, err := sqlcToEntityList(sqlRows)
	if err != nil {
		return page.Page[[]payment.Payment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountPaymentsByStatus(ctx, models.PaymentStatus(status))
	if err != nil {
		return page.Page[[]payment.Payment]{}, r.dbError(OpCount, fmt.Sprintf("failed to count payments with status %s", status), err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(payments, *pageMetadata), nil
}

func (r *SQLCPaymentRepository) ListOverduePayments(ctx context.Context, pageInput page.PageInput) (page.Page[[]payment.Payment], error) {
	params := sqlc.ListOverduePaymentsParams{
		Limit:  int32(pageInput.PageSize),
		Offset: r.calculateOffset(pageInput),
	}

	sqlRows, err := r.queries.ListOverduePayments(ctx, params)
	if err != nil {
		return page.Page[[]payment.Payment]{}, r.dbError(OpSelect, ErrMsgListOverduePayments, err)
	}

	payments, err := sqlcToEntityList(sqlRows)
	if err != nil {
		return page.Page[[]payment.Payment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountOverduePayments(ctx)
	if err != nil {
		return page.Page[[]payment.Payment]{}, r.dbError(OpCount, "failed to count overdue payments", err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(payments, *pageMetadata), nil
}

func (r *SQLCPaymentRepository) ListPaymentsByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageInput) (page.Page[[]payment.Payment], error) {
	params := sqlc.ListPaymentsByDateRangeParams{
		CreatedAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		CreatedAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Limit:       int32(pageInput.PageSize),
		Offset:      r.calculateOffset(pageInput),
	}

	sqlRows, err := r.queries.ListPaymentsByDateRange(ctx, params)
	if err != nil {
		return page.Page[[]payment.Payment]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list payments between %s and %s", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")), err)
	}

	payments, err := sqlcToEntityList(sqlRows)
	if err != nil {
		return page.Page[[]payment.Payment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountPaymentsByDateRange(ctx, sqlc.CountPaymentsByDateRangeParams{
		CreatedAt:   pgtype.Timestamptz{Time: startDate, Valid: true},
		CreatedAt_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return page.Page[[]payment.Payment]{}, r.dbError(OpCount, fmt.Sprintf("failed to count payments between %s and %s", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")), err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(payments, *pageMetadata), nil
}

func (r *SQLCPaymentRepository) Search(ctx context.Context, pageInput page.PageInput, searchCriteria map[string]any) (page.Page[[]payment.Payment], error) {
	return page.Page[[]payment.Payment]{}, r.dbError(OpSelect, ErrMsgSearchPayments, errors.New("search method not implemented"))
}

func (r *SQLCPaymentRepository) SoftDelete(ctx context.Context, id valueobject.PaymentID) error {
	if err := r.queries.SoftDeletePayment(ctx, int32(id.Value())); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("failed to soft delete payment with ID %d", id), err)
	}
	return nil
}

func (r *SQLCPaymentRepository) create(ctx context.Context, payment *payment.Payment) error {
	createParams := entityToCreateParams(payment)

	_, err := r.queries.CreatePayment(ctx, createParams)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreatePayment, err)
	}

	return nil
}

func (r *SQLCPaymentRepository) update(ctx context.Context, payment *payment.Payment) error {
	params := mapDomainToUpdateParams(payment)

	_, err := r.queries.UpdatePayment(ctx, params)
	if err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("failed to update payment with ID %d", payment.ID().Value()), err)
	}

	return nil
}
