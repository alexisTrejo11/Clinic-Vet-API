package userController

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserCommandController struct {
	validator *validator.Validate
}

func NewUserCommandController(validator *validator.Validate) *UserCommandController {
	return &UserCommandController{
		validator: validator,
	}
}

func (c *UserCommandController) CreateUser(ctx *gin.Context) {
	// Implementation for creating a user
}

func (c *UserCommandController) UpdateUser(ctx *gin.Context) {
	// Implementation for updating a user
}

func (c *UserCommandController) BanUser(ctx *gin.Context) {
	// Implementation for retrieving a user
}

func (c *UserCommandController) UnBanUser(ctx *gin.Context) {
	// Implementation for unbanning a user
}

func (c *UserCommandController) DeleteUser(ctx *gin.Context) {
	// Implementation for deleting a user
}
