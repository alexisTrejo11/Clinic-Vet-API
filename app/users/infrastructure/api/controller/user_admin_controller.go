package userController

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	userApplication "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application"
	userCommand "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/command"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAdminController struct {
	validator  *validator.Validate
	dispatcher *userApplication.CommandDispatcher
}

func NewUserAdminController(validator *validator.Validate, dispatcher *userApplication.CommandDispatcher) *UserAdminController {
	return &UserAdminController{
		validator:  validator,
		dispatcher: dispatcher,
	}
}

func (c *UserAdminController) GetUserById(ctx *gin.Context) {
	userId, err := shared.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	apiResponse.Success(ctx, gin.H{"user_id": userId})
}

func (c *UserAdminController) SearchUsers(ctx *gin.Context) {
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
	idInt, err := shared.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
	}

	userId, _ := user.NewUserId(idInt)
	command := userCommand.ChangeUserStatusCommand{
		UserId: userId,
		Status: user.UserStatusBanned,
		CTX:    context.Background(),
	}

	result := c.dispatcher.Dispatch(command)
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, result.Message)
}

func (c *UserAdminController) UnBanUser(ctx *gin.Context) {
	idInt, err := shared.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
	}

	userId, _ := user.NewUserId(idInt)
	command := userCommand.ChangeUserStatusCommand{
		UserId: userId,
		Status: user.UserStatusActive,
		CTX:    context.Background(),
	}

	result := c.dispatcher.Dispatch(command)
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, result.Message)
}

func (c *UserAdminController) DeleteUser(ctx *gin.Context) {
	idInt, err := shared.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
	}

	userId, _ := user.NewUserId(idInt)
	command := userCommand.DeleteUserCommand{
		UserId:     userId,
		SoftDelete: true,
		CTX:        context.Background(),
	}

	result := c.dispatcher.Dispatch(command)
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, result.Message)
}
