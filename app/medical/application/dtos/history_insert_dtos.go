package mhDTOs

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type VisitReason string
type VisitType string
type PetCondition string

const (
	RoutineCheckup VisitReason = "Routine Checkup"
	Vaccination    VisitReason = "Vaccination"
	Injury         VisitReason = "Injury"
	Illness        VisitReason = "Illness"
)

const (
	PhysicalExam   VisitType = "Physical Exam"
	Surgery        VisitType = "Surgery"
	EmergencyVisit VisitType = "Emergency Visit"
	FollowUp       VisitType = "Follow-up"
)

const (
	Stable   PetCondition = "Stable"
	Critical PetCondition = "Critical"
	Fair     PetCondition = "Fair"
)

type MedicalHistoryCreate struct {
	PetId       int          `json:"petId" binding:"required"`
	OwnerId     int          `json:"ownerId" binding:"required"`
	Date        time.Time    `json:"date" binding:"required"`
	Diagnosis   string       `json:"diagnosis" binding:"required,min=3,max=255"`
	Treatment   string       `json:"treatment" binding:"required,min=3,max=255"`
	Condition   PetCondition `json:"condition" binding:"required,validPetCondition"`
	VisitReason VisitReason  `json:"visitReason" binding:"required,validVisitReason"`
	VisitType   VisitType    `json:"visitType" binding:"required,validVisitType"`
	Notes       *string      `json:"notes" binding:"omitempty,max=1000"`
	VetId       int          `json:"vetId" binding:"required"`
}

type MedicalHistoryUpdate struct {
	PetId       *int          `json:"petId" binding:"omitempty,gt=0"`
	Date        *time.Time    `json:"date" binding:"omitempty"`
	VisitReason *VisitReason  `json:"visitReason" binding:"omitempty,validVisitReason"`
	VisitType   *VisitType    `json:"visitType" binding:"omitempty,validVisitType"`
	Notes       *string       `json:"notes" binding:"omitempty,max=1000"`
	Diagnosis   *string       `json:"diagnosis" binding:"omitempty,min=3,max=255"`
	Treatment   *string       `json:"treatment" binding:"omitempty,min=3,max=255"`
	Condition   *PetCondition `json:"condition" binding:"omitempty,validPetCondition"`
	VetId       *int          `json:"vetId" binding:"omitempty,gt=0"`
	OwnerId     *int          `json:"ownerId" binding:"omitempty,gt=0"`
}

func IsValidVisitReason(fl validator.FieldLevel) bool {
	reason := VisitReason(fl.Field().String())
	switch reason {
	case RoutineCheckup, Vaccination, Injury, Illness:
		return true
	default:
		return false
	}
}

func IsValidVisitType(fl validator.FieldLevel) bool {
	visitType := VisitType(fl.Field().String())
	switch visitType {
	case PhysicalExam, Surgery, EmergencyVisit, FollowUp:
		return true
	default:
		return false
	}
}

func IsValidPetCondition(fl validator.FieldLevel) bool {
	condition := PetCondition(fl.Field().String())
	switch condition {
	case Stable, Critical, Fair:
		return true
	default:
		return false
	}
}
