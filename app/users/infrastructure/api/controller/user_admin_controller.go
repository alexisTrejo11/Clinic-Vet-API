package userDomainController

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	userApplication "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application"
	userCommand "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/command"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
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

// GetUserById retrieves a user by their ID.
// @Summary Get a user by ID
// @Description Retrieves a single user record by their unique ID.
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} apiResponse.APIResponse{data=object} "User found"
// @Failure 400 {object} apiResponse.APIResponse "Invalid URL parameter"
// @Router /v1/users/{id} [get]
func (c *UserAdminController) GetUserById(ctx *gin.Context) {
	userId, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	apiResponse.Success(ctx, gin.H{"user_id": userId})
}

// CreateUser creates a new userDomain.
// @Summary Create a new user
// @Description Creates a new user record with the provided data.
// @Tags users
// @Accept json
// @Produce json
// @Param user body userDtos.CreateUserRequest true "User creation request"
// @Success 201 {object} apiResponse.APIResponse{message=string,id=int} "User created successfully"
// @Failure 400 {object} apiResponse.APIResponse "Invalid request body or validation error"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /v1/users [post]
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

// BanUser bans a userDomain.
// @Summary Ban a user
// @Description Bans a user by setting their status to 'banned'.
// @Tags admin
// @Param id path int true "User ID"
// @Success 200 {object} apiResponse.APIResponse{data=string} "User banned successfully"
// @Failure 400 {object} apiResponse.APIResponse "Invalid URL parameter"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /v1/admin/users/{id}/ban [post]
func (c *UserAdminController) BanUser(ctx *gin.Context) {
	userId, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
	}

	command := userCommand.ChangeUserStatusCommand{
		UserId: userId,
		Status: userDomain.UserStatusBanned,
		CTX:    context.Background(),
	}

	result := c.dispatcher.Dispatch(command)
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, result.Message)
}

// UnBanUser unbans a userDomain.
// @Summary Unban a user
// @Description Unbans a user by setting their status to 'active'.
// @Tags admin
// @Param id path int true "User ID"
// @Success 200 {object} apiResponse.APIResponse{data=string} "User unbanned successfully"
// @Failure 400 {object} apiResponse.APIResponse "Invalid URL parameter"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /v1/admin/users/{id}/unban [post]
func (c *UserAdminController) UnBanUser(ctx *gin.Context) {
	userId, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
	}

	command := userCommand.ChangeUserStatusCommand{
		UserId: userId,
		Status: userDomain.UserStatusActive,
		CTX:    context.Background(),
	}

	result := c.dispatcher.Dispatch(command)
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, result.Message)
}

// DeleteUser soft deletes a userDomain.
// @Summary Soft delete a user
// @Description Soft deletes a user record by marking it as deleted.
// @Tags admin
// @Param id path int true "User ID"
// @Success 200 {object} apiResponse.APIResponse{data=string} "User soft deleted successfully"
// @Failure 400 {object} apiResponse.APIResponse "Invalid URL parameter"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /v1/admin/users/{id} [delete]
func (c *UserAdminController) DeleteUser(ctx *gin.Context) {
	userId, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

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

func (c *UserAdminController) SearchUsers(ctx *gin.Context) {
}
