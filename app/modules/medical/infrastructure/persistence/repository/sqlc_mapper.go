package repositoryimpl

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToEntity(sqlRow sqlc.MedicalSession) (medical.MedicalSession, error) {
	medSessionID := valueobject.NewMedSessionID(uint(sqlRow.ID))
	petID := valueobject.NewPetID(uint(sqlRow.PetID))
	employeeID := valueobject.NewEmployeeID(uint(sqlRow.EmployeeID))
	customerID := valueobject.NewCustomerID(uint(sqlRow.CustomerID))

	// Parsear enums
	visitType, err := enum.ParseVisitType(sqlRow.VisitType)
	if err != nil {
		return medical.MedicalSession{}, fmt.Errorf("invalid visit type: %w", err)
	}

	var condition enum.PetCondition
	if sqlRow.Condition.Valid {
		condition, err = enum.ParsePetCondition(sqlRow.Condition.String)
		if err != nil {
			return medical.MedicalSession{}, fmt.Errorf("invalid pet condition: %w", err)
		}
	} else {
		condition = enum.PetConditionStable
	}

	visitReason := enum.VisitReasonRoutineCheckup
	if sqlRow.IsEmergency.Valid && sqlRow.IsEmergency.Bool {
		visitReason = enum.VisitReasonEmergency
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
		err := json.Unmarshal([]byte(sqlRow.Symptoms.String), &symptoms)
		if err != nil {
			return medical.MedicalSession{}, fmt.Errorf("invalid symptoms format: %w", err)
		}
	}

	var medications []string
	if sqlRow.Medications.Valid {
		err := json.Unmarshal([]byte(sqlRow.Medications.String), &medications)
		if err != nil {
			return medical.MedicalSession{}, fmt.Errorf("invalid medications format: %w", err)
		}
	}

	var followUpDate *time.Time
	if sqlRow.FollowUpDate.Valid {
		followUpDate = &sqlRow.FollowUpDate.Time
	}

	medicalSession, err := medical.NewMedicalSession(
		medSessionID,
		petID,
		customerID,
		employeeID,
		medical.WithVisitReason(visitReason),
		medical.WithVisitType(visitType),
		medical.WithVisitDate(visitDate),
		medical.WithNotes(notes),
		medical.WithDiagnosis(sqlRow.Diagnosis.String),
		medical.WithTreatment(sqlRow.Treatment.String),
		medical.WithCondition(condition),
		medical.WithWeight(weight),
		medical.WithTemperature(temperature),
		medical.WithHeartRate(heartRate),
		medical.WithRespiratoryRate(respiratoryRate),
		medical.WithSymptoms(symptoms),
		medical.WithMedications(medications),
		medical.WithFollowUpDate(followUpDate),
	)
	if err != nil {
		return medical.MedicalSession{}, fmt.Errorf("failed to create medical session: %w", err)
	}

	return *medicalSession, nil
}

func ToEntities(medSessionList []sqlc.MedicalSession) ([]medical.MedicalSession, error) {
	domainList := make([]medical.MedicalSession, len(medSessionList))

	for i, sqlRow := range medSessionList {
		domainMedSession, err := ToEntity(sqlRow)
		if err != nil {
			return nil, err
		}
		domainList[i] = domainMedSession
	}

	return domainList, nil
}

func ToUpdateParams(medSession medical.MedicalSession) sqlc.UpdateMedicalSessionParams {
	params := sqlc.UpdateMedicalSessionParams{
		ID:         int32(medSession.ID().Value()),
		PetID:      int32(medSession.PetID().Value()),
		CustomerID: int32(medSession.CustomerID().Value()),
		EmployeeID: int32(medSession.EmployeeID().Value()),
		VisitDate:  pgtype.Timestamptz{Time: medSession.VisitDate(), Valid: true},
		Diagnosis:  pgtype.Text{String: medSession.Diagnosis(), Valid: medSession.Diagnosis() != ""},
		Treatment:  pgtype.Text{String: medSession.Treatment(), Valid: medSession.Treatment() != ""},
		VisitType:  medSession.VisitType().DisplayName(),
		Condition:  pgtype.Text{String: medSession.Condition().DisplayName(), Valid: true},
	}

	// Notes
	if medSession.Notes() != nil && *medSession.Notes() != "" {
		params.Notes = pgtype.Text{String: *medSession.Notes(), Valid: true}
	} else {
		params.Notes = pgtype.Text{Valid: false}
	}

	if medSession.Weight() != nil {
		weightInt := big.NewInt(medSession.Weight().Int())
		params.Weight = pgtype.Numeric{Int: weightInt, Valid: true}
	} else {
		params.Weight = pgtype.Numeric{Valid: false}
	}

	if medSession.Temperature() != nil {
		params.Temperature = pgtype.Numeric{Int: big.NewInt(medSession.Temperature().Int()), Valid: true}
	} else {
		params.Temperature = pgtype.Numeric{Valid: false}
	}

	if medSession.HeartRate() != nil {
		params.HeartRate = pgtype.Int4{Int32: int32(*medSession.HeartRate()), Valid: true}
	} else {
		params.HeartRate = pgtype.Int4{Valid: false}
	}

	if medSession.RespiratoryRate() != nil {
		params.RespiratoryRate = pgtype.Int4{Int32: int32(*medSession.RespiratoryRate()), Valid: true}
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
		PetID:      int32(medSession.PetID().Value()),
		CustomerID: int32(medSession.CustomerID().Value()),
		EmployeeID: int32(medSession.EmployeeID().Value()),
		VisitType:  medSession.VisitType().DisplayName(),
		VisitDate:  pgtype.Timestamptz{Time: medSession.VisitDate(), Valid: true},
		Diagnosis:  pgtype.Text{String: medSession.Diagnosis(), Valid: medSession.Diagnosis() != ""},
		Treatment:  pgtype.Text{String: medSession.Treatment(), Valid: medSession.Treatment() != ""},
		Condition:  pgtype.Text{String: medSession.Condition().DisplayName(), Valid: true},
	}

	if medSession.Notes() != nil && *medSession.Notes() != "" {
		params.Notes = pgtype.Text{String: *medSession.Notes(), Valid: true}
	} else {
		params.Notes = pgtype.Text{Valid: false}
	}

	if medSession.Weight() != nil {
		params.Weight = pgtype.Numeric{Int: big.NewInt(medSession.Weight().Int()), Valid: true}
	} else {
		params.Weight = pgtype.Numeric{Valid: false}
	}

	if medSession.Temperature() != nil {
		params.Temperature = pgtype.Numeric{Int: big.NewInt(medSession.Temperature().Int()), Valid: true}
	} else {
		params.Temperature = pgtype.Numeric{Valid: false}
	}

	if medSession.HeartRate() != nil {
		params.HeartRate = pgtype.Int4{Int32: int32(*medSession.HeartRate()), Valid: true}
	} else {
		params.HeartRate = pgtype.Int4{Valid: false}
	}

	if medSession.RespiratoryRate() != nil {
		params.RespiratoryRate = pgtype.Int4{Int32: int32(*medSession.RespiratoryRate()), Valid: true}
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
