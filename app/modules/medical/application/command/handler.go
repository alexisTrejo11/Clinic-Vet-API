package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type MedicalSessionCommandHandlers interface {
	CreateMedicalSession(ctx context.Context, command CreateMedSessionCommand) cqrs.CommandResult
	UpdateMedicalSession(ctx context.Context, command UpdateMedSessionCommand) cqrs.CommandResult
	SoftDeleteMedicalSession(ctx context.Context, command SoftDeleteMedSessionCommand) cqrs.CommandResult
	HardDeleteMedicalSession(ctx context.Context, command HardDeleteMedSessionCommand) cqrs.CommandResult
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

func (h *medicalSessionCommandHandlers) CreateMedicalSession(ctx context.Context, command CreateMedSessionCommand) cqrs.CommandResult {
	entity := command.ToEntity()
	if err := h.repo.Save(ctx, &entity); err != nil {
		return errorCreateResult(msgErrorProcessingData, err)
	}

	return successCreateResult(entity)
}

func (h *medicalSessionCommandHandlers) UpdateMedicalSession(ctx context.Context, command UpdateMedSessionCommand) cqrs.CommandResult {
	existingEntity, err := h.repo.FindByID(ctx, command.ID)
	if err != nil {
		return errorUpdateResult(msgErrorProcessingData, err)
	}

	updatedEntity := applyCommandUpdates(command, *existingEntity)
	if err := h.repo.Save(ctx, updatedEntity); err != nil {
		return errorUpdateResult(msgErrorProcessingData, err)
	}

	return successUpdateResult(*updatedEntity)
}

func (h *medicalSessionCommandHandlers) SoftDeleteMedicalSession(ctx context.Context, command SoftDeleteMedSessionCommand) cqrs.CommandResult {
	if err := h.valdiateExistingMedSession(ctx, command.ID); err != nil {
		return errorDeleteResult(command.ID, msgMedicalSessionNotFound, err)
	}

	if err := h.repo.SoftDelete(ctx, command.ID); err != nil {
		return errorDeleteResult(command.ID, msgErrorProcessingData, err)
	}

	return successDeleteResult(command.ID, msgMedicalSessionSoftDeleted)
}

func (h *medicalSessionCommandHandlers) HardDeleteMedicalSession(ctx context.Context, command HardDeleteMedSessionCommand) cqrs.CommandResult {
	if err := h.valdiateExistingMedSession(ctx, command.ID); err != nil {
		return errorDeleteResult(command.ID, msgMedicalSessionNotFound, err)
	}

	if err := h.repo.HardDelete(ctx, command.ID); err != nil {
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
