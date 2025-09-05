package user

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type Profile struct {
	UserID         valueobject.UserID
	OwnerID        *int
	VeterinarianID *int
	PhotoURL       string
	Bio            string
	Address        *Address
	DateOfBirth    *time.Time
	JoinedAt       time.Time
}
