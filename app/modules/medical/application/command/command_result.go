package command

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/medical"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
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

type CreateMedHistResult struct {
	ID        valueobject.MedHistoryID
	CreatedAt time.Time
	Success   bool
	Message   string
	Error     error
}

type UpdateMedHistResult struct {
	ID        valueobject.MedHistoryID
	UpdatedAt time.Time
	Success   bool
	Message   string
	Error     error
}

type DeleteMedHistResult struct {
	ID      valueobject.MedHistoryID
	Success bool
	Message string
	Error   error
}

func successCreateResult(entity medical.MedicalHistory) *CreateMedHistResult {
	return &CreateMedHistResult{
		ID:        entity.ID(),
		CreatedAt: entity.CreatedAt(),
		Success:   true,
		Message:   msgMedicalHistoryCreated,
		Error:     nil,
	}
}

func errorCreateResult(message string, err error) *CreateMedHistResult {
	return &CreateMedHistResult{
		Success: false,
		Message: message,
		Error:   err,
	}
}

func successUpdateResult(entity medical.MedicalHistory) *UpdateMedHistResult {
	return &UpdateMedHistResult{
		ID:        entity.ID(),
		UpdatedAt: entity.UpdatedAt(),
		Success:   true,
		Message:   msgMedicalHistoryUpdated,
		Error:     nil,
	}
}

func errorUpdateResult(message string, err error) *UpdateMedHistResult {
	return &UpdateMedHistResult{
		Success: false,
		Message: message,
		Error:   err,
	}
}

func successDeleteResult(id valueobject.MedHistoryID, message string) *DeleteMedHistResult {
	return &DeleteMedHistResult{
		ID:      id,
		Success: true,
		Message: message,
		Error:   nil,
	}
}

func errorDeleteResult(id valueobject.MedHistoryID, message string, err error) *DeleteMedHistResult {
	return &DeleteMedHistResult{
		ID:      id,
		Success: false,
		Message: message,
		Error:   err,
	}
}
