package appointmentRepository

import (
	"context"
	"time"

	appointDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostgresAppointmentRepository struct {
	queries *sqlc.Queries
}

func NewPostgresAppointmentRepository(queries *sqlc.Queries) appointDomain.AppointmentRepository {
	return &PostgresAppointmentRepository{
		queries: queries,
	}
}

func (r *PostgresAppointmentRepository) ListAll(ctx context.Context, pageInput page.PageData) (*page.Page[[]appointDomain.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinments(ctx, sqlc.ListAppoinmentsParams{
		Offset: int32(pageInput.PageNumber)*int32(pageInput.PageSize) - 1,
		Limit:  int32(pageInput.PageSize),
	})

	if err != nil {
		return nil, err
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return nil, nil
	}

	queryTotalCount, err := r.queries.CountAppoinments(ctx)
	if err != nil {
		return nil, err
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)

	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *PostgresAppointmentRepository) GetById(ctx context.Context, appointmentId int) (*appointDomain.Appointment, error) {
	sqlRow, err := r.queries.GetAppoinmentByID(ctx, int32(appointmentId))
	if err != nil {
		return nil, err
	}

	appointment, err := sqlRowToDomain(sqlRow)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (r *PostgresAppointmentRepository) Search(ctx context.Context, pageInput page.PageData, searchCriteria map[string]interface{}) (*page.Page[[]appointDomain.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinments(ctx, sqlc.ListAppoinmentsParams{
		Offset: 0,
		Limit:  int32(pageInput.PageSize),
	})

	if err != nil {
		return nil, err
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return nil, nil
	}

	queryTotalCount, err := r.queries.CountAppoinments(ctx)
	if err != nil {
		return nil, err
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)

	return page.NewPage(appointments, *pageMetadata), nil

}

func (r *PostgresAppointmentRepository) ListByVetId(ctx context.Context, ownerId int, pageInput page.PageData) (*page.Page[[]appointDomain.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByVeterinarianID(ctx, sqlc.ListAppoinmentsByVeterinarianIDParams{
		VeterinarianID: int32(ownerId),
		Offset:         int32(pageInput.PageNumber)*int32(pageInput.PageSize) - 1,
		Limit:          int32(pageInput.PageNumber),
	})

	if err != nil {
		return nil, err
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return nil, err
	}

	queryTotalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, int32(ownerId))
	if err != nil {
		return nil, err
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)

	return page.NewPage(appointments, *pageMetadata), nil

}
func (r *PostgresAppointmentRepository) ListByPetId(ctx context.Context, ownerId int, pageInput page.PageData) (*page.Page[[]appointDomain.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByVeterinarianID(ctx, sqlc.ListAppoinmentsByVeterinarianIDParams{
		VeterinarianID: int32(ownerId),
		Offset:         int32(pageInput.PageNumber)*int32(pageInput.PageSize) - 1,
		Limit:          int32(pageInput.PageNumber),
	})

	if err != nil {
		return nil, err
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return nil, err
	}

	queryTotalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, int32(ownerId))
	if err != nil {
		return nil, err
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)

	return page.NewPage(appointments, *pageMetadata), nil

}
func (r *PostgresAppointmentRepository) ListByOwnerId(ctx context.Context, ownerId int, pageInput page.PageData) (*page.Page[[]appointDomain.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByOwnerID(ctx, sqlc.ListAppoinmentsByOwnerIDParams{
		OwnerID: int32(ownerId),
		Offset:  int32(pageInput.PageNumber)*int32(pageInput.PageSize) - 1,
		Limit:   int32(pageInput.PageSize),
	})

	if err != nil {
		return nil, err
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return nil, err
	}

	queryTotalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, int32(ownerId))
	if err != nil {
		return nil, err
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)

	return page.NewPage(appointments, *pageMetadata), nil

}
func (r *PostgresAppointmentRepository) ListByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageData) (*page.Page[[]appointDomain.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByDateRange(ctx, sqlc.ListAppoinmentsByDateRangeParams{
		ScheduleDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Offset:         int32(pageInput.PageNumber)*int32(pageInput.PageSize) - 1,
		Limit:          int32(pageInput.PageSize),
	})

	if err != nil {
		return nil, err
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return nil, err
	}

	queryTotalCount, err := r.queries.CountAppoinmentsByDateRange(ctx, sqlc.CountAppoinmentsByDateRangeParams{
		ScheduleDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})

	if err != nil {
		return nil, err
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)

	return page.NewPage(appointments, *pageMetadata), nil

}

func (r *PostgresAppointmentRepository) Save(ctx context.Context, appointment *appointDomain.Appointment) error {
	if appointment.GetId().GetValue() == 0 {
		return r.create(ctx, appointment)
	}

	return r.update(ctx, appointment)
}
func (r *PostgresAppointmentRepository) Delete(appointmentId int) error {
	if err := r.queries.DeleteAppoinment(context.Background(), int32(appointmentId)); err != nil {
		return err
	}
	return nil
}

func (r *PostgresAppointmentRepository) create(ctx context.Context, appointment *appointDomain.Appointment) error {
	result, err := r.queries.CreateAppoinment(ctx, domainToCreateParams(appointment))
	if err != nil {
		return err
	}

	appointmentId, err := appointDomain.NewAppointmentId(result.ID)
	if err != nil {
		return err
	}

	appointment.SetId(appointmentId)

	return nil
}

func (r *PostgresAppointmentRepository) update(ctx context.Context, appointment *appointDomain.Appointment) error {
	_, err := r.queries.UpdateAppoinment(ctx, domainToUpdateParams(appointment))

	appointment.SetUpdatedAt(time.Now())
	return err
}
