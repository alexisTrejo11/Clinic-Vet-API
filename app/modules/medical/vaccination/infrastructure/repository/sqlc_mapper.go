package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/sqlc"
)

func (r *SqlcPetVaccinationRepository) SqlcRowToEntity(row sqlc.PetVaccination) medical.PetVaccination {
	return *medical.NewPetVaccinationBuilder().
		WithID(valueobject.NewVaccinationID(uint(row.ID))).
		WithAdministeredBy(r.pgMap.PgInt4.ToEmployeeID(row.AdministeredBy)).
		WithPetID(valueobject.NewPetID(uint(row.PetID))).
		WithVaccineName(row.VaccineName).
		WithBatchNumber(r.pgMap.PgText.ToString(row.BatchNumber)).
		WithAdministeredDate(r.pgMap.PgDate.ToTime(row.AdministeredDate)).
		WithNextDueDate(r.pgMap.PgDate.ToTimePtr(row.NextDueDate)).
		WithNotes(r.pgMap.PgText.ToStringPtr(row.Notes)).
		WithVaccineType(row.VaccineType).
		WithTimeStamps(r.pgMap.PgTimestamptz.ToTime(row.CreatedAt), r.pgMap.PgTimestamptz.ToTime(row.UpdatedAt)).
		Build()
}

func (r *SqlcPetVaccinationRepository) SqlcRowsToEntities(rows []sqlc.PetVaccination) []medical.PetVaccination {
	entities := make([]medical.PetVaccination, len(rows))
	for i, row := range rows {
		entities[i] = r.SqlcRowToEntity(row)
	}
	return entities
}

func (r *SqlcPetVaccinationRepository) EntityToCreateParams(entity medical.PetVaccination) sqlc.CreatePetVaccinationParams {
	return sqlc.CreatePetVaccinationParams{
		PetID:            entity.PetID().Int32(),
		VaccineName:      entity.VaccineName(),
		AdministeredDate: r.pgMap.TimeToPgDate(entity.AdministeredDate()),
		AdministeredBy:   r.pgMap.PgInt4.FromUint(entity.AdministeredBy().Value()),
		Notes:            r.pgMap.PgText.FromStringPtr(entity.Notes()),
		NextDueDate:      r.pgMap.PgDate.FromTimePtr(entity.NextDueDate()),
		BatchNumber:      r.pgMap.PgText.FromString(entity.BatchNumber()),
		VaccineType:      entity.VaccineType(),
	}
}

func (r *SqlcPetVaccinationRepository) EntityToUpdateParams(entity medical.PetVaccination) sqlc.UpdateVaccinationParams {
	return sqlc.UpdateVaccinationParams{
		ID:               entity.ID().Int32(),
		PetID:            entity.PetID().Int32(),
		VaccineName:      entity.VaccineName(),
		AdministeredDate: r.pgMap.TimeToPgDate(entity.AdministeredDate()),
		AdministeredBy:   r.pgMap.PgInt4.FromUint(entity.AdministeredBy().Value()),
		Notes:            r.pgMap.PgText.FromStringPtr(entity.Notes()),
		NextDueDate:      r.pgMap.PgDate.FromTimePtr(entity.NextDueDate()),
		BatchNumber:      r.pgMap.PgText.FromString(entity.BatchNumber()),
		VaccineType:      entity.VaccineType(),
	}
}
