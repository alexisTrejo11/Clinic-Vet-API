package command

import (
	"context"
	"errors"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
)

type DewormCreateCommand struct {
	petID            valueobject.PetID
	administeredBy   valueobject.EmployeeID
	medicationName   string
	administeredDate time.Time
	nextDueDate      *time.Time
	notes            *string
}

func NewDewormCreateCommand(
	petID uint,
	administeredBy uint,
	medicationName string,
	administeredDate time.Time,
	nextDueDate *time.Time,
	notes *string,
) DewormCreateCommand {
	return DewormCreateCommand{
		petID:            valueobject.NewPetID(petID),
		administeredBy:   valueobject.NewEmployeeID(administeredBy),
		medicationName:   medicationName,
		administeredDate: administeredDate,
		nextDueDate:      nextDueDate,
		notes:            notes,
	}
}

// PetDewormingService??????
func (h *DewormCommandHandler) HandleCreate(ctx context.Context, cmd DewormCreateCommand) cqrs.CommandResult {
	if err := cmd.Validate(); err != nil {
		return *cqrs.FailureResult("command validtion error", err)
	}

	if err := h.validatePetEmployeeExistence(ctx, cmd); err != nil {
		return *cqrs.FailureResult("entity validation error", err)
	}

	entity := cmd.toEntity()
	dewormCreated, err := h.dewormRepo.Save(ctx, *entity)
	if err != nil {
		return *cqrs.FailureResult("failed to create deworming record", err)
	}

	return *cqrs.SuccessCreateResult(dewormCreated.ID().String(), "deworming record created successfully")
}

func (h *DewormCommandHandler) validatePetEmployeeExistence(ctx context.Context, cmd DewormCreateCommand) error {
	if exists, err := h.petRepo.ExistsByID(ctx, cmd.petID); err != nil {
		return err
	} else if !exists {
		return errors.New("pet does not exist")
	}

	if exists, err := h.employeeRepo.ExistsByID(ctx, cmd.administeredBy); err != nil {
		return err
	} else if !exists {
		return errors.New("employee does not exist")
	}

	return nil
}

func (c *DewormCreateCommand) toEntity() *medical.PetDeworming {
	return medical.NewPetDewormingBuilder().
		WithPetID(c.petID).
		WithAdministeredBy(c.administeredBy).
		WithMedicationName(c.medicationName).
		WithAdministeredDate(c.administeredDate).
		WithNextDueDate(c.nextDueDate).
		WithNotes(c.notes).
		Build()
}

func (c *DewormCreateCommand) Validate() error {
	if c.administeredBy.IsZero() {
		return errors.New("administeredBy is required")
	}
	if c.petID.IsZero() {
		return errors.New("petID is required")
	}

	return nil
}
