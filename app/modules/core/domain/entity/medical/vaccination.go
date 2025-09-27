package medical

import (
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type PetVaccination struct {
	id               vo.VaccinationID
	petID            vo.PetID
	vaccineName      string
	administeredDate time.Time
	nextDueDate      *time.Time
	administeredBy   *vo.EmployeeID
	notes            *string
	createdAt        time.Time
}

type PetVaccinationBuilder struct{ petVaccination *PetVaccination }

func NewPetVaccinationBuilder() *PetVaccinationBuilder {
	return &PetVaccinationBuilder{petVaccination: &PetVaccination{}}
}

func (b *PetVaccinationBuilder) WithID(id vo.VaccinationID) PetVaccinationBuilder {
	b.petVaccination.id = id
	return *b
}

func (b *PetVaccinationBuilder) WithPetID(petID vo.PetID) PetVaccinationBuilder {
	b.petVaccination.petID = petID
	return *b
}

func (b *PetVaccinationBuilder) WithVaccineName(vaccineName string) PetVaccinationBuilder {
	b.petVaccination.vaccineName = vaccineName
	return *b
}

func (b *PetVaccinationBuilder) WithAdministeredDate(administeredDate time.Time) PetVaccinationBuilder {
	b.petVaccination.administeredDate = administeredDate
	return *b
}

func (b *PetVaccinationBuilder) WithNextDueDate(nextDueDate *time.Time) PetVaccinationBuilder {
	b.petVaccination.nextDueDate = nextDueDate
	return *b
}

func (b *PetVaccinationBuilder) WithAdministeredBy(administeredBy *vo.EmployeeID) PetVaccinationBuilder {
	b.petVaccination.administeredBy = administeredBy
	return *b
}

func (b *PetVaccinationBuilder) WithNotes(notes *string) PetVaccinationBuilder {
	b.petVaccination.notes = notes
	return *b
}

func (b *PetVaccinationBuilder) WithCreatedAt(createdAt time.Time) PetVaccinationBuilder {
	b.petVaccination.createdAt = createdAt
	return *b
}

func (pv *PetVaccination) ID() vo.VaccinationID           { return pv.id }
func (pv *PetVaccination) PetID() vo.PetID                { return pv.petID }
func (pv *PetVaccination) VaccineName() string            { return pv.vaccineName }
func (pv *PetVaccination) AdministeredDate() time.Time    { return pv.administeredDate }
func (pv *PetVaccination) NextDueDate() *time.Time        { return pv.nextDueDate }
func (pv *PetVaccination) AdministeredBy() *vo.EmployeeID { return pv.administeredBy }
func (pv *PetVaccination) Notes() *string                 { return pv.notes }
func (pv *PetVaccination) CreatedAt() time.Time           { return pv.createdAt }
