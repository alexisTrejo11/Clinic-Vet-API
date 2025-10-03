// Package repository implements the AppointmentRepository interface using SQLC for database operations.
package repository

import (
	appt "clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"

	"clinic-vet-api/app/shared/mapper"
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
	pgMap   *mapper.SqlcFieldMapper
}

func NewSqlcAppointmentRepository(queries *sqlc.Queries) repository.AppointmentRepository {
	return &SqlcAppointmentRepository{
		queries: queries,
		pgMap:   mapper.NewSqlcFieldMapper(),
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

	sqlcParams := sqlc.FindAppointmentsBySpecParams{
		Column1: r.pgMap.Primitive.Int32PtrToInt32(params.ApptID),                          // ID
		Column2: r.pgMap.Primitive.Int32PtrToInt32(params.CustomerID),                      // CustomerID
		Column3: r.pgMap.Primitive.Int32PtrToInt32(params.EmployeeID),                      // EmployeeID
		Column4: r.pgMap.Primitive.Int32PtrToInt32(params.PetID),                           // PetID
		Column5: service,                                                                   // Service
		Column6: status,                                                                    // Status
		Column7: pgtype.Timestamp(r.pgMap.PgTimestamptz.FromTimePtr(params.StartDate)),     // StartDate
		Column8: pgtype.Timestamp(r.pgMap.PgTimestamptz.FromTimePtr(params.EndDate)),       // EndDate
		Column9: pgtype.Timestamp(r.pgMap.PgTimestamptz.FromTimePtr(params.ScheduledDate)), // ScheduledDate
		Limit:   params.Limit,
		Offset:  params.Offset,
	}

	rows, err := r.queries.FindAppointmentsBySpec(ctx, sqlcParams)
	if err != nil {
		return p.Page[appt.Appointment]{}, err
	}

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

	appointments := toEntities(rows)

	total, err := r.queries.CountAppointmentsBySpec(ctx, countParams)
	if err != nil {
		return p.Page[appt.Appointment]{}, err
	}

	// Fix pagination to have default values if not set
	pagination := p.PaginationRequest{
		PageSize: 10,
		Page:     1,
	}

	return p.NewPage(appointments, total, pagination), nil
}

func (r *SqlcAppointmentRepository) FindByID(ctx context.Context, id valueobject.AppointmentID) (appt.Appointment, error) {
	sqlRow, err := r.queries.FindAppointmentByID(ctx, id.Int32())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appt.Appointment{}, r.notFoundError("id", id.String())
		}
		return appt.Appointment{}, r.dbError("select", "failed to get appointment by ID", err)
	}

	return *sqlcToEntity(sqlRow), nil
}

func (r *SqlcAppointmentRepository) ExistsByID(ctx context.Context, id valueobject.AppointmentID) (bool, error) {
	exists, err := r.queries.ExistsAppointmentID(ctx, id.Int32())
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
	if err := r.queries.DeleteAppointment(ctx, id.Int32()); err != nil {
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

func (r *SqlcAppointmentRepository) update(ctx context.Context, appointment *appt.Appointment) error {
	params := appointmentToUpdateParams(appointment)

	_, err := r.queries.UpdateAppointment(ctx, params)
	if err != nil {
		return r.dbError("update", fmt.Sprintf("failed to update appointment with ID %d", appointment.ID().Value()), err)
	}

	return nil
}

func (r *SqlcAppointmentRepository) Count(ctx context.Context, spec specification.ApptSearchSpecification) (int64, error) {
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
		Column1: r.pgMap.Primitive.Int32PtrToInt32(params.ApptID),                          // ID
		Column2: r.pgMap.Primitive.Int32PtrToInt32(params.CustomerID),                      // CustomerID
		Column3: r.pgMap.Primitive.Int32PtrToInt32(params.EmployeeID),                      // EmployeeID
		Column4: r.pgMap.Primitive.Int32PtrToInt32(params.PetID),                           // PetID
		Column5: service,                                                                   // Service
		Column6: status,                                                                    // Status
		Column7: pgtype.Timestamp(r.pgMap.PgTimestamptz.FromTimePtr(params.StartDate)),     // StartDate
		Column8: pgtype.Timestamp(r.pgMap.PgTimestamptz.FromTimePtr(params.EndDate)),       // EndDate
		Column9: pgtype.Timestamp(r.pgMap.PgTimestamptz.FromTimePtr(params.ScheduledDate)), // ScheduledDate
	}
	total, err := r.queries.CountAppointmentsBySpec(ctx, sqlcParams)
	if err != nil {
		return 0, r.dbError("count", "failed to count appointments", err)
	}
	return total, nil
}
