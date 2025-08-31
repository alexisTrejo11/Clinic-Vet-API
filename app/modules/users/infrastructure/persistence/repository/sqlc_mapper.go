package persistence

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

func MapUserFromSQLC(sqlRow sqlc.User) (*entity.User, error) {
	userID, err := valueobject.NewUserID(int(sqlRow.ID))
	if err != nil {
		return nil, err
	}

	email, err := valueobject.NewEmail(sqlRow.Email.String)
	if err != nil {
		return nil, err
	}

	phone, err := valueobject.NewPhoneNumber(sqlRow.PhoneNumber.String)
	if err != nil {
		return nil, err
	}

	roleVal, _ := sqlRow.Role.Value()
	roleStr, _ := roleVal.(string)
	role := enum.UserRoleFromString(roleStr)

	statusVal, _ := sqlRow.Status.Value()
	statusStr, _ := statusVal.(string)
	status := enum.UserStatusFromString(statusStr)

	user, err := entity.NewUser(
		userID,
		email,
		phone,
		sqlRow.Password.String,
		role,
		status,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func MapUsersFromSQLC(sqlRows []sqlc.User) ([]entity.User, error) {
	users := make([]entity.User, 0, len(sqlRows))
	for i, sqlRow := range sqlRows {
		user, err := MapUserFromSQLC(sqlRow)
		if err != nil {
			return nil, err
		}
		users[i] = *user
	}

	return users, nil
}
