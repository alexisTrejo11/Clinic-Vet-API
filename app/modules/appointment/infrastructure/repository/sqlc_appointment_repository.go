// Package repository implements the AppointmentRepository interface using SQLC for database operations.
package repository

import (
	appt "clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"

	p "clinic-vet-api/app/shared/page"
	"clinic-vet-api/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (r *SqlcAppointmentRepository) Find(ctx context.Context, spec specification.ApptSearchSpecification) (p.Page[appt.Appointment], error) {
	params := spec.ToSQLCParams()
	var service string
	if params.Service != nil {
		service = *params.Service
	}

	var status string
	if params.Status != nil {
		status = *params.Status
	}

	// TODO: NIL FIX
	sqlcParams := sqlc.FindAppointmentsBySpecParams{
		Column1: int32(*params.ApptID),                                                             // ID
		Column2: int32(*params.CustomerID),                                                         // CustomerID
		Column3: int32(*params.EmployeeID),                                                         // EmployeeID
		Column4: int32(*params.PetID),                                                              // PetID
		Column5: service,                                                                           // Service
		Column6: status,                                                                            // Status
		Column7: pgtype.Timestamp{Time: *params.StartDate, Valid: params.StartDate != nil},         // StartDate
		Column8: pgtype.Timestamp{Time: *params.EndDate, Valid: params.EndDate != nil},             // EndDate
		Column9: pgtype.Timestamp{Time: *params.ScheduledDate, Valid: params.ScheduledDate != nil}, // ScheduledDate
		Limit:   params.Limit,
		Offset:  params.Offset,
	}

	rows, err := r.queries.FindAppointmentsBySpec(ctx, sqlcParams)
	if err != nil {
		return p.Page[appt.Appointment]{}, err
	}

	// Convert rows to domain entities
	countParams := sqlc.CountAppointmentsBySpecParams{
		Column1: sqlcParams.Column1, // ID
		Column2: sqlcParams.Column2, // CustomerID
		Column3: sqlcParams.Column3, // EmployeeID
		Column4: sqlcParams.Column4, // PetID
		Column5: sqlcParams.Column5, // Service
		Column6: sqlcParams.Column6, // Status
		Column7: sqlcParams.Column7, // StartDate
		Column8: sqlcParams.Column8, // EndDate
		Column9: sqlcParams.Column9, // ScheduledDate
	}

	appointments, err := r.ToDomainEntities(rows)
	if err != nil {
		return p.Page[appt.Appointment]{}, err
	}

	total, err := r.queries.CountAppointmentsBySpec(ctx, countParams)
	if err != nil {
		return p.Page[appt.Appointment]{}, err
	}

	pagination := p.PageInput{
		Limit:  int(params.Limit),
		Offset: int(params.Offset),
	}
	return p.NewPage(appointments, *p.GetPageMetadata(int(total), pagination)), nil
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

// ExistsByID checks if an appointment exists by ID
func (r *SqlcAppointmentRepository) ExistsByID(ctx context.Context, id valueobject.AppointmentID) (bool, error) {
	exists, err := r.queries.ExistsAppointmentID(ctx, int32(id.Value()))
	if err != nil {
		return false, r.dbError("select", "failed to check appointment existence by ID", err)
	}
	return exists, nil
}

func (r *SqlcAppointmentRepository) Save(ctx context.Context, appointment *appt.Appointment) error {
	if appointment.ID().IsZero() {
		return r.create(ctx, appointment)
	}
	return r.update(ctx, appointment)
}

func (r *SqlcAppointmentRepository) Delete(ctx context.Context, id valueobject.AppointmentID, hard bool) error {
	if err := r.queries.DeleteAppointment(ctx, int32(id.Value())); err != nil {
		return r.dbError("delete", "failed to delete appointment", err)
	}
	return nil
}

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

// Count returns the total number of appointments matching the given specification.
func (r *SqlcAppointmentRepository) Count(ctx context.Context, spec specification.ApptSearchSpecification) (int, error) {
	params := spec.ToSQLCParams()
	var service string
	if params.Service != nil {
		service = *params.Service
	}
	var status string
	if params.Status != nil {
		status = *params.Status
	}
	sqlcParams := sqlc.CountAppointmentsBySpecParams{
		Column1: int32(*params.ApptID),                                                             // ID
		Column2: int32(*params.CustomerID),                                                         // CustomerID
		Column3: int32(*params.EmployeeID),                                                         // EmployeeID
		Column4: int32(*params.PetID),                                                              // PetID
		Column5: service,                                                                           // Service
		Column6: status,                                                                            // Status
		Column7: pgtype.Timestamp{Time: *params.StartDate, Valid: params.StartDate != nil},         // StartDate
		Column8: pgtype.Timestamp{Time: *params.EndDate, Valid: params.EndDate != nil},             // EndDate
		Column9: pgtype.Timestamp{Time: *params.ScheduledDate, Valid: params.ScheduledDate != nil}, // ScheduledDate
	}
	total, err := r.queries.CountAppointmentsBySpec(ctx, sqlcParams)
	if err != nil {
		return 0, r.dbError("count", "failed to count appointments", err)
	}
	return int(total), nil
}
