package command

import (
	"context"
	"errors"
	"time"

	med "clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
)

type DewormUpdateCommand struct {
	ID               valueobject.DewormID
	PetID            *valueobject.PetID
	AdministeredBy   *valueobject.EmployeeID
	MedicationName   *string
	AdministeredDate *time.Time
	NextDueDate      *time.Time
	Notes            *string
}

func (h *DewormCommandHandler) HandleUpdate(ctx context.Context, cmd DewormUpdateCommand) cqrs.CommandResult {
	if err := cmd.Validate(); err != nil {
		return *cqrs.FailureResult("command validation failed", err)
	}

	existingDeworm, err := h.dewormRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return *cqrs.FailureResult("failed to retrieve existing deworm record", err)
	}

	updatedDeworm := updateDewormFields(cmd, *existingDeworm)
	if err := h.dewormRepo.Save(ctx, &updatedDeworm); err != nil {
		return *cqrs.FailureResult("Failed to update deworm record", err)
	}

	return *cqrs.SuccessResult("deworm record updated successfully")
}

func updateDewormFields(cmd DewormUpdateCommand, deworm med.PetDeworming) med.PetDeworming {
	dewormBuilder := med.NewPetDewormingBuilder().
		WithID(deworm.ID())

	if cmd.PetID != nil {
		dewormBuilder = dewormBuilder.WithPetID(*cmd.PetID)
	}

	if cmd.AdministeredBy != nil {
		dewormBuilder = dewormBuilder.WithAdministeredBy(*cmd.AdministeredBy)
	}

	if cmd.MedicationName != nil {
		dewormBuilder = dewormBuilder.WithMedicationName(*cmd.MedicationName)
	}

	if cmd.AdministeredDate != nil {
		dewormBuilder = dewormBuilder.WithAdministeredDate(*cmd.AdministeredDate)
	}
	if cmd.NextDueDate != nil {
		dewormBuilder = dewormBuilder.WithNextDueDate(cmd.NextDueDate)
	}
	if cmd.Notes != nil {
		dewormBuilder = dewormBuilder.WithNotes(cmd.Notes)
	}

	return *dewormBuilder.Build()
}

func (cmd *DewormUpdateCommand) Validate() error {
	if cmd.ID.IsZero() {
		return errors.New("deworm id is required if provided")
	}

	if cmd.PetID != nil && cmd.PetID.IsZero() {
		return errors.New("pet id is required if provided")
	}

	if cmd.AdministeredBy != nil && cmd.AdministeredBy.IsZero() {
		return errors.New("administered by (employee id) is required if provided")
	}

	if cmd.MedicationName != nil && *cmd.MedicationName == "" {
		return errors.New("medication name cannot be empty if provided")
	}

	if cmd.AdministeredDate != nil && cmd.AdministeredDate.IsZero() {
		return errors.New("administered date is required if provided")
	}

	return nil
}
