// Package command contains the command definitions and validation logic for medical history operations.
package command

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"errors"
	"time"
)

type CreateMedSessionCommand struct {
	CustomerID valueobject.CustomerID
	EmployeeID valueobject.EmployeeID
	VisitDate  time.Time
	VisitType  enum.VisitType
	Diagnosis  string
	Service    enum.ClinicService
	Notes      *string
	PetDetails PetSummary
}

type PetSummary struct {
	PetID           valueobject.PetID
	Weight          *valueobject.Decimal
	HeartRate       *int
	RespiratoryRate *int
	Temperature     *valueobject.Decimal
	Diagnosis       string
	Treatment       string
	Condition       enum.PetCondition
	Medications     []string
	FollowUpDate    *time.Time
	Symptoms        []string
}

type UpdateMedSessionCommand struct {
	ID        valueobject.MedSessionID
	Diagnosis *string
	VisitType *enum.VisitType
	Service   *enum.ClinicService
	Notes     *string
	Condition *enum.PetCondition
	Treatment *string
	Date      *time.Time
}

type SoftDeleteMedSessionCommand struct {
	ID valueobject.MedSessionID
}

type HardDeleteMedSessionCommand struct {
	ID valueobject.MedSessionID
}

func ValidateUpdateCommand(command UpdateMedSessionCommand) error {
	if command.ID.IsZero() {
		return errors.New("el ID del historial médico es requerido")
	}

	if command.Date == nil && command.VisitType == nil && command.Service == nil &&
		command.Diagnosis == nil && command.Treatment == nil && command.Condition == nil &&
		command.Notes == nil {
		return errors.New("at least one field must be provided for update")
	}

	return nil
}
