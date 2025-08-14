package vetDomain

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
)

type Veterinarian struct {
	id              int
	name            valueObjects.PersonName
	photo           string
	licenseNumber   string
	specialty       VetSpecialty
	yearsExperience int
	consultationFee *shared.Money
	isActive        bool
	userID          *int
	scheduleJSON    string
	schedule        *Schedule
	createdAt       time.Time
	updatedAt       time.Time
}

// Getters
func (v *Veterinarian) GetID() int {
	return v.id
}

func (v *Veterinarian) GetName() valueObjects.PersonName {
	return v.name
}

func (v *Veterinarian) GetPhoto() string {
	return v.photo
}

func (v *Veterinarian) GetLicenseNumber() string {
	return v.licenseNumber
}

func (v *Veterinarian) GetSpecialty() VetSpecialty {
	return v.specialty
}

func (v *Veterinarian) GetYearsExperience() int {
	return v.yearsExperience
}

func (v *Veterinarian) GetConsultationFee() *shared.Money {
	return v.consultationFee
}

func (v *Veterinarian) GetIsActive() bool {
	return v.isActive
}

func (v *Veterinarian) GetUserID() *int {
	return v.userID
}

func (v *Veterinarian) GetScheduleJSON() string {
	return v.scheduleJSON
}

func (v *Veterinarian) GetSchedule() *Schedule {
	return v.schedule
}

func (v *Veterinarian) GetCreatedAt() time.Time {
	return v.createdAt
}

func (v *Veterinarian) GetUpdatedAt() time.Time {
	return v.updatedAt
}

// Setters
func (v *Veterinarian) SetName(name valueObjects.PersonName) {
	v.name = name
}

func (v *Veterinarian) SetPhoto(photo string) {
	v.photo = photo
}

func (v *Veterinarian) SetLicenseNumber(licenseNumber string) {
	v.licenseNumber = licenseNumber
}

func (v *Veterinarian) SetSpecialty(specialty VetSpecialty) {
	v.specialty = specialty
}

func (v *Veterinarian) SetYearsExperience(yearsExperience int) {
	v.yearsExperience = yearsExperience
}

func (v *Veterinarian) SetConsultationFee(consultationFee *shared.Money) {
	v.consultationFee = consultationFee
}

func (v *Veterinarian) SetIsActive(isActive bool) {
	v.isActive = isActive
}

func (v *Veterinarian) SetUserID(userID *int) {
	v.userID = userID
}

func (v *Veterinarian) SetScheduleJSON(scheduleJSON string) {
	v.scheduleJSON = scheduleJSON
}

func (v *Veterinarian) SetSchedule(schedule *Schedule) {
	v.schedule = schedule
}

func (v *Veterinarian) SetUpdatedAt(updatedAt time.Time) {
	v.updatedAt = updatedAt
}

func (v *Veterinarian) SetCreatedAt(createdAt time.Time) {
	v.createdAt = createdAt
}

func (v *Veterinarian) SetID(id int) {
	v.id = id
}
