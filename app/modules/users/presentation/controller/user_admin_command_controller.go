package controller

import (
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/modules/users/application/usecase/command"
	"clinic-vet-api/app/modules/users/presentation/dto"
	"clinic-vet-api/app/shared/cqrs"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAdminController struct {
	validator  *validator.Validate
	commandBus cqrs.CommandBus
	queryBus   cqrs.QueryBus
}

func NewUserAdminController(validator *validator.Validate, commandBus cqrs.CommandBus, queryBus cqrs.QueryBus) *UserAdminController {
	return &UserAdminController{
		validator:  validator,
		commandBus: commandBus,
		queryBus:   queryBus,
	}
}

// CreateUser creates a new user.
// @Summary Create a new user (Admin only)
// @Description Creates a new user account with the specified role and permissions. Restricted to administrators.
// @Tags Admin Users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param user body dto.AdminCreateUserRequest true "User creation data"
// @Success 201 {object} response.APIResponse{data=valueobject.UserID} "User created successfully"
// @Success 202 {object} response.APIResponse{data=valueobject.UserID} "User created and awaiting verification"
// @Failure 400 {object} response.APIResponse "Invalid request data"
// @Failure 401 {object} response.APIResponse "Unauthorized - Admin privileges required"
// @Failure 409 {object} response.APIResponse "User already exists"
// @Failure 422 {object} response.APIResponse "Validation error"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /v1/admin/users [post]
func (controller *UserAdminController) CreateUser(c *gin.Context) {
	var requestData dto.AdminCreateUserRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(requestData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	createUserCommand, err := requestData.ToCommand(c.Request.Context())
	if err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	result := controller.commandBus.Execute(createUserCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Created(c, result.ID, "User created successfully")
}

// BanUser bans a user.
// @Summary Ban a user (Admin only)
// @Description Bans a user by setting their status to 'banned'. Restricted to administrators.
// @Tags Admin Users
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param id path string true "User ID (UUID format)"
// @Success 200 {object} response.APIResponse{data=string} "User banned successfully"
// @Failure 400 {object} response.APIResponse "Invalid user ID format"
// @Failure 401 {object} response.APIResponse "Unauthorized - Admin privileges required"
// @Failure 404 {object} response.APIResponse "User not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /v1/admin/users/{id}/ban [post]
func (controller *UserAdminController) BanUser(c *gin.Context) {
	userID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command, err := command.NewChangeUserStatusCommand(c.Request.Context(), userID, enum.UserStatusBanned.DisplayName())
	if err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	result := controller.commandBus.Execute(command)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, nil, result.Message)
}

// UnbanUser unbans a user.
// @Summary Unban a user (Admin only)
// @Description Unbans a user by setting their status to 'active'. Restricted to administrators.
// @Tags Admin Users
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param id path string true "User ID (UUID format)"
// @Success 200 {object} response.APIResponse{data=string} "User unbanned successfully"
// @Failure 400 {object} response.APIResponse "Invalid user ID format"
// @Failure 401 {object} response.APIResponse "Unauthorized - Admin privileges required"
// @Failure 404 {object} response.APIResponse "User not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /v1/admin/users/{id}/unban [post]
func (controller *UserAdminController) UnbanUser(c *gin.Context) {
	userID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command, err := command.NewChangeUserStatusCommand(c.Request.Context(), userID, enum.UserStatusActive.DisplayName())
	if err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	result := controller.commandBus.Execute(command)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, nil, result.Message)
}

// DeleteUser soft deletes a user.
// @Summary Soft delete a user (Admin only)
// @Description Performs a soft delete of a user account. Restricted to administrators.
// @Tags Admin Users
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param id path string true "User ID (UUID format)"
// @Success 200 {object} response.APIResponse{data=string} "User deleted successfully"
// @Failure 400 {object} response.APIResponse "Invalid user ID format"
// @Failure 401 {object} response.APIResponse "Unauthorized - Admin privileges required"
// @Failure 404 {object} response.APIResponse "User not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /v1/admin/users/{id} [delete]
func (controller *UserAdminController) DeleteUser(c *gin.Context) {
	userID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	deleteUserCommand := command.NewDeleteUserCommand(c.Request.Context(), userID, true)
	result := controller.commandBus.Execute(deleteUserCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, nil, result.Message)
}
