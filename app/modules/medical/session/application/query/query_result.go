package query

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type MedSessionResult struct {
	ID               valueobject.MedSessionID
	EmployeeID       valueobject.EmployeeID
	VisitDate        time.Time
	VisitType        enum.VisitType
	ClinicService    enum.ClinicService
	Notes            *string
	PetDetailsResult PetDetailsResult
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type PetDetailsResult struct {
	PetID           valueobject.PetID
	Weight          *valueobject.Decimal
	HeartRate       *int32
	RespiratoryRate *int32
	Temperature     *valueobject.Decimal
	Condition       enum.PetCondition
	Diagnosis       string
	Medications     []string
	Symptoms        []string
	Treatment       string
	FollowUpDate    *time.Time
}

type MedSessionDetailResult struct {
	ID        valueobject.MedSessionID
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
