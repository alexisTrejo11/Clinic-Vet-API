package dto

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user/address"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/application/usecase"
)

// @Description Represents the request body for updating a user's profile.
type UpdateProfileRequest struct {
	// The unique ID of the userDomain. (required)
	UserID int `json:"user_id" validate:"required"`
	// A brief biography of the userDomain. (optional, max 500 characters)
	Bio *string `json:"bio" validate:"min=0,max=500"`
	// The URL of the user's profile photo. (optional, must be a valid URL)
	PhotoURL *string `json:"photo_url" validate:"omitempty,url"`
	// The name of the userDomain. (optional)
	Name *string `json:"name" validate:"omitempty"`
	// The user's address. (optional)
	Address *AddressRequest `json:"address" validate:"omitempty"`
}

func (request *UpdateProfileRequest) ToProfileUpdateDTO(id int) usecase.UpdateProfileData {
	userID, _ := valueobject.NewUserID(id)
	updateData := usecase.UpdateProfileData{
		UserID:     userID,
		Bio:        request.Bio,
		ProfilePic: request.PhotoURL,
	}

	if request.Name != nil {
		updateData.Name = &valueobject.PersonName{
			FirstName: *request.Name,
			LastName:  *request.Name,
		}
	}

	if request.Address != nil {
		country := valueobject.Country(request.Address.Country)
		updateData.Address = &address.Address{
			Street:              request.Address.Street,
			City:                request.Address.City,
			State:               request.Address.State,
			ZipCode:             request.Address.ZipCode,
			Country:             country,
			BuildingType:        valueobject.BuildingType(request.Address.BuildingType),
			BuildingOuterNumber: request.Address.BuildingOuterNumber,
			BuildingInnerNumber: request.Address.BuildingInnerNumber,
		}
	}

	return updateData
}
