package sqlcVetRepo

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

func SqlcVetToDomain(sql sqlc.Veterinarian) *vetDomain.Veterinarian {
	name, _ := shared.NewPersonName(sql.FirstName, sql.LastName)

	var isActive bool
	if sql.IsActive.Valid {
		isActive = sql.IsActive.Bool
	}

	var userID *uint
	if sql.UserID.Valid {
		uid := uint(sql.UserID.Int32)
		userID = &uid
	}

	var createdAt, updatedAt time.Time
	if sql.CreatedAt.Valid {
		createdAt = sql.CreatedAt.Time
	}
	if sql.UpdatedAt.Valid {
		updatedAt = sql.UpdatedAt.Time
	}

	return &vetDomain.Veterinarian{
		ID:               uint(sql.ID),
		Name:             name,
		Photo:            sql.Photo,
		LicenseNumber:    sql.LicenseNumber,
		Specialty:        vetDomain.VetSpecialtyFromString(string(sql.Speciality)),
		YearsExperience:  uint(sql.YearsOfExperience),
		ConsultationFee:  nil,
		IsActive:         isActive,
		UserID:           userID,
		WorkDaysSchedule: []vetDomain.WorkDaySchedule{},
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
	}
}
