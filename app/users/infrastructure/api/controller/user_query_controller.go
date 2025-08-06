package userController

import (
	userUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserQueryController struct {
	validator *validator.Validate
	useCases  userUsecase.UserUseCases
}

func NewUserQueryController(validator *validator.Validate, useCases userUsecase.UserUseCases) *UserQueryController {
	return &UserQueryController{
		validator: validator,
		useCases:  useCases,
	}
}

func (c *UserQueryController) GetUserByID(ctx *gin.Context) {
}

func (c *UserQueryController) SearchUsers(ctx *gin.Context) {
}

func (c *UserQueryController) GetUserByEmail(ctx *gin.Context) {

}

func (c *UserQueryController) GetUserByPhone(phone string) {

}
