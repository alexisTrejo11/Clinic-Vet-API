package repositoryimpl

import (
	"fmt"
	"time"

	"clinic-vet-api/app/core/domain/entity/medical"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToEntity(sqlRow sqlc.MedicalHistory) (medical.MedicalHistory, error) {
	medHistID := valueobject.NewMedHistoryID(uint(sqlRow.ID))
	petID := valueobject.NewPetID(uint(sqlRow.PetID))
	employeeID := valueobject.NewEmployeeID(uint(sqlRow.EmployeeID))
	customerID := valueobject.NewCustomerID(uint(sqlRow.CustomerID))

	visitType, err := enum.ParseVisitType(sqlRow.VisitType)
	if err != nil {
		return medical.MedicalHistory{}, fmt.Errorf("invalid visit type: %w", err)
	}

	var condition enum.PetCondition
	if sqlRow.Condition.Valid {
		condition, err = enum.ParsePetCondition(sqlRow.Condition.String)
		if err != nil {
			return medical.MedicalHistory{}, fmt.Errorf("invalid pet condition: %w", err)
		}
	} else {
		// Valor por defecto si es NULL
		condition = enum.PetConditionStable
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

	medicalHistory, err := medical.NewMedicalHistory(
		medHistID,
		petID,
		customerID,
		employeeID,
		medical.WithVisitReason(enum.VisitReasonEmergency),
		medical.WithVisitType(visitType),
		medical.WithVisitDate(visitDate),
		medical.WithNotes(*notes),
		medical.WithDiagnosis(sqlRow.Diagnosis.String),
		medical.WithTreatment(sqlRow.Treatment.String),
		medical.WithCondition(condition),
	)
	if err != nil {
		return medical.MedicalHistory{}, fmt.Errorf("failed to create medical history: %w", err)
	}

	return *medicalHistory, nil
}

func ToEntities(medHistList []sqlc.MedicalHistory) ([]medical.MedicalHistory, error) {
	domainList := make([]medical.MedicalHistory, len(medHistList))

	for i, sqlRow := range medHistList {
		domainMedHist, err := ToEntity(sqlRow)
		if err != nil {
			return nil, err
		}
		domainList[i] = domainMedHist
	}

	return domainList, nil
}

func ToCreateParams(medHist medical.MedicalHistory) sqlc.SaveMedicalHistoryParams {
	var notes string
	if medHist.Notes() != nil {
		notes = *medHist.Notes()
	}

	params := sqlc.SaveMedicalHistoryParams{
		PetID:      int32(medHist.PetID().Value()),
		CustomerID: int32(medHist.CustomerID().Value()),
		EmployeeID: int32(medHist.EmployeeID().Value()),
		VisitType:  medHist.VisitType().DisplayName(),
		VisitDate:  pgtype.Timestamptz{Time: medHist.VisitDate(), Valid: true},
		Diagnosis:  pgtype.Text{String: medHist.Diagnosis(), Valid: true},
		Treatment:  pgtype.Text{String: medHist.Treatment(), Valid: true},
		Condition:  pgtype.Text{String: medHist.Condition().DisplayName(), Valid: true},
	}

	if notes != "" {
		params.Notes = pgtype.Text{String: notes, Valid: true}
	} else {
		params.Notes = pgtype.Text{String: "", Valid: false}
	}

	return params
}

func ToUpdateParams(medHistory medical.MedicalHistory) sqlc.UpdateMedicalHistoryParams {
	return sqlc.UpdateMedicalHistoryParams{
		ID:         int32(medHistory.ID().Value()),
		PetID:      int32(medHistory.PetID().Value()),
		CustomerID: int32(medHistory.CustomerID().Value()),
		EmployeeID: int32(medHistory.EmployeeID().Value()),
		VisitDate:  pgtype.Timestamptz{Time: medHistory.VisitDate(), Valid: true},
		Diagnosis:  pgtype.Text{String: medHistory.Diagnosis(), Valid: true},
		Treatment:  pgtype.Text{String: medHistory.Treatment(), Valid: true},
		Notes: pgtype.Text{String: func() string {
			if medHistory.Notes() != nil {
				return *medHistory.Notes()
			} else {
				return ""
			}
		}(), Valid: medHistory.Notes() != nil},
		VisitType: medHistory.VisitType().DisplayName(),
		Condition: pgtype.Text{String: medHistory.Condition().DisplayName(), Valid: true},
	}
}
