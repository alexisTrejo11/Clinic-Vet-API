package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	med "clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/specification"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/mapper"
	"clinic-vet-api/app/shared/page"
	p "clinic-vet-api/app/shared/page"
	"clinic-vet-api/sqlc"
)

type SqlcPetDeworming struct {
	queries *sqlc.Queries
	mapper  mapper.SqlcFieldMapper
}

func NewSqlcPetDeworming(queries *sqlc.Queries) repository.DewormRepository {
	return &SqlcPetDeworming{
		queries: queries,
		mapper:  mapper.SqlcFieldMapper{},
	}
}

func (r *SqlcPetDeworming) Save(ctx context.Context, deworming med.PetDeworming) (med.PetDeworming, error) {
	if deworming.ID().IsZero() {
		return r.create(ctx, deworming)
	}
	return r.update(ctx, deworming)
}

func (r *SqlcPetDeworming) FindByID(ctx context.Context, id vo.DewormID) (*med.PetDeworming, error) {
	result, err := r.queries.FindDewormingByID(ctx, id.Int32())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPetDewormingNotFound
		}
		return nil, fmt.Errorf("failed to get pet deworming by ID: %w", err)
	}

	return r.mapRowToDomain(result), nil
}

func (r *SqlcPetDeworming) FindBySpecification(ctx context.Context, spec specification.PetDewormSpecification) (p.Page[med.PetDeworming], error) {
	params := r.specToSqlc(spec)
	results, err := r.queries.FindDewormingsBySpec(ctx, params)
	if err != nil {
		return p.Page[med.PetDeworming]{}, fmt.Errorf("failed to find pet dewormings by specification: %w", err)
	}

	total, err := r.CountBySpecification(ctx, spec)
	if err != nil {
		return p.Page[med.PetDeworming]{}, fmt.Errorf("failed to count pet dewormings by specification: %w", err)
	}

	dewormings := r.mapRowsToDomains(results)
	pagination := page.PaginationRequest{
		PageSize: *spec.Limit,
		Page:     *spec.Offset / *spec.Limit + 1,
	}
	return page.NewPage(dewormings, total, pagination), nil
}

func (r *SqlcPetDeworming) Delete(ctx context.Context, id vo.DewormID, isHard bool) error {
	err := r.queries.SoftDeleteDeworming(ctx, id.Int32())
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrPetDewormingNotFound
		}
		return fmt.Errorf("failed to delete pet deworming: %w", err)
	}

	return nil
}

func (r *SqlcPetDeworming) CountBySpecification(ctx context.Context, spec specification.PetDewormSpecification) (int64, error) {
	params := r.ToCountSQLParams(spec)

	count, err := r.queries.CountDewormingsBySpec(ctx, params)
	if err != nil {
		return 0, fmt.Errorf("failed to count pet dewormings by specification: %w", err)
	}

	return count, nil
}

func (r *SqlcPetDeworming) create(ctx context.Context, deworming med.PetDeworming) (med.PetDeworming, error) {
	params := r.domainToCreateParams(deworming)
	result, err := r.queries.CreateDeworming(ctx, params)
	if err != nil {
		return med.PetDeworming{}, fmt.Errorf("failed to create pet deworming: %w", err)
	}

	return *r.mapRowToDomain(result), nil
}

func (r *SqlcPetDeworming) update(ctx context.Context, deworming med.PetDeworming) (med.PetDeworming, error) {
	params := r.domainToUpdateParams(deworming)
	result, err := r.queries.UpdateDeworming(ctx, params)
	if err != nil {
		return med.PetDeworming{}, fmt.Errorf("failed to update pet deworming: %w", err)
	}

	return *r.mapRowToDomain(result), nil
}

func (r *SqlcPetDeworming) FindByIDAndPetID(ctx context.Context, dewormationID vo.DewormID, petID vo.PetID) (*med.PetDeworming, error) {
	result, err := r.queries.FindDewormingByIDAndPetID(ctx, sqlc.FindDewormingByIDAndPetIDParams{
		ID:    dewormationID.Int32(),
		PetID: int32(petID.Value()),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPetDewormingNotFound
		}
		return nil, fmt.Errorf("failed to get pet deworming by ID and pet ID: %w", err)
	}

	return r.mapRowToDomain(result), nil
}

func (r *SqlcPetDeworming) FindByIDAndEmployeeID(ctx context.Context, dewormationID vo.DewormID, employeeID vo.EmployeeID) (*med.PetDeworming, error) {
	result, err := r.queries.FindDewormingByIDAndEmployeeID(ctx, sqlc.FindDewormingByIDAndEmployeeIDParams{
		ID:             dewormationID.Int32(),
		AdministeredBy: r.mapper.UintToPgInt4(employeeID.Value()),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPetDewormingNotFound
		}
		return nil, fmt.Errorf("failed to get pet deworming by ID and employee ID: %w", err)
	}

	return r.mapRowToDomain(result), nil
}

func (r *SqlcPetDeworming) FindByEmployeeID(ctx context.Context, employeeID vo.EmployeeID, pagination page.PaginationRequest) (p.Page[med.PetDeworming], error) {
	results, err := r.queries.FindDewormingsByEmployeeID(ctx, sqlc.FindDewormingsByEmployeeIDParams{
		AdministeredBy: r.mapper.UintToPgInt4(employeeID.Value()),
		Limit:          pagination.PageSize,
		Offset:         (pagination.Page - 1) * pagination.PageSize,
	})
	if err != nil {
		return p.Page[med.PetDeworming]{}, fmt.Errorf("failed to get pet dewormings by employee ID: %w", err)
	}

	total, err := r.queries.CountDewormingsByEmployeeID(ctx, r.mapper.UintToPgInt4(employeeID.Value()))
	if err != nil {
		return p.Page[med.PetDeworming]{}, fmt.Errorf("failed to count pet dewormings by employee ID: %w", err)
	}

	dewormings := r.mapRowsToDomains(results)
	return page.NewPage(dewormings, total, pagination), nil
}
func (r *SqlcPetDeworming) FindByPetID(ctx context.Context, petID vo.PetID, pagination page.PaginationRequest) (p.Page[med.PetDeworming], error) {
	results, err := r.queries.FindDewormingsByPetID(ctx, sqlc.FindDewormingsByPetIDParams{
		PetID:  petID.Int32(),
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	})
	if err != nil {
		return p.Page[med.PetDeworming]{}, fmt.Errorf("failed to get pet dewormings by pet ID: %w", err)
	}

	total, err := r.queries.CountDewormingsByPetID(ctx, int32(petID.Value()))
	if err != nil {
		return p.Page[med.PetDeworming]{}, fmt.Errorf("failed to count pet dewormings by pet ID: %w", err)
	}

	dewormings := r.mapRowsToDomains(results)
	return page.NewPage(dewormings, total, pagination), nil
}

func (r *SqlcPetDeworming) FindByPetIDs(ctx context.Context, petIDs []vo.PetID, pagination p.PaginationRequest) (p.Page[med.PetDeworming], error) {
	intPetIDs := make([]int32, len(petIDs))
	for i, id := range petIDs {
		intPetIDs[i] = id.Int32()
	}

	results, err := r.queries.FindDewormingsByPetIDs(ctx, sqlc.FindDewormingsByPetIDsParams{
		Column1: intPetIDs,
		Limit:   pagination.Limit(),
		Offset:  pagination.Offset(),
	})
	if err != nil {
		return p.Page[med.PetDeworming]{}, fmt.Errorf("failed to get pet dewormings by pet IDs: %w", err)
	}

	total, err := r.queries.CountDewormingsByPetIDs(ctx, intPetIDs)
	if err != nil {
		return p.Page[med.PetDeworming]{}, fmt.Errorf("failed to count pet dewormings by pet IDs: %w", err)
	}

	dewormings := r.mapRowsToDomains(results)
	return page.NewPage(dewormings, total, pagination), nil
}

func (r *SqlcPetDeworming) FindByDateRange(ctx context.Context, startDate, endDate time.Time, pagination page.PaginationRequest) (p.Page[med.PetDeworming], error) {
	results, err := r.queries.FindDewormingsByDateRange(ctx, sqlc.FindDewormingsByDateRangeParams{
		AdministeredDate:   r.mapper.TimeToPgDate(startDate),
		AdministeredDate_2: r.mapper.TimeToPgDate(endDate),
		Limit:              pagination.Limit(),
		Offset:             pagination.Offset(),
	})
	if err != nil {
		return p.Page[med.PetDeworming]{}, fmt.Errorf("failed to get pet dewormings by date range: %w", err)
	}

	total, err := r.queries.CountDewormingsByDateRange(ctx,
		sqlc.CountDewormingsByDateRangeParams{
			AdministeredDate:   r.mapper.TimeToPgDate(startDate),
			AdministeredDate_2: r.mapper.TimeToPgDate(endDate),
		})
	if err != nil {
		return p.Page[med.PetDeworming]{}, fmt.Errorf("failed to count pet dewormings by date range: %w", err)
	}

	dewormings := r.mapRowsToDomains(results)
	return page.NewPage(dewormings, total, pagination), nil
}
