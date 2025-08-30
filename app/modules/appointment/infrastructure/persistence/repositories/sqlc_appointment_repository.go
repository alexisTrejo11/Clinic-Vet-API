package sqlcrepository

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCAppointmentRepository struct {
	queries *sqlc.Queries
}

func NewSQLCAppointmentRepository(queries *sqlc.Queries) repository.AppointmentRepository {
	return &SQLCAppointmentRepository{
		queries: queries,
	}
}

func (r *SQLCAppointmentRepository) ListAll(ctx context.Context, pageInput page.PageData) (page.Page[[]entity.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinments(ctx, sqlc.ListAppoinmentsParams{
		Offset: int32(pageInput.PageNumber-1) * int32(pageInput.PageSize),
		Limit:  int32(pageInput.PageSize),
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	queryTotalCount, err := r.queries.CountAppoinments(ctx)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)

	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) GetByID(ctx context.Context, appointmentID valueobject.AppointmentID) (entity.Appointment, error) {
	sqlRow, err := r.queries.GetAppoinmentByID(ctx, int32(appointmentID.GetValue()))
	if err != nil {
		return entity.Appointment{}, err
	}

	appointment, err := sqlRowToDomain(sqlRow)
	if err != nil {
		return entity.Appointment{}, err
	}

	return *appointment, nil
}

func (r *SQLCAppointmentRepository) Search(ctx context.Context, pageInput page.PageData, searchCriteria map[string]interface{}) (page.Page[[]entity.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinments(ctx, sqlc.ListAppoinmentsParams{
		Offset: int32(pageInput.PageNumber-1) * int32(pageInput.PageSize),
		Limit:  int32(pageInput.PageSize),
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	queryTotalCount, err := r.queries.CountAppoinments(ctx)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)

	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) ListByVetID(ctx context.Context, vetID valueobject.VetID, pageInput page.PageData) (page.Page[[]entity.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByVeterinarianID(ctx, sqlc.ListAppoinmentsByVeterinarianIDParams{
		VeterinarianID: pgtype.Int4{Int32: int32(vetID.GetValue()), Valid: true},
		Offset:         int32(pageInput.PageNumber-1) * int32(pageInput.PageSize),
		Limit:          int32(pageInput.PageNumber),
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	queryTotalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, pgtype.Int4{Int32: int32(vetID.GetValue()), Valid: true})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) ListByPetID(ctx context.Context, petID valueobject.PetID, pageInput page.PageData) (page.Page[[]entity.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByVeterinarianID(ctx, sqlc.ListAppoinmentsByVeterinarianIDParams{
		VeterinarianID: pgtype.Int4{Int32: int32(petID.GetValue()), Valid: true},
		Offset:         int32(pageInput.PageNumber-1) * int32(pageInput.PageSize),
		Limit:          int32(pageInput.PageNumber),
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	queryTotalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, pgtype.Int4{Int32: int32(petID.GetValue()), Valid: true})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) ListByOwnerID(ctx context.Context, ownerID valueobject.OwnerID, pageInput page.PageData) (page.Page[[]entity.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByOwnerID(ctx, sqlc.ListAppoinmentsByOwnerIDParams{
		OwnerID: int32(ownerID.GetValue()),
		Offset:  int32(pageInput.PageNumber-1) * int32(pageInput.PageSize),
		Limit:   int32(pageInput.PageSize),
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	queryTotalCount, err := r.queries.CountAppoinmentsByVeterinarianID(ctx, pgtype.Int4{Int32: int32(ownerID.GetValue()), Valid: true})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) ListByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageData) (page.Page[[]entity.Appointment], error) {
	sqlRows, err := r.queries.ListAppoinmentsByDateRange(ctx, sqlc.ListAppoinmentsByDateRangeParams{
		ScheduleDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
		Offset:         int32(pageInput.PageNumber-1) * int32(pageInput.PageSize),
		Limit:          int32(pageInput.PageSize),
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	appointments, err := sqlRowsToDomainList(sqlRows)
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	queryTotalCount, err := r.queries.CountAppoinmentsByDateRange(ctx, sqlc.CountAppoinmentsByDateRangeParams{
		ScheduleDate:   pgtype.Timestamptz{Time: startDate, Valid: true},
		ScheduleDate_2: pgtype.Timestamptz{Time: endDate, Valid: true},
	})
	if err != nil {
		return page.Page[[]entity.Appointment]{}, AppointmentDBError(err.Error())
	}

	pageMetadata := page.GetPageMetadata(int(queryTotalCount), pageInput)
	return page.NewPage(appointments, *pageMetadata), nil
}

func (r *SQLCAppointmentRepository) Save(ctx context.Context, appointment *entity.Appointment) error {
	if appointment.GetID().GetValue() == 0 {
		return r.create(ctx, appointment)
	}
	return r.update(ctx, appointment)
}

func (r *SQLCAppointmentRepository) Delete(appointmentID valueobject.AppointmentID) error {
	if err := r.queries.DeleteAppoinment(context.Background(), int32(appointmentID.GetValue())); err != nil {
		return AppointmentDeleteDBErr(err.Error())
	}
	return nil
}

func (r *SQLCAppointmentRepository) create(ctx context.Context, appointment *entity.Appointment) error {
	result, err := r.queries.CreateAppoinment(ctx, domainToCreateParams(appointment))
	if err != nil {
		return AppointmentInsertDBErr(err.Error())
	}

	appointmentID, err := valueobject.NewAppointmentID(int(result.ID))
	if err != nil {
		return err
	}

	appointment.SetID(appointmentID)
	return nil
}

func (r *SQLCAppointmentRepository) update(ctx context.Context, appointment *entity.Appointment) error {
	_, err := r.queries.UpdateAppoinment(ctx, domainToUpdateParams(appointment))
	if err != nil {
		return AppointmentUpdateDBErr(err.Error())
	}

	appointment.SetUpdatedAt(time.Now())
	return nil
}
