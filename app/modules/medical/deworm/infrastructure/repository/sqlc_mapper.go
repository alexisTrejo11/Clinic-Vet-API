package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/sqlc"
	"fmt"
)

var (
	ErrPetDewormingNotFound = fmt.Errorf("pet deworming not found")
)

func (r *SqlcPetDeworming) mapRowToDomain(row sqlc.PetDeworming) *medical.PetDeworming {
	return medical.NewPetDewormingBuilder().
		WithID(valueobject.NewDewormID(uint(row.ID))).
		WithPetID(valueobject.NewPetID(uint(row.PetID))).
		WithMedicationName(row.MedicationName).
		WithAdministeredDate(row.AdministeredDate.Time).
		WithCreatedAt(row.CreatedAt.Time).
		WithAdministeredBy(valueobject.NewEmployeeID(uint(row.AdministeredBy.Int32))).
		WithNextDueDate(r.mapper.PgDate.ToTimePtr(row.NextDueDate)).
		WithNotes(r.mapper.PgText.ToStringPtr(row.Notes)).
		Build()
}

func (r *SqlcPetDeworming) mapRowsToDomains(rows []sqlc.PetDeworming) []medical.PetDeworming {
	if len(rows) == 0 {
		return []medical.PetDeworming{}
	}

	dewormings := make([]medical.PetDeworming, len(rows))
	for i, row := range rows {
		dewormings[i] = *r.mapRowToDomain(row)
	}
	return dewormings
}

func (r *SqlcPetDeworming) domainToCreateParams(deworming medical.PetDeworming) sqlc.CreateDewormingParams {
	params := &sqlc.CreateDewormingParams{
		PetID:            int32(deworming.PetID().Value()),
		MedicationName:   deworming.MedicationName(),
		AdministeredDate: r.mapper.TimeToPgDate(deworming.AdministeredDate()),
		AdministeredBy:   r.mapper.UintToPgInt4(deworming.AdministeredBy().Value()),
		Notes:            r.mapper.StringPtrToPgText(deworming.Notes()),
		NextDueDate:      r.mapper.TimePtrToPgDate(deworming.NextDueDate()),
	}
	return *params
}

func (r *SqlcPetDeworming) domainToUpdateParams(deworming medical.PetDeworming) sqlc.UpdateDewormingParams {
	params := &sqlc.UpdateDewormingParams{
		ID:               int32(deworming.ID().Value()),
		MedicationName:   deworming.MedicationName(),
		AdministeredDate: r.mapper.TimeToPgDate(deworming.AdministeredDate()),
		AdministeredBy:   r.mapper.UintToPgInt4(deworming.AdministeredBy().Value()),
		Notes:            r.mapper.StringPtrToPgText(deworming.Notes()),
		NextDueDate:      r.mapper.TimePtrToPgDate(deworming.NextDueDate()),
	}
	return *params
}

func (r *SqlcPetDeworming) specToSqlc(spec specification.PetDewormSpecification) sqlc.FindDewormingsBySpecParams {
	params := &sqlc.FindDewormingsBySpecParams{
		LimitVal:              r.mapper.Primitive.Int32PtrToInt32(spec.Limit),
		OffsetVal:             r.mapper.Primitive.Int32PtrToInt32(spec.Offset),
		PetID:                 r.mapper.PetIDPtrToInt32(spec.PetID),
		AdministeredBy:        r.mapper.EmployeeIDPtrToInt32(spec.AdministeredBy),
		MedicationName:        r.mapper.StringPtrToString(spec.MedicationName),
		AdministeredDateFrom:  r.mapper.TimePtrToPgDate(spec.AdministeredDateFrom),
		AdministeredDateTo:    r.mapper.TimePtrToPgDate(spec.AdministeredDateTo),
		AdministeredDateExact: r.mapper.TimePtrToPgDate(spec.AdministeredDateExact),
		NextDueDateFrom:       r.mapper.TimePtrToPgDate(spec.NextDueDateFrom),
		NextDueDateTo:         r.mapper.TimePtrToPgDate(spec.NextDueDateTo),
		NextDueDateExact:      r.mapper.TimePtrToPgDate(spec.NextDueDateExact),
		CreatedAtFrom:         r.mapper.TimePtrToPgTypestamp(spec.CreatedAtFrom),
		CreatedAtTo:           r.mapper.TimePtrToPgTypestamp(spec.CreatedAtTo),
	}

	return *params
}

func (r *SqlcPetDeworming) ToCountSQLParams(spec specification.PetDewormSpecification) sqlc.CountDewormingsBySpecParams {
	params := &sqlc.CountDewormingsBySpecParams{
		PetID:                 r.mapper.PetIDPtrToInt32(spec.PetID),
		AdministeredBy:        r.mapper.EmployeeIDPtrToInt32(spec.AdministeredBy),
		MedicationName:        r.mapper.StringPtrToString(spec.MedicationName),
		AdministeredDateFrom:  r.mapper.TimePtrToPgDate(spec.AdministeredDateFrom),
		AdministeredDateTo:    r.mapper.TimePtrToPgDate(spec.AdministeredDateTo),
		AdministeredDateExact: r.mapper.TimePtrToPgDate(spec.AdministeredDateExact),
		NextDueDateFrom:       r.mapper.TimePtrToPgDate(spec.NextDueDateFrom),
		NextDueDateTo:         r.mapper.TimePtrToPgDate(spec.NextDueDateTo),
		NextDueDateExact:      r.mapper.TimePtrToPgDate(spec.NextDueDateExact),
		CreatedAtFrom:         r.mapper.TimePtrToPgTypestamp(spec.CreatedAtFrom),
		CreatedAtTo:           r.mapper.TimePtrToPgTypestamp(spec.CreatedAtTo),
	}

	return *params
}
