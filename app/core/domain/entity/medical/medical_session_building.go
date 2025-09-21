package medical

import (
	"context"
	"time"

	"clinic-vet-api/app/core/domain/entity/base"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	domainerr "clinic-vet-api/app/core/error"
)

type MedicalSessionOptions func(*MedicalSession) error

func WithVisitReason(reason enum.VisitReason) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.visitReason = reason
		return nil
	}
}

func WithVisitType(vtype enum.VisitType) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.visitType = vtype
		return nil
	}
}

func WithVisitDate(vdate time.Time) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.visitDate = vdate
		return nil
	}
}

func WithDiagnosis(diagnosis string) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.diagnosis = diagnosis
		return nil
	}
}

func WithTreatment(treatment string) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.treatment = treatment
		return nil
	}
}

func WithNotes(notes *string) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.notes = notes
		return nil
	}
}

func WithEmployeeID(eid valueobject.EmployeeID) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.employeeID = eid
		return nil
	}
}

// TODO: Implement Field
func WithSymptoms(symptoms []string) MedicalSessionOptions {
	return func(mh *MedicalSession) error {

		return nil
	}
}

func WithFollowUpDate(followUpDate *time.Time) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		return nil
	}
}
func WithTimestamp(createdAt, updatedAt time.Time) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.SetTimeStamps(createdAt, updatedAt)
		return nil
	}
}

func WithCondition(condition enum.PetCondition) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.condition = condition
		return nil
	}
}

func WithWeight(weight *valueobject.Decimal) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.weight = weight
		return nil
	}
}

func WithHeartRate(heartRate *int) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.heartRate = heartRate
		return nil
	}
}

func WithRespiratoryRate(respiratoryRate *int) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.respiratoryRate = respiratoryRate
		return nil
	}
}

func WithTemperature(temperature *valueobject.Decimal) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.temperature = temperature
		return nil
	}
}

func WithMedications(medications []string) MedicalSessionOptions {
	return func(mh *MedicalSession) error {
		mh.Medications = medications
		return nil
	}
}

func NewMedicalSession(
	medhistoryID valueobject.MedSessionID,
	petID valueobject.PetID,
	customerID valueobject.CustomerID,
	employeeID valueobject.EmployeeID,
	opts ...MedicalSessionOptions,
) (*MedicalSession, error) {
	mh := &MedicalSession{
		Entity: base.NewEntity(medhistoryID, time.Now(), time.Now(), 1),
		PetSessionSummary: PetSessionSummary{
			petID: petID,
		},
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

func CreateMedicalSession(
	ctx context.Context,
	petID valueobject.PetID,
	customerID valueobject.CustomerID,
	employeeID valueobject.EmployeeID,
	opts ...MedicalSessionOptions,
) (*MedicalSession, error) {
	mh := &MedicalSession{
		Entity: base.CreateEntity(valueobject.MedSessionID{}),
		PetSessionSummary: PetSessionSummary{
			petID: petID,
		},
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

func (mh *MedicalSession) Validate(ctx context.Context) error {
	operation := "Medical-Session Validate"
	if mh.VisitDate().IsZero() {
		return domainerr.MissingFieldError(ctx, "visitDate", "visitDate cannot be zero", operation)
	}

	if mh.VisitReason() == "" {
		return domainerr.MissingFieldError(ctx, "visitReason", "visitReason cannot be empty", operation)
	}

	if mh.VisitType() == "" {
		return domainerr.MissingFieldError(ctx, "visitType", "visitType cannot be empty", operation)
	}

	if !mh.VisitReason().IsValid() {
		return domainerr.InvalidEnumValue(ctx, "visitReason", "medical history", string(mh.VisitReason()), operation)
	}

	if !mh.VisitType().IsValid() {
		return domainerr.InvalidEnumValue(ctx, "visitType", "medical history", string(mh.VisitType()), operation)
	}

	if mh.Weight() != nil {
		if mh.Weight().IsNegative() || mh.Weight().IsZero() {
			return domainerr.InvalidFieldValue(ctx, "weight", "medical history", "weight must be a positive value", operation)
		}
	}

	if mh.Temperature() != nil {
		if mh.Temperature().IsNegative() || mh.Temperature().IsZero() {
			return domainerr.InvalidFieldValue(ctx, "temperature", "medical history", "temperature must be a positive value", operation)
		}
	}

	if mh.HeartRate() != nil {
		if *mh.HeartRate() <= 0 {
			return domainerr.InvalidFieldValue(ctx, "heartRate", "medical history", "heartRate must be a positive value", operation)
		}
	}

	if mh.RespiratoryRate() != nil {
		if *mh.RespiratoryRate() <= 0 {
			return domainerr.InvalidFieldValue(ctx, "respiratoryRate", "medical history", "respiratoryRate must be a positive value", operation)
		}
	}

	if mh.VisitDate().After(time.Now().AddDate(1, 0, 0)) {
		return domainerr.BusinessRuleError(ctx, "visitDate", "medical history", "visitDate cannot be 1 year in the future", operation)
	}

	if mh.VisitDate().Before(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return domainerr.BusinessRuleError(ctx, "visitDate", "medical history", "visitDate cannot be before year 2015", operation)
	}

	if mh.Condition() != "" && !mh.Condition().IsValid() {
		return domainerr.InvalidEnumValue(ctx, "condition", "medical history", string(mh.Condition()), operation)
	}

	if mh.Medications != nil {
		if len(mh.Medications) > 100 {
			return domainerr.InvalidFieldValue(ctx, "medications", "medical history", "medications cannot have more than 100 items", operation)
		}
	}

	return nil
}
