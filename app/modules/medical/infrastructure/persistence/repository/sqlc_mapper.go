package repositoryimpl

import (
	"encoding/json"
	"math/big"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToEntity(sqlRow sqlc.MedicalSession) medical.MedicalSession {
	var condition enum.PetCondition
	if sqlRow.Condition.Valid {
		condition = enum.PetCondition(sqlRow.Condition.String)
	}

	var notes *string
	if sqlRow.Notes.Valid {
		notes = &sqlRow.Notes.String
	}

	var visitDate time.Time
	if sqlRow.VisitDate.Valid {
		visitDate = sqlRow.VisitDate.Time
	} else {
		visitDate = time.Now()
	}

	var weight *valueobject.Decimal
	if sqlRow.Weight.Valid {
		weightVal := valueobject.NewDecimalFromInt(sqlRow.Weight.Int.Int64())
		weight = &weightVal
	}

	var temperature *valueobject.Decimal
	if sqlRow.Temperature.Valid {
		if sqlRow.Temperature.Int != nil {
			tempVal := valueobject.NewDecimalFromInt(sqlRow.Temperature.Int.Int64())
			temperature = &tempVal
		}
	}

	var heartRate *int
	if sqlRow.HeartRate.Valid {
		hr := int(sqlRow.HeartRate.Int32)
		heartRate = &hr
	}

	var respiratoryRate *int
	if sqlRow.RespiratoryRate.Valid {
		rr := int(sqlRow.RespiratoryRate.Int32)
		respiratoryRate = &rr
	}

	var symptoms []string
	if sqlRow.Symptoms.Valid {
		json.Unmarshal([]byte(sqlRow.Symptoms.String), &symptoms)

	}

	var medications []string
	if sqlRow.Medications.Valid {
		json.Unmarshal([]byte(sqlRow.Medications.String), &medications)
	}

	var followUpDate *time.Time
	if sqlRow.FollowUpDate.Valid {
		followUpDate = &sqlRow.FollowUpDate.Time
	}

	petDetails := medical.NewPetSessionSummaryBuilder().
		WithPetID(valueobject.NewPetID(uint(sqlRow.PetID))).
		WithWeight(weight).
		WithTemperature(temperature).
		WithHeartRate(heartRate).
		WithRespiratoryRate(respiratoryRate).
		WithDiagnosis(sqlRow.Diagnosis.String).
		WithTreatment(sqlRow.Treatment.String).
		WithCondition(condition).
		WithMedications(medications).
		WithFollowUpDate(followUpDate).
		WithSymptoms(symptoms).
		Build()

	return *medical.NewMedicalSessionBuilder().
		WithID(valueobject.NewMedSessionID(uint(sqlRow.ID))).
		WithEmployeeID(valueobject.NewEmployeeID(uint(sqlRow.EmployeeID))).
		WithCustomerID(valueobject.NewCustomerID(uint(sqlRow.CustomerID))).
		WithVisitType(enum.VisitType(sqlRow.VisitType)).
		WithVisitDate(visitDate).
		WithNotes(notes).
		WithPetDetails(*petDetails).
		Build()
}

func ToEntities(medSessionList []sqlc.MedicalSession) ([]medical.MedicalSession, error) {
	domainList := make([]medical.MedicalSession, len(medSessionList))

	for i, sqlRow := range medSessionList {
		domainList[i] = ToEntity(sqlRow)
	}

	return domainList, nil
}

func ToUpdateParams(medSession medical.MedicalSession) sqlc.UpdateMedicalSessionParams {
	params := sqlc.UpdateMedicalSessionParams{
		ID:         int32(medSession.ID().Value()),
		PetID:      int32(medSession.PetDetails().PetID().Value()),
		CustomerID: int32(medSession.CustomerID().Value()),
		EmployeeID: int32(medSession.EmployeeID().Value()),
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
		params.HeartRate = pgtype.Int4{Int32: int32(*medSession.PetDetails().HeartRate()), Valid: true}
	}

	if medSession.PetDetails().RespiratoryRate() != nil {
		params.RespiratoryRate = pgtype.Int4{Int32: int32(*medSession.PetDetails().RespiratoryRate()), Valid: true}
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

func ToCreateParams(medSession medical.MedicalSession) sqlc.SaveMedicalSessionParams {
	params := sqlc.SaveMedicalSessionParams{
		PetID:      int32(medSession.PetDetails().PetID().Value()),
		CustomerID: int32(medSession.CustomerID().Value()),
		EmployeeID: int32(medSession.EmployeeID().Value()),
		VisitType:  medSession.VisitType().DisplayName(),
		VisitDate:  pgtype.Timestamptz{Time: medSession.VisitDate(), Valid: true},
	}

	if medSession.PetDetails().Diagnosis() != "" {
		params.Diagnosis = pgtype.Text{String: medSession.PetDetails().Diagnosis(), Valid: true}
	} else {
		params.Diagnosis = pgtype.Text{Valid: false}
	}

	if medSession.PetDetails().Treatment() != "" {
		params.Treatment = pgtype.Text{String: medSession.PetDetails().Treatment(), Valid: true}
	} else {
		params.Treatment = pgtype.Text{Valid: false}
	}

	if medSession.PetDetails().Condition().IsValid() {
		params.Condition = pgtype.Text{String: medSession.PetDetails().Condition().String(), Valid: true}
	}
	// Notes

	if medSession.Notes() != nil && *medSession.Notes() != "" {
		params.Notes = pgtype.Text{String: *medSession.Notes(), Valid: true}
	} else {
		params.Notes = pgtype.Text{Valid: false}
	}

	if medSession.PetDetails().Weight() != nil {
		params.Weight = pgtype.Numeric{Int: big.NewInt(medSession.PetDetails().Weight().Int()), Valid: true}
	} else {
		params.Weight = pgtype.Numeric{Valid: false}
	}

	if medSession.PetDetails().Temperature() != nil {
		params.Temperature = pgtype.Numeric{Int: big.NewInt(medSession.PetDetails().Temperature().Int()), Valid: true}
	} else {
		params.Temperature = pgtype.Numeric{Valid: false}
	}

	if medSession.PetDetails().HeartRate() != nil {
		params.HeartRate = pgtype.Int4{Int32: int32(*medSession.PetDetails().HeartRate()), Valid: true}
	} else {
		params.HeartRate = pgtype.Int4{Valid: false}
	}

	if medSession.PetDetails().RespiratoryRate() != nil {
		params.RespiratoryRate = pgtype.Int4{Int32: int32(*medSession.PetDetails().RespiratoryRate()), Valid: true}
	} else {
		params.RespiratoryRate = pgtype.Int4{Valid: false}
	}

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
	return params
}
