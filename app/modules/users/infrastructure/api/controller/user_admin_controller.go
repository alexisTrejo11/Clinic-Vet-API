package controller

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/application/usecase"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/application/usecase/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAdminController struct {
	validator  *validator.Validate
	dispatcher *usecase.CommandDispatcher
}

func NewUserAdminController(validator *validator.Validate, dispatcher *usecase.CommandDispatcher) *UserAdminController {
	return &UserAdminController{
		validator:  validator,
		dispatcher: dispatcher,
	}
}

// GetUserByID retrieves a user by their ID.
// @Summary Get a user by ID
// @Description Retrieves a single user record by their unique ID.
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.APIResponse{data=object} "User found"
// @Failure 400 {object} response.APIResponse "Invalid URL parameter"
// @Router /v1/users/{id} [get]
func (c *UserAdminController) GetUserByID(ctx *gin.Context) {
	userID, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		response.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	response.Success(ctx, gin.H{"user_id": userID})
}

// CreateUser creates a new user.
// @Summary Create a new user
// @Description Creates a new user record with the provided data.
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User creation request"
// @Success 201 {object} response.APIResponse{message=string,id=int} "User created successfully"
// @Failure 400 {object} response.APIResponse "Invalid request body or validation error"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /v1/users [post]
func (c *UserAdminController) CreateUser(ctx *gin.Context) {
	var request CreateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.validator.Struct(request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Handle missing fields fields
	command := command.CreateUserCommand{
		Email:          request.Email,
		Password:       request.Password,
		PhoneNumber:    request.PhoneNumber,
		Role:           request.Role,
		Address:        request.Address,
		OwnerID:        request.OwnerID,
		VeterinarianID: request.VeterinarianID,
	}

	result := c.dispatcher.Dispatch(command)
	if !result.IsSuccess {
		ctx.JSON(500, gin.H{"error": result.Message})
		return
	}

	ctx.JSON(201, gin.H{"message": result.Message, "id": result.ID})
}

// BanUser bans a user..
// @Summary Ban a user
// @Description Bans a user by setting their status to 'banned'.
// @Tags admin
// @Param id path int true "User ID"
// @Success 200 {object} response.APIResponse{data=string} "User banned successfully"
// @Failure 400 {object} response.APIResponse "Invalid URL parameter"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /v1/admin/users/{id}/ban [post]
func (c *UserAdminController) BanUser(ctx *gin.Context) {
	userID, err := shared.ParseParamToEntityID(ctx, "id", "user")
	if err != nil {
		response.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
	}

	command := command.ChangeUserStatusCommand{
		UserID: userID.(valueobject.UserID),
		Status: enum.UserStatusBanned,
		CTX:    context.Background(),
	}

	result := c.dispatcher.Dispatch(command)
	if !result.IsSuccess {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, result.Message)
}

// UnBanUser unbans a user.
// @Summary Unban a user
// @Description Unbans a user by setting their status to 'active'.
// @Tags admin
// @Param id path int true "User ID"
// @Success 200 {object} response.APIResponse{data=string} "User unbanned successfully"
// @Failure 400 {object} response.APIResponse "Invalid URL parameter"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /v1/admin/users/{id}/unban [post]
func (c *UserAdminController) UnBanUser(ctx *gin.Context) {
	userID, err := shared.ParseParamToEntityID(ctx, "id", "user")
	if err != nil {
		response.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
	}

	command := command.ChangeUserStatusCommand{
		UserID: userID.(valueobject.UserID),
		Status: enum.UserStatusActive,
		CTX:    context.Background(),
	}

	result := c.dispatcher.Dispatch(command)
	if !result.IsSuccess {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, result.Message)
}

// DeleteUser soft deletes a user.
// @Summary Soft delete a user
// @Desc
// @Tags admin
// @Param id path int true "User ID"
// @Success 200 {object} response.APIResponse{data=string} "User soft deleted successfully"
// @Failure 400 {object} response.APIResponse "Invalid URL parameter"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /v1/admin/users/{id} [delete]
func (c *UserAdminController) DeleteUser(ctx *gin.Context) {
	id, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		response.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	userID, _ := valueobject.NewUserID(id)
	command := command.DeleteUserCommand{
		UserID:     userID,
		SoftDelete: true,
		CTX:        context.Background(),
	}

	result := c.dispatcher.Dispatch(command)
	if !result.IsSuccess {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, result.Message)
}

func (c *UserAdminController) SearchUsers(ctx *gin.Context) {
}
