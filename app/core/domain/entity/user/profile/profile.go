// Package profile contains the Profile entity definition.
package profile

import (
	"time"

	"clinic-vet-api/app/core/domain/valueobject"
)

type Profile struct {
	ID             uint
	UserID         valueobject.UserID
	OwnerID        *int
	VeterinarianID *int
	PhotoURL       string
	Bio            string
	DateOfBirth    *time.Time
	JoinedAt       time.Time
}
