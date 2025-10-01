package repository

import (
	med "clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/specification"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"context"
	"time"

	p "clinic-vet-api/app/shared/page"
)

type MedicalSessionRepository interface {
	FindByID(ctx context.Context, medSessionID vo.MedSessionID) (*med.MedicalSession, error)
	FindByIDAndCustomerID(ctx context.Context, medSessionID vo.MedSessionID, customerID vo.CustomerID) (*med.MedicalSession, error)
	FindByIDAndPetID(ctx context.Context, medSessionID vo.MedSessionID, petID vo.PetID) (*med.MedicalSession, error)
	FindByIDAndEmployeeID(ctx context.Context, medSessionID vo.MedSessionID, employeeID vo.EmployeeID) (*med.MedicalSession, error)

	FindByEmployeeID(ctx context.Context, employeeID vo.EmployeeID, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error)
	FindByPetID(ctx context.Context, petID vo.PetID, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error)
	FindByCustomerID(ctx context.Context, customerID vo.CustomerID, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error)

	FindBySpecification(ctx context.Context, spec specification.MedicalSessionSpecification) (p.Page[med.MedicalSession], error)

	FindRecentByPetID(ctx context.Context, petID vo.PetID, limit int) ([]med.MedicalSession, error)
	FindByDateRange(ctx context.Context, startDate, endDate time.Time, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error)
	FindByPetAndDateRange(ctx context.Context, petID vo.PetID, startDate, endDate time.Time) ([]med.MedicalSession, error)
	FindByDiagnosis(ctx context.Context, diagnosis string, pagination p.PaginationRequest) (p.Page[med.MedicalSession], error)

	ExistsByID(ctx context.Context, medSessionID vo.MedSessionID) (bool, error)
	ExistsByPetAndDate(ctx context.Context, petID vo.PetID, date time.Time) (bool, error)

	Save(ctx context.Context, medSession *med.MedicalSession) error
	Delete(ctx context.Context, medSessionID vo.MedSessionID, isHard bool) error
}

type DewormRepository interface {
	FindByID(ctx context.Context, dewormationID vo.DewormID) (*med.PetDeworming, error)
	FindByIDAndPetID(ctx context.Context, dewormationID vo.DewormID, petID vo.PetID) (*med.PetDeworming, error)
	FindByIDAndEmployeeID(ctx context.Context, dewormationID vo.DewormID, employeeID vo.EmployeeID) (*med.PetDeworming, error)

	FindByPetID(ctx context.Context, petID vo.PetID, pagination p.PaginationRequest) (p.Page[med.PetDeworming], error)
	FindByPetIDs(ctx context.Context, petIDs []vo.PetID, pagination p.PaginationRequest) (p.Page[med.PetDeworming], error)
	FindByEmployeeID(ctx context.Context, employeeID vo.EmployeeID, pagination p.PaginationRequest) (p.Page[med.PetDeworming], error)
	FindByDateRange(ctx context.Context, startDate, endDate time.Time, pagination p.PaginationRequest) (p.Page[med.PetDeworming], error)

	Save(ctx context.Context, dewormation med.PetDeworming) (med.PetDeworming, error)
	Delete(ctx context.Context, dewormationID vo.DewormID, isHard bool) error
}

type VaccinationRepository interface {
	FindByID(ctx context.Context, vaccinationID vo.VaccinationID) (*med.PetVaccination, error)
	FindByIDAndPetID(ctx context.Context, vaccinationID vo.VaccinationID, petID vo.PetID) (*med.PetVaccination, error)
	FindByIDAndEmployeeID(ctx context.Context, vaccinationID vo.VaccinationID, employeeID vo.EmployeeID) (*med.PetVaccination, error)
	FindByPetID(ctx context.Context, petID vo.PetID, pagination p.PaginationRequest) (p.Page[med.PetVaccination], error)
	FindByEmployeeID(ctx context.Context, employeeID vo.EmployeeID, pagination p.PaginationRequest) (p.Page[med.PetVaccination], error)
	FindByDateRange(ctx context.Context, startDate, endDate time.Time, pagination p.PaginationRequest) (p.Page[med.PetVaccination], error)

	FindRecentByPetID(ctx context.Context, petID vo.PetID, days int) ([]med.PetVaccination, error)
	FindAllByPetID(ctx context.Context, petID vo.PetID) ([]med.PetVaccination, error)

	Save(ctx context.Context, vaccination med.PetVaccination) (med.PetVaccination, error)
	Delete(ctx context.Context, vaccinationID vo.VaccinationID, isHard bool) error
}
