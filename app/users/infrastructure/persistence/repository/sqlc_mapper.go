package sqlcUserRepo

import (
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

func MapUserFromSQLC(sqlRow sqlc.User) (*userDomain.User, error) {
	userId, err := userDomain.NewUserId(sqlRow.ID)
	if err != nil {
		return nil, err
	}

	email, err := userDomain.NewEmail(sqlRow.Email.String)
	if err != nil {
		return nil, err
	}

	phone, err := userDomain.NewPhoneNumber(sqlRow.PhoneNumber.String)
	if err != nil {
		return nil, err
	}

	roleStr, _ := sqlRow.Role.(string)
	role := userDomain.UserRoleFromString(roleStr)
	statusStr, _ := sqlRow.Status.(string)
	status := userDomain.UserStatusFromString(statusStr)

	user, err := userDomain.NewUser(
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

func MapUsersFromSQLC(sqlRows []sqlc.User) ([]userDomain.User, error) {
	users := make([]userDomain.User, 0, len(sqlRows))
	for i, sqlRow := range sqlRows {
		user, err := MapUserFromSQLC(sqlRow)
		if err != nil {
			return nil, err
		}
		users[i] = *user
	}

	return users, nil
}
