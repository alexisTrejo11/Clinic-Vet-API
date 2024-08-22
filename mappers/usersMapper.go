package mappers

import (
	"fmt"
	"time"

	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func MapSignUpDTOToParams(userSignUpDTO DTOs.UserSignUpDTO) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Name:  userSignUpDTO.Name + " " + userSignUpDTO.LastName,
		Email: userSignUpDTO.Email,
		Role:  "Common-Owner",
	}
}

func MapSignUpDataToCreateOwnerParams(userId int32, userSignUpDTO DTOs.UserSignUpDTO) (*sqlc.CreateOwnerParams, error) {
	birthday, err := time.Parse("2006-01-02", userSignUpDTO.Birthday)
	if err != nil {
		return nil, fmt.Errorf("invalid birthday format: %v", err)
	}

	ownerCreateArgs := sqlc.CreateOwnerParams{
		Photo:    pgtype.Text{String: userSignUpDTO.Photo, Valid: userSignUpDTO.Photo != ""}, // NULL if empty
		Name:     userSignUpDTO.Name,
		LastName: userSignUpDTO.LastName,
		Birthday: pgtype.Date{Time: birthday, Valid: true},
		UserID:   pgtype.Int4{Int32: int32(userId), Valid: true},
	}

	return &ownerCreateArgs, nil
}

func MapUserSqlcToDTO(user *sqlc.User) DTOs.UserDTO {
	return DTOs.UserDTO{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.PhoneNumber,
		Role:  user.Role,
	}
}
