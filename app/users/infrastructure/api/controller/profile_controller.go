package userController

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
	userApplication "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api/dtos"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	useCases userApplication.ProfileUseCases
}

func NewProfileController(useCases userApplication.ProfileUseCases) *ProfileController {
	return &ProfileController{
		useCases: useCases,
	}
}

func (c *ProfileController) GetUserProfile(ctx *gin.Context) {
	userId, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	profile, err := c.useCases.GetUserProfile(ctx, userId)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, profile)
}

func (c *ProfileController) UpdateUserProfile(ctx *gin.Context) {
	var request userDtos.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		apiResponse.BadRequest(ctx, err)
		return
	}

	userIdInt, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	profileUpdateData := mapRequestToProfileUpdate(request, userIdInt)

	if err := c.useCases.UpdateProfileUseCase(ctx, profileUpdateData); err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.NoContent(ctx)
}

func mapRequestToProfileUpdate(request userDtos.UpdateProfileRequest, userId int) userApplication.ProfileUpdate {
	updateData := userApplication.ProfileUpdate{
		UserId:     userId,
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
		country := user.Country(request.Address.Country)
		if !country.IsValid() {
			country = user.Mexico
		}

		updateData.Address = &user.Address{
			Street:              request.Address.Street,
			City:                request.Address.City,
			State:               request.Address.State,
			ZipCode:             request.Address.ZipCode,
			Country:             country,
			BuildingType:        user.BuildingType(request.Address.BuildingType),
			BuildingOuterNumber: request.Address.BuildingOuterNumber,
			BuildingInnerNumber: request.Address.BuildingInnerNumber,
		}
	}

	return updateData
}
