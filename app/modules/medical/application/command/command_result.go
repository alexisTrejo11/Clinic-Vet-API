package command

import (
	"clinic-vet-api/app/core/domain/entity/medical"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
)

const (
	msgMedicalHistoryCreated         = "Medical history created successfully"
	msgMedicalHistoryUpdated         = "Medical history updated successfully"
	msgMedicalHistorySoftDeleted     = "Medical history deleted successfully (soft delete)"
	msgMedicalHistoryHardDeleted     = "Medical history permanently deleted"
	msgMedicalHistoryNotFound        = "Medical history not found"
	msgMedicalHistoryDateConflict    = "A medical history already exists for this pet on the specified date"
	msgMedicalHistoryNewDateConflict = "A medical history already exists for this pet on the new specified date"
	msgErrorProcessingData           = "Error processing data: "
)

func successCreateResult(entity medical.MedicalHistory) cqrs.CommandResult {
	return *cqrs.SuccessResult(entity.ID().String(), msgMedicalHistoryCreated)
}

func errorCreateResult(message string, err error) cqrs.CommandResult {
	return *cqrs.FailureResult(message, err)
}

func successUpdateResult(entity medical.MedicalHistory) cqrs.CommandResult {
	return *cqrs.SuccessResult(entity.ID().String(), msgMedicalHistoryUpdated)
}

func errorUpdateResult(message string, err error) cqrs.CommandResult {
	return *cqrs.FailureResult(message, err)
}
func successDeleteResult(id valueobject.MedHistoryID, message string) cqrs.CommandResult {
	return *cqrs.SuccessResult(id.String(), message)
}
func errorDeleteResult(id valueobject.MedHistoryID, message string, err error) cqrs.CommandResult {
	return *cqrs.FailureResult(message, err)
}
