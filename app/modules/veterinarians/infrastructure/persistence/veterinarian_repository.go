// Package persistence includes the implementation of repositories using SQLC
package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	vet "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/veterinarian"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SqlcVetRepository struct {
	queries *sqlc.Queries
	pool    *pgxpool.Pool
}

func NewSqlcVetRepository(queries *sqlc.Queries, pool *pgxpool.Pool) repository.VetRepository {
	return &SqlcVetRepository{
		queries: queries,
		pool:    pool,
	}
}

func (r *SqlcVetRepository) Search(ctx context.Context, spec specification.VetSearchSpecification) (page.Page[[]vet.Veterinarian], error) {
	query, params := spec.ToSQL()

	// Execute the query
	rows, err := r.pool.Query(ctx, query, params...)
	if err != nil {
		return page.Page[[]vet.Veterinarian]{}, r.dbError(OpSelect, "failed to search veterinarians", err)
	}
	defer rows.Close()

	// Iterate through the rows and scan into Veterinarian structs
	var vets []vet.Veterinarian
	for rows.Next() {
		var veterinarian vet.Veterinarian
		err := r.scanVetFromRow(rows, &veterinarian)
		if err != nil {
			return page.Page[[]vet.Veterinarian]{}, r.wrapConversionError(err)
		}
		vets = append(vets, veterinarian)
	}

	if err := rows.Err(); err != nil {
		return page.Page[[]vet.Veterinarian]{}, r.dbError(OpSelect, "error iterating search results", err)
	}

	// Get total count for pagination
	totalCount, err := r.getTotalCountWithFilters(ctx, &spec)
	if err != nil {
		return page.Page[[]vet.Veterinarian]{}, err
	}

	// Handle pagination
	pageMetadata := page.GetPageMetadata(totalCount, page.PageInput{
		PageNumber: spec.GetPagination().Page,
		PageSize:   spec.GetPagination().PageSize,
	})

	return page.NewPage(vets, *pageMetadata), nil
}

func (r *SqlcVetRepository) GetByID(ctx context.Context, id valueobject.VetID) (vet.Veterinarian, error) {
	sqlVet, err := r.queries.GetVeterinarianById(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return vet.Veterinarian{}, r.notFoundError("id", id.String())
		}
		return vet.Veterinarian{}, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgGetVet, id.Value()), err)
	}

	veterinarian, err := SqlcVetToDomain(sqlVet)
	if err != nil {
		return vet.Veterinarian{}, r.wrapConversionError(err)
	}

	return *veterinarian, nil
}

func (r *SqlcVetRepository) GetByUserID(ctx context.Context, userID valueobject.UserID) (vet.Veterinarian, error) {
	sqlVet, err := r.queries.GetVeterinarianById(ctx, int32(userID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return vet.Veterinarian{}, r.notFoundError("user_id", userID.String())
		}
		return vet.Veterinarian{}, r.dbError(OpSelect, fmt.Sprintf("%s with user ID %d", ErrMsgGetVetByUserID, userID.Value()), err)
	}

	veterinarian, err := SqlcVetToDomain(sqlVet)
	if err != nil {
		return vet.Veterinarian{}, r.wrapConversionError(err)
	}
	return *veterinarian, nil
}

func (r *SqlcVetRepository) Save(ctx context.Context, veterinarian *vet.Veterinarian) error {
	if veterinarian.ID().IsZero() {
		return r.create(ctx, veterinarian)
	}
	return r.update(ctx, veterinarian)
}

func (r *SqlcVetRepository) SoftDelete(ctx context.Context, id valueobject.VetID) error {
	if err := r.queries.SoftDeleteVeterinarian(ctx, int32(id.Value())); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("%s with ID %d", ErrMsgSoftDeleteVet, id.Value()), err)
	}
	return nil
}

func (r *SqlcVetRepository) Exists(ctx context.Context, id valueobject.VetID) (bool, error) {
	_, err := r.queries.GetVeterinarianById(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgCheckVetExists, id.Value()), err)
	}
	return true, nil
}

func (r *SqlcVetRepository) create(ctx context.Context, vet *vet.Veterinarian) error {
	createParams := vetToCreateParams(vet)
	_, err := r.queries.CreateVeterinarian(ctx, *createParams)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreateVet, err)
	}

	return nil
}

func (r *SqlcVetRepository) update(ctx context.Context, vet *vet.Veterinarian) error {
	updateParams := vetToUpdateParams(vet)
	_, err := r.queries.UpdateVeterinarian(ctx, *updateParams)
	if err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("%s with ID %d", ErrMsgUpdateVet, vet.ID().Value()), err)
	}
	return nil
}
