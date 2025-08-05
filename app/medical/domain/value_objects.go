package mhDomain

import (
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type MedHistoryId struct {
	id shared.IntegerId
}

func NewMedHistoryId(value any) (MedHistoryId, error) {
	id, err := shared.NewIntegerId(value)
	if err != nil {
		return MedHistoryId{}, fmt.Errorf("invalid MedHistoryId: %w", err)
	}

	return MedHistoryId{id: id}, nil
}

func (m MedHistoryId) GetValue() int {
	return m.id.GetValue()
}

func (m MedHistoryId) String() string {
	return m.id.String()
}

func (m MedHistoryId) Equals(other MedHistoryId) bool {
	return m.id.GetValue() == other.id.GetValue()
}

type VisitReason string

const (
	RoutineCheckup VisitReason = "Routine Checkup"
	Vaccination    VisitReason = "Vaccination"
	Injury         VisitReason = "Injury"
	Illness        VisitReason = "Illness"
)

type VisitType string

const (
	PhysicalExam   VisitType = "Physical Exam"
	Surgery        VisitType = "Surgery"
	EmergencyVisit VisitType = "Emergency Visit"
	FollowUp       VisitType = "Follow-up"
)

type PetCondition string

const (
	stable   PetCondition = "Stable"
	critical PetCondition = "Critical"
	fair     PetCondition = "Fair"
)

func (v VisitReason) ToString() string {
	return string(v)
}

func (v VisitType) ToString() string {
	return string(v)
}

func (p PetCondition) ToString() string {
	return string(p)
}

func (v VisitReason) IsValid() bool {
	switch v {
	case RoutineCheckup, Vaccination, Injury, Illness:
		return true
	default:
		return false
	}
}

func (v VisitType) IsValid() bool {
	switch v {
	case PhysicalExam, Surgery, EmergencyVisit, FollowUp:
		return true
	default:
		return false
	}
}

func (p PetCondition) IsValid() bool {
	switch p {
	case stable, critical, fair:
		return true
	default:
		return false
	}
}

func (v VisitReason) Validate() error {
	if !v.IsValid() {
		return fmt.Errorf("invalid visit reason: %s", v)
	}
	return nil
}

func (v VisitType) Validate() error {
	if !v.IsValid() {
		return fmt.Errorf("invalid visit type: %s", v)
	}
	return nil
}

func (p PetCondition) Validate() error {
	if !p.IsValid() {
		return fmt.Errorf("invalid pet condition: %s", p)
	}
	return nil
}

func NewVisitReason(value string) (VisitReason, error) {
	reason := VisitReason(value)
	if !reason.IsValid() {
		return "", fmt.Errorf("invalid visit reason: %s", value)
	}
	return reason, nil
}

func NewVisitType(value string) (VisitType, error) {
	visitType := VisitType(value)
	if !visitType.IsValid() {
		return "", fmt.Errorf("invalid visit type: %s", value)
	}
	return visitType, nil
}

func NewPetCondition(value string) (PetCondition, error) {
	condition := PetCondition(value)
	if !condition.IsValid() {
		return "", fmt.Errorf("invalid pet condition: %s", value)
	}
	return condition, nil
}
