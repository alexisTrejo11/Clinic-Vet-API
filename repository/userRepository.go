package repository

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type UserRepository interface {
	CreateUser(arg sqlc.CreateUserParams) error
	GetUser(userId int32) (*sqlc.User, error)
	UpdateUser(arg sqlc.UpdateUserParams) error
	DeleteUser(userId int32) error
	CheckEmailExists(email string) bool
}

type UserRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewUserRepository(queries *sqlc.Queries) UserRepository {
	return UserRepositoryImpl{
		queries: queries,
	}
}

func (ur UserRepositoryImpl) CreateUser(arg sqlc.CreateUserParams) error {
	_, err := ur.queries.CreateUser(context.Background(), arg)
	if err != nil {
		return err
	}

	return nil
}

func (ur UserRepositoryImpl) GetUser(userId int32) (*sqlc.User, error) {
	user, err := ur.queries.GetUserByID(context.Background(), userId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur UserRepositoryImpl) UpdateUser(arg sqlc.UpdateUserParams) error {
	err := ur.queries.UpdateUser(context.Background(), arg)
	if err != nil {
		return err
	}

	return nil
}

func (ur UserRepositoryImpl) DeleteUser(userId int32) error {
	err := ur.queries.DeleteUser(context.Background(), userId)
	if err != nil {
		return err
	}
	return nil
}

func (ur UserRepositoryImpl) CheckEmailExists(email string) bool {
	isEmailExisting, _ := ur.queries.CheckEmailExists(context.Background(), email)
	return isEmailExisting
}
