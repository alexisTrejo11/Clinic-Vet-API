package controller

import (
	"errors"

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

func (c *UserQueryController) GetUserByEmail(ctx *gin.Context) {
}

func (c *UserQueryController) GetUserByPhone(ctx *gin.Context) {
	phone := ctx.Param("phone")
	if phone == "" {
		apiResponse.RequestURLParamError(ctx, errors.New("phone number cannot be empty"), "phone", phone)
		return
	}

	apiResponse.Success(ctx, gin.H{"phone": phone})
}
