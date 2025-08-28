package mhDTOs

import "time"

type MedHistResponse struct {
	Id          int
	PetId       int
	Date        time.Time
	Diagnosis   string
	VisitType   string
	VisitReason string
	Notes       *string
	Condition   string
	Treatment   string
	VetId       int
	OwnerId     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MedHistResponseDetail struct {
	Id           int
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
	Id        int
	FirstName string
	LastName  string
}

type VetDetails struct {
	Id        int
	FirstName string
	Specialty string
	LastName  string
}

type PetDetails struct {
	Id      int
	Name    string
	Species string
	Breed   string
	Age     int
	Weight  float64
}
