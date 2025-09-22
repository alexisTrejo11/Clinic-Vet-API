package controller

import (
	"context"

	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/users/application/usecase"
	authError "clinic-vet-api/app/shared/error/auth"
	"clinic-vet-api/app/shared/response"

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
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		authError.UnauthorizedCTXError()
		return
	}

	profile, err := controller.useCases.GetUserProfile(context.Background(), user.UserID, user.CustomerID, user.EmployeeID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, profile, "User Profile")
}
