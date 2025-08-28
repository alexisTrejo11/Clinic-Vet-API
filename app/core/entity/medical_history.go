package entity

import (
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

type MedicalHistory struct {
	Id          valueobject.MedHistoryID
	PetId       valueobject.PetID
	OwnerId     int
	VisitReason enum.VisitReason
	VisitType   enum.VisitType
	VisitDate   time.Time
	Notes       *string
	Diagnosis   string
	Treatment   string
	Condition   enum.PetCondition
	VetId       valueobject.VetID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (mh *MedicalHistory) ValidateBusinessRules() error {
	if mh.VisitDate.IsZero() {
		return errors.New("invalid date")
	}

	if mh.VisitDate.After(time.Now().AddDate(1, 0, 0)) {
		return errors.New("date cannot be one year in the future")
	}

	if mh.VisitDate.Before(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return errors.New("date cannot be before the year 2015")
	}

	return nil
}

func (mh *MedicalHistory) SetId(id int) {
	mh.Id, _ = valueobject.NewMedHistoryID(id)
}
