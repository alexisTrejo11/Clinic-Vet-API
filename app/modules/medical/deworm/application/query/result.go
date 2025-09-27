package query

import (
	"time"

	med "clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type DewormResult struct {
	ID               valueobject.DewormID
	PetID            valueobject.PetID
	AdministeredBy   valueobject.EmployeeID
	MedicationName   string
	AdministeredDate time.Time
	NextDueDate      *time.Time
	Notes            *string
	CreatedAt        time.Time
}

func toDewormResult(d med.PetDeworming) DewormResult {
	result := &DewormResult{
		ID:               d.ID(),
		PetID:            d.PetID(),
		MedicationName:   d.MedicationName(),
		AdministeredDate: d.AdministeredDate(),
		NextDueDate:      d.NextDueDate(),
		AdministeredBy:   d.AdministeredBy(),
		Notes:            d.Notes(),
		CreatedAt:        d.CreatedAt(),
	}

	return *result
}
