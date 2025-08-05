package mhDomain

import (
	"errors"
	"time"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

type MedicalHistory struct {
	Id          MedHistoryId
	PetId       petDomain.PetId
	OwnerId     int
	VisitReason VisitReason
	VisitType   VisitType
	VisitDate   time.Time
	Notes       *string
	Diagnosis   string
	Treatment   string
	Condition   PetCondition
	VetId       vetDomain.VetId
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
	mh.Id, _ = NewMedHistoryId(id)
}
