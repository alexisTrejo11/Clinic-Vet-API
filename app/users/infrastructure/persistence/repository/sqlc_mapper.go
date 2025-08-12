package sqlcUserRepo

import (
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

func MapUserFromSQLC(sqlRow sqlc.User) (*user.User, error) {
	userId, err := user.NewUserId(sqlRow.ID)
	if err != nil {
		return nil, err
	}

	email, err := user.NewEmail(sqlRow.Email.String)
	if err != nil {
		return nil, err
	}

	phone, err := user.NewPhoneNumber(sqlRow.PhoneNumber.String)
	if err != nil {
		return nil, err
	}

	roleVal, _ := sqlRow.Role.Value()
	roleStr, _ := roleVal.(string)
	role := user.UserRoleFromString(roleStr)

	statusVal, _ := sqlRow.Status.Value()
	statusStr, _ := statusVal.(string)
	status := user.UserStatusFromString(statusStr)

	user, err := user.NewUser(
		userId,
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

func MapUsersFromSQLC(sqlRows []sqlc.User) ([]user.User, error) {
	users := make([]user.User, 0, len(sqlRows))
	for i, sqlRow := range sqlRows {
		user, err := MapUserFromSQLC(sqlRow)
		if err != nil {
			return nil, err
		}
		users[i] = *user
	}

	return users, nil
}
