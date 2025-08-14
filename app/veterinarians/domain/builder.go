package vetDomain

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
)

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

func (vb *VeterinarianBuilder) WithName(name valueObjects.PersonName) *VeterinarianBuilder {
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

func (vb *VeterinarianBuilder) WithSpecialty(specialty VetSpecialty) *VeterinarianBuilder {
	vb.vet.specialty = specialty
	return vb
}

func (vb *VeterinarianBuilder) WithYearsExperience(yearsExperience int) *VeterinarianBuilder {
	vb.vet.yearsExperience = yearsExperience
	return vb
}

func (vb *VeterinarianBuilder) WithConsultationFee(consultationFee *shared.Money) *VeterinarianBuilder {
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

func (vb *VeterinarianBuilder) WithSchedule(schedule *Schedule) *VeterinarianBuilder {
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
