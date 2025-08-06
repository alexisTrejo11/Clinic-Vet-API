package userCommand

import (
	"time"

	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type CreateUserCommand struct {
	Email       string
	Password    string
	FirstName   string
	LastName    string
	Gender      string
	Phone       string
	Address     string
	Role        string
	Status      string
	DateOfBirth time.Time
	Profile     CreateProfile
}

type CreateProfile struct {
	FirstName   string
	LastName    string
	Gender      string
	ProfilePic  string
	Bio         string
	DateOfBirth time.Time
	Address     string
}

type UpdateProfileCommand struct {
	UserId      userDomain.UserId
	FirstName   *string
	LastName    *string
	Gender      *string
	ProfilePic  *string
	Bio         *string
	DateOfBirth *time.Time
	Address     *string
}
