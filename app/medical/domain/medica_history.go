package domain

import "time"

type MedicalHistory struct {
	ID          int
	PetID       int
	Date        time.Time
	Description *string
	VetID       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
