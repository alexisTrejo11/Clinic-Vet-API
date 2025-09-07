// Package profile contains the Profile entity definition.
package profile

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user/address"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type Profile struct {
	UserID         valueobject.UserID
	OwnerID        *int
	VeterinarianID *int
	PhotoURL       string
	Bio            string
	Address        *address.Address
	DateOfBirth    *time.Time
	JoinedAt       time.Time
}
