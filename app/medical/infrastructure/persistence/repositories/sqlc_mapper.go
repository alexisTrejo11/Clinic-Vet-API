package sqlcMedHistoryRepo

import (
	mhDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/domain"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToDomain(medHist sqlc.MedicalHistory) (mhDomain.MedicalHistory, error) {
	medHistId, err := mhDomain.NewMedHistoryId(medHist.ID)
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	petId, err := petDomain.NewPetId(medHist.PetID)
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	vetId, err := vetDomain.NewVeterinarianId(medHist.VeterinarianID)
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	visitType, err := mhDomain.NewVisitType(medHist.VisitType)
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	condition, err := mhDomain.NewPetCondition(medHist.Condition.String)
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	reason, err := mhDomain.NewVisitReason("Injury")
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	var notes string
	if medHist.Notes.Valid {
		notes = medHist.Notes.String
	} else {
		notes = ""
	}

	return mhDomain.MedicalHistory{
		Id:          medHistId,
		PetId:       petId,
		OwnerId:     int(medHist.OwnerID),
		VetId:       vetId,
		VisitDate:   medHist.VisitDate.Time,
		Diagnosis:   medHist.Diagnosis.String,
		Treatment:   medHist.Treatment.String,
		Notes:       &notes,
		VisitType:   visitType,
		VisitReason: reason,
		Condition:   condition,
		CreatedAt:   medHist.CreatedAt.Time,
		UpdatedAt:   medHist.UpdatedAt.Time,
	}, nil
}

func ToDomainList(medHistList []sqlc.MedicalHistory) ([]mhDomain.MedicalHistory, error) {
	domainList := make([]mhDomain.MedicalHistory, len(medHistList))

	for i, medHist := range medHistList {
		domainMedHist, err := ToDomain(medHist)
		if err != nil {
			return nil, err
		}
		domainList[i] = domainMedHist
	}

	return domainList, nil
}

func ToCreateParams(medHist mhDomain.MedicalHistory) sqlc.CreateMedicalHistoryParams {
	var notes string
	if *medHist.Notes != *medHist.Notes {
		notes = *medHist.Notes
	}

	params := sqlc.CreateMedicalHistoryParams{
		PetID:          int32(medHist.PetId.GetValue()),
		OwnerID:        int32(medHist.OwnerId),
		VeterinarianID: int32(medHist.VetId.GetValue()),
		VisitType:      medHist.VisitType.ToString(),
		VisitDate:      pgtype.Timestamptz{Time: medHist.VisitDate, Valid: true},
		Diagnosis:      pgtype.Text{String: medHist.Diagnosis, Valid: true},
		Treatment:      pgtype.Text{String: medHist.Treatment, Valid: true},
		Condition:      pgtype.Text{String: medHist.Condition.ToString(), Valid: true},
	}

	if notes != "" {
		params.Notes = pgtype.Text{String: notes, Valid: true}
	} else {
		params.Notes = pgtype.Text{String: "", Valid: false}
	}

	return params
}
