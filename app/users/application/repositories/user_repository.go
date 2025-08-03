package userRepository

import (
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type UserRepository interface {
	Save(user *userDomain.User) error
	FindByEmail(email string) (*userDomain.User, error)
	FindByPhone(email string) (*userDomain.User, error)
	FindByID(id int) (*userDomain.User, error)
	ActivateUser(id int) error
	Delete(id int) error
}
