package enum

import (
	"fmt"
	"strings"
)

// VisitReason represents the reason for a veterinary visit
type VisitReason string

const (
	VisitReasonRoutineCheckup VisitReason = "routine_checkup"
	VisitReasonVaccination    VisitReason = "vaccination"
	VisitReasonInjury         VisitReason = "injury"
	VisitReasonIllness        VisitReason = "illness"
	VisitReasonDental         VisitReason = "dental"
	VisitReasonGrooming       VisitReason = "grooming"
	VisitReasonBehavior       VisitReason = "behavior"
	VisitReasonNutrition      VisitReason = "nutrition"
	VisitReasonFollowUp       VisitReason = "follow_up"
	VisitReasonEmergency      VisitReason = "emergency"
)

// VisitReason constants and methods
var (
	ValidVisitReasons = []VisitReason{
		VisitReasonRoutineCheckup,
		VisitReasonVaccination,
		VisitReasonInjury,
		VisitReasonIllness,
		VisitReasonDental,
		VisitReasonGrooming,
		VisitReasonBehavior,
		VisitReasonNutrition,
		VisitReasonFollowUp,
		VisitReasonEmergency,
	}

	visitReasonMap = map[string]VisitReason{
		"routine_checkup": VisitReasonRoutineCheckup,
		"routine checkup": VisitReasonRoutineCheckup,
		"checkup":         VisitReasonRoutineCheckup,
		"vaccination":     VisitReasonVaccination,
		"injury":          VisitReasonInjury,
		"illness":         VisitReasonIllness,
		"dental":          VisitReasonDental,
		"grooming":        VisitReasonGrooming,
		"behavior":        VisitReasonBehavior,
		"nutrition":       VisitReasonNutrition,
		"follow_up":       VisitReasonFollowUp,
		"follow up":       VisitReasonFollowUp,
		"followup":        VisitReasonFollowUp,
		"emergency":       VisitReasonEmergency,
	}

	visitReasonDisplayNames = map[VisitReason]string{
		VisitReasonRoutineCheckup: "Routine Checkup",
		VisitReasonVaccination:    "Vaccination",
		VisitReasonInjury:         "Injury",
		VisitReasonIllness:        "Illness",
		VisitReasonDental:         "Dental Care",
		VisitReasonGrooming:       "Grooming",
		VisitReasonBehavior:       "Behavior Consultation",
		VisitReasonNutrition:      "Nutrition Consultation",
		VisitReasonFollowUp:       "Follow-up Visit",
		VisitReasonEmergency:      "Emergency Visit",
	}
)

func (vr VisitReason) IsValid() bool {
	_, exists := visitReasonMap[string(vr)]
	return exists
}

func ParseVisitReason(reason string) (VisitReason, error) {
	normalized := normalizeInput(reason)
	if val, exists := visitReasonMap[normalized]; exists {
		return val, nil
	}
	return "", fmt.Errorf("invalid visit reason: %s", reason)
}

func MustParseVisitReason(reason string) VisitReason {
	parsed, err := ParseVisitReason(reason)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (vr VisitReason) String() string {
	return string(vr)
}

func (vr VisitReason) DisplayName() string {
	if displayName, exists := visitReasonDisplayNames[vr]; exists {
		return displayName
	}
	return "Unknown Reason"
}

func (vr VisitReason) Values() []VisitReason {
	return ValidVisitReasons
}

func (vr VisitReason) IsRoutine() bool {
	return vr == VisitReasonRoutineCheckup ||
		vr == VisitReasonVaccination ||
		vr == VisitReasonGrooming
}

func (vr VisitReason) IsUrgent() bool {
	return vr == VisitReasonEmergency ||
		vr == VisitReasonInjury ||
		vr == VisitReasonIllness
}

// VisitType represents the type of veterinary visit
type VisitType string

const (
	VisitTypePhysicalExam   VisitType = "physical_exam"
	VisitTypeSurgery        VisitType = "surgery"
	VisitTypeEmergencyVisit VisitType = "emergency_visit"
	VisitTypeFollowUp       VisitType = "follow_up"
	VisitTypeConsultation   VisitType = "consultation"
	VisitTypeProcedure      VisitType = "procedure"
	VisitTypeVaccination    VisitType = "vaccination"
	VisitTypeDental         VisitType = "dental"
	VisitTypeGrooming       VisitType = "grooming"
)

// VisitType constants and methods
var (
	ValidVisitTypes = []VisitType{
		VisitTypePhysicalExam,
		VisitTypeSurgery,
		VisitTypeEmergencyVisit,
		VisitTypeFollowUp,
		VisitTypeConsultation,
		VisitTypeProcedure,
		VisitTypeVaccination,
		VisitTypeDental,
		VisitTypeGrooming,
	}

	visitTypeMap = map[string]VisitType{
		"physical_exam":   VisitTypePhysicalExam,
		"physical exam":   VisitTypePhysicalExam,
		"exam":            VisitTypePhysicalExam,
		"surgery":         VisitTypeSurgery,
		"emergency_visit": VisitTypeEmergencyVisit,
		"emergency visit": VisitTypeEmergencyVisit,
		"emergency":       VisitTypeEmergencyVisit,
		"follow_up":       VisitTypeFollowUp,
		"follow up":       VisitTypeFollowUp,
		"followup":        VisitTypeFollowUp,
		"consultation":    VisitTypeConsultation,
		"consult":         VisitTypeConsultation,
		"procedure":       VisitTypeProcedure,
		"vaccination":     VisitTypeVaccination,
		"dental":          VisitTypeDental,
		"grooming":        VisitTypeGrooming,
	}

	visitTypeDisplayNames = map[VisitType]string{
		VisitTypePhysicalExam:   "Physical Examination",
		VisitTypeSurgery:        "Surgery",
		VisitTypeEmergencyVisit: "Emergency Visit",
		VisitTypeFollowUp:       "Follow-up Visit",
		VisitTypeConsultation:   "Consultation",
		VisitTypeProcedure:      "Medical Procedure",
		VisitTypeVaccination:    "Vaccination",
		VisitTypeDental:         "Dental Procedure",
		VisitTypeGrooming:       "Grooming Service",
	}

	visitTypeDurations = map[VisitType]int{
		VisitTypePhysicalExam:   30,
		VisitTypeSurgery:        120,
		VisitTypeEmergencyVisit: 60,
		VisitTypeFollowUp:       20,
		VisitTypeConsultation:   45,
		VisitTypeProcedure:      60,
		VisitTypeVaccination:    15,
		VisitTypeDental:         45,
		VisitTypeGrooming:       60,
	}
)

func (vt VisitType) IsValid() bool {
	_, exists := visitTypeMap[string(vt)]
	return exists
}

func ParseVisitType(visitType string) (VisitType, error) {
	normalized := normalizeInput(visitType)
	if val, exists := visitTypeMap[normalized]; exists {
		return val, nil
	}
	return "", fmt.Errorf("invalid visit type: %s", visitType)
}

func MustParseVisitType(visitType string) VisitType {
	parsed, err := ParseVisitType(visitType)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (vt VisitType) String() string {
	return string(vt)
}

func (vt VisitType) DisplayName() string {
	if displayName, exists := visitTypeDisplayNames[vt]; exists {
		return displayName
	}
	return "Unknown Visit Type"
}

func (vt VisitType) Values() []VisitType {
	return ValidVisitTypes
}

func (vt VisitType) DefaultDuration() int {
	if duration, exists := visitTypeDurations[vt]; exists {
		return duration
	}
	return 30 // default duration in minutes
}

func (vt VisitType) RequiresMedicalStaff() bool {
	return vt != VisitTypeGrooming
}

func (vt VisitType) IsSurgical() bool {
	return vt == VisitTypeSurgery || vt == VisitTypeDental
}

// PetCondition represents the condition of a pet during a visit
type PetCondition string

const (
	PetConditionStable   PetCondition = "stable"
	PetConditionCritical PetCondition = "critical"
	PetConditionFair     PetCondition = "fair"
	PetConditionGood     PetCondition = "good"
	PetConditionSerious  PetCondition = "serious"
	PetConditionGuarded  PetCondition = "guarded"
)

// PetCondition constants and methods
var (
	ValidPetConditions = []PetCondition{
		PetConditionStable,
		PetConditionCritical,
		PetConditionFair,
		PetConditionGood,
		PetConditionSerious,
		PetConditionGuarded,
	}

	petConditionMap = map[string]PetCondition{
		"stable":   PetConditionStable,
		"critical": PetConditionCritical,
		"fair":     PetConditionFair,
		"good":     PetConditionGood,
		"serious":  PetConditionSerious,
		"guarded":  PetConditionGuarded,
	}

	petConditionDisplayNames = map[PetCondition]string{
		PetConditionStable:   "Stable",
		PetConditionCritical: "Critical",
		PetConditionFair:     "Fair",
		PetConditionGood:     "Good",
		PetConditionSerious:  "Serious",
		PetConditionGuarded:  "Guarded",
	}

	petConditionSeverity = map[PetCondition]int{
		PetConditionGood:     1,
		PetConditionStable:   2,
		PetConditionFair:     3,
		PetConditionGuarded:  4,
		PetConditionSerious:  5,
		PetConditionCritical: 6,
	}
)

func (pc PetCondition) IsValid() bool {
	_, exists := petConditionMap[string(pc)]
	return exists
}

func ParsePetCondition(condition string) (PetCondition, error) {
	normalized := normalizeInput(condition)
	if val, exists := petConditionMap[normalized]; exists {
		return val, nil
	}
	return "", fmt.Errorf("invalid pet condition: %s", condition)
}

func MustParsePetCondition(condition string) PetCondition {
	parsed, err := ParsePetCondition(condition)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (pc PetCondition) String() string {
	return string(pc)
}

func (pc PetCondition) DisplayName() string {
	if displayName, exists := petConditionDisplayNames[pc]; exists {
		return displayName
	}
	return "Unknown Condition"
}

func (pc PetCondition) Values() []PetCondition {
	return ValidPetConditions
}

func (pc PetCondition) Severity() int {
	if severity, exists := petConditionSeverity[pc]; exists {
		return severity
	}
	return 0
}

func (pc PetCondition) IsCritical() bool {
	return pc == PetConditionCritical || pc == PetConditionSerious
}

func (pc PetCondition) IsStable() bool {
	return pc == PetConditionStable || pc == PetConditionGood || pc == PetConditionFair
}

func (pc PetCondition) RequiresImmediateAttention() bool {
	return pc.Severity() >= 4 // Guarded, Serious, Critical
}

// Utility functions
func normalizeInput(input string) string {
	input = strings.TrimSpace(strings.ToLower(input))
	input = strings.ReplaceAll(input, " ", "_")
	return input
}

func GetAllVisitReasons() []VisitReason {
	return ValidVisitReasons
}

func GetAllVisitTypes() []VisitType {
	return ValidVisitTypes
}

func GetAllPetConditions() []PetCondition {
	return ValidPetConditions
}

func GetUrgentVisitReasons() []VisitReason {
	return []VisitReason{
		VisitReasonEmergency,
		VisitReasonInjury,
		VisitReasonIllness,
	}
}

func GetRoutineVisitReasons() []VisitReason {
	return []VisitReason{
		VisitReasonRoutineCheckup,
		VisitReasonVaccination,
		VisitReasonGrooming,
		VisitReasonDental,
	}
}
