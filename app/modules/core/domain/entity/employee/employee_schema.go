// Package employee defines the Employee entity and its business logic.
package employee

import (
	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	domainerr "clinic-vet-api/app/modules/core/error"
	"context"
	"fmt"
	"time"
)

type Employee struct {
	base.Entity[valueobject.EmployeeID]
	base.Person
	photo           string
	licenseNumber   string
	specialty       enum.VetSpecialty
	yearsExperience int
	consultationFee *valueobject.Money
	isActive        bool
	userID          *valueobject.UserID
	schedule        *valueobject.Schedule
}

func (v *Employee) ID() valueobject.EmployeeID {
	return v.Entity.ID()
}

func (v *Employee) Name() valueobject.PersonName {
	return v.Person.Name()
}

func (v *Employee) Photo() string {
	return v.photo
}

func (v *Employee) LicenseNumber() string {
	return v.licenseNumber
}

func (v *Employee) Specialty() enum.VetSpecialty {
	return v.specialty
}

func (v *Employee) YearsExperience() int {
	return v.yearsExperience
}

func (v *Employee) ConsultationFee() *valueobject.Money {
	return v.consultationFee
}

func (v *Employee) IsActive() bool {
	return v.isActive
}

func (v *Employee) UserID() *valueobject.UserID {
	return v.userID
}

func (v *Employee) Schedule() *valueobject.Schedule {
	return v.schedule
}

func (v *Employee) AssignUser(ctx context.Context, userID valueobject.UserID) error {
	if v.userID != nil {
		return domainerr.ConflictError(ctx, "userID", fmt.Sprintf("employee %s is already assigned to a user", v.ID().String()), "assinging user to employee")
	}
	v.userID = &userID

	return nil
}

func (v *Employee) SetID(id valueobject.EmployeeID) {
	v.Entity.SetID(id)
}

func (v *Employee) CreatedAt() time.Time {
	return v.Entity.CreatedAt()
}

func (v *Employee) UpdatedAt() time.Time {
	return v.Entity.UpdatedAt()
}

func (v *Employee) SetTimeStamps(createdAt, updatedAt time.Time) {
	v.Entity.SetTimeStamps(createdAt, updatedAt)
}
