package repository

import (
	"context"
	"time"

	"clinic-vet-api/app/core/domain/entity/medical"
	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
	p "clinic-vet-api/app/shared/page"
)

type MedicalHistoryRepository interface {
	FindBySpecification(ctx context.Context, spec specification.MedicalHistorySpecification) (p.Page[medical.MedicalHistory], error)
	FindByID(ctx context.Context, medicalHistoryID valueobject.MedHistoryID) (medical.MedicalHistory, error)

	FindAll(ctx context.Context, pageInput p.PageInput) (p.Page[medical.MedicalHistory], error)
	FindByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID, pageInput p.PageInput) (p.Page[medical.MedicalHistory], error)
	FindByPetID(ctx context.Context, petID valueobject.PetID, pageInput p.PageInput) (p.Page[medical.MedicalHistory], error)
	FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID, pageInput p.PageInput) (p.Page[medical.MedicalHistory], error)

	FindRecentByPetID(ctx context.Context, petID valueobject.PetID, limit int) (medical.MedicalHistory, error)
	FindByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput p.PageInput) (p.Page[medical.MedicalHistory], error)
	FindByPetAndDateRange(ctx context.Context, petID valueobject.PetID, startDate, endDate time.Time) (medical.MedicalHistory, error)
	FindByDiagnosis(ctx context.Context, diagnosis string, pageInput p.PageInput) (p.Page[medical.MedicalHistory], error)

	ExistsByID(ctx context.Context, medicalHistoryID valueobject.MedHistoryID) (bool, error)
	ExistsByPetAndDate(ctx context.Context, petID valueobject.PetID, date time.Time) (bool, error)

	Save(ctx context.Context, medHistory *medical.MedicalHistory) error
	Update(ctx context.Context, medHistory *medical.MedicalHistory) error
	SoftDelete(ctx context.Context, medicalHistoryID valueobject.MedHistoryID) error
	HardDelete(ctx context.Context, medicalHistoryID valueobject.MedHistoryID) error

	CountByPetID(ctx context.Context, petID valueobject.PetID) (int64, error)
	CountByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID) (int64, error)
	CountByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (int64, error)
	CountByDateRange(ctx context.Context, startDate, endDate time.Time) (int64, error)
}
