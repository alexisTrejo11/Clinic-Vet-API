package repositoryimpl

import (
	"fmt"

	"clinic-vet-api/app/core/domain/entity/user"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func sqlcRowToEntity(sqlRow sqlc.User) (user.User, error) {
	userID := valueobject.NewUserID(uint(sqlRow.ID))

	email, err := valueobject.NewEmail(sqlRow.Email.String)
	if err != nil {
		return user.User{}, fmt.Errorf("invalid email: %w", err)
	}

	userRole, err := enum.ParseUserRole(string(sqlRow.Role))
	if err != nil {
		return user.User{}, fmt.Errorf("invalid user role: %w", err)
	}

	userStatus, err := enum.ParseUserStatus(string(sqlRow.Status))
	if err != nil {
		return user.User{}, fmt.Errorf("invalid user status: %w", err)
	}

	opts := []user.UserOption{
		user.WithEmail(email),
	}

	if sqlRow.PhoneNumber.Valid && sqlRow.PhoneNumber.String != "" {
		phone, err := valueobject.NewPhoneNumber(sqlRow.PhoneNumber.String)
		if err != nil {
			return user.User{}, fmt.Errorf("invalid phone number: %w", err)
		}
		opts = append(opts, user.WithPhoneNumber(&phone))
	}

	if sqlRow.Password.Valid && sqlRow.Password.String != "" {
		opts = append(opts, user.WithPassword(sqlRow.Password.String))
	}
	if sqlRow.LastLogin.Valid {
		opts = append(opts, user.WithLastLoginAt(sqlRow.LastLogin.Time))
	}
	if sqlRow.EmployeeID.Valid {
		employeeID := valueobject.NewEmployeeID(uint(sqlRow.EmployeeID.Int32))
		opts = append(opts, user.WithEmployeeID(employeeID))
	}

	if sqlRow.CreatedAt.Valid {
		opts = append(opts, user.WithJoinedAt(sqlRow.CreatedAt.Time))
	}

	userEntity, err := user.NewUser(userID, userRole, userStatus, opts...)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to create user entity: %w", err)
	}

	return userEntity, nil
}

func sqlcRowsToEntities(sqlRows []sqlc.User) ([]user.User, error) {
	users := make([]user.User, 0, len(sqlRows))
	for i, sqlRow := range sqlRows {
		user, err := sqlcRowToEntity(sqlRow)
		if err != nil {
			return nil, err
		}
		users[i] = user
	}

	return users, nil
}

func entityToCreateParams(user *user.User) (*sqlc.CreateUserParams, error) {
	var employeeID pgtype.Int4
	if user.EmployeeID() != nil {
		employeeID = pgtype.Int4{Int32: int32(user.EmployeeID().Value()), Valid: true}
	} else {
		employeeID = pgtype.Int4{Valid: false}
	}

	var customerID pgtype.Int4
	if user.CustomerID() != nil {
		customerID = pgtype.Int4{Int32: int32(user.CustomerID().Value()), Valid: true}
	} else {
		customerID = pgtype.Int4{Valid: false}
	}

	return &sqlc.CreateUserParams{
		Email:       pgtype.Text{String: user.Email().String(), Valid: true},
		PhoneNumber: pgtype.Text{String: user.PhoneNumber().String(), Valid: true},
		Password:    pgtype.Text{String: user.HashedPassword(), Valid: true},
		Role:        models.UserRole(user.Role().String()),
		Status:      models.UserStatus(user.Status().String()),
		EmployeeID:  employeeID,
		ProfileID:   pgtype.Int4{Valid: false},
		CustomerID:  customerID,
	}, nil
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
