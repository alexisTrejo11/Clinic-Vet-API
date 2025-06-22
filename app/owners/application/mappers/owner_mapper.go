package mappers

/*
type OwnerMapper struct {
}

func (OwnerMapper) MapSignUpDataToCreateParams(userId int32, userSignUpDTO DTOs.UserSignUpDTO) (*sqlc.CreateOwnerParams, error) {
	birthday, err := time.Parse("2006-01-02", userSignUpDTO.Birthday)
	if err != nil {
		return nil, fmt.Errorf("invalid birthday format: %v", err)
	}

	genre := userSignUpDTO.Genre
	genreStr := string(genre)

	ownerCreateArgs := sqlc.CreateOwnerParams{
		Photo:    pgtype.Text{String: userSignUpDTO.Photo, Valid: userSignUpDTO.Photo != ""},
		Name:     userSignUpDTO.Name,
		LastName: userSignUpDTO.LastName,
		Genre:    pgtype.Text{String: genreStr, Valid: userSignUpDTO.Genre != ""},
		Birthday: pgtype.Date{Time: birthday, Valid: true},
		UserID:   pgtype.Int4{Int32: int32(userId), Valid: true},
	}

	return &ownerCreateArgs, nil
}

func (OwnerMapper) MapOwnerUpdateDtoToEntity(ownerUpdateDTO *DTOs.OwnerUpdateDTO, existingOwner sqlc.Owner) sqlc.UpdateOwnerParams {
	params := sqlc.UpdateOwnerParams{
		ID: ownerUpdateDTO.Id,
	}

	params.Name = coalesceString(ownerUpdateDTO.Name, existingOwner.Name)
	params.Photo = coalescePgText(ownerUpdateDTO.Photo, existingOwner.Photo)

	return params
}

func coalesceString(newVal, existingVal string) string {
	if newVal != "" {
		return newVal
	}
	return existingVal
}

func coalescePgText(newVal string, existingVal pgtype.Text) pgtype.Text {
	if newVal != "" {
		return pgtype.Text{String: newVal, Valid: true}
	}
	return existingVal
}
*/
