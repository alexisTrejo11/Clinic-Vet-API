// Package veterinarian defines the Veterinarian entity and its business logic.
package veterinarian

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type Veterinarian struct {
	base.Entity[valueobject.VetID]
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

func (v *Veterinarian) ID() valueobject.VetID {
	return v.Entity.ID()
}

func (v *Veterinarian) Name() valueobject.PersonName {
	return v.Person.Name()
}

func (v *Veterinarian) Photo() string {
	return v.photo
}

func (v *Veterinarian) LicenseNumber() string {
	return v.licenseNumber
}

func (v *Veterinarian) Specialty() enum.VetSpecialty {
	return v.specialty
}

func (v *Veterinarian) YearsExperience() int {
	return v.yearsExperience
}

func (v *Veterinarian) ConsultationFee() *valueobject.Money {
	return v.consultationFee
}

func (v *Veterinarian) IsActive() bool {
	return v.isActive
}

func (v *Veterinarian) UserID() *valueobject.UserID {
	return v.userID
}

func (v *Veterinarian) Schedule() *valueobject.Schedule {
	return v.schedule
}
