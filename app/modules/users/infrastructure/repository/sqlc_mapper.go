package repositoryimpl

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func sqlcUseWithJoin(row any) user.User {
	switch v := row.(type) {
	case sqlc.FindUserByEmailRow:
		return sqlcRowWithJoinToEntity(v)
	case sqlc.FindUserByIDRow:
		return sqlcRowWithJoinToEntity(v)
	case sqlc.FindUserByPhoneNumberRow:
		return sqlcRowWithJoinToEntity(v)
	case sqlc.User:
		return sqlcRowToEntity(v)
	default:
		return user.User{}
	}
}

func sqlcRowWithJoinToEntity(row any) user.User {
	var (
		id          int32
		email       pgtype.Text
		role        models.UserRole
		status      models.UserStatus
		phoneNumber pgtype.Text
		password    pgtype.Text
		lastLogin   pgtype.Timestamptz
		createdAt   pgtype.Timestamptz
		customerID  pgtype.Int4
		employeeID  pgtype.Int4
	)

	switch v := row.(type) {
	case sqlc.FindUserByEmailRow:
		id = v.ID
		email = v.Email
		role = v.Role
		status = v.Status
		phoneNumber = v.PhoneNumber
		password = v.Password
		lastLogin = v.LastLogin
		createdAt = v.CreatedAt
		customerID = v.CustomerID
		employeeID = v.EmployeeID
	case sqlc.FindUserByIDRow:
		id = v.ID
		email = v.Email
		role = v.Role
		status = v.Status
		phoneNumber = v.PhoneNumber
		password = v.Password
		lastLogin = v.LastLogin
		createdAt = v.CreatedAt
		customerID = v.CustomerID
		employeeID = v.EmployeeID
	case sqlc.FindUserByPhoneNumberRow:
		id = v.ID
		email = v.Email
		role = v.Role
		status = v.Status
		phoneNumber = v.PhoneNumber
		password = v.Password
		lastLogin = v.LastLogin
		createdAt = v.CreatedAt
		customerID = v.CustomerID
		employeeID = v.EmployeeID
	default:
		return user.User{}
	}

	emailVO := valueobject.NewEmailNoErr(email.String)
	userBuilder := user.NewUserBuilder()

	if customerID.Valid {
		customerIDVO := valueobject.NewCustomerID(uint(customerID.Int32))
		userBuilder.WithCustomerID(&customerIDVO)
	}

	if employeeID.Valid {
		employeeIDVO := valueobject.NewEmployeeID(uint(employeeID.Int32))
		userBuilder.WithEmployeeID(&employeeIDVO)
	}

	if phoneNumber.Valid && phoneNumber.String != "" {
		phoneNumberVO := valueobject.NewPhoneNumberNoErr(phoneNumber.String)
		userBuilder.WithPhoneNumber(&phoneNumberVO)
	}

	userBuilder.WithPassword(password.String)
	userBuilder.WithLastLoginAt(lastLogin.Time)
	userBuilder.WithTimeStamps(createdAt.Time, time.Time{})
	userBuilder.WithID(valueobject.NewUserID(uint(id)))
	userBuilder.WithEmail(emailVO)
	userBuilder.WithRole(enum.UserRole(string(role)))
	userBuilder.WithStatus(enum.UserStatus(string(status)))

	userEntity := userBuilder.Build()
	return *userEntity
}

func sqlcRowToEntity(sqlRow sqlc.User) user.User {
	userBuilder := user.NewUserBuilder()

	if sqlRow.PhoneNumber.Valid && sqlRow.PhoneNumber.String != "" {
		phoneNumber := valueobject.NewPhoneNumberNoErr(sqlRow.PhoneNumber.String)
		userBuilder.WithPhoneNumber(&phoneNumber)
	}

	if sqlRow.Password.Valid {
		userBuilder.WithPassword(sqlRow.Password.String)
	}

	if sqlRow.LastLogin.Valid {
		userBuilder.WithLastLoginAt(sqlRow.LastLogin.Time)
	}

	userBuilder.WithID(valueobject.NewUserID(uint(sqlRow.ID)))
	userBuilder.WithRole(enum.UserRole(string(sqlRow.Role)))
	userBuilder.WithStatus(enum.UserStatus(string(sqlRow.Status)))

	userEntity := userBuilder.Build()
	return *userEntity
}

func sqlcRowsToEntities(sqlRows []sqlc.User) ([]user.User, error) {
	users := make([]user.User, len(sqlRows))
	for i, sqlRow := range sqlRows {
		users[i] = sqlcRowToEntity(sqlRow)
	}

	return users, nil
}

func entityToCreateParams(user *user.User) *sqlc.CreateUserParams {
	return &sqlc.CreateUserParams{
		Email:       pgtype.Text{String: user.Email().String(), Valid: user.Email().String() != ""},
		PhoneNumber: pgtype.Text{String: user.PhoneNumber().String(), Valid: user.PhoneNumber().String() != ""},
		Password:    pgtype.Text{String: user.HashedPassword(), Valid: user.HashedPassword() != ""},
		Role:        models.UserRole(user.Role().String()),
		Status:      models.UserStatus(user.Status().String()),
	}
}

func entityToUpdateParams(user *user.User) *sqlc.UpdateUserParams {
	return &sqlc.UpdateUserParams{
		ID:          int32(user.ID().Value()),
		Email:       pgtype.Text{String: user.Email().String(), Valid: true},
		PhoneNumber: pgtype.Text{String: user.PhoneNumber().String(), Valid: true},
		Password:    pgtype.Text{String: user.HashedPassword(), Valid: true},
		Status:      models.UserStatus(user.Status().String()),
		Role:        models.UserRole(user.Role()),
	}
}
