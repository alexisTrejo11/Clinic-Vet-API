// Package controller implements all the controller for auth module
package controller

import (
	"clinic-vet-api/app/middleware"
	cmdBus "clinic-vet-api/app/modules/auth/application/command"
	sessionCmd "clinic-vet-api/app/modules/auth/application/command/session"
	"clinic-vet-api/app/modules/auth/presentation/dto"
	autherror "clinic-vet-api/app/shared/error/auth"
	ginutils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	validator *validator.Validate
	bus       cmdBus.AuthCommandBus
}

func NewAuthController(validator *validator.Validate, bus cmdBus.AuthCommandBus) *AuthController {
	return &AuthController{
		validator: validator,
		bus:       bus,
	}
}

func (ctrl *AuthController) CustomerSignup(c *gin.Context) {
	var reuqestBodyData dto.CustomerRequestSingup
	if err := ginutils.BindAndValidateBody(c, &reuqestBodyData, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command, err := reuqestBodyData.ToCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CustomerRegister(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result, "User")
}

func (ctrl *AuthController) EmployeeSignup(c *gin.Context) {
	var requestBodyData dto.EmployeeRequestRegister
	if err := ginutils.BindAndValidateBody(c, &requestBodyData, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command, err := requestBodyData.ToCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	createResult := ctrl.bus.StaffRegister(c.Request.Context(), command)
	if !createResult.IsSuccess() {
		response.ApplicationError(c, createResult.Error())
		return
	}

	response.Success(c, nil, createResult.Message())
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var requestlogin dto.RequestLogin
	if err := ginutils.BindAndValidateBody(c, &requestlogin, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	loginCommand := requestlogin.ToCommand()
	result := ctrl.bus.Login(c.Request.Context(), *loginCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, result.Session(), result.Message())
}

func (ctrl *AuthController) Logout(c *gin.Context) {
	id, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	var requestLogout dto.RequestLogout
	if err := ginutils.BindAndValidateBody(c, &requestLogout, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command, err := requestLogout.ToCommand(id.Value())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.RevokeUserSession(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (ctrl *AuthController) LogoutAll(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	command := sessionCmd.NewRevokeAllUserSessionsCommand(userID.Value())
	result := ctrl.bus.RevokeAllUserSessions(c.Request.Context(), *command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}
