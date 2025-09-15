package controller

import (
	"context"

	"clinic-vet-api/app/modules/users/application/usecase"
	"clinic-vet-api/app/modules/users/presentation/dto"
	authError "clinic-vet-api/app/shared/error/auth"
	"clinic-vet-api/app/shared/response"
	"clinic-vet-api/middleware"
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

func (controller *ProfileController) GetUserProfile(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		authError.UnauthorizedCTXError()
		return
	}

	profile, err := controller.useCases.GetUserProfile(context.Background(), userID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, profile, "User Profile")
}

func (controller *ProfileController) UpdateUserProfile(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		authError.UnauthorizedCTXError()
		return
	}

	var requestData dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		response.BadRequest(c, err)
		return
	}

	profileUpdateData := requestData.ToProfileUpdateDTO(userID)
	if err := controller.useCases.UpdateProfile(c.Request.Context(), profileUpdateData); err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Updated(c, nil, "User Profile")
}
