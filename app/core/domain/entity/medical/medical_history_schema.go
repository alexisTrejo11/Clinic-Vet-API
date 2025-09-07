// Package medical contains entities related to medical information.
package medical

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type MedicalHistory struct {
	base.Entity[valueobject.MedHistoryID]
	petID       valueobject.PetID
	ownerID     valueobject.OwnerID
	visitReason enum.VisitReason
	visitType   enum.VisitType
	visitDate   time.Time
	notes       *string
	diagnosis   string
	treatment   string
	condition   enum.PetCondition
	vetID       valueobject.VetID
}

func (mh *MedicalHistory) ID() valueobject.MedHistoryID {
	return mh.Entity.ID()
}

func (mh *MedicalHistory) PetID() valueobject.PetID {
	return mh.petID
}

func (mh *MedicalHistory) OwnerID() valueobject.OwnerID {
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
	return mh.Entity.CreatedAt()
}

func (mh *MedicalHistory) UpdatedAt() time.Time {
	return mh.Entity.UpdatedAt()
}
