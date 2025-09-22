package repositoryimpl

import (
	"context"
	"errors"

	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/sqlc"
)

type SQLCProfileRepository struct {
	queries *sqlc.Queries
}

func NewSQLCProfileRepository(queries *sqlc.Queries) repository.ProfileRepository {
	return &SQLCProfileRepository{
		queries: queries,
	}
}

func (r *SQLCProfileRepository) GetMapByUserID(ctx context.Context, userID valueobject.UserID, role enum.UserRole) (map[string]any, error) {
	profileMap := make(map[string]any)
	if role.IsStaff() {
		sqlcRow, err := r.queries.GetUserEmployeeProfile(ctx, int32(userID.Value()))
		if err != nil {
			return nil, err
		}
		profileMap["employee_id"] = sqlcRow.ID
		profileMap["email"] = sqlcRow.Email.String
		profileMap["phone_number"] = sqlcRow.PhoneNumber.String
		profileMap["status"] = sqlcRow.Status
		profileMap["role"] = sqlcRow.Role
		profileMap["first_name"] = sqlcRow.FirstName
		profileMap["last_name"] = sqlcRow.LastName
		profileMap["photo"] = sqlcRow.Photo
		profileMap["license_number"] = sqlcRow.LicenseNumber
		profileMap["speciality"] = sqlcRow.Speciality
		profileMap["years_of_experience"] = sqlcRow.YearsOfExperience
		profileMap["is_active"] = sqlcRow.IsActive
		profileMap["user_id"] = sqlcRow.UserID.Int32

		if sqlcRow.LastLogin.Valid {
			profileMap["last_login"] = sqlcRow.LastLogin.Time
		} else {
			profileMap["last_login"] = nil
		}
		profileMap["created_at"] = sqlcRow.CreatedAt.Time
		if sqlcRow.UpdatedAt.Valid {
			profileMap["updated_at"] = sqlcRow.UpdatedAt.Time
		}

		return profileMap, nil

	} else if role == enum.UserRoleCustomer {
		sqlcRow, err := r.queries.GetUserCustomerProfile(ctx, int32(userID.Value()))
		if err != nil {
			return nil, err
		}
		profileMap["customer_id"] = sqlcRow.ID
		profileMap["email"] = sqlcRow.Email.String
		profileMap["phone_number"] = sqlcRow.PhoneNumber.String
		profileMap["status"] = sqlcRow.Status
		profileMap["role"] = sqlcRow.Role
		profileMap["first_name"] = sqlcRow.FirstName
		profileMap["last_name"] = sqlcRow.LastName
		profileMap["photo"] = sqlcRow.Photo
		profileMap["is_active"] = sqlcRow.IsActive
		profileMap["user_id"] = sqlcRow.UserID.Int32
	}
	return nil, errors.New("unsupported role for profile retrieval")
}
