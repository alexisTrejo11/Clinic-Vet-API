package controller

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/users/application/usecase/command"
	"clinic-vet-api/app/modules/users/infrastructure/bus"
	"clinic-vet-api/app/modules/users/presentation/dto"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAdminController struct {
	validator *validator.Validate
	bus       *bus.UserBus
}

func NewUserAdminController(validator *validator.Validate, bus *bus.UserBus) *UserAdminController {
	return &UserAdminController{
		validator: validator,
		bus:       bus,
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
func (ctrl *UserAdminController) CreateUser(c *gin.Context) {
	var requestData dto.AdminCreateUserRequest
	if err := ginUtils.ShouldBindAndValidateBody(c, &requestData, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	createUserCommand, err := requestData.ToCommand()
	if err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	result := ctrl.bus.CreateUser(c.Request.Context(), createUserCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
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
func (ctrl *UserAdminController) BanUser(c *gin.Context) {
	userID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command, err := command.NewChangeUserStatusCommand(userID, enum.UserStatusBanned.DisplayName())
	if err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	result := ctrl.bus.ChangeUserStatus(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
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
func (ctrl *UserAdminController) UnbanUser(c *gin.Context) {
	userID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command, err := command.NewChangeUserStatusCommand(userID, enum.UserStatusActive.String())
	if err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	result := ctrl.bus.ChangeUserStatus(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
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
func (ctrl *UserAdminController) DeleteUser(c *gin.Context) {
	userID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	deleteUserCommand := command.NewDeleteUserCommand(userID, true)
	result := ctrl.bus.DeleteUser(c.Request.Context(), deleteUserCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}
