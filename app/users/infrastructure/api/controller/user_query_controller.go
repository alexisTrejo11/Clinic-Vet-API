package userController

import (
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserQueryController struct {
	validator *validator.Validate
}

func NewUserQueryController(validator *validator.Validate) *UserQueryController {
	return &UserQueryController{
		validator: validator,
	}
}

func (c *UserQueryController) GetUserById(ctx *gin.Context) {
	userId, err := shared.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	apiResponse.Ok(ctx, gin.H{"user_id": userId})
	return
}

func (c *UserQueryController) SearchUsers(ctx *gin.Context) {
}

func (c *UserQueryController) GetUserByEmail(ctx *gin.Context) {

}

func (c *UserQueryController) GetUserByPhone(ctx *gin.Context) {
	phone := ctx.Param("phone")
	if phone == "" {
		apiResponse.RequestURLParamError(ctx, errors.New("phone number cannot be empty"), "phone", phone)
		return
	}

	apiResponse.Ok(ctx, gin.H{"phone": phone})
	return

}
