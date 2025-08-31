package controller

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/application/usecase"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	useCases usecase.ProfileUseCases
}

func NewProfileController(useCases usecase.ProfileUseCases) *ProfileController {
	return &ProfileController{
		useCases: useCases,
	}
}

func (c *ProfileController) GetUserProfile(ctx *gin.Context) {
	id, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	userID, _ := valueobject.NewUserID(id)
	profile, err := c.useCases.GetUserProfile(ctx, userID)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, profile)
}

func (c *ProfileController) UpdateUserProfile(ctx *gin.Context) {
	var request UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		apiResponse.BadRequest(ctx, err)
		return
	}

	userIDInt, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	profileUpdateData := mapRequestToProfileUpdate(request, userIDInt)

	if err := c.useCases.UpdateProfileUseCase(ctx, profileUpdateData); err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.NoContent(ctx)
}

func mapRequestToProfileUpdate(request UpdateProfileRequest, id int) usecase.ProfileUpdate {
	userID, _ := valueobject.NewUserID(id)
	updateData := usecase.ProfileUpdate{
		UserID:     userID,
		Bio:        request.Bio,
		ProfilePic: request.PhotoURL,
	}

	if request.Name != nil {
		updateData.Name = &valueObjects.PersonName{
			FirstName: *request.Name,
			LastName:  *request.Name,
		}
	}

	if request.Address != nil {
		country := valueobject.Country(request.Address.Country)
		updateData.Address = &entity.Address{
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
