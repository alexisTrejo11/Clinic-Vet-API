package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"context"
	"time"

	p "clinic-vet-api/app/shared/page"
)

type MedicalSessionRepository interface {
	FindByID(ctx context.Context, medicalSessionID valueobject.MedSessionID) (*medical.MedicalSession, error)
	FindByIDAndCustomerID(ctx context.Context, medicalSessionID valueobject.MedSessionID, customerID valueobject.CustomerID) (*medical.MedicalSession, error)
	FindByIDAndPetID(ctx context.Context, medicalSessionID valueobject.MedSessionID, petID valueobject.PetID) (*medical.MedicalSession, error)
	FindByIDAndEmployeeID(ctx context.Context, medicalSessionID valueobject.MedSessionID, employeeID valueobject.EmployeeID) (*medical.MedicalSession, error)

	FindByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID, PaginationRequest p.PaginationRequest) (p.Page[medical.MedicalSession], error)
	FindByPetID(ctx context.Context, petID valueobject.PetID, PaginationRequest p.PaginationRequest) (p.Page[medical.MedicalSession], error)
	FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID, PaginationRequest p.PaginationRequest) (p.Page[medical.MedicalSession], error)

	FindBySpecification(ctx context.Context, spec specification.MedicalSessionSpecification) (p.Page[medical.MedicalSession], error)

	FindRecentByPetID(ctx context.Context, petID valueobject.PetID, limit int) ([]medical.MedicalSession, error)
	FindByDateRange(ctx context.Context, startDate, endDate time.Time, PaginationRequest p.PaginationRequest) (p.Page[medical.MedicalSession], error)
	FindByPetAndDateRange(ctx context.Context, petID valueobject.PetID, startDate, endDate time.Time) ([]medical.MedicalSession, error)
	FindByDiagnosis(ctx context.Context, diagnosis string, PaginationRequest p.PaginationRequest) (p.Page[medical.MedicalSession], error)

	ExistsByID(ctx context.Context, medicalSessionID valueobject.MedSessionID) (bool, error)
	ExistsByPetAndDate(ctx context.Context, petID valueobject.PetID, date time.Time) (bool, error)

	Save(ctx context.Context, medSession *medical.MedicalSession) error
	SoftDelete(ctx context.Context, medicalSessionID valueobject.MedSessionID) error
	HardDelete(ctx context.Context, medicalSessionID valueobject.MedSessionID) error

	CountByPetID(ctx context.Context, petID valueobject.PetID) (int64, error)
	CountByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID) (int64, error)
	CountByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (int64, error)
	CountByDateRange(ctx context.Context, startDate, endDate time.Time) (int64, error)
}
