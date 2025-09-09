// Package event contains domain events for the Clinic Vet API application.
package event

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type CreateUserEvent struct {
	UserID valueobject.UserID `json:"user_id"`
	Role   enum.UserRole      `json:"role"`

	// Personal Data
	Name      valueobject.PersonName `json:"name"`
	Gender    enum.PersonGender      `json:" gender"`
	Birthdate time.Time              `json:"birthdate"`
	Location  string                 `json:"location"`

	// Vet Data (only for vets)
	LicenseNumber *string            `json:"license_number,omitempty"`
	Specialty     *enum.VetSpecialty `json:"specialty,omitempty"`

	// Profile Data
	Bio        string `json:"bio"`
	ProfilePic string `json:"profile_pic"`
}
