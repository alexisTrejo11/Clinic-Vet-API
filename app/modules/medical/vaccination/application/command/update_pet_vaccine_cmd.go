package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type UpdateVaccinationCommand struct {
	VaccinationID    valueobject.VaccinationID
	VaccineName      *string
	VaccineType      *string
	AdministeredDate *time.Time
	AdministeredBy   *valueobject.EmployeeID
	BatchNumber      *string
	Notes            *string
	NextDueDate      *time.Time
}

func (cmd *UpdateVaccinationCommand) ToUpdateEntity(vaccination *medical.PetVaccination) *medical.PetVaccination {
	builder := medical.NewPetVaccinationBuilder().
		WithID(vaccination.ID()).
		WithPetID(vaccination.PetID()).
		WithVaccineName(vaccination.VaccineName()).
		WithVaccineType(vaccination.VaccineType()).
		WithAdministeredDate(vaccination.AdministeredDate()).
		WithAdministeredBy(vaccination.AdministeredBy()).
		WithBatchNumber(vaccination.BatchNumber()).
		WithNextDueDate(vaccination.NextDueDate()).
		WithNotes(vaccination.Notes())

	if cmd.VaccineName != nil {
		builder = builder.WithVaccineName(*cmd.VaccineName)
	}

	if cmd.VaccineType != nil {
		builder = builder.WithVaccineType(*cmd.VaccineType)
	}
	if cmd.AdministeredDate != nil {
		builder = builder.WithAdministeredDate(*cmd.AdministeredDate)
	}
	if cmd.BatchNumber != nil {
		builder = builder.WithBatchNumber(*cmd.BatchNumber)
	}
	if cmd.Notes != nil {
		builder = builder.WithNotes(cmd.Notes)
	}
	if cmd.NextDueDate != nil {
		builder = builder.WithNextDueDate(cmd.NextDueDate)
	}

	return builder.Build()
}
