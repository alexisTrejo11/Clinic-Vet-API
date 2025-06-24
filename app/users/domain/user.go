package userDomain

import (
	"time"

	userEnums "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/enum"
	valueObjects "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/valueobjects"
)

type User struct {
	ID          uint
	Email       valueObjects.Email
	PhoneNumber valueObjects.PhoneNumber
	Password    string
	Role        userEnums.UserRole
	IsActive    bool
	LastLoginAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
