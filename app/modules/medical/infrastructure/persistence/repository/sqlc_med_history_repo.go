// Package repositoryimpl implements the MedicalHistoryRepository using SQLC for database operations.
package repositoryimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	med "clinic-vet-api/app/core/domain/entity/medical"
	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	p "clinic-vet-api/app/shared/page"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCMedHistRepository struct {
	queries *sqlc.Queries
}

func NewSQLCMedHistRepository(queries *sqlc.Queries) repository.MedicalHistoryRepository {
	return &SQLCMedHistRepository{
		queries: queries,
	}
}

func (r *SQLCMedHistRepository) FindBySpecification(ctx context.Context, spec specification.MedicalHistorySpecification) (p.Page[med.MedicalHistory], error) {
	return p.Page[med.MedicalHistory]{}, errors.New("FindBySpecification not implemented")
}

func (r *SQLCMedHistRepository) FindByID(ctx context.Context, medHistID valueobject.MedHistoryID) (med.MedicalHistory, error) {
	sqlcMedHist, err := r.queries.FindMedicalHistoryByID(ctx, int32(medHistID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return med.MedicalHistory{}, r.notFoundError("id", fmt.Sprintf("%d", medHistID.Value()))
		}
		return med.MedicalHistory{}, r.dbError("select", fmt.Sprintf("failed to find medical history with ID %d", medHistID.Value()), err)
	}

	medHist, err := ToEntity(sqlcMedHist)
	if err != nil {
		return med.MedicalHistory{}, r.wrapConversionError(err)
	}

	return medHist, nil
}

func (r *SQLCMedHistRepository) FindAll(ctx context.Context, pageInput p.PageInput) (p.Page[med.MedicalHistory], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize

	// Get medical histories
	medHistRows, err := r.queries.FindAllMedicalHistory(ctx, sqlc.FindAllMedicalHistoryParams{
		Limit:  int32(pageInput.PageSize),
		Offset: int32(offset),
	})
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.dbError("select", "failed to find all medical history", err)
	}

	// Get total count
	total, err := r.queries.CountAllMedicalHistory(ctx)
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.dbError("select", "failed to count all medical history", err)
	}

	medicalHistories, err := ToEntities(medHistRows)
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.wrapConversionError(err)
	}

	paginationMetadata := p.GetPageMetadata(int(total), pageInput)
	return p.NewPage(medicalHistories, *paginationMetadata), nil
}

func (r *SQLCMedHistRepository) FindByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID, pageInput p.PageInput) (p.Page[med.MedicalHistory], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize

	// Get medical histories
	medHistRows, err := r.queries.FindMedicalHistoryByEmployeeID(ctx, sqlc.FindMedicalHistoryByEmployeeIDParams{
		EmployeeID: int32(employeeID.Value()),
		Limit:      int32(pageInput.PageSize),
		Offset:     int32(offset),
	})
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.dbError("select", fmt.Sprintf("failed to find medical history for employee ID %d", employeeID.Value()), err)
	}

	// Get total count
	total, err := r.queries.CountMedicalHistoryByEmployeeID(ctx, int32(employeeID.Value()))
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.dbError("select", fmt.Sprintf("failed to count medical history for employee ID %d", employeeID.Value()), err)
	}

	medicalHistories, err := ToEntities(medHistRows)
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.wrapConversionError(err)
	}

	paginationMetadata := p.GetPageMetadata(int(total), pageInput)
	return p.NewPage(medicalHistories, *paginationMetadata), nil
}

func (r *SQLCMedHistRepository) FindByPetID(ctx context.Context, petID valueobject.PetID, pageInput p.PageInput) (p.Page[med.MedicalHistory], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize

	// Get medical histories
	medHistRows, err := r.queries.FindMedicalHistoryByPetID(ctx, sqlc.FindMedicalHistoryByPetIDParams{
		PetID:  int32(petID.Value()),
		Limit:  int32(pageInput.PageSize),
		Offset: int32(offset),
	})
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.dbError("select", fmt.Sprintf("failed to find medical history for pet ID %d", petID.Value()), err)
	}

	// Get total count
	total, err := r.queries.CountMedicalHistoryByPetID(ctx, int32(petID.Value()))
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.dbError("select", fmt.Sprintf("failed to count medical history for pet ID %d", petID.Value()), err)
	}

	medicalHistories, err := ToEntities(medHistRows)
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.wrapConversionError(err)
	}

	paginationMetadata := p.GetPageMetadata(int(total), pageInput)
	return p.NewPage(medicalHistories, *paginationMetadata), nil
}

func (r *SQLCMedHistRepository) FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID, pageInput p.PageInput) (p.Page[med.MedicalHistory], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize

	// Get medical histories
	medHistRows, err := r.queries.FindMedicalHistoryByCustomerID(ctx, sqlc.FindMedicalHistoryByCustomerIDParams{
		CustomerID: int32(customerID.Value()),
		Limit:      int32(pageInput.PageSize),
		Offset:     int32(offset),
	})
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.dbError("select", fmt.Sprintf("failed to find medical history for customer ID %d", customerID.Value()), err)
	}

	// Get total count
	total, err := r.queries.CountMedicalHistoryByCustomerID(ctx, int32(customerID.Value()))
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.dbError("select", fmt.Sprintf("failed to count medical history for customer ID %d", customerID.Value()), err)
	}

	medicalHistories, err := ToEntities(medHistRows)
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.wrapConversionError(err)
	}

	paginationMetadata := p.GetPageMetadata(int(total), pageInput)
	return p.NewPage(medicalHistories, *paginationMetadata), nil
}

func (r *SQLCMedHistRepository) FindRecentByPetID(ctx context.Context, petID valueobject.PetID, limit int) ([]med.MedicalHistory, error) {
	// Get recent medical histories
	medHistRows, err := r.queries.FindRecentMedicalHistoryByPetID(ctx, sqlc.FindRecentMedicalHistoryByPetIDParams{
		PetID: int32(petID.Value()),
		Limit: int32(limit),
	})
	if err != nil {
		return nil, r.dbError("select", fmt.Sprintf("failed to find recent medical history for pet ID %d", petID.Value()), err)
	}

	// Convert rows to entities
	var medicalHistories []med.MedicalHistory
	for _, row := range medHistRows {
		medHistEntity, err := ToEntity(row)
		if err != nil {
			return nil, r.wrapConversionError(err)
		}
		medicalHistories = append(medicalHistories, medHistEntity)
	}

	return medicalHistories, nil
}

func (r *SQLCMedHistRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput p.PageInput) (p.Page[med.MedicalHistory], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize

	// Get medical histories
	medHistRows, err := r.queries.FindMedicalHistoryByDateRange(ctx, sqlc.FindMedicalHistoryByDateRangeParams{
		VisitDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		VisitDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Limit:       int32(pageInput.PageSize),
		Offset:      int32(offset),
	})
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.dbError("select", "failed to find medical history by date range", err)
	}

	// Get total count
	total, err := r.queries.CountMedicalHistoryByDateRange(ctx, sqlc.CountMedicalHistoryByDateRangeParams{
		VisitDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		VisitDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.dbError("select", "failed to count medical history by date range", err)
	}

	medicalHistories, err := ToEntities(medHistRows)
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.wrapConversionError(err)
	}

	paginationMetadata := p.GetPageMetadata(int(total), pageInput)
	return p.NewPage(medicalHistories, *paginationMetadata), nil
}

func (r *SQLCMedHistRepository) FindByPetAndDateRange(ctx context.Context, petID valueobject.PetID, startDate, endDate time.Time) ([]med.MedicalHistory, error) {
	// Get medical histories
	medHistRows, err := r.queries.FindMedicalHistoryByPetAndDateRange(ctx, sqlc.FindMedicalHistoryByPetAndDateRangeParams{
		PetID:       int32(petID.Value()),
		VisitDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		VisitDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return nil, r.dbError("select", fmt.Sprintf("failed to find medical history for pet ID %d and date range", petID.Value()), err)
	}

	// Convert rows to entities
	var medicalHistories []med.MedicalHistory
	for _, row := range medHistRows {
		medHistEntity, err := ToEntity(row)
		if err != nil {
			return nil, r.wrapConversionError(err)
		}
		medicalHistories = append(medicalHistories, medHistEntity)
	}

	return medicalHistories, nil
}

func (r *SQLCMedHistRepository) FindByDiagnosis(ctx context.Context, diagnosis string, pageInput p.PageInput) (p.Page[med.MedicalHistory], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize

	// Get medical histories
	medHistRows, err := r.queries.FindMedicalHistoryByDiagnosis(ctx, sqlc.FindMedicalHistoryByDiagnosisParams{
		Column1: pgtype.Text{String: diagnosis, Valid: true},
		Limit:   int32(pageInput.PageSize),
		Offset:  int32(offset),
	})
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.dbError("select", "failed to find medical history by diagnosis", err)
	}

	// Get total count
	total, err := r.queries.CountMedicalHistoryByDiagnosis(ctx, pgtype.Text{String: diagnosis, Valid: true})
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.dbError("select", "failed to count medical history by diagnosis", err)
	}

	medicalHistories, err := ToEntities(medHistRows)
	if err != nil {
		return p.Page[med.MedicalHistory]{}, r.wrapConversionError(err)
	}

	paginationMetadata := p.GetPageMetadata(int(total), pageInput)
	return p.NewPage(medicalHistories, *paginationMetadata), nil
}

func (r *SQLCMedHistRepository) ExistsByID(ctx context.Context, medicalHistoryID valueobject.MedHistoryID) (bool, error) {
	exists, err := r.queries.ExistsMedicalHistoryByID(ctx, int32(medicalHistoryID.Value()))
	if err != nil {
		return false, r.dbError("select", "failed to check medical history existence by ID", err)
	}
	return exists, nil
}

func (r *SQLCMedHistRepository) ExistsByPetAndDate(ctx context.Context, petID valueobject.PetID, date time.Time) (bool, error) {
	exists, err := r.queries.ExistsMedicalHistoryByPetAndDate(ctx, sqlc.ExistsMedicalHistoryByPetAndDateParams{
		PetID: int32(petID.Value()),
		Date:  pgtype.Timestamptz{Time: date, Valid: true},
	})
	if err != nil {
		return false, r.dbError("select", "failed to check medical history existence by pet and date", err)
	}
	return exists, nil
}

func (r *SQLCMedHistRepository) Save(ctx context.Context, medHistory *med.MedicalHistory) error {
	if medHistory.ID().IsZero() {
		return r.create(ctx, medHistory)
	}
	return r.update(ctx, medHistory)
}

func (r *SQLCMedHistRepository) Update(ctx context.Context, medHistory *med.MedicalHistory) error {
	return r.update(ctx, medHistory)
}

func (r *SQLCMedHistRepository) SoftDelete(ctx context.Context, medicalHistoryID valueobject.MedHistoryID) error {
	if err := r.queries.SoftDeleteMedicalHistory(ctx, int32(medicalHistoryID.Value())); err != nil {
		return r.dbError("delete", "failed to soft delete medical history", err)
	}
	return nil
}

func (r *SQLCMedHistRepository) HardDelete(ctx context.Context, medicalHistoryID valueobject.MedHistoryID) error {
	if err := r.queries.HardDeleteMedicalHistory(ctx, int32(medicalHistoryID.Value())); err != nil {
		return r.dbError("delete", "failed to hard delete medical history", err)
	}
	return nil
}

func (r *SQLCMedHistRepository) CountByPetID(ctx context.Context, petID valueobject.PetID) (int64, error) {
	count, err := r.queries.CountMedicalHistoryByPetID(ctx, int32(petID.Value()))
	if err != nil {
		return 0, r.dbError("select", "failed to count medical history by pet ID", err)
	}
	return count, nil
}

func (r *SQLCMedHistRepository) CountByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID) (int64, error) {
	count, err := r.queries.CountMedicalHistoryByEmployeeID(ctx, int32(employeeID.Value()))
	if err != nil {
		return 0, r.dbError("select", "failed to count medical history by employee ID", err)
	}
	return count, nil
}

func (r *SQLCMedHistRepository) CountByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (int64, error) {
	count, err := r.queries.CountMedicalHistoryByCustomerID(ctx, int32(customerID.Value()))
	if err != nil {
		return 0, r.dbError("select", "failed to count medical history by customer ID", err)
	}
	return count, nil
}

func (r *SQLCMedHistRepository) CountByDateRange(ctx context.Context, startDate, endDate time.Time) (int64, error) {
	count, err := r.queries.CountMedicalHistoryByDateRange(ctx, sqlc.CountMedicalHistoryByDateRangeParams{
		VisitDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		VisitDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return 0, r.dbError("select", "failed to count medical history by date range", err)
	}
	return count, nil
}

func (r *SQLCMedHistRepository) create(ctx context.Context, medHistory *med.MedicalHistory) error {
	params := ToCreateParams(*medHistory)
	_, err := r.queries.SaveMedicalHistory(ctx, params)
	if err != nil {
		return r.dbError("insert", "failed to create medical history", err)
	}

	return nil
}

func (r *SQLCMedHistRepository) update(ctx context.Context, medHistory *med.MedicalHistory) error {
	params := ToUpdateParams(*medHistory)

	_, err := r.queries.UpdateMedicalHistory(ctx, params)
	if err != nil {
		return r.dbError("update", fmt.Sprintf("failed to update medical history with ID %d", medHistory.ID().Value()), err)
	}

	return nil
}
