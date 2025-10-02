// Package controller implements all the controller for auth module
package controller

import (
	"clinic-vet-api/app/middleware"
	service "clinic-vet-api/app/modules/account/auth/local/application"
	"clinic-vet-api/app/modules/account/auth/local/presentation/dto"
	sc "clinic-vet-api/app/modules/account/auth/session/application/command"
	autherror "clinic-vet-api/app/shared/error/auth"
	ginutils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	validator *validator.Validate
	service   service.AuthApplicationFacade
}

func NewAuthController(validator *validator.Validate, service service.AuthApplicationFacade) *AuthController {
	return &AuthController{
		validator: validator,
		service:   service,
	}
}

func (ctrl *AuthController) CustomerSignup(c *gin.Context) {
	var reuqestBodyData dto.CustomerRequestSingup
	if err := ginutils.ShouldBindAndValidateBody(c, &reuqestBodyData, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command, err := reuqestBodyData.ToCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.service.RegisterCustomer(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result.ID(), "User")
}

func (ctrl *AuthController) EmployeeSignup(c *gin.Context) {
	var requestBodyData dto.EmployeeRequestRegister
	if err := ginutils.ShouldBindAndValidateBody(c, &requestBodyData, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command, err := requestBodyData.ToCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	createResult := ctrl.service.RegisterEmployee(c.Request.Context(), command)
	if !createResult.IsSuccess() {
		response.ApplicationError(c, createResult.Error())
		return
	}

	response.Success(c, nil, createResult.Message())
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var requestlogin dto.RequestLogin
	if err := ginutils.ShouldBindAndValidateBody(c, &requestlogin, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	loginMetadata := ginutils.NewLoginMetadata(c)
	loginCommand, err := requestlogin.ToCommand(loginMetadata)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.service.Login(c.Request.Context(), loginCommand)
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
	if err := ginutils.ShouldBindAndValidateBody(c, &requestLogout, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command, err := requestLogout.ToCommand(id.Value())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.service.Logout(c.Request.Context(), command)
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

	command := sc.NewRevokeAllUserSessionsCommand(userID.Value())
	result := ctrl.service.LogoutAll(c.Request.Context(), *command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}
