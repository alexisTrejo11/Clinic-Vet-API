package medical

import (
	"context"
	"time"

	"clinic-vet-api/app/core/domain/entity/base"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	domainerr "clinic-vet-api/app/core/error"
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

func WithTimestamp(createdAt, updatedAt time.Time) MedicalHistoryOptions {
	return func(mh *MedicalHistory) error {
		mh.SetTimeStamps(createdAt, updatedAt)
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
	customerID valueobject.CustomerID,
	employeeID valueobject.EmployeeID,
	opts ...MedicalHistoryOptions,
) (*MedicalHistory, error) {
	mh := &MedicalHistory{
		Entity:     base.NewEntity(medhistoryID, time.Now(), time.Now(), 1),
		petID:      petID,
		customerID: customerID,
		employeeID: employeeID,
	}

	for _, opt := range opts {
		if err := opt(mh); err != nil {
			return nil, err
		}
	}

	return mh, nil
}

func CreateMedicalHistory(
	ctx context.Context,
	petID valueobject.PetID,
	customerID valueobject.CustomerID,
	employeeID valueobject.EmployeeID,
	opts ...MedicalHistoryOptions,
) (*MedicalHistory, error) {
	mh := &MedicalHistory{
		Entity:     base.CreateEntity(valueobject.MedHistoryID{}),
		petID:      petID,
		customerID: customerID,
		employeeID: employeeID,
	}

	for _, opt := range opts {
		if err := opt(mh); err != nil {
			return nil, err
		}
	}

	if err := mh.Validate(ctx); err != nil {
		return nil, err
	}

	return mh, nil
}

func (mh *MedicalHistory) Validate(ctx context.Context) error {
	operation := "Medical-History Validate"
	if mh.VisitDate().IsZero() {
		return domainerr.MissingFieldError(ctx, "visitDate", "visitDate cannot be zero", operation)
	}

	if mh.VisitDate().After(time.Now().AddDate(1, 0, 0)) {
		return domainerr.BusinessRuleError(ctx, "visitDate", "medical history", "visitDate cannot be 1 year in the future", operation)
	}

	if mh.VisitDate().Before(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return domainerr.BusinessRuleError(ctx, "visitDate", "medical history", "visitDate cannot be before year 2015", operation)
	}

	return nil
}
