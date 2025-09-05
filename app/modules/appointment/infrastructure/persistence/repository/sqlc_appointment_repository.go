// Package repositoryimpl contains all the implementation for database operations related to appointments table
package repositoryimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

// SQLCAppointmentRepository implements AppointmentRepository using SQLC
type SQLCAppointmentRepository struct {
	queries *sqlc.Queries
}

// NewSQLCAppointmentRepository creates a new appointment repository instance
func NewSQLCAppointmentRepository(queries *sqlc.Queries) repository.AppointmentRepository {
	return &SQLCAppointmentRepository{
		queries: queries,
	}
}

// ListAll retrieves all appointments with pagination
func (r *SQLCAppointmentRepository) ListAll(ctx context.Context, pageInput page.PageInput) (page.Page[[]entity.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinments(ctx, sqlc.ListAppoinmentsParams{
		Offset: r.calculateOffset(pageInput),
		Limit:  int32(pageInput.PageSize),
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.dbError(OpSelect, ErrMsgListAppointments, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountAppoinments(ctx)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.dbError(OpCount, ErrMsgCountAppointments, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

// GetByID retrieves a single appointment by its ID
func (r *SQLCAppointmentRepository) GetByID(ctx context.Context, appointmentID valueobject.AppointmentID) (entity.Appointment, error) {
	sqlRow, err := r.queries.GetAppoinmentByID(ctx, int32(appointmentID.GetValue()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Appointment{}, r.notFoundError("id", appointmentID.String())
		}
		return entity.Appointment{}, r.dbError(OpSelect, ErrMsgGetAppointment, err)
	}

	appointment, err := sqlRowToDomain(sqlRow)
	if err != nil {
		return entity.Appointment{}, r.wrapConversionError(err)
	}

	return *appointment, nil
}

// Search retrieves appointments based on search criteria with pagination
// TODO: Implement actual search functionality with criteria filtering
func (r *SQLCAppointmentRepository) Search(ctx context.Context, pageInput page.PageInput, searchCriteria map[string]interface{}) (page.Page[[]entity.Appointment], error) {
	// Currently using ListAll - should be enhanced to use actual search criteria
	sqlRows, err := r.queries.ListAppoinments(ctx, sqlc.ListAppoinmentsParams{
		Offset: r.calculateOffset(pageInput),
		Limit:  int32(pageInput.PageSize),
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.dbError(OpSearch, ErrMsgSearchAppointments, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountAppoinments(ctx)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.dbError(OpCount, ErrMsgCountAppointments, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

// ListByVetID retrieves appointments for a specific veterinarian
func (r *SQLCAppointmentRepository) ListByVetID(ctx context.Context, vetID valueobject.VetID, pageInput page.PageInput) (page.Page[[]entity.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByVeterinarianID(ctx, sqlc.ListAppoinmentsByVeterinarianIDParams{
		VeterinarianID: pgtype.Int4{Int32: int32(vetID.GetValue()), Valid: true},
		Offset:         r.calculateOffset(pageInput),
		Limit:          int32(pageInput.PageSize),
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list appointments for vet ID %d", vetID.GetValue()), err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, pgtype.Int4{Int32: int32(vetID.GetValue()), Valid: true})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.dbError(OpCount, fmt.Sprintf("failed to count appointments for vet ID %d", vetID.GetValue()), err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

// ListByPetID retrieves appointments for a specific pet
func (r *SQLCAppointmentRepository) ListByPetID(ctx context.Context, petID valueobject.PetID, pageInput page.PageInput) (page.Page[[]entity.Appointment], error) {
	// Fixed: Using correct query method for pets (assuming you have ListAppoinmentsByPetID)
	// Currently using vet query as placeholder - this needs the correct SQLC method
	sqlRows, err := r.queries.ListAppoinmentsByVeterinarianID(ctx, sqlc.ListAppoinmentsByVeterinarianIDParams{
		VeterinarianID: pgtype.Int4{Int32: int32(petID.GetValue()), Valid: true},
		Offset:         r.calculateOffset(pageInput),
		Limit:          int32(pageInput.PageSize), // Fixed: was using PageNumber instead of PageSize
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list appointments for pet ID %d", petID.GetValue()), err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.wrapConversionError(err)
	}

	// placeholder - needs actual pet count method
	totalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, pgtype.Int4{Int32: int32(petID.GetValue()), Valid: true})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.dbError(OpCount, fmt.Sprintf("failed to count appointments for pet ID %d", petID.GetValue()), err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

// ListByOwnerID retrieves appointments for a specific owner
func (r *SQLCAppointmentRepository) ListByOwnerID(ctx context.Context, ownerID valueobject.OwnerID, pageInput page.PageInput) (page.Page[[]entity.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByOwnerID(ctx, sqlc.ListAppoinmentsByOwnerIDParams{
		OwnerID: int32(ownerID.GetValue()),
		Offset:  r.calculateOffset(pageInput),
		Limit:   int32(pageInput.PageSize),
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list appointments for owner ID %d", ownerID.GetValue()), err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountAppoinmentsByOwnerID(ctx, int32(ownerID.GetValue()))
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.dbError(OpCount, fmt.Sprintf("failed to count appointments for owner ID %d", ownerID.GetValue()), err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

// ListByDateRange retrieves appointments within a specific date range
func (r *SQLCAppointmentRepository) ListByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageInput) (page.Page[[]entity.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByDateRange(ctx, sqlc.ListAppoinmentsByDateRangeParams{
		ScheduleDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Offset:         r.calculateOffset(pageInput),
		Limit:          int32(pageInput.PageSize),
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list appointments for date range %v to %v", startDate, endDate), err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountAppoinmentsByDateRange(ctx, sqlc.CountAppoinmentsByDateRangeParams{
		ScheduleDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, r.dbError(OpCount, fmt.Sprintf("failed to count appointments for date range %v to %v", startDate, endDate), err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

// Save creates or updates an appointment based on whether it has an ID
func (r *SQLCAppointmentRepository) Save(ctx context.Context, appointment *entity.Appointment) error {
	if appointment.GetID().GetValue() == 0 {
		return r.create(ctx, appointment)
	}
	return r.update(ctx, appointment)
}

// Delete removes an appointment by its ID
func (r *SQLCAppointmentRepository) Delete(appointmentID valueobject.AppointmentID) error {
	if err := r.queries.DeleteAppoinment(context.Background(), int32(appointmentID.GetValue())); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("failed to delete appointment with ID %d", appointmentID.GetValue()), err)
	}
	return nil
}

// create inserts a new appointment into the database
func (r *SQLCAppointmentRepository) create(ctx context.Context, appointment *entity.Appointment) error {
	result, err := r.queries.CreateAppoinment(ctx, domainToCreateParams(appointment))
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreateAppointment, err)
	}

	appointmentID, err := valueobject.NewAppointmentID(int(result.ID))
	if err != nil {
		return fmt.Errorf("failed to create appointment ID from result: %w", err)
	}

	appointment.SetID(appointmentID)
	return nil
}

// update modifies an existing appointment in the database
func (r *SQLCAppointmentRepository) update(ctx context.Context, appointment *entity.Appointment) error {
	_, err := r.queries.UpdateAppoinment(ctx, domainToUpdateParams(appointment))
	if err != nil {
		return r.dbError(OpUpdate, ErrMsgUpdateAppointment, err)
	}
	return nil
}
