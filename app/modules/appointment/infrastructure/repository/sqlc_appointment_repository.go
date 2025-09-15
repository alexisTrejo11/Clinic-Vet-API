// Package repository implements the AppointmentRepository interface using SQLC for database operations.
package repository

import (
	appt "clinic-vet-api/app/core/domain/entity/appointment"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/page"
	p "clinic-vet-api/app/shared/page"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcAppointmentRepository struct {
	queries *sqlc.Queries
}

func NewSqlcAppointmentRepository(queries *sqlc.Queries) repository.AppointmentRepository {
	return &SqlcAppointmentRepository{
		queries: queries,
	}
}

func (r *SqlcAppointmentRepository) FindBySpecification(ctx context.Context, spec specification.ApptSearchSpecification) (p.Page[appt.Appointment], error) {
	return p.Page[appt.Appointment]{}, errors.New("not implemented")
}

// FindByID finds an appointment by ID
func (r *SqlcAppointmentRepository) FindByID(ctx context.Context, id valueobject.AppointmentID) (appt.Appointment, error) {
	sqlRow, err := r.queries.FindAppointmentByID(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appt.Appointment{}, r.notFoundError("id", id.String())
		}
		return appt.Appointment{}, r.dbError("select", "failed to get appointment by ID", err)
	}

	appointmentEntity, err := sqlRowToAppointment(sqlRow)
	if err != nil {
		return appt.Appointment{}, r.wrapConversionError(err)
	}

	return *appointmentEntity, nil
}

// FindByIDAndCustomerID finds an appointment by ID and customer ID
func (r *SqlcAppointmentRepository) FindByIDAndCustomerID(ctx context.Context, id valueobject.AppointmentID, customerID valueobject.CustomerID) (appt.Appointment, error) {
	sqlRow, err := r.queries.FindAppointmentByIDAndCustomerID(ctx, sqlc.FindAppointmentByIDAndCustomerIDParams{
		ID:         int32(id.Value()),
		CustomerID: int32(customerID.Value()),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appt.Appointment{}, r.notFoundError("id and customer ID", fmt.Sprintf("%s, %s", id.String(), customerID.String()))
		}
		return appt.Appointment{}, r.dbError("select", "failed to get appointment by ID and customer ID", err)
	}

	appointment, err := sqlRowToAppointment(sqlRow)
	if err != nil {
		return appt.Appointment{}, r.wrapConversionError(err)
	}

	return *appointment, nil
}

// FindByIDAndEmployeeID finds an appointment by ID and employee ID
func (r *SqlcAppointmentRepository) FindByIDAndEmployeeID(ctx context.Context, id valueobject.AppointmentID, employeeID valueobject.EmployeeID) (appt.Appointment, error) {
	sqlRow, err := r.queries.FindAppointmentByIDAndEmployeeID(ctx, sqlc.FindAppointmentByIDAndEmployeeIDParams{
		ID:         int32(id.Value()),
		EmployeeID: pgtype.Int4{Int32: int32(employeeID.Value()), Valid: true},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appt.Appointment{}, r.notFoundError("id and employee ID", fmt.Sprintf("%s, %s", id.String(), employeeID.String()))
		}
		return appt.Appointment{}, r.dbError("select", "failed to get appointment by ID and employee ID", err)
	}

	appointment, err := sqlRowToAppointment(sqlRow)
	if err != nil {
		return appt.Appointment{}, r.wrapConversionError(err)
	}

	return *appointment, nil
}

// FindAll finds all appointments with pagination
func (r *SqlcAppointmentRepository) FindAll(ctx context.Context, pageInput p.PageInput) (page.Page[appt.Appointment], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize
	limit := pageInput.PageSize

	appointmentRows, err := r.queries.FindAppointments(ctx, sqlc.FindAppointmentsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return p.Page[appt.Appointment]{}, r.dbError("select", "failed to list appointments", err)
	}

	total, err := r.queries.CountAppointments(ctx)
	if err != nil {
		return p.Page[appt.Appointment]{}, r.dbError("select", "failed to count appointments", err)
	}

	appointments, err := sqlRowsToAppointments(appointmentRows)
	if err != nil {
		return p.Page[appt.Appointment]{}, r.wrapConversionError(err)
	}

	return page.NewPage(appointments, *page.GetPageMetadata(int(total), pageInput)), nil
}

// FindByEmployeeID finds appointments by employee ID with pagination
func (r *SqlcAppointmentRepository) FindByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID, pageInput p.PageInput) (page.Page[appt.Appointment], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize
	limit := pageInput.PageSize

	// Find appointments
	appointmentRows, err := r.queries.FindAppointmentsByEmployeeID(ctx, sqlc.FindAppointmentsByEmployeeIDParams{
		EmployeeID: pgtype.Int4{Int32: int32(employeeID.Value()), Valid: true},
		Limit:      int32(limit),
		Offset:     int32(offset),
	})
	if err != nil {
		return p.Page[appt.Appointment]{}, r.dbError("select", "failed to list appointments by employee ID", err)
	}

	// Find total count
	total, err := r.queries.CountAppointmentsByEmployeeID(ctx, pgtype.Int4{Int32: int32(employeeID.Value()), Valid: true})
	if err != nil {
		return p.Page[appt.Appointment]{}, r.dbError("select", "failed to count appointments by employee ID", err)
	}

	appointments, err := sqlRowsToAppointments(appointmentRows)
	if err != nil {
		return p.Page[appt.Appointment]{}, r.wrapConversionError(err)
	}

	return page.NewPage(appointments, *page.GetPageMetadata(int(total), pageInput)), nil
}

// FindByPetID finds appointments by pet ID with pagination
func (r *SqlcAppointmentRepository) FindByPetID(ctx context.Context, petID valueobject.PetID, pageInput p.PageInput) (page.Page[appt.Appointment], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize
	limit := pageInput.PageSize

	// Find appointments
	appointmentRows, err := r.queries.FindAppointmentsByPetID(ctx, sqlc.FindAppointmentsByPetIDParams{
		PetID:  int32(petID.Value()),
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return p.Page[appt.Appointment]{}, r.dbError("select", "failed to list appointments by pet ID", err)
	}

	// Find total count
	total, err := r.queries.CountAppointmentsByPetID(ctx, int32(petID.Value()))
	if err != nil {
		return p.Page[appt.Appointment]{}, r.dbError("select", "failed to count appointments by pet ID", err)
	}

	appointments, err := sqlRowsToAppointments(appointmentRows)
	if err != nil {
		return p.Page[appt.Appointment]{}, r.wrapConversionError(err)
	}

	return page.NewPage(appointments, *page.GetPageMetadata(int(total), pageInput)), nil
}

// FindByCustomerID finds appointments by customer ID with pagination
func (r *SqlcAppointmentRepository) FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID, pageInput p.PageInput) (page.Page[appt.Appointment], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize
	limit := pageInput.PageSize

	// find appointments
	appointmentrows, err := r.queries.FindAppointmentsByCustomerID(ctx, sqlc.FindAppointmentsByCustomerIDParams{
		CustomerID: int32(customerID.Value()),
		Limit:      int32(limit),
		Offset:     int32(offset),
	})
	if err != nil {
		return p.Page[appt.Appointment]{}, r.dbError("select", "failed to list appointments by customer id", err)
	}

	// find total count
	total, err := r.queries.CountAppointmentsByCustomerID(ctx, int32(customerID.Value()))
	if err != nil {
		return p.Page[appt.Appointment]{}, r.dbError("select", "failed to count appointments by customer id", err)
	}

	appointments, err := sqlRowsToAppointments(appointmentrows)
	if err != nil {
		return p.Page[appt.Appointment]{}, r.wrapConversionError(err)
	}

	return page.NewPage(appointments, *page.GetPageMetadata(int(total), pageInput)), nil
}

// findbydaterange finds appointments within a date range with pagination
func (r *SqlcAppointmentRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput p.PageInput) (page.Page[appt.Appointment], error) {
	offset := (pageInput.Page - 1) * pageInput.PageSize
	limit := pageInput.PageSize

	appointmentRows, err := r.queries.FindAppointmentsByDateRange(ctx, sqlc.FindAppointmentsByDateRangeParams{
		ScheduleDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Limit:          int32(limit),
		Offset:         int32(offset),
	})
	if err != nil {
		return p.Page[appt.Appointment]{}, r.dbError("select", "failed to list appointments by date range", err)
	}

	total, err := r.queries.CountAppointmentsByDateRange(ctx, sqlc.CountAppointmentsByDateRangeParams{
		ScheduleDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return p.Page[appt.Appointment]{}, r.dbError("select", "failed to count appointments by date range", err)
	}

	appointments, err := sqlRowsToAppointments(appointmentRows)
	if err != nil {
		return p.Page[appt.Appointment]{}, r.wrapConversionError(err)
	}

	paginationMetadata := page.GetPageMetadata(int(total), pageInput)
	return page.NewPage(appointments, *paginationMetadata), nil
}

// ExistsByID checks if an appointment exists by ID
func (r *SqlcAppointmentRepository) ExistsByID(ctx context.Context, id valueobject.AppointmentID) (bool, error) {
	exists, err := r.queries.ExistsAppointmentID(ctx, int32(id.Value()))
	if err != nil {
		return false, r.dbError("select", "failed to check appointment existence by ID", err)
	}
	return exists, nil
}

// ExistsConflictingAppointment checks for conflicting appointments
func (r *SqlcAppointmentRepository) ExistsConflictingAppointment(ctx context.Context, employeeID valueobject.EmployeeID, startTime, endTime time.Time) (bool, error) {
	exists, err := r.queries.ExistsConflictingAppointment(ctx, sqlc.ExistsConflictingAppointmentParams{
		EmployeeID:     pgtype.Int4{Int32: int32(employeeID.Value()), Valid: true},
		ScheduleDate:   pgtype.Timestamptz{Time: startTime, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: endTime, Valid: true},
	})
	if err != nil {
		return false, r.dbError("select", "failed to check for conflicting appointments", err)
	}
	return exists, nil
}

// Save creates or updates an appointment
func (r *SqlcAppointmentRepository) Save(ctx context.Context, appointment *appt.Appointment) error {
	if appointment.ID().IsZero() {
		return r.create(ctx, appointment)
	}
	return r.update(ctx, appointment)
}

// Update updates an existing appointment
func (r *SqlcAppointmentRepository) Update(ctx context.Context, appointment *appt.Appointment) error {
	return r.update(ctx, appointment)
}

// Delete deletes an appointment
func (r *SqlcAppointmentRepository) Delete(ctx context.Context, id valueobject.AppointmentID) error {
	if err := r.queries.DeleteAppointment(ctx, int32(id.Value())); err != nil {
		return r.dbError("delete", "failed to delete appointment", err)
	}
	return nil
}

// CountByStatus counts appointments by status
func (r *SqlcAppointmentRepository) CountByStatus(ctx context.Context, status enum.AppointmentStatus) (int64, error) {
	count, err := r.queries.CountAppointmentsByStatus(ctx, models.AppointmentStatus(status))
	if err != nil {
		return 0, r.dbError("select", "failed to count appointments by status", err)
	}
	return count, nil
}

// CountByEmployeeID counts appointments by employee ID
func (r *SqlcAppointmentRepository) CountByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID) (int64, error) {
	count, err := r.queries.CountAppointmentsByEmployeeID(ctx, pgtype.Int4{Int32: int32(employeeID.Value()), Valid: true})
	if err != nil {
		return 0, r.dbError("select", "failed to count appointments by employee ID", err)
	}
	return count, nil
}

// create inserts a new appointment
func (r *SqlcAppointmentRepository) create(ctx context.Context, appointment *appt.Appointment) error {
	params := appointmentToCreateParams(appointment)

	_, err := r.queries.CreateAppointment(ctx, params)
	if err != nil {
		return r.dbError("insert", "failed to create appointment", err)
	}

	return nil
}

// update modifies an existing appointment
func (r *SqlcAppointmentRepository) update(ctx context.Context, appointment *appt.Appointment) error {
	params := appointmentToUpdateParams(appointment)

	_, err := r.queries.UpdateAppointment(ctx, params)
	if err != nil {
		return r.dbError("update", fmt.Sprintf("failed to update appointment with ID %d", appointment.ID().Value()), err)
	}

	return nil
}
