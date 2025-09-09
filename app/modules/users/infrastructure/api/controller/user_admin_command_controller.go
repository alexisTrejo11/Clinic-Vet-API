package controller

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/application/usecase/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/infrastructure/api/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	ginUtils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/gin_utils"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
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
func (controller *UserAdminController) CreateUser(c *gin.Context) {
	var requestData dto.CreateUserRequest
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
		response.ApplicationError(c, err)
		return
	}

	response.Created(c, result)
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
func (controller *UserAdminController) BanUser(c *gin.Context) {
	userID, err := ginUtils.ParseParamToEntityID(c, "id", "user")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
	}

	command := command.ChangeUserStatusCommand{
		UserID: userID.(valueobject.UserID),
		Status: enum.UserStatusBanned,
		CTX:    context.Background(),
	}

	result := controller.commandBus.Execute(command)
	if !result.IsSuccess {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, result.Message)
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
func (controller *UserAdminController) UnBanUser(c *gin.Context) {
	userID, err := ginUtils.ParseParamToEntityID(c, "id", "user")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command := command.ChangeUserStatusCommand{
		UserID: userID.(valueobject.UserID),
		Status: enum.UserStatusActive,
		CTX:    context.Background(),
	}

	result := controller.commandBus.Execute(command)
	if !result.IsSuccess {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, result.Message)
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
func (controller *UserAdminController) DeleteUser(c *gin.Context) {
	id, err := ginUtils.ParseParamToInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	deleteUserCommand, err := command.NewDeleteUserCommand(c.Request.Context(), id, true)
	if err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	result := controller.commandBus.Execute(deleteUserCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, result.Message)
}
