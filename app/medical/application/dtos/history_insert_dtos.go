package dtos

import "time"

type MedicalHistoryCreate struct {
	PetID       int32
	Date        time.Time
	Description string
	VetID       int32
}

type MedicalHistoryUpdate struct {
	ID          int32
	PetID       int32
	Date        time.Time
	Description string
	VetID       int32
}
