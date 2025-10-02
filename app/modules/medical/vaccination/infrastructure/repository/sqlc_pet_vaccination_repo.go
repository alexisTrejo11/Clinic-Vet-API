package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/mapper"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/sqlc"
	"context"
	"time"
)

type SqlcPetVaccinationRepository struct {
	queries *sqlc.Queries
	pgMap   mapper.SqlcFieldMapper
}

func NewSqlcPetVaccinationRepository(queries *sqlc.Queries, pgMap mapper.SqlcFieldMapper) repository.VaccinationRepository {
	return &SqlcPetVaccinationRepository{queries: queries, pgMap: pgMap}
}

func (r *SqlcPetVaccinationRepository) FindAllByPetID(ctx context.Context, petID valueobject.PetID) ([]medical.PetVaccination, error) {
	sqlcRows, err := r.queries.FindAllVaccinationsByPetID(ctx, petID.Int32())
	if err != nil {
		return nil, err
	}
	return r.SqlcRowsToEntities(sqlcRows), nil
}

// FindByDateRange implements repository.VaccinationRepository.
func (r *SqlcPetVaccinationRepository) FindByDateRange(ctx context.Context, startDate time.Time, endDate time.Time, pagination page.PaginationRequest) (page.Page[medical.PetVaccination], error) {
	sqlcRows, err := r.queries.FindVaccinationsByDateRange(ctx, sqlc.FindVaccinationsByDateRangeParams{
		AdministeredDate:   r.pgMap.TimeToPgDate(startDate),
		AdministeredDate_2: r.pgMap.TimeToPgDate(endDate),
		Limit:              pagination.Limit(),
		Offset:             pagination.Offset(),
	})
	if err != nil {
		return page.Page[medical.PetVaccination]{}, err
	}

	total, err := r.queries.CountVaccinationsByDateRange(ctx, sqlc.CountVaccinationsByDateRangeParams{
		AdministeredDate:   r.pgMap.TimeToPgDate(startDate),
		AdministeredDate_2: r.pgMap.TimeToPgDate(endDate),
	})
	if err != nil {
		return page.Page[medical.PetVaccination]{}, err
	}

	vaccines := r.SqlcRowsToEntities(sqlcRows)
	return page.NewPage(vaccines, total, pagination), nil
}

// FindByEmployeeID implements repository.VaccinationRepository.
func (r *SqlcPetVaccinationRepository) FindByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID, pagination page.PaginationRequest) (page.Page[medical.PetVaccination], error) {
	panic("unimplemented")
}

func (r *SqlcPetVaccinationRepository) FindByID(ctx context.Context, vaccinationID valueobject.VaccinationID) (*medical.PetVaccination, error) {
	row, err := r.queries.FindVaccinationByID(ctx, vaccinationID.Int32())
	if err != nil {
		return nil, err
	}
	entity := r.SqlcRowToEntity(row)
	return &entity, nil
}

// FindByIDAndEmployeeID implements repository.VaccinationRepository.
func (r *SqlcPetVaccinationRepository) FindByIDAndEmployeeID(ctx context.Context, vaccinationID valueobject.VaccinationID, employeeID valueobject.EmployeeID) (*medical.PetVaccination, error) {
	row, err := r.queries.FindVaccinationByIDAndAdministeredBy(ctx, sqlc.FindVaccinationByIDAndAdministeredByParams{
		ID:             vaccinationID.Int32(),
		AdministeredBy: r.pgMap.PgInt4.FromUint(employeeID.Value()),
	})
	if err != nil {
		return nil, err
	}
	entity := r.SqlcRowToEntity(row)
	return &entity, nil
}

func (r *SqlcPetVaccinationRepository) FindByIDAndPetID(ctx context.Context, vaccinationID valueobject.VaccinationID, petID valueobject.PetID) (*medical.PetVaccination, error) {
	row, err := r.queries.FindByIDAndPetID(ctx, sqlc.FindByIDAndPetIDParams{
		ID:    vaccinationID.Int32(),
		PetID: petID.Int32(),
	})
	if err != nil {
		return nil, err
	}
	entity := r.SqlcRowToEntity(row)
	return &entity, nil
}

func (r *SqlcPetVaccinationRepository) FindByPetID(ctx context.Context, petID valueobject.PetID, pagination page.PaginationRequest) (page.Page[medical.PetVaccination], error) {
	sqlcRows, err := r.queries.FindVaccinationsByPetID(ctx, sqlc.FindVaccinationsByPetIDParams{
		PetID:  petID.Int32(),
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	})
	if err != nil {
		return page.Page[medical.PetVaccination]{}, err
	}

	total, err := r.queries.CountVaccinationsByPetID(ctx, petID.Int32())
	if err != nil {
		return page.Page[medical.PetVaccination]{}, err
	}

	vaccines := r.SqlcRowsToEntities(sqlcRows)
	return page.NewPage(vaccines, total, pagination), nil
}

func (r *SqlcPetVaccinationRepository) FindByPetIDs(ctx context.Context, petIDs []valueobject.PetID, pagination page.PaginationRequest) (page.Page[medical.PetVaccination], error) {
	ids := make([]int32, len(petIDs))
	for i, id := range petIDs {
		ids[i] = id.Int32()
	}

	rows, err := r.queries.FindVaccinationsByPetIDs(ctx, sqlc.FindVaccinationsByPetIDsParams{
		Column1: ids,
		Limit:   pagination.Limit(),
		Offset:  pagination.Offset(),
	})
	if err != nil {
		return page.Page[medical.PetVaccination]{}, err
	}

	total, err := r.queries.CountVaccinationsByPetIDs(ctx, ids)
	if err != nil {
		return page.Page[medical.PetVaccination]{}, err
	}

	vaccines := r.SqlcRowsToEntities(rows)
	return page.NewPage(vaccines, total, pagination), nil
}

func (r *SqlcPetVaccinationRepository) FindRecentByPetID(ctx context.Context, petID valueobject.PetID, days int) ([]medical.PetVaccination, error) {
	panic("unimplemented")
}

func (r *SqlcPetVaccinationRepository) Save(ctx context.Context, vaccination medical.PetVaccination) (medical.PetVaccination, error) {
	if vaccination.ID().IsZero() {
		// Create
		params := r.EntityToCreateParams(vaccination)
		row, err := r.queries.CreatePetVaccination(ctx, params)
		if err != nil {
			return medical.PetVaccination{}, err
		}
		createdEntity := r.SqlcRowToEntity(row)
		return createdEntity, nil
	} else {
		// Update
		params := r.EntityToUpdateParams(vaccination)
		row, err := r.queries.UpdateVaccination(ctx, params)
		if err != nil {
			return medical.PetVaccination{}, err
		}
		updatedEntity := r.SqlcRowToEntity(row)
		return updatedEntity, nil
	}
}

func (r *SqlcPetVaccinationRepository) Delete(ctx context.Context, vaccinationID valueobject.VaccinationID) error {
	return r.queries.DeleteVaccination(ctx, vaccinationID.Int32())
}
