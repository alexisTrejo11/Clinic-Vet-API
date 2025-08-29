package dto

import "time"

type MedHistResponse struct {
	ID          int
	PetID       int
	Date        time.Time
	Diagnosis   string
	VisitType   string
	VisitReason string
	Notes       *string
	Condition   string
	Treatment   string
	VetID       int
	OwnerID     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MedHistResponseDetail struct {
	ID           int
	Pet          PetDetails
	Owner        OwnerDetails
	Date         time.Time
	Diagnosis    string
	Notes        string
	Treatment    string
	Veterinarian VetDetails
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type OwnerDetails struct {
	ID        int
	FirstName string
	LastName  string
}

type VetDetails struct {
	ID        int
	FirstName string
	Specialty string
	LastName  string
}

type PetDetails struct {
	ID      int
	Name    string
	Species string
	Breed   string
	Age     int
	Weight  float64
}
