package userController

import (
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

func (c *UserQueryController) GetUserByID(ctx *gin.Context) {
}

func (c *UserQueryController) SearchUsers(ctx *gin.Context) {
}

func (c *UserQueryController) GetUserByEmail(ctx *gin.Context) {

}

func (c *UserQueryController) GetUserByPhone(phone string) {

}
