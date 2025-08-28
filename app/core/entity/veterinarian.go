package entity

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

type Veterinarian struct {
	id              int
	name            valueobject.PersonName
	photo           string
	licenseNumber   string
	specialty       enum.VetSpecialty
	yearsExperience int
	consultationFee *valueobject.Money
	isActive        bool
	userID          *int
	scheduleJSON    string
	schedule        *valueobject.Schedule
	createdAt       time.Time
	updatedAt       time.Time
}

// Getters
func (v *Veterinarian) GetID() int {
	return v.id
}

func (v *Veterinarian) GetName() valueobject.PersonName {
	return v.name
}

func (v *Veterinarian) GetPhoto() string {
	return v.photo
}

func (v *Veterinarian) GetLicenseNumber() string {
	return v.licenseNumber
}

func (v *Veterinarian) GetSpecialty() enum.VetSpecialty {
	return v.specialty
}

func (v *Veterinarian) GetYearsExperience() int {
	return v.yearsExperience
}

func (v *Veterinarian) GetConsultationFee() *valueobject.Money {
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

func (v *Veterinarian) GetSchedule() *valueobject.Schedule {
	return v.schedule
}

func (v *Veterinarian) GetCreatedAt() time.Time {
	return v.createdAt
}

func (v *Veterinarian) GetUpdatedAt() time.Time {
	return v.updatedAt
}

// Setters
func (v *Veterinarian) SetName(name valueobject.PersonName) {
	v.name = name
}

func (v *Veterinarian) SetPhoto(photo string) {
	v.photo = photo
}

func (v *Veterinarian) SetLicenseNumber(licenseNumber string) {
	v.licenseNumber = licenseNumber
}

func (v *Veterinarian) SetSpecialty(specialty enum.VetSpecialty) {
	v.specialty = specialty
}

func (v *Veterinarian) SetYearsExperience(yearsExperience int) {
	v.yearsExperience = yearsExperience
}

func (v *Veterinarian) SetConsultationFee(consultationFee *valueobject.Money) {
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

func (v *Veterinarian) SetSchedule(schedule *valueobject.Schedule) {
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

type VeterinarianBuilder struct {
	vet *Veterinarian
}

func NewVeterinarianBuilder() *VeterinarianBuilder {
	return &VeterinarianBuilder{vet: &Veterinarian{}}
}

func (vb *VeterinarianBuilder) WithID(id int) *VeterinarianBuilder {
	vb.vet.id = id
	return vb
}

func (vb *VeterinarianBuilder) WithName(name valueobject.PersonName) *VeterinarianBuilder {
	vb.vet.name = name
	return vb
}

func (vb *VeterinarianBuilder) WithPhoto(photo string) *VeterinarianBuilder {
	vb.vet.photo = photo
	return vb
}

func (vb *VeterinarianBuilder) WithLicenseNumber(licenseNumber string) *VeterinarianBuilder {
	vb.vet.licenseNumber = licenseNumber
	return vb
}

func (vb *VeterinarianBuilder) WithSpecialty(specialty enum.VetSpecialty) *VeterinarianBuilder {
	vb.vet.specialty = specialty
	return vb
}

func (vb *VeterinarianBuilder) WithYearsExperience(yearsExperience int) *VeterinarianBuilder {
	vb.vet.yearsExperience = yearsExperience
	return vb
}

func (vb *VeterinarianBuilder) WithConsultationFee(consultationFee *valueobject.Money) *VeterinarianBuilder {
	vb.vet.consultationFee = consultationFee
	return vb
}

func (vb *VeterinarianBuilder) WithIsActive(isActive bool) *VeterinarianBuilder {
	vb.vet.isActive = isActive
	return vb
}

func (vb *VeterinarianBuilder) WithUserID(userID *int) *VeterinarianBuilder {
	vb.vet.userID = userID
	return vb
}

func (vb *VeterinarianBuilder) WithScheduleJSON(scheduleJSON string) *VeterinarianBuilder {
	vb.vet.scheduleJSON = scheduleJSON
	return vb
}

func (vb *VeterinarianBuilder) WithSchedule(schedule *valueobject.Schedule) *VeterinarianBuilder {
	vb.vet.schedule = schedule
	return vb
}

func (vb *VeterinarianBuilder) WithCreatedAt(createdAt time.Time) *VeterinarianBuilder {
	vb.vet.createdAt = createdAt
	return vb
}

func (vb *VeterinarianBuilder) WithUpdatedAt(updatedAt time.Time) *VeterinarianBuilder {
	vb.vet.updatedAt = updatedAt
	return vb
}

func (vb *VeterinarianBuilder) Build() *Veterinarian {
	return vb.vet
}
