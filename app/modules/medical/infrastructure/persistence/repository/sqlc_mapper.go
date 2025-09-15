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

func ToDomain(sqlRow sqlc.MedicalHistory) (medical.MedicalHistory, error) {
	medHistID, err := valueobject.NewMedHistoryID(int(sqlRow.ID))
	if err != nil {
		return medical.MedicalHistory{}, fmt.Errorf("invalid medical history ID: %w", err)
	}

	petID, err := valueobject.NewPetID(int(sqlRow.PetID))
	if err != nil {
		return medical.MedicalHistory{}, fmt.Errorf("invalid pet ID: %w", err)
	}

	vetID, err := valueobject.NewVetID(int(sqlRow.VeterinarianID))
	if err != nil {
		return medical.MedicalHistory{}, fmt.Errorf("invalid vet ID: %w", err)
	}

	ownerID, err := valueobject.NewOwnerID(int(sqlRow.OwnerID))
	if err != nil {
		return medical.MedicalHistory{}, fmt.Errorf("invalid owner ID: %w", err)
	}

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
		ownerID,
		vetID,
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

func ToDomainList(medHistList []sqlc.MedicalHistory) ([]medical.MedicalHistory, error) {
	domainList := make([]medical.MedicalHistory, len(medHistList))

	for i, sqlRow := range medHistList {
		domainMedHist, err := ToDomain(sqlRow)
		if err != nil {
			return nil, err
		}
		domainList[i] = domainMedHist
	}

	return domainList, nil
}

func ToCreateParams(medHist medical.MedicalHistory) sqlc.CreateMedicalHistoryParams {
	var notes string
	if medHist.Notes() != nil {
		notes = *medHist.Notes()
	}

	params := sqlc.CreateMedicalHistoryParams{
		PetID:          int32(medHist.PetID().Value()),
		OwnerID:        int32(medHist.OwnerID().Value()),
		VeterinarianID: int32(medHist.VetID().Value()),
		VisitType:      medHist.VisitType().DisplayName(),
		VisitDate:      pgtype.Timestamptz{Time: medHist.VisitDate(), Valid: true},
		Diagnosis:      pgtype.Text{String: medHist.Diagnosis(), Valid: true},
		Treatment:      pgtype.Text{String: medHist.Treatment(), Valid: true},
		Condition:      pgtype.Text{String: medHist.Condition().DisplayName(), Valid: true},
	}

	if notes != "" {
		params.Notes = pgtype.Text{String: notes, Valid: true}
	} else {
		params.Notes = pgtype.Text{String: "", Valid: false}
	}

	return params
}

func entityToUpdateParam(medHistory medical.MedicalHistory, notes pgtype.Text) sqlc.UpdateMedicalHistoryParams {
	return sqlc.UpdateMedicalHistoryParams{
		ID:             int32(medHistory.ID().Value()),
		PetID:          int32(medHistory.PetID().Value()),
		OwnerID:        int32(medHistory.OwnerID().Value()),
		VeterinarianID: int32(medHistory.VetID().Value()),
		VisitDate:      pgtype.Timestamptz{Time: medHistory.VisitDate(), Valid: true},
		Diagnosis:      pgtype.Text{String: medHistory.Diagnosis(), Valid: true},
		Treatment:      pgtype.Text{String: medHistory.Treatment(), Valid: true},
		Notes:          notes,
		VisitType:      medHistory.VisitType().DisplayName(),
		Condition:      pgtype.Text{String: medHistory.Condition().DisplayName(), Valid: true},
	}
}
