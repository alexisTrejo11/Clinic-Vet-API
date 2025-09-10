// Package persistence includes the implementation of repositories using SQLC
package repositoryimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/customer"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SqlcEmployeeRepository struct {
	queries *sqlc.Queries
	pool    *pgxpool.Pool
}

func NewSqlcEmployeeRepository(queries *sqlc.Queries, pool *pgxpool.Pool) repository.VetRepository {
	return &SqlcEmployeeRepository{
		queries: queries,
		pool:    pool,
	}
}

func (r *SqlcEmployeeRepository) Search(ctx context.Context, spec specification.VetSearchSpecification) (page.Page[[]customer.Customer], error) {
	query, params := spec.ToSQL()

	// Execute the query
	rows, err := r.pool.Query(ctx, query, params...)
	if err != nil {
		return page.Page[[]customer.Customer]{}, r.dbError(OpSelect, "failed to search veterinarians", err)
	}
	defer rows.Close()

	// Iterate through the rows and scan into Customer structs
	var vets []customer.Customer
	for rows.Next() {
		var veterinarian customer.Customer
		err := r.scanEmployeeFromRow(rows, &veterinarian)
		if err != nil {
			return page.Page[[]customer.Customer]{}, r.wrapConversionError(err)
		}
		vets = append(vets, veterinarian)
	}

	if err := rows.Err(); err != nil {
		return page.Page[[]customer.Customer]{}, r.dbError(OpSelect, "error iterating search results", err)
	}

	// Get total count for pagination
	totalCount, err := r.getTotalCountWithFilters(ctx, &spec)
	if err != nil {
		return page.Page[[]customer.Customer]{}, err
	}

	// Handle pagination
	pageMetadata := page.GetPageMetadata(totalCount, page.PageInput{
		PageNumber: spec.GetPagination().Page,
		PageSize:   spec.GetPagination().PageSize,
	})

	return page.NewPage(vets, *pageMetadata), nil
}

func (r *SqlcEmployeeRepository) GetByID(ctx context.Context, id valueobject.EmployeeID) (customer.Customer, error) {
	sqlEmployee, err := r.queries.GetVeterinarianById(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customer.Customer{}, r.notFoundError("id", id.String())
		}
		return customer.Customer{}, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgGetVet, id.Value()), err)
	}

	veterinarian, err := SqlcEmployeeToDomain(sqlVet)
	if err != nil {
		return customer.Customer{}, r.wrapConversionError(err)
	}

	return *veterinarian, nil
}

func (r *SqlcEmployeeRepository) GetByUserID(ctx context.Context, userID valueobject.UserID) (customer.Customer, error) {
	sqlEmployee, err := r.queries.GetVeterinarianById(ctx, int32(userID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customer.Customer{}, r.notFoundError("user_id", userID.String())
		}
		return customer.Customer{}, r.dbError(OpSelect, fmt.Sprintf("%s with user ID %d", ErrMsgGetVetByUserID, userID.Value()), err)
	}

	veterinarian, err := SqlcEmployeeToDomain(sqlEmployee)
	if err != nil {
		return customer.Customer{}, r.wrapConversionError(err)
	}
	return *veterinarian, nil
}

func (r *SqlcEmployeeRepository) Save(ctx context.Context, veterinarian *customer.Customer) error {
	if veterinarian.ID().IsZero() {
		return r.create(ctx, veterinarian)
	}
	return r.update(ctx, veterinarian)
}

func (r *SqlcEmployeeRepository) SoftDelete(ctx context.Context, id valueobject.EmployeeID) error {
	if err := r.queries.SoftDeleteEmployeeerinarian(ctx, int32(id.Value())); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("%s with ID %d", ErrMsgSoftDeleteEmployee, id.Value()), err)
	}
	return nil
}

func (r *SqlcEmployeeRepository) Exists(ctx context.Context, id valueobject.EmployeeID) (bool, error) {
	_, err := r.queries.GetEmployeeerinarianById(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgCheckEmployeeExists, id.Value()), err)
	}
	return true, nil
}

func (r *SqlcEmployeeRepository) create(ctx context.Context, customer *customer.Customer) error {
	createParams := vetToCreateParams(customer)
	_, err := r.queries.CreateEmployeeerinarian(ctx, *createParams)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreateEmployee, err)
	}

	return nil
}

func (r *SqlcEmployeeRepository) update(ctx context.Context, customer *customer.Customer) error {
	updateParams := vetToUpdateParams(customer)
	_, err := r.queries.UpdateEmployeeerinarian(ctx, *updateParams)
	if err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("%s with ID %d", ErrMsgUpdateEmployee, customer.ID().Value()), err)
	}
	return nil
}
