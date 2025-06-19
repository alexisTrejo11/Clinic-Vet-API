package dtos

import "time"

type MedicalHistoryResponse struct {
	ID          int32
	PetID       int32
	Date        time.Time
	Description string
	VetID       int32
}

type MedicalHistoryResponseNamed struct {
	Pet          string
	Date         time.Time
	Description  string
	Veterinarian string
}
