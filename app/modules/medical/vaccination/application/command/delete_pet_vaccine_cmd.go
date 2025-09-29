package command

import (
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type DeleteVaccinationCommand struct {
	VaccinationID vo.VaccinationID
	Reason        string
	DeletedBy     vo.EmployeeID
}

type ScheduleVaccineApptCommand struct {
	PetID                vo.PetID
	VaccineName          string
	ScheduledDate        time.Time
	ScheduledBy          vo.EmployeeID
	AssignedVeterinarian vo.EmployeeID
	Notes                *string
}

type RecordAdverseReactionCommand struct {
	VaccinationID  vo.VaccinationID
	PetID          vo.PetID
	ReactionType   AdverseReactionType
	Severity       ReactionSeverity
	Symptoms       []string
	OnsetTime      time.Time // Tiempo después de la vacunación
	TreatmentGiven *string
	ReportedBy     vo.EmployeeID
	Notes          *string
}

type AdverseReactionType string

const (
	AdverseReactionAllergic      AdverseReactionType = "allergic"
	AdverseReactionAnaphylaxis   AdverseReactionType = "anaphylaxis"
	AdverseReactionLocalSwelling AdverseReactionType = "local_swelling"
	AdverseReactionFever         AdverseReactionType = "fever"
	AdverseReactionLethargy      AdverseReactionType = "lethargy"
	AdverseReactionVomiting      AdverseReactionType = "vomiting"
	AdverseReactionOther         AdverseReactionType = "other"
)

type ReactionSeverity string

const (
	ReactionSeverityMild     ReactionSeverity = "mild"
	ReactionSeverityModerate ReactionSeverity = "moderate"
	ReactionSeveritySevere   ReactionSeverity = "severe"
	ReactionSeverityCritical ReactionSeverity = "critical"
)

type BatchVaccinationCommand struct {
	PetIDs           []vo.PetID
	VaccineName      string
	VaccineType      string
	AdministeredDate time.Time
	AdministeredBy   vo.EmployeeID
	BatchNumber      string
	Notes            *string
}

type CancelScheduledVaccinationCommand struct {
	AppointmentID vo.VaccinationID
	PetID         vo.PetID
	Reason        string
	CancelledBy   vo.EmployeeID
}

type MarkVaccinationAsCompletedCommand struct {
	AppointmentID vo.VaccinationID
	PetID         vo.PetID
	ActualDate    time.Time
	BatchNumber   string
	Notes         *string
	CompletedBy   vo.EmployeeID
}

type UpdateVaccinationProtocolCommand struct {
	PetID               vo.PetID
	ExcludedVaccines    []string
	ExclusionReasons    map[string]string
	SpecialInstructions *string
	ModifiedBy          vo.EmployeeID
}

type RecalculateVaccinationScheduleCommand struct {
	PetID                  vo.PetID
	ReasonForRecalculation string
	RequestedBy            vo.EmployeeID
}
