package dto

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/application/usecase"
)

// @Description Represents the request body for updating a user's profile.
type UpdateProfileRequest struct {
	// A brief biography of the userDomain. (optional, max 500 characters)
	Bio *string `json:"bio" validate:"min=0,max=500"`
	// The URL of the user's profile photo. (optional, must be a valid URL)
	PhotoURL *string `json:"photo_url" validate:"omitempty,url"`
}

func (request *UpdateProfileRequest) ToProfileUpdateDTO(userID valueobject.UserID) usecase.UpdateProfileData {
	updateData := usecase.UpdateProfileData{
		UserID:     userID,
		Bio:        request.Bio,
		ProfilePic: request.PhotoURL,
	}
	return updateData
}
