// Package medical contains entities related to medical information.
package medical

import (
	"time"

	"clinic-vet-api/app/core/domain/entity/base"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
)

type MedicalSession struct {
	base.Entity[valueobject.MedSessionID]
	customerID  valueobject.CustomerID
	visitReason enum.VisitReason
	visitType   enum.VisitType
	visitDate   time.Time
	notes       *string
	employeeID  valueobject.EmployeeID
	PetSessionSummary
}

type PetSessionSummary struct {
	petID           valueobject.PetID
	weight          *valueobject.Decimal
	heartRate       *int
	respiratoryRate *int
	temperature     *valueobject.Decimal
	diagnosis       string
	treatment       string
	condition       enum.PetCondition
	medications     []string
}

func (mh *MedicalSession) ID() valueobject.MedSessionID {
	return mh.Entity.ID()
}

func (mh *MedicalSession) Medications() []string {
	return mh.medications
}

func (mh *MedicalSession) PetID() valueobject.PetID {
	return mh.petID
}

func (mh *MedicalSession) CustomerID() valueobject.CustomerID {
	return mh.customerID
}

func (mh *MedicalSession) VisitReason() enum.VisitReason {
	return mh.visitReason
}

func (mh *MedicalSession) VisitType() enum.VisitType {
	return mh.visitType
}

func (mh *MedicalSession) VisitDate() time.Time {
	return mh.visitDate
}

func (mh *MedicalSession) Notes() *string {
	return mh.notes
}

func (mh *MedicalSession) Diagnosis() string {
	return mh.diagnosis
}

func (mh *MedicalSession) Treatment() string {
	return mh.treatment
}

func (mh *MedicalSession) Condition() enum.PetCondition {
	return mh.condition
}

func (mh *MedicalSession) EmployeeID() valueobject.EmployeeID {
	return mh.employeeID
}

func (mh *MedicalSession) CreatedAt() time.Time {
	return mh.Entity.CreatedAt()
}

func (mh *MedicalSession) UpdatedAt() time.Time {
	return mh.Entity.UpdatedAt()
}

func (mh *MedicalSession) HeartRate() *int {
	return mh.heartRate
}

func (mh *MedicalSession) RespiratoryRate() *int {
	return mh.respiratoryRate
}

func (mh *MedicalSession) Temperature() *valueobject.Decimal {
	return mh.temperature
}

func (mh *MedicalSession) Weight() *valueobject.Decimal {
	return mh.weight
}
