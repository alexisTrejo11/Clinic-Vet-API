package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	apperror "clinic-vet-api/app/shared/error/application"
)

const (
	msgMedicalSessionCreated         = "Medical history created successfully"
	msgMedicalSessionUpdated         = "Medical history updated successfully"
	msgMedicalSessionSoftDeleted     = "Medical history deleted successfully (soft delete)"
	msgMedicalSessionHardDeleted     = "Medical history permanently deleted"
	msgMedicalSessionNotFound        = "Medical history not found"
	msgMedicalSessionDateConflict    = "A medical history already exists for this pet on the specified date"
	msgMedicalSessionNewDateConflict = "A medical history already exists for this pet on the new specified date"
	msgErrorProcessingData           = "Error processing data: "
)

func MedicalNotFoundErr(id valueobject.MedSessionID) error {
	return apperror.EntityNotFoundValidationError("MedicalSession", "id", id.String())
}

func successCreateResult(entity medical.MedicalSession) cqrs.CommandResult {
	return *cqrs.SuccessResult(entity.ID().String(), msgMedicalSessionCreated)
}

func errorCreateResult(message string, err error) cqrs.CommandResult {
	return *cqrs.FailureResult(message, err)
}

func successUpdateResult(entity medical.MedicalSession) cqrs.CommandResult {
	return *cqrs.SuccessResult(entity.ID().String(), msgMedicalSessionUpdated)
}

func errorUpdateResult(message string, err error) cqrs.CommandResult {
	return *cqrs.FailureResult(message, err)
}
func successDeleteResult(id valueobject.MedSessionID, message string) cqrs.CommandResult {
	return *cqrs.SuccessResult(id.String(), message)
}
func errorDeleteResult(id valueobject.MedSessionID, message string, err error) cqrs.CommandResult {
	return *cqrs.FailureResult(message, err)
}
