package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type RegisterVaccinationCommand struct {
	PetID            vo.PetID
	VaccineName      string
	VaccineType      string
	AdministeredDate time.Time
	AdministeredBy   vo.EmployeeID
	BatchNumber      string
	Notes            *string
	NextDueDate      *time.Time
}

func (cmd *RegisterVaccinationCommand) ToEntity(nextDueDate *time.Time) medical.PetVaccination {
	vaccination := medical.NewPetVaccinationBuilder().
		WithPetID(cmd.PetID).
		WithVaccineName(cmd.VaccineName).
		WithVaccineType(cmd.VaccineType).
		WithAdministeredDate(cmd.AdministeredDate).
		WithAdministeredBy(cmd.AdministeredBy).
		WithBatchNumber(cmd.BatchNumber).
		WithNextDueDate(nextDueDate).
		WithNotes(cmd.Notes).
		WithCreatedAt(time.Now()).
		Build()
	return *vaccination
}
