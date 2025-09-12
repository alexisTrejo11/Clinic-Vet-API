// Package controller implements all the controller for auth module
package controller

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/infrastructure/bus"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/presentation/dto"
	autherror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/auth"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/alexisTrejo11/Clinic-Vet-API/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	validator *validator.Validate
	bus       *bus.AuthBus
}

func NewAuthController(validator *validator.Validate, bus *bus.AuthBus) *AuthController {
	return &AuthController{
		validator: validator,
		bus:       bus,
	}
}

func (controller *AuthController) CustomerSignup(c *gin.Context) {
	var reuqestBodyData dto.CustomerRequestSingup

	if err := c.ShouldBindBodyWithJSON(&reuqestBodyData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&reuqestBodyData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	signupCommand, err := reuqestBodyData.ToCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := controller.bus.CustomerRegister(signupCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result, "User")
}

func (controller *AuthController) Login(c *gin.Context) {
	var requestlogin *dto.RequestLogin
	if err := c.ShouldBindBodyWithJSON(&requestlogin); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&requestlogin); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	loginCommand := requestlogin.ToCommand()
	result := controller.bus.Login(*loginCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, result.Session(), result.Message())
}

func (controller *AuthController) Logout(c *gin.Context) {
	id, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	var requestLogout dto.RequestLogout
	if err := c.ShouldBindBodyWithJSON(&requestLogout); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&requestLogout); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	logoutCommand, err := requestLogout.ToCommand(c.Request.Context(), id.Value())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := controller.bus.Logout(logoutCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (controller *AuthController) LogoutAll(c *gin.Context) {
	id, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	logoutAllCommand := command.NewLogoutAllCommand(c.Request.Context(), id.Value())
	result := controller.bus.LogoutAll(*logoutAllCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}
