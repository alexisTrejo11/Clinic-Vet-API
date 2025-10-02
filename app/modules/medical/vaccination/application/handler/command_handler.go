package handler

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	c "clinic-vet-api/app/modules/medical/vaccination/application/command"
	"clinic-vet-api/app/shared/cqrs"
	"context"
	"errors"
	"time"
)

var (
	FailFindingVaccinationMsg     = "failed to find vaccination"
	FailSavingVaccinationMsg      = "failed to save vaccination"
	FailDeletingVaccinationMsg    = "failed to delete vaccination"
	FailCalculatingNextDueDateMsg = "failed to calculate next due date"
	FailVaccineConflictMsg        = "vaccine conflict detected"
	FailVaccineValidationMsg      = "vaccine validation failed"

	SuccesRegisteringVaccinationMsg = "vaccination registered successfully"
	SuccesUpdatingVaccinationMsg    = "vaccination updated successfully"
	SuccesDeletingVaccinationMsg    = "vaccination deleted successfully"
)

type PetVaccineCmdHandler struct {
	petRepo            repository.PetRepository
	vaccinationService *service.VaccinationScheduleService
	vaccinationRepo    repository.VaccinationRepository
}

func NewPetVaccineCmdHandler(
	petRepo repository.PetRepository,
	vaccinationRepo repository.VaccinationRepository,
	vaccinationService *service.VaccinationScheduleService,
) *PetVaccineCmdHandler {
	return &PetVaccineCmdHandler{
		petRepo:            petRepo,
		vaccinationRepo:    vaccinationRepo,
		vaccinationService: vaccinationService,
	}
}

func (h *PetVaccineCmdHandler) HandleRegister(ctx context.Context, cmd c.RegisterVaccinationCommand) cqrs.CommandResult {
	pet, err := h.petRepo.FindByID(ctx, cmd.PetID)
	if err != nil {
		return *cqrs.FailureResult(FailFindingVaccinationMsg, err)
	} else if !pet.IsActive() {
		return *cqrs.FailureResult("Inactive Pet", errors.New("pet is not active"))
	}

	if err := h.vaccinationService.ValidateVaccination(&pet, cmd.VaccineName, cmd.AdministeredDate); err != nil {
		return *cqrs.FailureResult(FailVaccineValidationMsg, err)
	}

	recentVaccinations, err := h.vaccinationRepo.FindRecentByPetID(ctx, cmd.PetID, 30)
	if err != nil {
		return *cqrs.FailureResult(FailFindingVaccinationMsg, err)
	}

	if err := h.vaccinationService.CheckVaccinationConflicts(cmd.VaccineName, cmd.AdministeredDate, recentVaccinations); err != nil {
		return *cqrs.FailureResult(FailVaccineConflictMsg, err)
	}

	var nextDueDate *time.Time
	if cmd.NextDueDate != nil {
		nextDueDate = cmd.NextDueDate
	} else {
		nextDueDate, err = h.calculateNextDueDate(ctx, cmd)
		if err != nil {
			return *cqrs.FailureResult(FailCalculatingNextDueDateMsg, err)
		}
	}

	vaccination := cmd.ToEntity(nextDueDate)
	vaccineCreated, err := h.vaccinationRepo.Save(ctx, vaccination)
	if err != nil {
		return *cqrs.FailureResult(FailSavingVaccinationMsg, err)
	}

	return *cqrs.SuccessCreateResult(vaccineCreated.ID().String(), SuccesRegisteringVaccinationMsg)
}

func (h *PetVaccineCmdHandler) HandleUpdate(ctx context.Context, cmd c.UpdateVaccinationCommand) cqrs.CommandResult {
	vaccination, err := h.vaccinationRepo.FindByID(ctx, cmd.VaccinationID)
	if err != nil {
		return *cqrs.FailureResult(FailFindingVaccinationMsg, err)
	}

	updatedVaccination := cmd.ToUpdateEntity(vaccination)
	_, err = h.vaccinationRepo.Save(ctx, *updatedVaccination)
	if err != nil {
		return *cqrs.FailureResult(FailSavingVaccinationMsg, err)
	}

	return *cqrs.SuccessResult(SuccesUpdatingVaccinationMsg)
}

func (h *PetVaccineCmdHandler) HandleDelete(ctx context.Context, vaccinationID vo.VaccinationID) cqrs.CommandResult {
	_, err := h.vaccinationRepo.FindByID(ctx, vaccinationID)
	if err != nil {
		return *cqrs.FailureResult(FailFindingVaccinationMsg, err)
	}

	if err := h.vaccinationRepo.Delete(ctx, vaccinationID); err != nil {
		return *cqrs.FailureResult(FailDeletingVaccinationMsg, err)
	}

	return *cqrs.SuccessResult(SuccesDeletingVaccinationMsg)
}

func (h *PetVaccineCmdHandler) calculateNextDueDate(ctx context.Context, cmd c.RegisterVaccinationCommand) (*time.Time, error) {
	var nextDueDate *time.Time
	allVaccinations, _ := h.vaccinationRepo.FindAllByPetID(ctx, cmd.PetID)
	previousCount := 0
	for _, v := range allVaccinations {
		if v.VaccineName() == cmd.VaccineName {
			previousCount++
		}
	}

	nextDueDate, err := h.vaccinationService.CalculateNextVaccination(cmd.VaccineName, cmd.AdministeredDate, previousCount)
	if err != nil {
		return nil, err
	}

	return nextDueDate, nil
}

func (h *PetVaccineCmdHandler) HandleGenerateVaccPlan(ctx context.Context, cmd c.GenerateVaccinationPlanCommand) (medical.VaccinationPlan, error) {
	pet, err := h.petRepo.FindByID(ctx, cmd.PetID)
	if err != nil {
		return medical.VaccinationPlan{}, err
	}

	existingVaccinations, err := h.vaccinationRepo.FindAllByPetID(ctx, cmd.PetID)
	if err != nil {
		return medical.VaccinationPlan{}, err
	}

	plan := h.vaccinationService.GenerateVaccinationPlan(&pet, cmd.SinceStartPlan, existingVaccinations)
	return plan, nil
}
