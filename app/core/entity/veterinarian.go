package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

type Veterinarian struct {
	id              valueobject.VetID
	name            valueobject.PersonName
	photo           string
	licenseNumber   string
	specialty       enum.VetSpecialty
	yearsExperience int
	consultationFee *valueobject.Money
	isActive        bool
	userID          *valueobject.UserID
	scheduleJSON    string
	schedule        *valueobject.Schedule
	createdAt       time.Time
	updatedAt       time.Time
}

func (v *Veterinarian) GetID() valueobject.VetID {
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

func (v *Veterinarian) GetUserID() *valueobject.UserID {
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

func (v *Veterinarian) SetUserID(userID valueobject.UserID) {
	v.userID = &userID
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

func (v *Veterinarian) SetID(id valueobject.VetID) {
	v.id = id
}

type VeterinarianBuilder struct {
	vet *Veterinarian
}

func NewVeterinarianBuilder() *VeterinarianBuilder {
	return &VeterinarianBuilder{vet: &Veterinarian{}}
}

func (vb *VeterinarianBuilder) WithID(id valueobject.VetID) *VeterinarianBuilder {
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

func (vb *VeterinarianBuilder) WithUserID(userID valueobject.UserID) *VeterinarianBuilder {
	vb.vet.userID = &userID
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

// Validation
const (
	CLINIC_OPENING_HOUR  = 9
	CLINIC_CLOSING_HOUR  = 20
	TOTAL_WEEK_DAYS      = 7
	MAX_WORK_DAYS        = TOTAL_WEEK_DAYS - 1
	MIN_LICENSE_LENGTH   = 8
	MAX_LICENSE_LENGTH   = 20
	MAX_EXPERIENCE_YEARS = 60
	MAX_BREAK_HOURS      = 2
)

func (v *Veterinarian) ValidateInsert() error {
	service := VetValidatorService{vet: v}
	var errs []error

	if err := service.validateLicenseNumber(); err != nil {
		errs = append(errs, err)
	}
	if err := service.validateYearsOfExperience(); err != nil {
		errs = append(errs, err)
	}
	if err := service.validateSchedule(); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

type VetValidatorService struct {
	vet *Veterinarian
}

func NewVetValidatorService(vet *Veterinarian) *VetValidatorService {
	return &VetValidatorService{vet: vet}
}

func (v *VetValidatorService) validateLicenseNumber() error {
	if len(v.vet.GetLicenseNumber()) < MIN_LICENSE_LENGTH || len(v.vet.GetLicenseNumber()) > MAX_LICENSE_LENGTH {
		return fmt.Errorf("veterinarian license number invalid length")
	}
	return nil
}

func (v *VetValidatorService) validateYearsOfExperience() error {
	if v.vet.GetYearsExperience() > MAX_EXPERIENCE_YEARS {
		return fmt.Errorf("years of experience seems unrealistic for a human career span")
	}
	return nil
}

func (v *VetValidatorService) validateSchedule() error {
	if v.vet.GetScheduleJSON() == "" {
		return nil
	}

	if err := json.Unmarshal([]byte(v.vet.GetScheduleJSON()), v.vet.GetSchedule()); err != nil {
		return fmt.Errorf("invalid schedule format: %v", err)
	}

	return v.vet.GetSchedule().Validate()
}

func (v *VetValidatorService) BeforeSave() error {
	scheduleBytes, err := json.Marshal(v.vet.GetSchedule())
	if err != nil {
		return err
	}
	v.vet.SetScheduleJSON(string(scheduleBytes))
	return nil
}

func (v *VetValidatorService) AfterFind() error {
	if v.vet.GetScheduleJSON() != "" {
		return json.Unmarshal([]byte(v.vet.GetScheduleJSON()), v.vet.GetSchedule())
	}
	return nil
}
