package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

type MedicalHistory struct {
	id          valueobject.MedHistoryID
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
	createdAt   time.Time
	updatedAt   time.Time
}

func NewMedicalHistory(
	id valueobject.MedHistoryID,
	petID valueobject.PetID,
	ownerID valueobject.OwnerID,
	visitReason enum.VisitReason,
	visitType enum.VisitType,
	visitDate time.Time,
	notes *string,
	diagnosis string,
	treatment string,
	condition enum.PetCondition,
	vetID valueobject.VetID,
	createdAt time.Time,
	updateAt time.Time,
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

type MedicalHistoryBuilder struct {
	medicalHistory MedicalHistory
	errors         []error
}

func NewMedicalHistoryBuilder() *MedicalHistoryBuilder {
	return &MedicalHistoryBuilder{
		medicalHistory: MedicalHistory{},
		errors:         []error{},
	}
}

func (b *MedicalHistoryBuilder) WithID(id int) *MedicalHistoryBuilder {
	medHistoryID, err := valueobject.NewMedHistoryID(id)
	if err != nil {
		b.errors = append(b.errors, err)
		return b
	}

	b.medicalHistory.id = medHistoryID
	return b
}

func (b *MedicalHistoryBuilder) WithPetID(petID int) *MedicalHistoryBuilder {
	petIDVO, err := valueobject.NewPetID(petID)
	if err != nil {
		b.errors = append(b.errors, err)
		return b
	}

	b.medicalHistory.petID = petIDVO
	return b
}

func (b *MedicalHistoryBuilder) WithVetID(vetID int) *MedicalHistoryBuilder {
	vetIDVO, err := valueobject.NewVetID(vetID)
	if err != nil {
		b.errors = append(b.errors, err)
		return b
	}

	b.medicalHistory.vetID = vetIDVO
	return b
}

func (b *MedicalHistoryBuilder) WithOwnerID(ownerID int) *MedicalHistoryBuilder {
	if ownerID <= 0 {
		b.errors = append(b.errors, errors.New("ownerID must be positive"))
		return b
	}
	return b
}

func (b *MedicalHistoryBuilder) WithVisitReason(visitReason string) *MedicalHistoryBuilder {
	reason, err := enum.NewVisitReason(visitReason)
	if err != nil {
		b.errors = append(b.errors, err)
		return b
	}

	b.medicalHistory.visitReason = reason
	return b
}

func (b *MedicalHistoryBuilder) WithVisitType(visitType string) *MedicalHistoryBuilder {
	vt, err := enum.NewVisitType(visitType)
	if err != nil {
		b.errors = append(b.errors, err)
		return b
	}

	b.medicalHistory.visitType = vt
	return b
}

func (b *MedicalHistoryBuilder) WithCondition(condition string) *MedicalHistoryBuilder {
	cond, err := enum.NewPetCondition(condition)
	if err != nil {
		b.errors = append(b.errors, err)
		return b
	}
	b.medicalHistory.condition = cond
	return b
}

func (b *MedicalHistoryBuilder) WithVisitDate(date time.Time) *MedicalHistoryBuilder {
	if date.IsZero() {
		b.errors = append(b.errors, errors.New("visit date cannot be zero"))
		return b
	}
	if date.After(time.Now()) {
		b.errors = append(b.errors, errors.New("visit date cannot be in the future"))
		return b
	}

	b.medicalHistory.visitDate = date
	return b
}

func (b *MedicalHistoryBuilder) WithDiagnosis(diagnosis string) *MedicalHistoryBuilder {
	if diagnosis == "" {
		b.errors = append(b.errors, errors.New("diagnosis cannot be empty"))
		return b
	}

	b.medicalHistory.diagnosis = diagnosis
	return b
}

func (b *MedicalHistoryBuilder) WithTreatment(treatment string) *MedicalHistoryBuilder {
	if treatment == "" {
		b.errors = append(b.errors, errors.New("treatment cannot be empty"))
		return b
	}

	b.medicalHistory.treatment = treatment
	return b
}

func (b *MedicalHistoryBuilder) WithNotes(notes *string) *MedicalHistoryBuilder {
	b.medicalHistory.notes = notes
	return b
}

func (b *MedicalHistoryBuilder) WithTimestamps(createdAt, updatedAt time.Time) *MedicalHistoryBuilder {
	if createdAt.IsZero() {
		b.medicalHistory.createdAt = time.Now()
	}
	if updatedAt.IsZero() {
		b.medicalHistory.updatedAt = time.Now()
	}
	// Setters en la entidad
	return b
}

func (b *MedicalHistoryBuilder) Build() (MedicalHistory, error) {
	if len(b.errors) > 0 {
		return MedicalHistory{}, b.combineErrors()
	}

	return b.medicalHistory, nil
}

func (b *MedicalHistoryBuilder) combineErrors() error {
	var errorMessages []string
	for _, err := range b.errors {
		errorMessages = append(errorMessages, err.Error())
	}
	return errors.New(strings.Join(errorMessages, "; "))
}

func (b *MedicalHistoryBuilder) HasErrors() bool {
	return len(b.errors) > 0
}

func (b *MedicalHistoryBuilder) GetErrors() []error {
	return b.errors
}
