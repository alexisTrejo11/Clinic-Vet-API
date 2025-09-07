package controller

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/application/usecase"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/infrastructure/api/dto"
	authError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/auth"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/alexisTrejo11/Clinic-Vet-API/middleware"
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
	idInt, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		authError.UnauthorizedCTXError()
		return
	}

	userID, _ := valueobject.NewUserID(idInt)
	profile, err := controller.useCases.GetUserProfile(context.Background(), userID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, profile)
}

func (controller *ProfileController) UpdateUserProfile(c *gin.Context) {
	idInt, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		authError.UnauthorizedCTXError()
		return
	}

	var requestData dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		response.BadRequest(c, err)
		return
	}

	profileUpdateData := requestData.ToProfileUpdateDTO(idInt)
	if err := controller.useCases.UpdateProfileUseCase(context.Background(), profileUpdateData); err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.NoContent(c)
}
