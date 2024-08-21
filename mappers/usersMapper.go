package mappers

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/sqlc"
)

func MapSignUpDTOToParams(userSignUpDTO DTOs.UserSignUpDTO) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Name:  userSignUpDTO.Name + " " + userSignUpDTO.LastName,
		Email: userSignUpDTO.Email,
		Role:  "Common-Owner",
	}
}
