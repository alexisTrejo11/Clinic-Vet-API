// Package employee defines the Employee entity and its business logic.
package employee

import (
	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type Employee struct {
	base.Entity[vo.EmployeeID]
	base.Person
	photo           string
	licenseNumber   string
	specialty       enum.VetSpecialty
	yearsExperience int
	consultationFee *vo.Money
	isActive        bool
	userID          *vo.UserID
	schedule        *vo.Schedule
}

type EmployeeBuilder struct{ emp *Employee }

func NewEmployeeBuilder() *EmployeeBuilder { return &EmployeeBuilder{emp: &Employee{}} }

func (b *EmployeeBuilder) WithID(id vo.EmployeeID) *EmployeeBuilder {
	b.emp.SetID(id)
	return b
}

func (b *EmployeeBuilder) WithPersonalData(person base.Person) *EmployeeBuilder {
	b.emp.Person = person
	return b
}

func (b *EmployeeBuilder) WithName(name valueobject.PersonName) *EmployeeBuilder {
	b.emp.SetName(name)
	return b
}

func (b *EmployeeBuilder) WithPhoto(photo string) *EmployeeBuilder {
	b.emp.photo = photo
	return b
}

func (b *EmployeeBuilder) WithLicenseNumber(licenseNumber string) *EmployeeBuilder {
	b.emp.licenseNumber = licenseNumber
	return b
}

func (b *EmployeeBuilder) WithSpecialty(specialty enum.VetSpecialty) *EmployeeBuilder {
	b.emp.specialty = specialty
	return b
}

func (b *EmployeeBuilder) WithYearsExperience(years int) *EmployeeBuilder {
	b.emp.yearsExperience = years
	return b
}

func (b *EmployeeBuilder) WithIsActive(isActive bool) *EmployeeBuilder {
	b.emp.isActive = isActive
	return b
}

func (b *EmployeeBuilder) WithUserID(userID *valueobject.UserID) *EmployeeBuilder {
	b.emp.userID = userID
	return b
}

func (b *EmployeeBuilder) WithSchedule(schedule *valueobject.Schedule) *EmployeeBuilder {
	b.emp.schedule = schedule
	return b
}

func (b *EmployeeBuilder) WithTimestamps(createdAt, updatedAt time.Time) *EmployeeBuilder {
	b.emp.SetTimeStamps(createdAt, updatedAt)
	return b
}

func (b *EmployeeBuilder) Build() *Employee {
	return b.emp
}

func (v *Employee) ID() vo.EmployeeID                              { return v.Entity.ID() }
func (v *Employee) Name() vo.PersonName                            { return v.Person.Name() }
func (v *Employee) Photo() string                                  { return v.photo }
func (v *Employee) LicenseNumber() string                          { return v.licenseNumber }
func (v *Employee) Specialty() enum.VetSpecialty                   { return v.specialty }
func (v *Employee) YearsExperience() int                           { return v.yearsExperience }
func (v *Employee) ConsultationFee() *vo.Money                     { return v.consultationFee }
func (v *Employee) IsActive() bool                                 { return v.isActive }
func (v *Employee) IsWithinWorkdayBreak(start, end time.Time) bool { return false }
func (v *Employee) UserID() *vo.UserID                             { return v.userID }
func (v *Employee) Schedule() *vo.Schedule                         { return v.schedule }
func (v *Employee) CreatedAt() time.Time                           { return v.Entity.CreatedAt() }
func (v *Employee) UpdatedAt() time.Time                           { return v.Entity.UpdatedAt() }
func (v *Employee) SetID(id vo.EmployeeID)                         { v.Entity.SetID(id) }
