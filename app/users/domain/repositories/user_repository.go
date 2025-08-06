package userRepository

import (
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type UserRepository interface {
	Save(user *userDomain.User) error
	Delete(id int) error

	GetByID(id int) (*userDomain.User, error)
	GetByEmail(email string) (*userDomain.User, error)
	GetByPhone(phone string) (*userDomain.User, error)
	ListByPhone(email string) (*userDomain.User, error)
	Search(query string) ([]*userDomain.User, error)

	ExistsByEmail(email string) (bool, error)
	ExistsByPhone(phone string) (bool, error)
	UpdateLastLogin(id int) error
}
