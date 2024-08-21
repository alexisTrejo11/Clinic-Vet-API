package mappers

import (
	dtos "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func MapSignupDTOToParams(userSignupDTO dtos.UserSignupDTO, userId int32) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Name:   userSignupDTO.Name + " " + userSignupDTO.LastName,
		Email:  userSignupDTO.Email,
		UserID: pgtype.Int4{userId, true},
		Role:   "Common-Owner",
	}
}
