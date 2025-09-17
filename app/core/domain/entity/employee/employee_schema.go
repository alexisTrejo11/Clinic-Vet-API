// Package employee defines the Employee entity and its business logic.
package employee

import (
	"clinic-vet-api/app/core/domain/entity/base"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"fmt"
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

func (v *Employee) AssignUser(userID valueobject.UserID) error {
	if v.userID != nil {
		return fmt.Errorf("employee already assigned to a user with ID %s", v.userID.String())
	}
	v.userID = &userID

	return nil
}
