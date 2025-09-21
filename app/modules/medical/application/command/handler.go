package command

import (
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type MedicalSessionCommandHandlers interface {
	CreateMedicalSession(command CreateMedSessionCommand) cqrs.CommandResult
	UpdateMedicalSession(command UpdateMedSessionCommand) cqrs.CommandResult
	SoftDeleteMedicalSession(command SoftDeleteMedSessionCommand) cqrs.CommandResult
	HardDeleteMedicalSession(command HardDeleteMedSessionCommand) cqrs.CommandResult
}

type medicalSessionCommandHandlers struct {
	repo repository.MedicalSessionRepository
}

func NewMedicalSessionCommandHandlers(repo repository.MedicalSessionRepository) MedicalSessionCommandHandlers {
	return &medicalSessionCommandHandlers{
		repo: repo,
	}
}

// Validate Schedule??

func (h *medicalSessionCommandHandlers) CreateMedicalSession(command CreateMedSessionCommand) cqrs.CommandResult {
	entity, err := ToEntityFromCreate(&command)
	if err != nil {
		return errorCreateResult(msgErrorProcessingData, err)
	}

	if err := h.repo.Save(command.CTX, entity); err != nil {
		return errorCreateResult(msgErrorProcessingData, err)
	}

	return successCreateResult(*entity)
}

func (h *medicalSessionCommandHandlers) UpdateMedicalSession(command UpdateMedSessionCommand) cqrs.CommandResult {
	existingEntity, err := h.repo.FindByID(command.CTX, command.ID)
	if err != nil {
		return errorUpdateResult(msgErrorProcessingData, err)
	}

	updatedEntity, err := ToEntityFromUpdate(&command, existingEntity)
	if err != nil {
		return errorUpdateResult(msgErrorProcessingData, err)
	}

	if err := h.repo.Save(command.CTX, updatedEntity); err != nil {
		return errorUpdateResult(msgErrorProcessingData, err)
	}

	return successUpdateResult(*updatedEntity)
}

func (h *medicalSessionCommandHandlers) SoftDeleteMedicalSession(command SoftDeleteMedSessionCommand) cqrs.CommandResult {
	if err := h.valdiateExistingMedSession(command.CTX, command.ID); err != nil {
		return errorDeleteResult(command.ID, msgMedicalSessionNotFound, err)
	}

	if err := h.repo.SoftDelete(command.CTX, command.ID); err != nil {
		return errorDeleteResult(command.ID, msgErrorProcessingData, err)
	}

	return successDeleteResult(command.ID, msgMedicalSessionSoftDeleted)
}

func (h *medicalSessionCommandHandlers) HardDeleteMedicalSession(command HardDeleteMedSessionCommand) cqrs.CommandResult {
	if err := h.valdiateExistingMedSession(command.CTX, command.ID); err != nil {
		return errorDeleteResult(command.ID, msgMedicalSessionNotFound, err)
	}

	if err := h.repo.HardDelete(command.CTX, command.ID); err != nil {
		return errorDeleteResult(command.ID, msgErrorProcessingData, err)
	}

	return successDeleteResult(command.ID, msgMedicalSessionHardDeleted)
}

func (h *medicalSessionCommandHandlers) valdiateExistingMedSession(ctx context.Context, medHistID valueobject.MedSessionID) error {
	exists, err := h.repo.ExistsByID(ctx, medHistID)
	if err != nil {
		return err
	}
	if !exists {
		return MedicalNotFoundErr(medHistID)
	}
	return nil
}
