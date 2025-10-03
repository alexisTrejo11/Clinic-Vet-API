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

func (r *SqlcUserRepository) WithJoinToEntity(row any) user.User {
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

	return *user.NewUserBuilder().
		WithEmail(valueobject.NewEmailNoErr(email.String)).
		WithCustomerID(r.pgMap.PgInt4.ToCustomerIDPtr(customerID)).
		WithEmployeeID(r.pgMap.PgInt4.ToEmployeeIDPtr(employeeID)).
		WithPhoneNumber(r.pgMap.PgText.ToPhoneNumberPtr(phoneNumber)).
		WithPassword(password.String).
		WithLastLoginAt(lastLogin.Time).
		WithTimeStamps(createdAt.Time, time.Time{}).
		WithID(valueobject.NewUserID(uint(id))).
		WithRole(enum.UserRole(string(role))).
		WithStatus(enum.UserStatus(string(status))).
		Build()
}

func (r *SqlcUserRepository) ToEntity(sqlRow sqlc.User) user.User {
	id := valueobject.NewUserID(uint(sqlRow.ID))
	role := enum.UserRole(string(sqlRow.Role))
	phoneNumber := r.pgMap.PgText.ToPhoneNumberPtr(sqlRow.PhoneNumber)
	userStatus := enum.UserStatus(string(sqlRow.Status))

	return *user.NewUserBuilder().
		WithID(id).
		WithPhoneNumber(phoneNumber).
		WithPassword(sqlRow.Password.String).
		WithLastLoginAt(sqlRow.LastLogin.Time).
		WithRole(role).
		WithStatus(userStatus).
		Build()
}

func (r *SqlcUserRepository) ToEntities(sqlRows []sqlc.User) []user.User {
	users := make([]user.User, len(sqlRows))
	for i, sqlRow := range sqlRows {
		users[i] = r.ToEntity(sqlRow)
	}

	return users
}

func (r *SqlcUserRepository) toCreateParams(user user.User) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Email:       r.pgMap.PgText.FromString(user.Email().String()),
		PhoneNumber: r.pgMap.PgText.FromString(user.PhoneNumber().String()),
		Password:    r.pgMap.PgText.FromString(user.HashedPassword()),
		Role:        models.UserRole(user.Role().String()),
		Status:      models.UserStatus(user.Status().String()),
	}
}

func (r *SqlcUserRepository) toUpdateParams(user user.User) sqlc.UpdateUserParams {
	return sqlc.UpdateUserParams{
		ID:          int32(user.ID().Value()),
		Email:       r.pgMap.PgText.FromString(user.Email().String()),
		PhoneNumber: r.pgMap.PgText.FromString(user.PhoneNumber().String()),
		Password:    r.pgMap.PgText.FromString(user.HashedPassword()),
		Status:      models.UserStatus(user.Status().String()),
		Role:        models.UserRole(user.Role()),
	}
}
