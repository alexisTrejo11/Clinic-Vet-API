package repositoryimpl

import (
	"fmt"

	"clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func sqlcUseWithJoin(row any) (user.User, error) {
	switch v := row.(type) {
	case sqlc.FindUserByEmailRow:
		return sqlcRowWithJoinToEntity(v)
	case sqlc.FindUserByIDRow:
		return sqlcRowWithJoinToEntity(v)
	case sqlc.FindUserByPhoneNumberRow:
		return sqlcRowWithJoinToEntity(v)
	case sqlc.User: // Para queries que solo devuelven user sin joins
		return sqlcRowToEntity(v)
	default:
		return user.User{}, fmt.Errorf("unsupported type: %T", v)
	}
}

func sqlcRowWithJoinToEntity(row any) (user.User, error) {
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
		return user.User{}, fmt.Errorf("unsupported type for join row: %T", v)
	}

	emailVO := valueobject.NewEmailNoErr(email.String)
	opts := []user.UserOption{
		user.WithEmail(emailVO),
	}

	if customerID.Valid {
		customerIDVO := valueobject.NewCustomerID(uint(customerID.Int32))
		opts = append(opts, user.WithCustomerID(customerIDVO))
	}

	if employeeID.Valid {
		employeeIDVO := valueobject.NewEmployeeID(uint(employeeID.Int32))
		opts = append(opts, user.WithEmployeeID(employeeIDVO))
	}

	if phoneNumber.Valid && phoneNumber.String != "" {
		phoneNumberVO := valueobject.NewPhoneNumberNoErr(phoneNumber.String)
		opts = append(opts, user.WithPhoneNumber(phoneNumberVO))
	}

	if password.Valid && password.String != "" {
		opts = append(opts, user.WithPassword(password.String))
	}

	if lastLogin.Valid {
		opts = append(opts, user.WithLastLoginAt(lastLogin.Time))
	}

	opts = append(opts, user.WithJoinedAt(createdAt.Time))

	userID := valueobject.NewUserID(uint(id))
	roleEnum := enum.UserRole(role)
	statusEnum := enum.UserStatus(status)

	userEntity, err := user.NewUser(userID, roleEnum, statusEnum, opts...)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to create user entity: %w", err)
	}

	return userEntity, nil
}

func sqlcRowToEntity(sqlRow sqlc.User) (user.User, error) {
	email := (valueobject.NewEmailNoErr(sqlRow.Email.String))
	opts := []user.UserOption{
		user.WithEmail(email),
	}

	if sqlRow.PhoneNumber.Valid && sqlRow.PhoneNumber.String != "" {
		phoneNumber := valueobject.NewPhoneNumberNoErr(sqlRow.PhoneNumber.String)
		opts = append(opts, user.WithPhoneNumber(phoneNumber))
	}

	if sqlRow.Password.Valid && sqlRow.Password.String != "" {
		opts = append(opts, user.WithPassword(sqlRow.Password.String))
	}

	if sqlRow.LastLogin.Valid {
		opts = append(opts, user.WithLastLoginAt(sqlRow.LastLogin.Time))
	}

	if sqlRow.CreatedAt.Valid {
		opts = append(opts, user.WithJoinedAt(sqlRow.CreatedAt.Time))
	}

	userID := valueobject.NewUserID(uint(sqlRow.ID))
	role := enum.UserRole(string(sqlRow.Role))
	status := enum.UserStatus(string(sqlRow.Status))
	userEntity, err := user.NewUser(userID, role, status, opts...)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to create user entity: %w", err)
	}

	return userEntity, nil
}

func sqlcRowsToEntities(sqlRows []sqlc.User) ([]user.User, error) {
	users := make([]user.User, len(sqlRows))
	for i, sqlRow := range sqlRows {
		user, err := sqlcRowToEntity(sqlRow)
		if err != nil {
			return nil, err
		}
		users[i] = user
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
