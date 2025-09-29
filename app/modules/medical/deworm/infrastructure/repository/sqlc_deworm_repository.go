package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/mapper"
	"clinic-vet-api/app/shared/page"
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

func (r *SqlcPetDeworming) Save(ctx context.Context, deworming medical.PetDeworming) (medical.PetDeworming, error) {
	if deworming.ID().IsZero() {
		return r.create(ctx, deworming)
	}
	return r.update(ctx, deworming)
}

func (r *SqlcPetDeworming) FindByID(ctx context.Context, id valueobject.DewormID) (*medical.PetDeworming, error) {
	result, err := r.queries.GetPetDewormingByID(ctx, id.Int32())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPetDewormingNotFound
		}
		return nil, fmt.Errorf("failed to get pet deworming by ID: %w", err)
	}

	return r.mapRowToDomain(result), nil
}

func (r *SqlcPetDeworming) FindByPetID(ctx context.Context, petID valueobject.PetID) ([]medical.PetDeworming, error) {
	results, err := r.queries.GetPetDewormingsByPetID(ctx, int32(petID.Value()))
	if err != nil {
		return nil, fmt.Errorf("failed to get pet dewormings by pet ID: %w", err)
	}

	dewormings := r.mapRowsToDomains(results)
	return dewormings, nil
}

func (r *SqlcPetDeworming) FindBySpecification(ctx context.Context, spec specification.PetDewormSpecification) (page.Page[medical.PetDeworming], error) {
	params := r.specToSqlc(spec)
	results, err := r.queries.FindPetDewormingsBySpec(ctx, params)
	if err != nil {
		return page.Page[medical.PetDeworming]{}, fmt.Errorf("failed to find pet dewormings by specification: %w", err)
	}

	total, err := r.CountBySpecification(ctx, spec)
	if err != nil {
		return page.Page[medical.PetDeworming]{}, fmt.Errorf("failed to count pet dewormings by specification: %w", err)
	}

	dewormings := r.mapRowsToDomains(results)
	pagination := page.PaginationRequest{
		PageSize: *spec.Limit,
		Page:     *spec.Offset / *spec.Limit + 1,
	}
	return page.NewPage(dewormings, total, pagination), nil
}

func (r *SqlcPetDeworming) Delete(ctx context.Context, id valueobject.DewormID, isHard bool) error {
	err := r.queries.DeletePetDeworming(ctx, id.Int32())
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

	count, err := r.queries.CountPetDewormingsBySpec(ctx, params)
	if err != nil {
		return 0, fmt.Errorf("failed to count pet dewormings by specification: %w", err)
	}

	return count, nil
}

func (r *SqlcPetDeworming) create(ctx context.Context, deworming medical.PetDeworming) (medical.PetDeworming, error) {
	params := r.domainToCreateParams(deworming)
	result, err := r.queries.CreatePetDeworming(ctx, params)
	if err != nil {
		return medical.PetDeworming{}, fmt.Errorf("failed to create pet deworming: %w", err)
	}

	return *r.mapRowToDomain(result), nil
}

func (r *SqlcPetDeworming) update(ctx context.Context, deworming medical.PetDeworming) (medical.PetDeworming, error) {
	params := r.domainToUpdateParams(deworming)
	result, err := r.queries.UpdatePetDeworming(ctx, params)
	if err != nil {
		return medical.PetDeworming{}, fmt.Errorf("failed to update pet deworming: %w", err)
	}

	return *r.mapRowToDomain(result), nil
}

func (r *SqlcPetDeworming) FindByIDAndPetID(ctx context.Context, dewormationID vo.DewormID, petID vo.PetID) (*med.PetDeworming, error) {
	result, err := r.queries.GetPetDewormingByIDAndPetID(ctx, sqlc.GetPetDewormingByIDAndPetIDParams{
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
	result, err := r.queries.GetPetDewormingByIDAndEmployeeID(ctx, sqlc.GetPetDewormingByIDAndEmployeeIDParams{
		ID:         dewormationID.Int32(),
		EmployeeID: int32(employeeID.Value()),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPetDewormingNotFound
		}
		return nil, fmt.Errorf("failed to get pet deworming by ID and employee ID: %w", err)
	}

	return r.mapRowToDomain(result), nil
}

func (r *SqlcPetDeworming) FindByEmployeeID(ctx context.Context, employeeID vo.EmployeeID, pagination page.PaginationRequest) (page.Page[med.PetDeworming], error) {
	results, err := r.queries.GetPetDewormingsByEmployeeID(ctx, sqlc.GetPetDewormingsByEmployeeIDParams{
		EmployeeID: int32(employeeID.Value()),
		Limit:      int32(pagination.PageSize),
		Offset:     int32((pagination.Page - 1) * pagination.PageSize),
	})
	if err != nil {
		return page.Page[med.PetDeworming]{}, fmt.Errorf("failed to get pet dewormings by employee ID: %w", err)
	}

	total, err := r.queries.CountPetDewormingsByEmployeeID(ctx, int32(employeeID.Value()))
	if err != nil {
		return page.Page[med.PetDeworming]{}, fmt.Errorf("failed to count pet dewormings by employee ID: %w", err)
	}

	dewormings := r.mapRowsToDomains(results)
	return page.NewPage(dewormings, total, pagination), nil
}
func (r *SqlcPetDeworming) FindByPetID(ctx context.Context, petID vo.PetID, pagination page.PaginationRequest) (page.Page[med.PetDeworming], error) {
	results, err := r.queries.GetPetDewormingsByPetIDWithPagination(ctx, sqlc.GetPetDewormingsByPetIDWithPaginationParams{
		PetID:  int32(petID.Value()),
		Limit:  int32(pagination.PageSize),
		Offset: int32((pagination.Page - 1) * pagination.PageSize),
	})
	if err != nil {
		return page.Page[med.PetDeworming]{}, fmt.Errorf("failed to get pet dewormings by pet ID: %w", err)
	}

	total, err := r.queries.CountPetDewormingsByPetID(ctx, int32(petID.Value()))
	if err != nil {
		return page.Page[med.PetDeworming]{}, fmt.Errorf("failed to count pet dewormings by pet ID: %w", err)
	}

	dewormings := r.mapRowsToDomains(results)
	return page.NewPage(dewormings, total, pagination), nil
}

func (r *SqlcPetDeworming) FindByDateRange(ctx context.Context, startDate, endDate time.Time, pagination page.PaginationRequest) (page.Page[med.PetDeworming], error) {
	results, err := r.queries.GetPetDewormingsByDateRange(ctx, sqlc.GetPetDewormingsByDateRangeParams{
		StartDate: r.mapper.TimeToPgDate(startDate),
		EndDate:   r.mapper.TimeToPgDate(endDate),
		Limit:     int32(pagination.PageSize),
		Offset:    int32((pagination.Page - 1) * pagination.PageSize),
	})
	if err != nil {
		return page.Page[med.PetDeworming]{}, fmt.Errorf("failed to get pet dewormings by date range: %w", err)
	}

	total, err := r.queries.CountPetDewormingsByDateRange(ctx,
		sqlc.CountPetDewormingsByDateRangeParams{
			StartDate: r.mapper.TimeToPgDate(startDate),
			EndDate:   r.mapper.TimeToPgDate(endDate),
		})
	if err != nil {
		return page.Page[med.PetDeworming]{}, fmt.Errorf("failed to count pet dewormings by date range: %w", err)
	}

	dewormings := r.mapRowsToDomains(results)
	return page.NewPage(dewormings, total, pagination), nil
}
