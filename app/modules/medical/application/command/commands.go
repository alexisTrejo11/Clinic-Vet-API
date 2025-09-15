// Package command contains the command definitions and validation logic for medical history operations.
package command

import (
	"context"
	"errors"
	"time"

	"clinic-vet-api/app/core/domain/valueobject"
)

type CreateMedHistCommand struct {
	PetID       valueobject.PetID
	CustomerID  valueobject.CustomerID
	EmployeeID  valueobject.EmployeeID
	Date        time.Time
	Diagnosis   string
	VisitType   string
	VisitReason string
	Notes       *string
	Condition   string
	Treatment   string
	CTX         context.Context
}

type UpdateMedHistCommand struct {
	ID          valueobject.MedHistoryID
	Diagnosis   *string
	VisitType   *string
	VisitReason *string
	Notes       *string
	Condition   *string
	Treatment   *string
	Date        *time.Time
	CTX         context.Context
}

type SoftDeleteMedHistCommand struct {
	ID  valueobject.MedHistoryID
	CTX context.Context
}

type HardDeleteMedHistCommand struct {
	ID  valueobject.MedHistoryID
	CTX context.Context
}

func ValidateCreateCommand(command CreateMedHistCommand) error {
	if command.PetID.IsZero() {
		return errors.New("el ID de la mascota es requerido")
	}
	if command.CustomerID.IsZero() {
		return errors.New("el ID del cliente es requerido")
	}
	if command.EmployeeID.IsZero() {
		return errors.New("el ID del empleado es requerido")
	}
	if command.Date.IsZero() {
		return errors.New("la fecha de visita es requerida")
	}
	if command.Diagnosis == "" {
		return errors.New("el diagnóstico es requerido")
	}
	if command.Treatment == "" {
		return errors.New("el tratamiento es requerido")
	}
	return nil
}

// ValidateUpdateCommand valida los datos del comando de actualización
func ValidateUpdateCommand(command UpdateMedHistCommand) error {
	if command.ID.IsZero() {
		return errors.New("el ID del historial médico es requerido")
	}

	if command.Date == nil && command.VisitType == nil && command.VisitReason == nil &&
		command.Diagnosis == nil && command.Treatment == nil && command.Condition == nil &&
		command.Notes == nil {
		return errors.New("debe proporcionar al menos un campo para actualizar")
	}

	return nil
}
