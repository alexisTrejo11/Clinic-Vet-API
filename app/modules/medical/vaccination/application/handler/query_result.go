package handler

import (
	med "clinic-vet-api/app/modules/core/domain/entity/medical"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type VaccinationResult struct {
	ID               vo.VaccinationID
	PetID            vo.PetID
	VaccineName      string
	AdministeredDate time.Time
	NextDueDate      *time.Time
	AdministeredBy   vo.EmployeeID
	Notes            *string
	CreatedAt        time.Time
}

func toVaccinationResult(d med.PetVaccination) VaccinationResult {
	result := &VaccinationResult{
		ID:               d.ID(),
		PetID:            d.PetID(),
		VaccineName:      d.VaccineName(),
		AdministeredDate: d.AdministeredDate(),
		NextDueDate:      d.NextDueDate(),
		AdministeredBy:   d.AdministeredBy(),
		Notes:            d.Notes(),
		CreatedAt:        d.CreatedAt(),
	}

	return *result
}
