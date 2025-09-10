// Package repositoryimpl contains implementation for database operations related to appointments
package repositoryimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	appt "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	vo "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repo "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCApptRepository struct {
	queries *sqlc.Queries
}

func NewSQLCApptRepository(queries *sqlc.Queries) repo.AppointmentRepository {
	return &SQLCApptRepository{queries: queries}
}

func (r *SQLCApptRepository) GetByIDAndCustomerID(ctx context.Context, id vo.AppointmentID, customerID vo.CustomerID) (appt.Appointment, error) {
	sqlRow, err := r.queries.GetAppointmentByIDAndCustomerID(ctx, sqlc.GetAppointmentByIDAndCustomerIDParams{
		ID:         int32(id.Value()),
		CustomerID: int32(customerID.Value()),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appt.Appointment{}, r.notFoundError("id", id.String())
		}
		return appt.Appointment{}, r.dbError(OpSelect, ErrMsgGetAppt, err)
	}

	appointment, err := sqlRowToDomain(sqlRow)
	if err != nil {
		return appt.Appointment{}, r.wrapConversionError(err)
	}

	return *appointment, nil
}

func (r *SQLCApptRepository) GetByIDAndEmployeeID(ctx context.Context, id vo.AppointmentID, employeeID vo.EmployeeID) (appt.Appointment, error) {
	sqlRow, err := r.queries.GetAppointmentByIDAndEmployeeID(ctx, sqlc.GetAppointmentByIDAndEmployeeIDParams{
		ID:         int32(id.Value()),
		EmployeeID: pgtype.Int4{Int32: int32(employeeID.Value()), Valid: true},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appt.Appointment{}, r.notFoundError("id", id.String())
		}
		return appt.Appointment{}, r.dbError(OpSelect, ErrMsgGetAppt, err)
	}

	appointment, err := sqlRowToDomain(sqlRow)
	if err != nil {
		return appt.Appointment{}, r.wrapConversionError(err)
	}

	return *appointment, nil
}

func (r *SQLCApptRepository) GetByID(ctx context.Context, id vo.AppointmentID) (appt.Appointment, error) {
	sqlRow, err := r.queries.GetAppoinmentByID(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appt.Appointment{}, r.notFoundError("id", id.String())
		}
		return appt.Appointment{}, r.dbError(OpSelect, ErrMsgGetAppt, err)
	}

	appointment, err := sqlRowToDomain(sqlRow)
	if err != nil {
		return appt.Appointment{}, r.wrapConversionError(err)
	}

	return *appointment, nil
}

func (r *SQLCApptRepository) ListByCustomerID(ctx context.Context, customerID vo.CustomerID, paginatio page.PageInput) ([]appt.Appointment, error) {
	sqlRows, err := r.queries.ListAppoinmentsByCustomerID(ctx, sqlc.ListAppoinmentsByCustomerIDParams{
		CustomerID: int32(customerID.Value()),
		Offset:     r.calculateOffset(paginatio),
		Limit:      int32(paginatio.PageSize),
	})
	if err != nil {
		msg := fmt.Sprintf("failed to list appointments for owner ID %d", customerID.Value())
		return nil, r.dbError(OpSelect, msg, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return nil, r.wrapConversionError(err)
	}

	return appointments, nil
}

func (r *SQLCApptRepository) ListAll(ctx context.Context, pageInput page.PageInput) (page.Page[[]appt.Appointment], error) {
	params := sqlc.ListAppoinmentsParams{
		Offset: r.calculateOffset(pageInput),
		Limit:  int32(pageInput.PageSize),
	}

	sqlRows, err := r.queries.ListAppoinments(ctx, params)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.dbError(OpSelect, ErrMsgListAppts, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountAppoinments(ctx)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.dbError(OpCount, ErrMsgCountAppts, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCApptRepository) Search(ctx context.Context, pageInput page.PageInput, criteria map[string]interface{}) (page.Page[[]appt.Appointment], error) {
	params := sqlc.ListAppoinmentsParams{
		Offset: r.calculateOffset(pageInput),
		Limit:  int32(pageInput.PageSize),
	}

	sqlRows, err := r.queries.ListAppoinments(ctx, params)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.dbError(OpSearch, ErrMsgSearchAppts, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountAppoinments(ctx)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.dbError(OpCount, ErrMsgCountAppts, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCApptRepository) ListByVetID(ctx context.Context, employeeID vo.EmployeeID, pageInput page.PageInput) (page.Page[[]appt.Appointment], error) {
	params := sqlc.ListAppoinmentsByEmployeeIDParams{
		EmployeeID: pgtype.Int4{Int32: int32(employeeID.Value()), Valid: true},
		Offset:     r.calculateOffset(pageInput),
		Limit:      int32(pageInput.PageSize),
	}

	sqlRows, err := r.queries.ListAppoinmentsByEmployeeID(ctx, params)
	if err != nil {
		msg := fmt.Sprintf("failed to list appointments for vet ID %d", employeeID.Value())
		return page.Page[[]appt.Appointment]{}, r.dbError(OpSelect, msg, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.wrapConversionError(err)
	}

	countParams := pgtype.Int4{Int32: int32(employeeID.Value()), Valid: true}
	totalCount, err := r.queries.CountAppoinmentsByEmployeeID(ctx, countParams)
	if err != nil {
		msg := fmt.Sprintf("failed to count appointments for vet ID %d", employeeID.Value())
		return page.Page[[]appt.Appointment]{}, r.dbError(OpCount, msg, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCApptRepository) ListByPetID(ctx context.Context, petID vo.PetID, pageInput page.PageInput) (page.Page[[]appt.Appointment], error) {
	params := sqlc.ListAppoinmentsByEmployeeIDParams{
		EmployeeID: pgtype.Int4{Int32: int32(petID.Value()), Valid: true},
		Offset:     r.calculateOffset(pageInput),
		Limit:      int32(pageInput.PageSize),
	}

	sqlRows, err := r.queries.ListAppoinmentsByEmployeeID(ctx, params)
	if err != nil {
		msg := fmt.Sprintf("failed to list appointments for pet ID %d", petID.Value())
		return page.Page[[]appt.Appointment]{}, r.dbError(OpSelect, msg, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.wrapConversionError(err)
	}

	countParams := pgtype.Int4{Int32: int32(petID.Value()), Valid: true}
	totalCount, err := r.queries.CountAppoinmentsByEmployeeID(ctx, countParams)
	if err != nil {
		msg := fmt.Sprintf("failed to count appointments for pet ID %d", petID.Value())
		return page.Page[[]appt.Appointment]{}, r.dbError(OpCount, msg, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

// ListByCustomerID retrieves appointments for a specific owner
func (r *SQLCApptRepository) ListByCustomerID(ctx context.Context, customerID vo.CustomerID, pageInput page.PageInput) (page.Page[[]appt.Appointment], error) {
	params := sqlc.ListAppoinmentsByCustomerIDParams{
		CustomerID: int32(customerID.Value()),
		Offset:     r.calculateOffset(pageInput),
		Limit:      int32(pageInput.PageSize),
	}

	sqlRows, err := r.queries.ListAppoinmentsByCustomerID(ctx, params)
	if err != nil {
		msg := fmt.Sprintf("failed to list appointments for owner ID %d", customerID.Value())
		return page.Page[[]appt.Appointment]{}, r.dbError(OpSelect, msg, err)
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appt.Appointment]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountAppoinmentsByCustomerID(ctx, int32(customerID.Value()))
	if err != nil {
		msg := fmt.Sprintf("failed to count appointments for owner ID %d", customerID.Value())
		return page.Page[[]appt.Appointment]{}, r.dbError(OpCount, msg, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

// ListByDateRange retrieves appointments within a specific date range
func (r *SQLCApptRepository) ListByDateRange(ctx context.Context, start, end time.Time, pageInput page.PageInput) (page.Page[[]appt.Appointment], error) {
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
func (r *SQLCApptRepository) Save(ctx context.Context, appt *appt.Appointment) error {
	if appt.ID().IsZero() {
		return r.create(ctx, appt)
	}
	return r.update(ctx, appt)
}

// Delete removes an appointment by its ID
func (r *SQLCApptRepository) Delete(id vo.AppointmentID) error {
	if err := r.queries.DeleteAppoinment(context.Background(), int32(id.Value())); err != nil {
		msg := fmt.Sprintf("failed to delete appointment with ID %d", id.Value())
		return r.dbError(OpDelete, msg, err)
	}
	return nil
}

// create inserts a new appointment
func (r *SQLCApptRepository) create(ctx context.Context, appt *appt.Appointment) error {
	params := domainToCreateParams(appt)
	_, err := r.queries.CreateAppoinment(ctx, params)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreateAppt, err)
	}
	return nil
}

// update modifies an existing appointment
func (r *SQLCApptRepository) update(ctx context.Context, appt *appt.Appointment) error {
	params := domainToUpdateParams(appt)
	_, err := r.queries.UpdateAppoinment(ctx, params)
	if err != nil {
		return r.dbError(OpUpdate, ErrMsgUpdateAppt, err)
	}
	return nil
}
