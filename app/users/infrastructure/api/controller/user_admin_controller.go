package userController

import (
	userApplication "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application"
	userCommand "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/command"
	userDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAdminController struct {
	validator  *validator.Validate
	dispatcher *userApplication.CommandDispatcher
}

func NewUserAdminController(validator *validator.Validate) *UserAdminController {
	return &UserAdminController{
		validator: validator,
	}
}

func (c *UserAdminController) CreateUser(ctx *gin.Context) {
	var request userDtos.CreateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.validator.Struct(request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Handle missing fields fields
	command := userCommand.CreateUserCommand{
		Email:          request.Email,
		Password:       request.Password,
		PhoneNumber:    request.PhoneNumber,
		Role:           request.Role,
		Address:        request.Address,
		OwnerId:        request.OwnerId,
		VeterinarianId: request.VeterinarianId,
	}

	result := c.dispatcher.Dispatch(command)
	if !result.IsSuccess {
		ctx.JSON(500, gin.H{"error": result.Message})
		return
	}

	ctx.JSON(201, gin.H{"message": result.Message, "id": result.Id})
}

func (c *UserAdminController) BanUser(ctx *gin.Context) {
	// Implementation for retrieving a user
}

func (c *UserAdminController) UnBanUser(ctx *gin.Context) {
	// Implementation for unbanning a user
}

func (c *UserAdminController) DeleteUser(ctx *gin.Context) {
	// Implementation for deleting a user
}
