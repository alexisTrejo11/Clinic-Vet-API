package repositoryimpl

import (
	"encoding/json"
	"math/big"

	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *SQLCMedSessionRepository) sqlcRowToEntity(sqlRow sqlc.MedicalSession) medical.MedicalSession {
	var symptoms []string
	if sqlRow.Symptoms.Valid {
		json.Unmarshal([]byte(sqlRow.Symptoms.String), &symptoms)
	}

	var medications []string
	if sqlRow.Medications.Valid {
		json.Unmarshal([]byte(sqlRow.Medications.String), &medications)
	}

	petDetails := medical.NewPetSessionSummaryBuilder().
		WithPetID(valueobject.NewPetID(uint(sqlRow.PetID))).
		WithWeight(r.pgMap.PgNumeric.ToDecimalPtr(sqlRow.Weight)).
		WithTemperature(r.pgMap.PgNumeric.ToDecimalPtr(sqlRow.Temperature)).
		WithHeartRate(r.pgMap.PgInt4.ToInt32Ptr(sqlRow.HeartRate)).
		WithRespiratoryRate(r.pgMap.PgInt4.ToInt32Ptr(sqlRow.RespiratoryRate)).
		WithDiagnosis(sqlRow.Diagnosis.String).
		WithTreatment(sqlRow.Treatment.String).
		WithCondition(enum.PetCondition(sqlRow.Condition.String)).
		WithMedications(medications).
		WithFollowUpDate(r.pgMap.PgTimestamptz.ToTimePtr(sqlRow.FollowUpDate)).
		WithSymptoms(symptoms).
		Build()

	return *medical.NewMedicalSessionBuilder().
		WithID(valueobject.NewMedSessionID(uint(sqlRow.ID))).
		WithEmployeeID(valueobject.NewEmployeeID(uint(sqlRow.EmployeeID))).
		WithCustomerID(valueobject.NewCustomerID(uint(sqlRow.CustomerID))).
		WithVisitType(enum.VisitType(sqlRow.VisitType)).
		WithVisitDate(r.pgMap.PgTimestamptz.ToTime(sqlRow.VisitDate)).
		WithNotes(r.pgMap.PgText.ToStringPtr(sqlRow.Notes)).
		WithPetDetails(*petDetails).
		WithTimeStamps(sqlRow.CreatedAt.Time, sqlRow.UpdatedAt.Time).
		Build()
}

func (r *SQLCMedSessionRepository) ToEntities(medSessionList []sqlc.MedicalSession) []medical.MedicalSession {
	domainList := make([]medical.MedicalSession, len(medSessionList))
	for i, sqlRow := range medSessionList {
		domainList[i] = r.sqlcRowToEntity(sqlRow)
	}
	return domainList
}

func (r *SQLCMedSessionRepository) toUpdateParams(medSession medical.MedicalSession) sqlc.UpdateMedicalSessionParams {
	params := sqlc.UpdateMedicalSessionParams{
		ID:         medSession.ID().Int32(),
		PetID:      medSession.PetDetails().PetID().Int32(),
		CustomerID: medSession.CustomerID().Int32(),
		EmployeeID: medSession.EmployeeID().Int32(),
		VisitDate:  pgtype.Timestamptz{Time: medSession.VisitDate(), Valid: true},
		VisitType:  medSession.VisitType().String(),
		Condition:  pgtype.Text{String: medSession.PetDetails().Condition().DisplayName(), Valid: true},
	}

	if medSession.PetDetails().Diagnosis() != "" {
		params.Diagnosis = pgtype.Text{String: medSession.PetDetails().Diagnosis(), Valid: true}
	}

	if medSession.PetDetails().Treatment() != "" {
		params.Treatment = pgtype.Text{String: medSession.PetDetails().Treatment(), Valid: true}
	}

	if medSession.PetDetails().Condition().IsValid() {
		params.Condition = pgtype.Text{String: medSession.PetDetails().Condition().String(), Valid: true}
	}

	// Notes
	if medSession.Notes() != nil && *medSession.Notes() != "" {
		params.Notes = pgtype.Text{String: *medSession.Notes(), Valid: true}
	}

	if medSession.PetDetails().Weight() != nil {
		weightInt := big.NewInt(medSession.PetDetails().Weight().Int())
		params.Weight = pgtype.Numeric{Int: weightInt, Valid: true}
	}

	if medSession.PetDetails().Temperature() != nil {
		params.Temperature = pgtype.Numeric{Int: big.NewInt(medSession.PetDetails().Temperature().Int()), Valid: true}
	}

	if medSession.PetDetails().HeartRate() != nil {
		params.HeartRate = pgtype.Int4{Int32: *medSession.PetDetails().HeartRate(), Valid: true}
	}

	if medSession.PetDetails().RespiratoryRate() != nil {
		params.RespiratoryRate = pgtype.Int4{Int32: *medSession.PetDetails().RespiratoryRate(), Valid: true}
	}

	/*
		// Arrays (JSON)
		if len(medSession.Symptoms()) > 0 {
			symptomsJSON, _ := json.Marshal(medSession.Symptoms())
			params.Symptoms = pgtype.Text{String: string(symptomsJSON), Valid: true}
		} else {
			params.Symptoms = pgtype.Text{Valid: false}
		}
	*/

	/*
		if len(medSession.Medications()) > 0 {
			medicationsJSON, _ := json.Marshal(medSession.Medications())
			params.Medications = pgtype.Text{String: string(medicationsJSON), Valid: true}
		} else {
			params.Medications = pgtype.Text{Valid: false}
		}
	*/

	/*
		// Follow-up date
		if medSession.FollowUpDate() != nil {
			params.FollowUpDate = pgtype.Timestamptz{Time: *medSession.FollowUpDate(), Valid: true}
		} else {
			params.FollowUpDate = pgtype.Timestamptz{Valid: false}
		}

		// IsEmergency
		isEmergency := medSession.VisitReason() == enum.VisitReasonEmergency
		params.IsEmergency = pgtype.Bool{Bool: isEmergency, Valid: true}
	*/

	return params
}

func (r *SQLCMedSessionRepository) toCreateParams(medSession medical.MedicalSession) sqlc.SaveMedicalSessionParams {
	/*
		// Arrays (JSON)
		if len(medSession.Symptoms()) > 0 {
			symptomsJSON, _ := json.Marshal(medSession.Symptoms())
			params.Symptoms = pgtype.Text{String: string(symptomsJSON), Valid: true}
		} else {
			params.Symptoms = pgtype.Text{Valid: false}
		}

		if len(medSession.Medications()) > 0 {
			medicationsJSON, _ := json.Marshal(medSession.Medications())
			params.Medications = pgtype.Text{String: string(medicationsJSON), Valid: true}
		} else {
			params.Medications = pgtype.Text{Valid: false}
		}

		// Follow-up date
		if medSession.FollowUpDate() != nil {
			params.FollowUpDate = pgtype.Timestamptz{Time: *medSession.FollowUpDate(), Valid: true}
		} else {
			params.FollowUpDate = pgtype.Timestamptz{Valid: false}
		}


		// IsEmergency
		isEmergency := medSession.VisitReason() == enum.VisitReasonEmergency
		params.IsEmergency = pgtype.Bool{Bool: isEmergency, Valid: true}
	*/
	return sqlc.SaveMedicalSessionParams{
		PetID:           medSession.PetDetails().PetID().Int32(),
		CustomerID:      medSession.CustomerID().Int32(),
		EmployeeID:      medSession.EmployeeID().Int32(),
		VisitType:       medSession.VisitType().DisplayName(),
		VisitDate:       r.pgMap.PgTimestamptz.FromTime(medSession.VisitDate()),
		Diagnosis:       r.pgMap.PgText.FromString(medSession.PetDetails().Diagnosis()),
		Treatment:       r.pgMap.PgText.FromString(medSession.PetDetails().Treatment()),
		Condition:       r.pgMap.PgText.FromString(medSession.PetDetails().Condition().String()),
		Notes:           r.pgMap.PgText.FromStringPtr(medSession.Notes()),
		Weight:          r.pgMap.PgNumeric.FromDecimalPtr(medSession.PetDetails().Weight()),
		Temperature:     r.pgMap.PgNumeric.FromDecimalPtr(medSession.PetDetails().Temperature()),
		HeartRate:       r.pgMap.PgInt4.FromInt32Ptr(medSession.PetDetails().HeartRate()),
		RespiratoryRate: r.pgMap.PgInt4.FromInt32Ptr(medSession.PetDetails().RespiratoryRate()),
	}
}
