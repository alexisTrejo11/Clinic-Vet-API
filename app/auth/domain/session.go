package authDomain

import (
	"time"

	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type Session struct {
	ID           uint
	UserID       uint
	RefreshToken string
	DeviceInfo   string
	IPAddress    string
	IsActive     bool
	ExpiresAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         userDomain.User
}
