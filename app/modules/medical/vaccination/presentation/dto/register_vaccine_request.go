package dto

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/medical/vaccination/application/command"
	"time"
)

type RegisterVaccineRequest struct {
	PetID            uint       `json:"pet_id" binding:"required"`
	VaccineName      string     `json:"vaccine_name" binding:"required"`
	VaccineType      string     `json:"vaccine_type" binding:"required"`
	AdministeredDate time.Time  `json:"administered_date" binding:"required"`
	BatchNumber      string     `json:"batch_number" binding:"required"`
	Notes            *string    `json:"notes,omitempty"`
	NextDueDate      *time.Time `json:"next_due_date,omitempty"`
}

func (r *RegisterVaccineRequest) ToCommand(administeredBy uint) (command.RegisterVaccinationCommand, error) {
	return command.RegisterVaccinationCommand{
		PetID:            valueobject.NewPetID(r.PetID),
		VaccineName:      r.VaccineName,
		VaccineType:      r.VaccineType,
		AdministeredDate: r.AdministeredDate,
		AdministeredBy:   valueobject.NewEmployeeID(administeredBy),
		BatchNumber:      r.BatchNumber,
		Notes:            r.Notes,
		NextDueDate:      r.NextDueDate,
	}, nil
}

type UpdateVaccineRequest struct {
	VaccineName      *string    `json:"vaccine_name,omitempty"`
	VaccineType      *string    `json:"vaccine_type,omitempty"`
	AdministeredDate *time.Time `json:"administered_date,omitempty"`
	BatchNumber      *string    `json:"batch_number,omitempty"`
	Notes            *string    `json:"notes,omitempty"`
	NextDueDate      *time.Time `json:"next_due_date,omitempty"`
}

func (r *UpdateVaccineRequest) ToCommand(vaccinationID uint, optEmployeeID uint) (command.UpdateVaccinationCommand, error) {
	return command.UpdateVaccinationCommand{
		VaccinationID:    valueobject.NewVaccinationID(vaccinationID),
		VaccineName:      r.VaccineName,
		VaccineType:      r.VaccineType,
		AdministeredDate: r.AdministeredDate,
		AdministeredBy:   valueobject.NewOptEmployeeID(&optEmployeeID),
		BatchNumber:      r.BatchNumber,
		Notes:            r.Notes,
		NextDueDate:      r.NextDueDate,
	}, nil
}

type RegisterVaccineResponse struct {
	ID               uint       `json:"id"`
	PetID            uint       `json:"pet_id"`
	VaccineName      string     `json:"vaccine_name"`
	VaccineType      string     `json:"vaccine_type"`
	AdministeredDate time.Time  `json:"administered_date"`
	AdministeredBy   uint       `json:"administered_by"`
	BatchNumber      string     `json:"batch_number"`
	Notes            *string    `json:"notes,omitempty"`
	NextDueDate      *time.Time `json:"next_due_date,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}
