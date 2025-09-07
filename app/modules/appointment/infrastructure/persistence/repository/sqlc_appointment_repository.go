// Package repositoryimpl contains implementation for database operations related to appointments
package repositoryimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	appt "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	vo "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repo "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCAppointmentRepository struct {
	queries *sqlc.Queries
}

func NewSQLCAppointmentRepository(queries *sqlc.Queries) repo.AppointmentRepository {
	return &SQLCAppointmentRepository{queries: queries}
}

func (r *SQLCAppointmentRepository) GetByIDAndOwnerID(ctx context.Context, id vo.AppointmentID, ownerID valueobject.OwnerID) (appt.Appointment, error) {
	sqlRow, err := r.queries.GetAppoinmentByIDAndOwnerID(ctx, sqlc.GetAppoinmentByIDAndOwnerIDParams{
		ID:      int32(id.Value()),
		OwnerID: int32(ownerID.Value()),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appt.Appointment{}, r.notFoundError("id", id.String())
		}
		return appt.Appointment{}, r.dbError(OpSelect, ErrMsgGetAppointment, err)
	}

	appointment, err := sqlRowToDomain(sqlRow)
	if err != nil {
		return appt.Appointment{}, r.wrapConversionError(err)
	}

	return *appointment, nil
}

func (r *SQLCAppointmentRepository) GetByIDAndVetID(ctx context.Context, id vo.AppointmentID, vetID valueobject.VetID) (appt.Appointment, error) {
	sqlRow, err := r.queries.GetAppoinmentByIDAndVetID(ctx, sqlc.GetAppoinmentByIDAndVetIDParams{
		ID:             int32(id.Value()),
		VeterinarianID: pgtype.Int4{Int32: int32(vetID.Value()), Valid: true},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appt.Appointment{}, r.notFoundError("id", id.String())
		}
		return appt.Appointment{}, r.dbError(OpSelect, ErrMsgGetAppointment, err)
	}

	appointment, err := sqlRowToDomain(sqlRow)
	if err != nil {
		return appt.Appointment{}, r.wrapConversionError(err)
	}

	return *appointment, nil
}

func (r *SQLCAppointmentRepository) GetByID(ctx context.Context, id vo.AppointmentID) (appt.Appointment, error) {
	sqlRow, err := r.queries.GetAppoinmentByID(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appt.Appointment{}, r.notFoundError("id", id.String())
		}
		return appt.Appointment{}, r.dbError(OpSelect, ErrMsgGetAppointment, err)
	}

	appointment, err := sqlRowToDomain(sqlRow)
	if err != nil {
		return appt.Appointment{}, r.wrapConversionError(err)
	}

	return *appointment, nil
}

func (r *SQLCAppointmentRepository) ListAll(ctx context.Context, pageInput page.PageInput) (page.Page[[]appt.Appointment], error) {
	params := sqlc.ListAppoinmentsParams{
		Offset: r.calculateOffset(pageInput),
		Limit:  int32(pageInput.PageSize),
	}

	sqlRows, err := r.queries.ListAppoinments(ctx, params)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.dbError(OpSelect, ErrMsgListAppointments, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountAppoinments(ctx)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.dbError(OpCount, ErrMsgCountAppointments, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) Search(ctx context.Context, pageInput page.PageInput, criteria map[string]interface{}) (page.Page[[]appt.Appointment], error) {
	params := sqlc.ListAppoinmentsParams{
		Offset: r.calculateOffset(pageInput),
		Limit:  int32(pageInput.PageSize),
	}

	sqlRows, err := r.queries.ListAppoinments(ctx, params)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.dbError(OpSearch, ErrMsgSearchAppointments, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountAppoinments(ctx)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.dbError(OpCount, ErrMsgCountAppointments, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) ListByVetID(ctx context.Context, vetID vo.VetID, pageInput page.PageInput) (page.Page[[]appt.Appointment], error) {
	params := sqlc.ListAppoinmentsByVeterinarianIDParams{
		VeterinarianID: pgtype.Int4{Int32: int32(vetID.Value()), Valid: true},
		Offset:         r.calculateOffset(pageInput),
		Limit:          int32(pageInput.PageSize),
	}

	sqlRows, err := r.queries.ListAppoinmentsByVeterinarianID(ctx, params)
	if err != nil {
		msg := fmt.Sprintf("failed to list appointments for vet ID %d", vetID.Value())
		return page.Page[[]appt.Appointment]{}, r.dbError(OpSelect, msg, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.wrapConversionError(err)
	}

	countParams := pgtype.Int4{Int32: int32(vetID.Value()), Valid: true}
	totalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, countParams)
	if err != nil {
		msg := fmt.Sprintf("failed to count appointments for vet ID %d", vetID.Value())
		return page.Page[[]appt.Appointment]{}, r.dbError(OpCount, msg, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) ListByPetID(ctx context.Context, petID vo.PetID, pageInput page.PageInput) (page.Page[[]appt.Appointment], error) {
	params := sqlc.ListAppoinmentsByVeterinarianIDParams{
		VeterinarianID: pgtype.Int4{Int32: int32(petID.Value()), Valid: true},
		Offset:         r.calculateOffset(pageInput),
		Limit:          int32(pageInput.PageSize),
	}

	sqlRows, err := r.queries.ListAppoinmentsByVeterinarianID(ctx, params)
	if err != nil {
		msg := fmt.Sprintf("failed to list appointments for pet ID %d", petID.Value())
		return page.Page[[]appt.Appointment]{}, r.dbError(OpSelect, msg, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.wrapConversionError(err)
	}

	countParams := pgtype.Int4{Int32: int32(petID.Value()), Valid: true}
	totalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, countParams)
	if err != nil {
		msg := fmt.Sprintf("failed to count appointments for pet ID %d", petID.Value())
		return page.Page[[]appt.Appointment]{}, r.dbError(OpCount, msg, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

// ListByOwnerID retrieves appointments for a specific owner
func (r *SQLCAppointmentRepository) ListByOwnerID(ctx context.Context, ownerID vo.OwnerID, pageInput page.PageInput) (page.Page[[]appt.Appointment], error) {
	params := sqlc.ListAppoinmentsByOwnerIDParams{
		OwnerID: int32(ownerID.Value()),
		Offset:  r.calculateOffset(pageInput),
		Limit:   int32(pageInput.PageSize),
	}

	sqlRows, err := r.queries.ListAppoinmentsByOwnerID(ctx, params)
	if err != nil {
		msg := fmt.Sprintf("failed to list appointments for owner ID %d", ownerID.Value())
		return page.Page[[]appt.Appointment]{}, r.dbError(OpSelect, msg, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountAppoinmentsByOwnerID(ctx, int32(ownerID.Value()))
	if err != nil {
		msg := fmt.Sprintf("failed to count appointments for owner ID %d", ownerID.Value())
		return page.Page[[]appt.Appointment]{}, r.dbError(OpCount, msg, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

// ListByDateRange retrieves appointments within a specific date range
func (r *SQLCAppointmentRepository) ListByDateRange(ctx context.Context, start, end time.Time, pageInput page.PageInput) (page.Page[[]appt.Appointment], error) {
	params := sqlc.ListAppoinmentsByDateRangeParams{
		ScheduleDate:   pgtype.Timestamptz{Time: start, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: end, Valid: true},
		Offset:         r.calculateOffset(pageInput),
		Limit:          int32(pageInput.PageSize),
	}

	sqlRows, err := r.queries.ListAppoinmentsByDateRange(ctx, params)
	if err != nil {
		msg := fmt.Sprintf("failed to list appointments from %s to %s", start.Format(time.DateOnly), end.Format(time.DateOnly))
		return page.Page[[]appt.Appointment]{}, r.dbError(OpSelect, msg, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.wrapConversionError(err)
	}

	countParams := sqlc.CountAppoinmentsByDateRangeParams{
		ScheduleDate:   pgtype.Timestamptz{Time: start, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: end, Valid: true},
	}

	totalCount, err := r.queries.CountAppoinmentsByDateRange(ctx, countParams)
	if err != nil {
		msg := fmt.Sprintf("failed to count appointments from %s to %s", start.Format(time.DateOnly), end.Format(time.DateOnly))
		return page.Page[[]appt.Appointment]{}, r.dbError(OpCount, msg, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

// Save creates or updates an appointment
func (r *SQLCAppointmentRepository) Save(ctx context.Context, appt *appt.Appointment) error {
	if appt.ID().IsZero() {
		return r.create(ctx, appt)
	}
	return r.update(ctx, appt)
}

// Delete removes an appointment by its ID
func (r *SQLCAppointmentRepository) Delete(id vo.AppointmentID) error {
	if err := r.queries.DeleteAppoinment(context.Background(), int32(id.Value())); err != nil {
		msg := fmt.Sprintf("failed to delete appointment with ID %d", id.Value())
		return r.dbError(OpDelete, msg, err)
	}
	return nil
}

// create inserts a new appointment
func (r *SQLCAppointmentRepository) create(ctx context.Context, appt *appt.Appointment) error {
	params := domainToCreateParams(appt)
	_, err := r.queries.CreateAppoinment(ctx, params)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreateAppointment, err)
	}
	return nil
}

// update modifies an existing appointment
func (r *SQLCAppointmentRepository) update(ctx context.Context, appt *appt.Appointment) error {
	params := domainToUpdateParams(appt)
	_, err := r.queries.UpdateAppoinment(ctx, params)
	if err != nil {
		return r.dbError(OpUpdate, ErrMsgUpdateAppointment, err)
	}
	return nil
}
