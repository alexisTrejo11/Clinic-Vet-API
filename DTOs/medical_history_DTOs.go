package DTOs

import "time"

type MedicalHistoryInsertDTO struct {
	PetID       int32
	Date        time.Time
	Description string
	VetID       int32
}

type MedicalHistoryDTO struct {
	ID          int32
	PetID       int32
	Date        time.Time
	Description string
	VetID       int32
}

type MedicalHistoryNamedDTO struct {
	Pet          string
	Date         time.Time
	Description  string
	Veterinarian string
}

type MedicalHistoryUpdateDTO struct {
	ID          int32
	PetID       int32
	Date        time.Time
	Description string
	VetID       int32
}
