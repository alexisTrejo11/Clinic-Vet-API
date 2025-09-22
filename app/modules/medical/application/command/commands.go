// Package command contains the command definitions and validation logic for medical history operations.
package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"context"
	"errors"
	"time"
)

type CreateMedSessionCommand struct {
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

type UpdateMedSessionCommand struct {
	ID          valueobject.MedSessionID
	Diagnosis   *string
	VisitType   *string
	VisitReason *string
	Notes       *string
	Condition   *string
	Treatment   *string
	Date        *time.Time
	CTX         context.Context
}

type SoftDeleteMedSessionCommand struct {
	ID  valueobject.MedSessionID
	CTX context.Context
}

type HardDeleteMedSessionCommand struct {
	ID  valueobject.MedSessionID
	CTX context.Context
}

// ValidateUpdateCommand valida los datos del comando de actualización
func ValidateUpdateCommand(command UpdateMedSessionCommand) error {
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
