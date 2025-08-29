package entity

import (
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

type MedicalHistory struct {
	id          valueobject.MedHistoryID
	petID       valueobject.PetID
	ownerID     int
	visitReason enum.VisitReason
	visitType   enum.VisitType
	visitDate   time.Time
	notes       *string
	diagnosis   string
	treatment   string
	condition   enum.PetCondition
	vetID       valueobject.VetID
	createdAt   time.Time
	updatedAt   time.Time
}

func NewMedicalHistory(
	id valueobject.MedHistoryID,
	petID valueobject.PetID,
	ownerID int,
	visitReason enum.VisitReason,
	visitType enum.VisitType,
	visitDate time.Time,
	notes *string,
	diagnosis string,
	treatment string,
	condition enum.PetCondition,
	vetID valueobject.VetID,
	created_at time.Time,
	update_at time.Time,
) *MedicalHistory {
	now := time.Now()
	return &MedicalHistory{
		id:          id,
		petID:       petID,
		ownerID:     ownerID,
		visitReason: visitReason,
		visitType:   visitType,
		visitDate:   visitDate,
		notes:       notes,
		diagnosis:   diagnosis,
		treatment:   treatment,
		condition:   condition,
		vetID:       vetID,
		createdAt:   now,
		updatedAt:   now,
	}
}

func (mh *MedicalHistory) ID() valueobject.MedHistoryID {
	return mh.id
}

func (mh *MedicalHistory) PetID() valueobject.PetID {
	return mh.petID
}

func (mh *MedicalHistory) OwnerID() int {
	return mh.ownerID
}

func (mh *MedicalHistory) VisitReason() enum.VisitReason {
	return mh.visitReason
}

func (mh *MedicalHistory) VisitType() enum.VisitType {
	return mh.visitType
}

func (mh *MedicalHistory) VisitDate() time.Time {
	return mh.visitDate
}

func (mh *MedicalHistory) Notes() *string {
	return mh.notes
}

func (mh *MedicalHistory) Diagnosis() string {
	return mh.diagnosis
}

func (mh *MedicalHistory) Treatment() string {
	return mh.treatment
}

func (mh *MedicalHistory) Condition() enum.PetCondition {
	return mh.condition
}

func (mh *MedicalHistory) VetID() valueobject.VetID {
	return mh.vetID
}

func (mh *MedicalHistory) CreatedAt() time.Time {
	return mh.createdAt
}

func (mh *MedicalHistory) UpdatedAt() time.Time {
	return mh.updatedAt
}

func (mh *MedicalHistory) ValidateBusinessRules() error {
	if mh.VisitDate().IsZero() {
		return errors.New("invalid date")
	}

	if mh.VisitDate().After(time.Now().AddDate(1, 0, 0)) {
		return errors.New("date cannot be one year in the future")
	}

	if mh.VisitDate().Before(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return errors.New("date cannot be before the year 2015")
	}

	return nil
}

func (mh *MedicalHistory) SetID(id int) {
	mh.id, _ = valueobject.NewMedHistoryID(id)
}
