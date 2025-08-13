package appointmentRepository

import (
	"context"
	"time"

	appoint "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCAppointmentRepository struct {
	queries *sqlc.Queries
}

func NewSQLCAppointmentRepository(queries *sqlc.Queries) appoint.AppointmentRepository {
	return &SQLCAppointmentRepository{
		queries: queries,
	}
}

func (r *SQLCAppointmentRepository) ListAll(ctx context.Context, pageInput page.PageData) (page.Page[[]appoint.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinments(ctx, sqlc.ListAppoinmentsParams{
		Offset: int32(pageInput.PageNumber-1) * int32(pageInput.PageSize),
		Limit:  int32(pageInput.PageSize),
	})

	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	queryTotalCount, err := r.queries.CountAppoinments(ctx)
	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)

	return *page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) GetById(ctx context.Context, appointmentId int) (appoint.Appointment, error) {
	sqlRow, err := r.queries.GetAppoinmentByID(ctx, int32(appointmentId))
	if err != nil {
		return appoint.Appointment{}, err
	}

	appointment, err := sqlRowToDomain(sqlRow)
	if err != nil {
		return appoint.Appointment{}, err
	}

	return *appointment, nil
}

func (r *SQLCAppointmentRepository) Search(ctx context.Context, pageInput page.PageData, searchCriteria map[string]interface{}) (page.Page[[]appoint.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinments(ctx, sqlc.ListAppoinmentsParams{
		Offset: int32(pageInput.PageNumber-1) * int32(pageInput.PageSize),
		Limit:  int32(pageInput.PageSize),
	})

	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	queryTotalCount, err := r.queries.CountAppoinments(ctx)
	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)

	return *page.NewPage(appointments, *pageMetadata), nil

}

func (r *SQLCAppointmentRepository) ListByVetId(ctx context.Context, ownerId int, pageInput page.PageData) (page.Page[[]appoint.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByVeterinarianID(ctx, sqlc.ListAppoinmentsByVeterinarianIDParams{
		VeterinarianID: int32(ownerId),
		Offset:         int32(pageInput.PageNumber-1) * int32(pageInput.PageSize),
		Limit:          int32(pageInput.PageNumber),
	})

	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	queryTotalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, int32(ownerId))
	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)
	return *page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) ListByPetId(ctx context.Context, ownerId int, pageInput page.PageData) (page.Page[[]appoint.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByVeterinarianID(ctx, sqlc.ListAppoinmentsByVeterinarianIDParams{
		VeterinarianID: int32(ownerId),
		Offset:         int32(pageInput.PageNumber-1) * int32(pageInput.PageSize),
		Limit:          int32(pageInput.PageNumber),
	})

	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	queryTotalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, int32(ownerId))
	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)
	return *page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) ListByOwnerId(ctx context.Context, ownerId int, pageInput page.PageData) (page.Page[[]appoint.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByOwnerID(ctx, sqlc.ListAppoinmentsByOwnerIDParams{
		OwnerID: int32(ownerId),
		Offset:  int32(pageInput.PageNumber-1) * int32(pageInput.PageSize),
		Limit:   int32(pageInput.PageSize),
	})

	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	queryTotalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, int32(ownerId))
	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)
	return *page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) ListByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageData) (page.Page[[]appoint.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByDateRange(ctx, sqlc.ListAppoinmentsByDateRangeParams{
		ScheduleDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Offset:         int32(pageInput.PageNumber-1) * int32(pageInput.PageSize),
		Limit:          int32(pageInput.PageSize),
	})

	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	queryTotalCount, err := r.queries.CountAppoinmentsByDateRange(ctx, sqlc.CountAppoinmentsByDateRangeParams{
		ScheduleDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})

	if err != nil {
		return page.Page[[]appoint.Appointment]{}, AppointmentDBError(err.Error())
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)
	return *page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) Save(ctx context.Context, appointment *appoint.Appointment) error {
	if appointment.GetId().GetValue() == 0 {
		return r.create(ctx, appointment)
	}
	return r.update(ctx, appointment)
}

func (r *SQLCAppointmentRepository) Delete(appointmentId int) error {
	if err := r.queries.DeleteAppoinment(context.Background(), int32(appointmentId)); err != nil {
		return AppointmentDeleteDBErr(err.Error())
	}
	return nil
}

func (r *SQLCAppointmentRepository) create(ctx context.Context, appointment *appoint.Appointment) error {
	result, err := r.queries.CreateAppoinment(ctx, domainToCreateParams(appointment))
	if err != nil {
		return AppointmentInsertDBErr(err.Error())
	}

	appointmentId, err := appoint.NewAppointmentId(result.ID)
	if err != nil {
		return err
	}

	appointment.SetId(appointmentId)
	return nil
}

func (r *SQLCAppointmentRepository) update(ctx context.Context, appointment *appoint.Appointment) error {
	_, err := r.queries.UpdateAppoinment(ctx, domainToUpdateParams(appointment))
	if err != nil {
		return AppointmentUpdateDBErr(err.Error())
	}

	appointment.SetUpdatedAt(time.Now())
	return nil
}
