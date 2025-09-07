package medical

import (
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type MedicalHistoryOptions func(*MedicalHistory) error

func WithVisitReason(reason enum.VisitReason) MedicalHistoryOptions {
	return func(mh *MedicalHistory) error {
		mh.visitReason = reason
		return nil
	}
}

func WithVisitType(vtype enum.VisitType) MedicalHistoryOptions {
	return func(mh *MedicalHistory) error {
		mh.visitType = vtype
		return nil
	}
}

func WithVisitDate(vdate time.Time) MedicalHistoryOptions {
	return func(mh *MedicalHistory) error {
		mh.visitDate = vdate
		return nil
	}
}

func WithDiagnosis(diagnosis string) MedicalHistoryOptions {
	return func(mh *MedicalHistory) error {
		mh.diagnosis = diagnosis
		return nil
	}
}

func WithTreatment(treatment string) MedicalHistoryOptions {
	return func(mh *MedicalHistory) error {
		mh.treatment = treatment
		return nil
	}
}

func WithNotes(notes string) MedicalHistoryOptions {
	return func(mh *MedicalHistory) error {
		mh.notes = &notes
		return nil
	}
}

func WithCondition(condition enum.PetCondition) MedicalHistoryOptions {
	return func(mh *MedicalHistory) error {
		mh.condition = condition
		return nil
	}
}

func NewMedicalHistory(
	medhistoryID valueobject.MedHistoryID,
	petID valueobject.PetID,
	ownerID valueobject.OwnerID,
	vetID valueobject.VetID,
	opts ...MedicalHistoryOptions,
) (*MedicalHistory, error) {
	mh := &MedicalHistory{
		Entity:  base.NewEntity(medhistoryID),
		petID:   petID,
		ownerID: ownerID,
		vetID:   vetID,
	}

	for _, opt := range opts {
		if err := opt(mh); err != nil {
			return nil, err
		}
	}

	if err := mh.Validate(); err != nil {
		return nil, err
	}

	return mh, nil
}

func (mh *MedicalHistory) Validate() error {
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
