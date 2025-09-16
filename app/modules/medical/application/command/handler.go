package command

import (
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
)

type MedicalHistoryCommandHandlers interface {
	CreateMedicalHistory(command CreateMedHistCommand) cqrs.CommandResult
	UpdateMedicalHistory(command UpdateMedHistCommand) cqrs.CommandResult
	SoftDeleteMedicalHistory(command SoftDeleteMedHistCommand) cqrs.CommandResult
	HardDeleteMedicalHistory(command HardDeleteMedHistCommand) cqrs.CommandResult
}

type medicalHistoryCommandHandlers struct {
	repo repository.MedicalHistoryRepository
}

func NewMedicalHistoryCommandHandlers(repo repository.MedicalHistoryRepository) MedicalHistoryCommandHandlers {
	return &medicalHistoryCommandHandlers{
		repo: repo,
	}
}

func (h *medicalHistoryCommandHandlers) CreateMedicalHistory(command CreateMedHistCommand) cqrs.CommandResult {
	exists, err := h.repo.ExistsByPetAndDate(command.CTX, command.PetID, command.Date)
	if err != nil {
		return errorCreateResult(msgErrorProcessingData, err)
	}
	if exists {
		return errorCreateResult(msgMedicalHistoryDateConflict, nil)
	}

	entity, err := ToEntityFromCreate(&command)
	if err != nil {
		return errorCreateResult(msgErrorProcessingData, err)
	}

	if err := h.repo.Save(command.CTX, entity); err != nil {
		return errorCreateResult(msgErrorProcessingData, err)
	}

	return successCreateResult(*entity)
}

func (h *medicalHistoryCommandHandlers) UpdateMedicalHistory(command UpdateMedHistCommand) cqrs.CommandResult {
	exists, err := h.repo.ExistsByID(command.CTX, command.ID)
	if err != nil {
		return errorUpdateResult(msgErrorProcessingData, err)
	}
	if !exists {
		return errorUpdateResult(msgMedicalHistoryNotFound, nil)
	}

	existingEntity, err := h.repo.FindByID(command.CTX, command.ID)
	if err != nil {
		return errorUpdateResult(msgErrorProcessingData, err)
	}

	if command.Date != nil && !command.Date.Equal(existingEntity.VisitDate()) {
		exists, err := h.repo.ExistsByPetAndDate(command.CTX, existingEntity.PetID(), *command.Date)
		if err != nil {
			return errorUpdateResult(msgErrorProcessingData, err)
		}
		if exists {
			return errorUpdateResult(msgMedicalHistoryNewDateConflict, nil)
		}
	}

	updatedEntity, err := ToEntityFromUpdate(&command, &existingEntity)
	if err != nil {
		return errorUpdateResult(msgErrorProcessingData, err)
	}

	if err := h.repo.Update(command.CTX, updatedEntity); err != nil {
		return errorUpdateResult(msgErrorProcessingData, err)
	}

	return successUpdateResult(*updatedEntity)
}

func (h *medicalHistoryCommandHandlers) SoftDeleteMedicalHistory(command SoftDeleteMedHistCommand) cqrs.CommandResult {
	exists, err := h.repo.ExistsByID(command.CTX, command.ID)
	if err != nil {
		return errorDeleteResult(command.ID, msgErrorProcessingData, err)
	}
	if !exists {
		return errorDeleteResult(command.ID, msgMedicalHistoryNotFound, err)
	}

	if err := h.repo.SoftDelete(command.CTX, command.ID); err != nil {
		return errorDeleteResult(command.ID, msgErrorProcessingData, err)
	}

	return successDeleteResult(command.ID, msgMedicalHistorySoftDeleted)
}

func (h *medicalHistoryCommandHandlers) HardDeleteMedicalHistory(command HardDeleteMedHistCommand) cqrs.CommandResult {
	exists, err := h.repo.ExistsByID(command.CTX, command.ID)
	if err != nil {
		return errorDeleteResult(command.ID, msgErrorProcessingData, err)
	}
	if !exists {
		return errorDeleteResult(command.ID, msgMedicalHistoryNotFound, nil)
	}

	if err := h.repo.HardDelete(command.CTX, command.ID); err != nil {
		return errorDeleteResult(command.ID, msgErrorProcessingData, err)
	}

	return successDeleteResult(command.ID, msgMedicalHistoryHardDeleted)
}
