package mhDTOs

import "time"

type MedicalHistoryCreate struct {
	PetId       int       `json:"petId" binding:"required"`
	OwnerId     int       `json:"ownerId" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	VisitReason string    `json:"visitReason" binding:"required,min=3,max=255"`
	VisitType   string    `json:"visitType" binding:"required,oneof='Check-up' 'Vaccination' 'Surgery' 'Emergency' 'Follow-up'"`
	Description *string   `json:"description" binding:"omitempty,max=1000"`
	VetId       int       `json:"vetId" binding:"required"`
}

type MedicalHistoryUpdate struct {
	PetId       *int       `json:"petId" binding:"omitempty,gt=0"`
	Date        *time.Time `json:"date" binding:"omitempty"`
	VisitReason *string    `json:"visitReason" binding:"omitempty,min=3,max=255"`
	VisitType   *string    `json:"visitType" binding:"omitempty,oneof='Check-up' 'Vaccination' 'Surgery' 'Emergency' 'Follow-up'"`
	Description *string    `json:"description" binding:"omitempty,max=1000"`
	VetId       *int       `json:"vetId" binding:"omitempty,gt=0"`
	OwnerId     *int       `json:"ownerId" binding:"omitempty,gt=0"`
}
