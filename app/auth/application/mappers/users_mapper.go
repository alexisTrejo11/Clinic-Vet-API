package mappers

import "example.com/at/backend/api-vet/app/container/sqlc"

type UserMappers struct {
}

func (UserMappers) MapSignUpDTOToCreateParams(name, lastName, email, phoneNumber, role string) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Name:        name + " " + lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Role:        role,
	}
}

func (UserMappers) MapSqlcEntityToDTO(user *sqlc.User) DTOs.UserDTO {
	return DTOs.UserDTO{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.PhoneNumber,
		Role:  user.Role,
	}
}
