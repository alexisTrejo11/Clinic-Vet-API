package query

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type MedHistoryResult struct {
	ID          valueobject.MedHistoryID
	PetID       valueobject.PetID
	CustomerID  valueobject.CustomerID
	EmployeeID  valueobject.EmployeeID
	Date        time.Time
	Diagnosis   string
	VisitType   string
	VisitReason string
	Notes       *string
	Condition   string
	Treatment   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MedHistoryDetailResult struct {
	ID        valueobject.MedHistoryID
	Date      time.Time
	Diagnosis string
	Notes     string
	Treatment string
	CreatedAt time.Time
	UpdatedAt time.Time
	Pet       PetResult
	Customer  CustomerResult
	Employee  EmployeeResult
}

type PetResult struct {
	ID      valueobject.PetID
	Name    string
	Species string
	Breed   string
	Age     int
	Weight  float64
}

type CustomerResult struct {
	ID        valueobject.CustomerID
	FirstName string
	LastName  string
	Email     string
}

type EmployeeResult struct {
	ID        valueobject.EmployeeID
	FirstName string
	LastName  string
	Specialty string
}
