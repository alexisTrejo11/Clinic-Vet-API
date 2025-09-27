// Package medical contains entities related to medical information.
package medical

import (
	"context"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	domainerr "clinic-vet-api/app/modules/core/error"
)

type MedicalSession struct {
	base.Entity[vo.MedSessionID]
	customerID vo.CustomerID
	visitType  enum.VisitType
	service    enum.ClinicService
	visitDate  time.Time
	notes      *string
	employeeID vo.EmployeeID
	petDetails PetSessionSummary
}

type PetSessionSummary struct {
	petID           vo.PetID
	weight          *vo.Decimal
	heartRate       *int
	respiratoryRate *int
	temperature     *vo.Decimal
	diagnosis       string
	treatment       string
	condition       enum.PetCondition
	medications     []string
	followUpDate    *time.Time
	symptoms        []string
}

type PetSessionSummaryBuilder struct{ petSession *PetSessionSummary }

func NewPetSessionSummaryBuilder() *PetSessionSummaryBuilder {
	return &PetSessionSummaryBuilder{petSession: &PetSessionSummary{
		medications: []string{},
		symptoms:    []string{},
	}}
}

type MedicalSessionBuilder struct{ medSession *MedicalSession }

func NewMedicalSessionBuilder() *MedicalSessionBuilder {
	return &MedicalSessionBuilder{medSession: &MedicalSession{
		petDetails: PetSessionSummary{
			medications: []string{},
			symptoms:    []string{},
		},
	}}
}

func (b *MedicalSessionBuilder) WithID(id vo.MedSessionID) *MedicalSessionBuilder {
	b.medSession.Entity = base.NewEntity(id, nil, nil, 0)
	return b
}

func (b *MedicalSessionBuilder) WithCustomerID(customerID vo.CustomerID) *MedicalSessionBuilder {
	b.medSession.customerID = customerID
	return b
}

func (b *MedicalSessionBuilder) WithVisitType(vtype enum.VisitType) *MedicalSessionBuilder {
	b.medSession.visitType = vtype
	return b
}

func (b *MedicalSessionBuilder) WithVisitDate(vdate time.Time) *MedicalSessionBuilder {
	b.medSession.visitDate = vdate
	return b
}

func (b *MedicalSessionBuilder) WithService(service enum.ClinicService) *MedicalSessionBuilder {
	b.medSession.service = service
	return b
}

func (b *MedicalSessionBuilder) WithNotes(notes *string) *MedicalSessionBuilder {
	b.medSession.notes = notes
	return b
}

func (b *MedicalSessionBuilder) WithEmployeeID(employeeID vo.EmployeeID) *MedicalSessionBuilder {
	b.medSession.employeeID = employeeID
	return b
}

func (b *MedicalSessionBuilder) WithPetDetails(petDetails PetSessionSummary) *MedicalSessionBuilder {
	b.medSession.petDetails = petDetails
	return b
}

func (b *MedicalSessionBuilder) Build() *MedicalSession {
	return b.medSession
}

func (b *PetSessionSummaryBuilder) WithPetID(petID vo.PetID) *PetSessionSummaryBuilder {
	b.petSession.petID = petID
	return b
}

func (b *PetSessionSummaryBuilder) WithWeight(weight *vo.Decimal) *PetSessionSummaryBuilder {
	b.petSession.weight = weight
	return b
}

func (b *PetSessionSummaryBuilder) WithTemperature(temperature *vo.Decimal) *PetSessionSummaryBuilder {
	b.petSession.temperature = temperature
	return b
}

func (b *PetSessionSummaryBuilder) WithHeartRate(heartRate *int) *PetSessionSummaryBuilder {
	b.petSession.heartRate = heartRate
	return b
}

func (b *PetSessionSummaryBuilder) WithRespiratoryRate(respiratoryRate *int) *PetSessionSummaryBuilder {
	b.petSession.respiratoryRate = respiratoryRate
	return b
}

func (b *PetSessionSummaryBuilder) WithDiagnosis(diagnosis string) *PetSessionSummaryBuilder {
	b.petSession.diagnosis = diagnosis
	return b
}

func (b *PetSessionSummaryBuilder) WithTreatment(treatment string) *PetSessionSummaryBuilder {
	b.petSession.treatment = treatment
	return b
}

func (b *PetSessionSummaryBuilder) WithCondition(condition enum.PetCondition) *PetSessionSummaryBuilder {
	b.petSession.condition = condition
	return b
}

func (b *PetSessionSummaryBuilder) WithMedications(medications []string) *PetSessionSummaryBuilder {
	b.petSession.medications = medications
	return b
}

func (b *PetSessionSummaryBuilder) WithFollowUpDate(followUpDate *time.Time) *PetSessionSummaryBuilder {
	b.petSession.followUpDate = followUpDate
	return b
}

func (b *PetSessionSummaryBuilder) WithSymptoms(symptoms []string) *PetSessionSummaryBuilder {
	b.petSession.symptoms = symptoms
	return b
}

func (b *PetSessionSummaryBuilder) Build() *PetSessionSummary {
	return b.petSession
}

// Getters

func (mh *MedicalSession) ID() vo.MedSessionID           { return mh.Entity.ID() }
func (mh *MedicalSession) PetDetails() PetSessionSummary { return mh.petDetails }
func (mh *MedicalSession) CustomerID() vo.CustomerID     { return mh.customerID }
func (mh *MedicalSession) VisitDate() time.Time          { return mh.visitDate }
func (mh *MedicalSession) Service() enum.ClinicService   { return mh.service }
func (mh *MedicalSession) Notes() *string                { return mh.notes }
func (mh *MedicalSession) VisitType() enum.VisitType     { return mh.visitType }
func (mh *MedicalSession) EmployeeID() vo.EmployeeID     { return mh.employeeID }
func (mh *MedicalSession) CreatedAt() time.Time          { return mh.Entity.CreatedAt() }
func (mh *MedicalSession) UpdatedAt() time.Time          { return mh.Entity.UpdatedAt() }

func (ps PetSessionSummary) PetID() vo.PetID              { return ps.petID }
func (ps PetSessionSummary) Weight() *vo.Decimal          { return ps.weight }
func (ps PetSessionSummary) Temperature() *vo.Decimal     { return ps.temperature }
func (ps PetSessionSummary) HeartRate() *int              { return ps.heartRate }
func (ps PetSessionSummary) RespiratoryRate() *int        { return ps.respiratoryRate }
func (ps PetSessionSummary) Diagnosis() string            { return ps.diagnosis }
func (ps PetSessionSummary) Treatment() string            { return ps.treatment }
func (ps PetSessionSummary) Condition() enum.PetCondition { return ps.condition }
func (ps PetSessionSummary) Medications() []string        { return ps.medications }
func (ps PetSessionSummary) FollowUpDate() *time.Time     { return ps.followUpDate }
func (ps PetSessionSummary) Symptoms() []string           { return ps.symptoms }

// Validator

func (mh *MedicalSession) ValidatePersist(ctx context.Context) error {
	operation := "Medical-Session Validate"
	if mh.visitDate.IsZero() {
		return domainerr.MissingFieldError(ctx, "visitDate", "visitDate cannot be zero", operation)
	}

	if !mh.service.IsValid() {
		return domainerr.InvalidEnumValue(ctx, "service", "medical history", string(mh.service), operation)
	}

	if !mh.visitType.IsValid() {
		return domainerr.InvalidEnumValue(ctx, "visitType", "medical history", string(mh.visitType), operation)
	}

	if mh.visitDate.After(time.Now().AddDate(1, 0, 0)) {
		return domainerr.BusinessRuleError(ctx, "visitDate", "medical history", "visitDate cannot be 1 year in the future", operation)
	}

	if mh.visitDate.Before(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return domainerr.BusinessRuleError(ctx, "visitDate", "medical history", "visitDate cannot be before year 2015", operation)
	}

	if err := mh.petDetails.Validate(ctx); err != nil {
		return err
	}

	return nil
}

func (ps *PetSessionSummary) Validate(ctx context.Context) error {
	operation := "Pet-Session-Summary Validate"
	if ps.condition != "" && !ps.condition.IsValid() {
		return domainerr.InvalidEnumValue(ctx, "condition", "medical history", string(ps.condition), operation)
	}

	if ps.weight != nil {
		if ps.weight.IsNegative() || ps.weight.IsZero() {
			return domainerr.InvalidFieldValue(ctx, "weight", "medical history", "weight must be a positive value", operation)
		}
	}

	if ps.temperature != nil {
		if ps.temperature.IsNegative() || ps.temperature.IsZero() {
			return domainerr.InvalidFieldValue(ctx, "temperature", "medical history", "temperature must be a positive value", operation)
		}
	}

	if ps.heartRate != nil {
		if *ps.heartRate <= 0 {
			return domainerr.InvalidFieldValue(ctx, "heartRate", "medical history", "heartRate must be a positive value", operation)
		}
	}

	if ps.respiratoryRate != nil {
		if *ps.respiratoryRate <= 0 {
			return domainerr.InvalidFieldValue(ctx, "respiratoryRate", "medical history", "respiratoryRate must be a positive value", operation)
		}
	}

	if len(ps.medications) > 100 {
		return domainerr.InvalidFieldValue(ctx, "medications", "medical history", "medications cannot have more than 100 items", operation)
	}

	return nil
}
