// Package controller implements all the controller for auth module
package controller

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/command"
	autherror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/auth"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/alexisTrejo11/Clinic-Vet-API/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	validator      *validator.Validate
	authCommandBus command.AuthCommandBus
}

func NewAuthController(
	validator *validator.Validate,
	authCommandBus command.AuthCommandBus,
) *AuthController {
	return &AuthController{
		validator:      validator,
		authCommandBus: authCommandBus,
	}
}

func (controller *AuthController) Signup(c *gin.Context) {
	var singupRequest RequestSignup

	if err := c.ShouldBindBodyWithJSON(&singupRequest); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&singupRequest); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	signupCommand := singupRequest.ToCommand()
	result := controller.authCommandBus.Dispatch(signupCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result)
}

func (controller *AuthController) Login(c *gin.Context) {
	var requestlogin *RequestLogin

	if err := c.ShouldBindBodyWithJSON(&requestlogin); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&requestlogin); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	loginCommand := requestlogin.ToCommand()

	result := controller.authCommandBus.Dispatch(loginCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, result.Session)
}

func (controller *AuthController) Logout(c *gin.Context) {
	id, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	var requestLogout RequestLogout

	if err := c.ShouldBindBodyWithJSON(&requestLogout); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&requestLogout); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))

		return
	}

	logoutCommand, err := requestLogout.ToCommand(id)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := controller.authCommandBus.Dispatch(logoutCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, result.Message)
}

func (controller *AuthController) LogoutAll(c *gin.Context) {
	id, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	logoutAllCommand, err := command.NewLogoutAllCommand(id)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := controller.authCommandBus.Dispatch(logoutAllCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, result.Message())
}
