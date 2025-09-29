package command

import (
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type GenerateVaccinationPlanCommand struct {
	PetID          vo.PetID
	SinceStartPlan time.Time
}
