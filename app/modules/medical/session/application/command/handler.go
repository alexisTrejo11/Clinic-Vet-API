package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type MedicalSessionCommandHandlers struct {
	repo repository.MedicalSessionRepository
}

func NewMedicalSessionCommandHandlers(repo repository.MedicalSessionRepository) *MedicalSessionCommandHandlers {
	return &MedicalSessionCommandHandlers{repo: repo}
}

func (h *MedicalSessionCommandHandlers) CreateMedicalSession(ctx context.Context, cmd CreateMedSessionCommand) cqrs.CommandResult {
	entity := cmd.ToEntity()
	if err := h.repo.Save(ctx, &entity); err != nil {
		return errorCreateResult(msgErrorProcessingData, err)
	}

	return successCreateResult(entity)
}

func (h *MedicalSessionCommandHandlers) UpdateMedicalSession(ctx context.Context, cmd UpdateMedSessionCommand) cqrs.CommandResult {
	existingEntity, err := h.repo.FindByID(ctx, cmd.ID)
	if err != nil {
		return errorUpdateResult(msgErrorProcessingData, err)
	}

	updatedEntity := applyCommandUpdates(cmd, *existingEntity)
	if err := h.repo.Save(ctx, updatedEntity); err != nil {
		return errorUpdateResult(msgErrorProcessingData, err)
	}

	return successUpdateResult(*updatedEntity)
}

func (h *MedicalSessionCommandHandlers) DeleteMedSessionCommand(ctx context.Context, cmd DeleteMedSessionCommand) cqrs.CommandResult {
	if err := h.valdiateExistingMedSession(ctx, cmd.ID); err != nil {
		return errorDeleteResult(cmd.ID, msgMedicalSessionNotFound, err)
	}

	if err := h.repo.Delete(ctx, cmd.ID, cmd.IsHardDelete); err != nil {
		return errorDeleteResult(cmd.ID, msgErrorProcessingData, err)
	}

	return successDeleteResult(cmd.ID, msgMedicalSessionSoftDeleted)
}
func (h *MedicalSessionCommandHandlers) valdiateExistingMedSession(ctx context.Context, medHistID valueobject.MedSessionID) error {
	exists, err := h.repo.ExistsByID(ctx, medHistID)
	if err != nil {
		return err
	}
	if !exists {
		return MedicalNotFoundErr(medHistID)
	}
	return nil
}
