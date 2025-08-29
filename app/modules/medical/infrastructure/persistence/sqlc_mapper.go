package persistence

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToDomain(sqlRow sqlc.MedicalHistory) (entity.MedicalHistory, error) {
	// Mapeo de Value Objects y Enums.
	medHistId, err := valueobject.NewMedHistoryID(sqlRow.ID)
	if err != nil {
		return entity.MedicalHistory{}, err
	}

	petId, err := valueobject.NewPetID(sqlRow.PetID)
	if err != nil {
		return entity.MedicalHistory{}, err
	}

	vetId, err := valueobject.NewVetID(sqlRow.VeterinarianID)
	if err != nil {
		return entity.MedicalHistory{}, err
	}

	visitType, err := enum.NewVisitType(sqlRow.VisitType)
	if err != nil {
		return entity.MedicalHistory{}, err
	}

	condition, err := enum.NewPetCondition(sqlRow.Condition.String)
	if err != nil {
		return entity.MedicalHistory{}, err
	}

	reason, err := enum.NewVisitReason("injury")
	if err != nil {
		return entity.MedicalHistory{}, err
	}

	var notes *string
	if sqlRow.Notes.Valid {
		notes = &sqlRow.Notes.String
	}

	medicalHistory := entity.NewMedicalHistory(
		medHistId,
		petId,
		int(sqlRow.OwnerID),
		reason,
		visitType,
		sqlRow.VisitDate.Time,
		notes,
		sqlRow.Diagnosis.String,
		sqlRow.Treatment.String,
		condition,
		vetId,
		sqlRow.CreatedAt.Time,
		sqlRow.UpdatedAt.Time,
	)

	return *medicalHistory, nil
}

func ToDomainList(medHistList []sqlc.MedicalHistory) ([]entity.MedicalHistory, error) {
	domainList := make([]entity.MedicalHistory, len(medHistList))

	for i, sqlRow := range medHistList {
		domainMedHist, err := ToDomain(sqlRow)
		if err != nil {
			return nil, err
		}
		domainList[i] = domainMedHist
	}

	return domainList, nil
}

func ToCreateParams(medHist entity.MedicalHistory) sqlc.CreateMedicalHistoryParams {
	var notes string
	if medHist.Notes != nil {
		notes = *medHist.Notes()
	}

	params := sqlc.CreateMedicalHistoryParams{
		PetID:          int32(medHist.PetID().GetValue()),
		OwnerID:        int32(medHist.OwnerID()),
		VeterinarianID: int32(medHist.VetID().GetValue()),
		VisitType:      medHist.VisitType().ToString(),
		VisitDate:      pgtype.Timestamptz{Time: medHist.VisitDate(), Valid: true},
		Diagnosis:      pgtype.Text{String: medHist.Diagnosis(), Valid: true},
		Treatment:      pgtype.Text{String: medHist.Treatment(), Valid: true},
		Condition:      pgtype.Text{String: medHist.Condition().ToString(), Valid: true},
	}

	if notes != "" {
		params.Notes = pgtype.Text{String: notes, Valid: true}
	} else {
		params.Notes = pgtype.Text{String: "", Valid: false}
	}

	return params
}
