package repository

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type UserRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewUserRepository(queries *sqlc.Queries) UserRepository {
	return UserRepositoryImpl{
		queries: queries,
	}
}

func (ur UserRepositoryImpl) CreateUser(arg sqlc.CreateUserParams) (*sqlc.CreateUserRow, error) {
	user, err := ur.queries.CreateUser(context.Background(), arg)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur UserRepositoryImpl) GetUserByID(userId int32) (*sqlc.User, error) {
	user, err := ur.queries.GetUserByID(context.Background(), userId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur UserRepositoryImpl) GetUserByEmail(email string) (*sqlc.User, error) {
	user, err := ur.queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur UserRepositoryImpl) GetUserByPhoneNumber(phone string) (*sqlc.User, error) {
	user, err := ur.queries.GetUserByPhoneNumber(context.Background(), phone)
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

func (ur UserRepositoryImpl) CheckPhoneNumberExists(phone_number string) bool {
	isPhoneExisting, _ := ur.queries.CheckPhoneNumberExists(context.Background(), phone_number)
	return isPhoneExisting
}

func (ur UserRepositoryImpl) UpdateUserLastLogin(userId int32) error {
	err := ur.queries.UpdateLastLogin(context.Background(), userId)
	return err
}
