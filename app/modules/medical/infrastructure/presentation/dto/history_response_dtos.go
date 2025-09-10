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
	CustomerID  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MedHistResponseDetail struct {
	ID        int
	Pet       PetDetails
	Customer  CustomerDetails
	Date      time.Time
	Diagnosis string
	Notes     string
	Treatment string
	Employee  EmployeeDetails
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CustomerDetails struct {
	ID        int
	FirstName string
	LastName  string
}

type EmployeeDetails struct {
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
