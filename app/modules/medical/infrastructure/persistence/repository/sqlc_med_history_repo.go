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

	p "clinic-vet-api/app/shared/page"

	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCMedSessionRepository struct {
	queries *sqlc.Queries
}

func NewSQLCMedSessionRepository(queries *sqlc.Queries) repository.MedicalSessionRepository {
	return &SQLCMedSessionRepository{
		queries: queries,
	}
}

func (r *SQLCMedSessionRepository) FindBySpecification(ctx context.Context, spec specification.MedicalSessionSpecification) (p.Page[med.MedicalSession], error) {
	return p.Page[med.MedicalSession]{}, errors.New("FindBySpecification not implemented")
}

func (r *SQLCMedSessionRepository) FindByIDAndCustomerID(ctx context.Context, medicalSessionID valueobject.MedSessionID, customerID valueobject.CustomerID) (*med.MedicalSession, error) {
	sqlcRow, err := r.queries.FindMedicalSessionByIDAndCustomerID(ctx, sqlc.FindMedicalSessionByIDAndCustomerIDParams{
		ID:         int32(medicalSessionID.Value()),
		CustomerID: int32(customerID.Value()),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, r.notFoundError("id", fmt.Sprintf("%d", medicalSessionID.Value()))
		}
		return nil, r.dbError("select", fmt.Sprintf("failed to find medical history with ID %d for customer ID %d", medicalSessionID.Value(), customerID.Value()), err)
	}

	medSession := ToEntity(sqlcRow)
	return &medSession, nil
}

func (r *SQLCMedSessionRepository) FindByIDAndEmployeeID(ctx context.Context, medicalSessionID valueobject.MedSessionID, employeeID valueobject.EmployeeID) (*med.MedicalSession, error) {
	sqlcRow, err := r.queries.FindMedicalSessionByIDAndEmployeeID(ctx, sqlc.FindMedicalSessionByIDAndEmployeeIDParams{
		ID:         int32(medicalSessionID.Value()),
		EmployeeID: int32(employeeID.Value()),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, r.notFoundError("id", fmt.Sprintf("%d", medicalSessionID.Value()))
		}
		return nil, r.dbError("select", fmt.Sprintf("failed to find medical history with ID %d for employee ID %d", medicalSessionID.Value(), employeeID.Value()), err)
	}

	medSession := ToEntity(sqlcRow)
	return &medSession, nil
}

func (r *SQLCMedSessionRepository) FindByIDAndPetID(ctx context.Context, medicalSessionID valueobject.MedSessionID, petID valueobject.PetID) (*med.MedicalSession, error) {
	sqlcRow, err := r.queries.FindMedicalSessionByIDAndPetID(ctx, sqlc.FindMedicalSessionByIDAndPetIDParams{
		ID:    int32(medicalSessionID.Value()),
		PetID: int32(petID.Value()),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, r.notFoundError("id", fmt.Sprintf("%d", medicalSessionID.Value()))
		}
		return nil, r.dbError("select", fmt.Sprintf("failed to find medical history with ID %d for pet ID %d", medicalSessionID.Value(), petID.Value()), err)
	}

	medSession := ToEntity(sqlcRow)
	return &medSession, nil
}

func (r *SQLCMedSessionRepository) FindByID(ctx context.Context, medSessionID valueobject.MedSessionID) (*med.MedicalSession, error) {
	sqlcMedSession, err := r.queries.FindMedicalSessionByID(ctx, int32(medSessionID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, r.notFoundError("id", fmt.Sprintf("%d", medSessionID.Value()))
		}
		return nil, r.dbError("select", fmt.Sprintf("failed to find medical history with ID %d", medSessionID.Value()), err)
	}

	medSession := ToEntity(sqlcMedSession)
	return &medSession, nil
}

func (r *SQLCMedSessionRepository) FindByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error) {
	medSessionRows, err := r.queries.FindMedicalSessionByEmployeeID(ctx, sqlc.FindMedicalSessionByEmployeeIDParams{
		EmployeeID: int32(employeeID.Value()),
		Limit:      int32(pagination.Limit()),
		Offset:     int32(pagination.Offset()),
	})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", fmt.Sprintf("failed to find medical history for employee ID %d", employeeID.Value()), err)
	}

	total, err := r.queries.CountMedicalSessionByEmployeeID(ctx, int32(employeeID.Value()))
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", fmt.Sprintf("failed to count medical history for employee ID %d", employeeID.Value()), err)
	}

	medicalSessions, err := ToEntities(medSessionRows)
	return p.NewPage(medicalSessions, int(total), pagination), nil
}

func (r *SQLCMedSessionRepository) FindByPetID(ctx context.Context, petID valueobject.PetID, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error) {
	medSessionRows, err := r.queries.FindMedicalSessionByPetID(ctx, sqlc.FindMedicalSessionByPetIDParams{
		PetID:  int32(petID.Value()),
		Limit:  int32(pagination.Limit()),
		Offset: int32(pagination.Offset()),
	})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", fmt.Sprintf("failed to find medical history for pet ID %d", petID.Value()), err)
	}

	total, err := r.queries.CountMedicalSessionByPetID(ctx, int32(petID.Value()))
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", fmt.Sprintf("failed to count medical history for pet ID %d", petID.Value()), err)
	}

	medicalSessions, err := ToEntities(medSessionRows)
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.wrapConversionError(err)
	}

	return p.NewPage(medicalSessions, int(total), pagination), nil

}

func (r *SQLCMedSessionRepository) FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error) {
	medSessionRows, err := r.queries.FindMedicalSessionByCustomerID(ctx, sqlc.FindMedicalSessionByCustomerIDParams{
		CustomerID: int32(customerID.Value()),
		Limit:      int32(pagination.Limit()),
		Offset:     int32(pagination.Offset()),
	})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", fmt.Sprintf("failed to find medical history for customer ID %d", customerID.Value()), err)
	}

	total, err := r.queries.CountMedicalSessionByCustomerID(ctx, int32(customerID.Value()))
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", fmt.Sprintf("failed to count medical history for customer ID %d", customerID.Value()), err)
	}

	medicalSessions, err := ToEntities(medSessionRows)
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.wrapConversionError(err)
	}

	return p.NewPage(medicalSessions, int(total), pagination), nil
}

func (r *SQLCMedSessionRepository) FindRecentByPetID(ctx context.Context, petID valueobject.PetID, limit int) ([]med.MedicalSession, error) {
	medSessionRows, err := r.queries.FindRecentMedicalSessionByPetID(ctx, sqlc.FindRecentMedicalSessionByPetIDParams{
		PetID: int32(petID.Value()),
		Limit: int32(limit),
	})
	if err != nil {
		return nil, r.dbError("select", fmt.Sprintf("failed to find recent medical history for pet ID %d", petID.Value()), err)
	}

	medicalSessions, err := ToEntities(medSessionRows)
	if err != nil {
		return nil, r.wrapConversionError(err)
	}

	return medicalSessions, nil
}

func (r *SQLCMedSessionRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error) {
	medSessionRows, err := r.queries.FindMedicalSessionByDateRange(ctx, sqlc.FindMedicalSessionByDateRangeParams{
		VisitDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		VisitDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Limit:       int32(pagination.Limit()),
		Offset:      int32(pagination.Offset()),
	})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", "failed to find medical history by date range", err)
	}

	total, err := r.queries.CountMedicalSessionByDateRange(ctx, sqlc.CountMedicalSessionByDateRangeParams{
		VisitDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		VisitDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", "failed to count medical history by date range", err)
	}

	medicalSessions, err := ToEntities(medSessionRows)
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.wrapConversionError(err)
	}

	return p.NewPage(medicalSessions, int(total), pagination), nil
}

func (r *SQLCMedSessionRepository) FindByPetAndDateRange(ctx context.Context, petID valueobject.PetID, startDate, endDate time.Time) ([]med.MedicalSession, error) {
	medSessionRows, err := r.queries.FindMedicalSessionByPetAndDateRange(ctx, sqlc.FindMedicalSessionByPetAndDateRangeParams{
		PetID:       int32(petID.Value()),
		VisitDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		VisitDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return nil, r.dbError("select", fmt.Sprintf("failed to find medical history for pet ID %d and date range", petID.Value()), err)
	}

	medicalSessions, err := ToEntities(medSessionRows)
	if err != nil {
		return nil, r.wrapConversionError(err)
	}

	return medicalSessions, nil
}

func (r *SQLCMedSessionRepository) FindByDiagnosis(ctx context.Context, diagnosis string, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error) {
	medSessionRows, err := r.queries.FindMedicalSessionByDiagnosis(ctx, sqlc.FindMedicalSessionByDiagnosisParams{
		Column1: pgtype.Text{String: diagnosis, Valid: true},
		Limit:   int32(pagination.Limit()),
		Offset:  int32(pagination.Offset()),
	})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", "failed to find medical history by diagnosis", err)
	}

	total, err := r.queries.CountMedicalSessionByDiagnosis(ctx, pgtype.Text{String: diagnosis, Valid: true})
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.dbError("select", "failed to count medical history by diagnosis", err)
	}

	medicalSessions, err := ToEntities(medSessionRows)
	if err != nil {
		return p.Page[med.MedicalSession]{}, r.wrapConversionError(err)
	}

	return p.NewPage(medicalSessions, int(total), pagination), nil
}

func (r *SQLCMedSessionRepository) ExistsByID(ctx context.Context, medicalSessionID valueobject.MedSessionID) (bool, error) {
	exists, err := r.queries.ExistsMedicalSessionByID(ctx, int32(medicalSessionID.Value()))
	if err != nil {
		return false, r.dbError("select", "failed to check medical history existence by ID", err)
	}
	return exists, nil
}

func (r *SQLCMedSessionRepository) ExistsByPetAndDate(ctx context.Context, petID valueobject.PetID, date time.Time) (bool, error) {
	exists, err := r.queries.ExistsMedicalSessionByPetAndDate(ctx, sqlc.ExistsMedicalSessionByPetAndDateParams{
		PetID: int32(petID.Value()),
		Date:  pgtype.Timestamptz{Time: date, Valid: true},
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

func (r *SQLCMedSessionRepository) SoftDelete(ctx context.Context, medicalSessionID valueobject.MedSessionID) error {
	if err := r.queries.SoftDeleteMedicalSession(ctx, int32(medicalSessionID.Value())); err != nil {
		return r.dbError("delete", "failed to soft delete medical history", err)
	}
	return nil
}

func (r *SQLCMedSessionRepository) HardDelete(ctx context.Context, medicalSessionID valueobject.MedSessionID) error {
	if err := r.queries.HardDeleteMedicalSession(ctx, int32(medicalSessionID.Value())); err != nil {
		return r.dbError("delete", "failed to hard delete medical history", err)
	}
	return nil
}

func (r *SQLCMedSessionRepository) CountByPetID(ctx context.Context, petID valueobject.PetID) (int64, error) {
	count, err := r.queries.CountMedicalSessionByPetID(ctx, int32(petID.Value()))
	if err != nil {
		return 0, r.dbError("select", "failed to count medical history by pet ID", err)
	}
	return count, nil
}

func (r *SQLCMedSessionRepository) CountByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID) (int64, error) {
	count, err := r.queries.CountMedicalSessionByEmployeeID(ctx, int32(employeeID.Value()))
	if err != nil {
		return 0, r.dbError("select", "failed to count medical history by employee ID", err)
	}
	return count, nil
}

func (r *SQLCMedSessionRepository) CountByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (int64, error) {
	count, err := r.queries.CountMedicalSessionByCustomerID(ctx, int32(customerID.Value()))
	if err != nil {
		return 0, r.dbError("select", "failed to count medical history by customer ID", err)
	}
	return count, nil
}

func (r *SQLCMedSessionRepository) CountByDateRange(ctx context.Context, startDate, endDate time.Time) (int64, error) {
	count, err := r.queries.CountMedicalSessionByDateRange(ctx, sqlc.CountMedicalSessionByDateRangeParams{
		VisitDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		VisitDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return 0, r.dbError("select", "failed to count medical history by date range", err)
	}
	return count, nil
}

func (r *SQLCMedSessionRepository) create(ctx context.Context, medSession *med.MedicalSession) error {
	params := ToCreateParams(*medSession)
	_, err := r.queries.SaveMedicalSession(ctx, params)
	if err != nil {
		return r.dbError("insert", "failed to create medical history", err)
	}

	return nil
}

func (r *SQLCMedSessionRepository) update(ctx context.Context, medSession *med.MedicalSession) error {
	params := ToUpdateParams(*medSession)

	_, err := r.queries.UpdateMedicalSession(ctx, params)
	if err != nil {
		return r.dbError("update", fmt.Sprintf("failed to update medical history with ID %d", medSession.ID().Value()), err)
	}

	return nil
}
