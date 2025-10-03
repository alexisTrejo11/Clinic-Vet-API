// Package repositoryimpl implements the MedicalSessionRepository using SQLC for database operations.
package repositoryimpl

import (
	"clinic-vet-api/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	med "clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"

	"clinic-vet-api/app/shared/mapper"
	p "clinic-vet-api/app/shared/page"
)

type SQLCMedSessionRepository struct {
	queries *sqlc.Queries
	pgMap   *mapper.SqlcFieldMapper
}

func NewSQLCMedSessionRepository(queries *sqlc.Queries) repository.MedicalSessionRepository {
	return &SQLCMedSessionRepository{
		queries: queries,
		pgMap:   mapper.NewSqlcFieldMapper(),
	}
}

func (r *SQLCMedSessionRepository) FindBySpecification(ctx context.Context, spec specification.MedicalSessionSpecification) (p.Page[med.MedicalSession], error) {
	return p.Page[med.MedicalSession]{}, errors.New("FindBySpecification not implemented")
}

func (r *SQLCMedSessionRepository) FindByIDAndCustomerID(ctx context.Context, medicalSessionID valueobject.MedSessionID, customerID valueobject.CustomerID) (*med.MedicalSession, error) {
	sqlcRow, err := r.queries.FindMedicalSessionByIDAndCustomerID(ctx, sqlc.FindMedicalSessionByIDAndCustomerIDParams{
		ID:         medicalSessionID.Int32(),
		CustomerID: customerID.Int32(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, r.notFoundError("id", fmt.Sprintf("%d", medicalSessionID.Int32()))
		}
		return nil, r.dbError("select", fmt.Sprintf("failed to find medical history with ID %d for customer ID %d", medicalSessionID.Value(), customerID.Int32()), err)
	}

	medSession := r.sqlcRowToEntity(sqlcRow)
	return &medSession, nil
}

func (r *SQLCMedSessionRepository) FindByIDAndEmployeeID(ctx context.Context, medicalSessionID valueobject.MedSessionID, employeeID valueobject.EmployeeID) (*med.MedicalSession, error) {
	sqlcRow, err := r.queries.FindMedicalSessionByIDAndEmployeeID(ctx, sqlc.FindMedicalSessionByIDAndEmployeeIDParams{
		ID:         medicalSessionID.Int32(),
		EmployeeID: employeeID.Int32(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, r.notFoundError("id", fmt.Sprintf("%d", medicalSessionID.Int32()))
		}
		return nil, r.dbError("select", fmt.Sprintf("failed to find medical history with ID %d for employee ID %d", medicalSessionID.Value(), employeeID.Int32()), err)
	}

	medSession := r.sqlcRowToEntity(sqlcRow)
	return &medSession, nil
}

func (r *SQLCMedSessionRepository) FindByIDAndPetID(ctx context.Context, medicalSessionID valueobject.MedSessionID, petID valueobject.PetID) (*med.MedicalSession, error) {
	sqlcRow, err := r.queries.FindMedicalSessionByIDAndPetID(ctx, sqlc.FindMedicalSessionByIDAndPetIDParams{
		ID:    medicalSessionID.Int32(),
		PetID: petID.Int32(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, r.notFoundError("id", fmt.Sprintf("%d", medicalSessionID.Int32()))
		}
		return nil, r.dbError("select", fmt.Sprintf("failed to find medical history with ID %d for pet ID %d", medicalSessionID.Value(), petID.Int32()), err)
	}

	medSession := r.sqlcRowToEntity(sqlcRow)
	return &medSession, nil
}

func (r *SQLCMedSessionRepository) FindByID(ctx context.Context, medSessionID valueobject.MedSessionID) (*med.MedicalSession, error) {
	sqlcRow, err := r.queries.FindMedicalSessionByID(ctx, medSessionID.Int32())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, r.notFoundError("id", fmt.Sprintf("%d", medSessionID.Int32()))
		}
		return nil, r.dbError("select", fmt.Sprintf("failed to find medical history with ID %d", medSessionID.Int32()), err)
	}

	medSession := r.sqlcRowToEntity(sqlcRow)
	return &medSession, nil
}

func (r *SQLCMedSessionRepository) FindByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error) {
	rows, err := r.queries.FindMedicalSessionByEmployeeID(ctx, sqlc.FindMedicalSessionByEmployeeIDParams{
		EmployeeID: employeeID.Int32(),
		Limit:      pagination.Limit(),
		Offset:     pagination.Offset(),
	})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", fmt.Sprintf("failed to find medical history for employee ID %d", employeeID.Int32()), err)
	}

	total, err := r.queries.CountMedicalSessionByEmployeeID(ctx, employeeID.Int32())
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", fmt.Sprintf("failed to count medical history for employee ID %d", employeeID.Int32()), err)
	}

	medicalSessions := r.ToEntities(rows)
	return p.NewPage(medicalSessions, total, pagination), nil
}

func (r *SQLCMedSessionRepository) FindByPetID(ctx context.Context, petID valueobject.PetID, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error) {
	rows, err := r.queries.FindMedicalSessionByPetID(ctx, sqlc.FindMedicalSessionByPetIDParams{
		PetID:  petID.Int32(),
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", fmt.Sprintf("failed to find medical history for pet ID %d", petID.Int32()), err)
	}

	total, err := r.queries.CountMedicalSessionByPetID(ctx, petID.Int32())
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", fmt.Sprintf("failed to count medical history for pet ID %d", petID.Int32()), err)
	}

	medicalSessions := r.ToEntities(rows)
	return p.NewPage(medicalSessions, total, pagination), nil

}

func (r *SQLCMedSessionRepository) FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error) {
	rows, err := r.queries.FindMedicalSessionByCustomerID(ctx, sqlc.FindMedicalSessionByCustomerIDParams{
		CustomerID: customerID.Int32(),
		Limit:      pagination.Limit(),
		Offset:     pagination.Offset(),
	})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", fmt.Sprintf("failed to find medical history for customer ID %d", customerID.Int32()), err)
	}

	total, err := r.queries.CountMedicalSessionByCustomerID(ctx, customerID.Int32())
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", fmt.Sprintf("failed to count medical history for customer ID %d", customerID.Int32()), err)
	}

	medicalSessions := r.ToEntities(rows)
	return p.NewPage(medicalSessions, total, pagination), nil
}

func (r *SQLCMedSessionRepository) FindRecentByPetID(ctx context.Context, petID valueobject.PetID, limit int) ([]med.MedicalSession, error) {
	rows, err := r.queries.FindRecentMedicalSessionByPetID(ctx, sqlc.FindRecentMedicalSessionByPetIDParams{
		PetID: petID.Int32(),
		Limit: int32(limit),
	})
	if err != nil {
		return nil, r.dbError("select", fmt.Sprintf("failed to find recent medical history for pet ID %d", petID.Int32()), err)
	}

	medicalSessions := r.ToEntities(rows)
	return medicalSessions, nil
}

func (r *SQLCMedSessionRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error) {
	rows, err := r.queries.FindMedicalSessionByDateRange(ctx, sqlc.FindMedicalSessionByDateRangeParams{
		VisitDate:   r.pgMap.PgTimestamptz.FromTime(startDate),
		VisitDate_2: r.pgMap.PgTimestamptz.FromTime(endDate),
		Limit:       pagination.Limit(),
		Offset:      pagination.Offset(),
	})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", "failed to find medical history by date range", err)
	}

	total, err := r.queries.CountMedicalSessionByDateRange(ctx, sqlc.CountMedicalSessionByDateRangeParams{
		VisitDate:   r.pgMap.PgTimestamptz.FromTime(startDate),
		VisitDate_2: r.pgMap.PgTimestamptz.FromTime(endDate),
	})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", "failed to count medical history by date range", err)
	}

	medicalSessions := r.ToEntities(rows)
	return p.NewPage(medicalSessions, total, pagination), nil
}

func (r *SQLCMedSessionRepository) FindByPetAndDateRange(ctx context.Context, petID valueobject.PetID, startDate, endDate time.Time) ([]med.MedicalSession, error) {
	rows, err := r.queries.FindMedicalSessionByPetAndDateRange(ctx, sqlc.FindMedicalSessionByPetAndDateRangeParams{
		PetID:       petID.Int32(),
		VisitDate:   r.pgMap.PgTimestamptz.FromTime(startDate),
		VisitDate_2: r.pgMap.PgTimestamptz.FromTime(endDate),
	})
	if err != nil {
		return nil, r.dbError("select", "failed to find medical history by pet and date range", err)
	}

	medicalSessions := r.ToEntities(rows)
	return medicalSessions, nil
}

func (r *SQLCMedSessionRepository) FindByDiagnosis(ctx context.Context, diagnosis string, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error) {
	rows, err := r.queries.FindMedicalSessionByDiagnosis(ctx, sqlc.FindMedicalSessionByDiagnosisParams{
		Column1: r.pgMap.PgText.FromString(diagnosis),
		Limit:   pagination.Limit(),
		Offset:  pagination.Offset(),
	})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", "failed to find medical history by diagnosis", err)
	}

	total, err := r.queries.CountMedicalSessionByDiagnosis(ctx, r.pgMap.PgText.FromString(diagnosis))
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", "failed to count medical history by diagnosis", err)
	}

	medicalSessions := r.ToEntities(rows)
	return p.NewPage(medicalSessions, total, pagination), nil
}

func (r *SQLCMedSessionRepository) ExistsByID(ctx context.Context, medicalSessionID valueobject.MedSessionID) (bool, error) {
	exists, err := r.queries.ExistsMedicalSessionByID(ctx, medicalSessionID.Int32())
	if err != nil {
		return false, r.dbError("select", "failed to check medical history existence by ID", err)
	}
	return exists, nil
}

func (r *SQLCMedSessionRepository) ExistsByPetAndDate(ctx context.Context, petID valueobject.PetID, date time.Time) (bool, error) {
	exists, err := r.queries.ExistsMedicalSessionByPetAndDate(ctx, sqlc.ExistsMedicalSessionByPetAndDateParams{
		PetID: petID.Int32(),
		Date:  r.pgMap.PgDate.FromTime(date),
	})
	if err != nil {
		return false, r.dbError("select", "failed to check medical history existence by pet and date", err)
	}
	return exists, nil
}

func (r *SQLCMedSessionRepository) Save(ctx context.Context, medSession *med.MedicalSession) error {
	if medSession.ID().IsZero() {
		return r.create(ctx, medSession)
	}
	return r.update(ctx, medSession)
}

func (r *SQLCMedSessionRepository) Delete(ctx context.Context, medicalSessionID valueobject.MedSessionID, isHarDelete bool) error {
	if isHarDelete {
		if err := r.queries.HardDeleteMedicalSession(ctx, medicalSessionID.Int32()); err != nil {
			return r.dbError("delete", "failed to hard delete medical history", err)
		}
		return nil
	}

	if err := r.queries.SoftDeleteMedicalSession(ctx, medicalSessionID.Int32()); err != nil {
		return r.dbError("delete", "failed to soft delete medical history", err)
	}

	return nil
}

func (r *SQLCMedSessionRepository) create(ctx context.Context, medSession *med.MedicalSession) error {
	params := r.toCreateParams(*medSession)
	_, err := r.queries.SaveMedicalSession(ctx, params)
	if err != nil {
		return r.dbError("insert", "failed to create medical history", err)
	}

	return nil
}

func (r *SQLCMedSessionRepository) update(ctx context.Context, medSession *med.MedicalSession) error {
	params := r.toUpdateParams(*medSession)
	_, err := r.queries.UpdateMedicalSession(ctx, params)
	if err != nil {
		return r.dbError("update", "failed to update medical history", err)
	}

	return nil
}
