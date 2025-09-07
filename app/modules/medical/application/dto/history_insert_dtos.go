// Package dto contains Data Transfer Objects for medical history operations.
package dto

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/go-playground/validator/v10"
)

type (
	VisitReason  string
	VisitType    string
	PetCondition string
)

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
	PetID       valueobject.PetID   `json:"petID" binding:"required"`
	OwnerID     valueobject.OwnerID `json:"ownerID" binding:"required"`
	Date        time.Time           `json:"date" binding:"required"`
	Diagnosis   string              `json:"diagnosis" binding:"required,min=3,max=255"`
	Treatment   string              `json:"treatment" binding:"required,min=3,max=255"`
	Condition   enum.PetCondition   `json:"condition" binding:"required,validPetCondition"`
	VisitReason enum.VisitReason    `json:"visitReason" binding:"required,validVisitReason"`
	VisitType   enum.VisitType      `json:"visitType" binding:"required,validVisitType"`
	Notes       *string             `json:"notes" binding:"omitempty,max=1000"`
	VetID       valueobject.VetID   `json:"vetID" binding:"required"`
}

type MedicalHistoryUpdate struct {
	PetID       *int          `json:"petID" binding:"omitempty,gt=0"`
	Date        *time.Time    `json:"date" binding:"omitempty"`
	VisitReason *VisitReason  `json:"visitReason" binding:"omitempty,validVisitReason"`
	VisitType   *VisitType    `json:"visitType" binding:"omitempty,validVisitType"`
	Notes       *string       `json:"notes" binding:"omitempty,max=1000"`
	Diagnosis   *string       `json:"diagnosis" binding:"omitempty,min=3,max=255"`
	Treatment   *string       `json:"treatment" binding:"omitempty,min=3,max=255"`
	Condition   *PetCondition `json:"condition" binding:"omitempty,validPetCondition"`
	VetID       *int          `json:"vetID" binding:"omitempty,gt=0"`
	OwnerID     *int          `json:"ownerID" binding:"omitempty,gt=0"`
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
