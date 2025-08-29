package controller

import (
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/command"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
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

func (c *AuthController) Signup(ctx *gin.Context) {
	var singupRequest RequestSignup

	if err := ctx.ShouldBindBodyWithJSON(&singupRequest); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&singupRequest); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	signupCommand := singupRequest.ToCommand()

	result := c.authCommandBus.Dispatch(signupCommand)
	if !result.IsSuccess() {
		apiResponse.ApplicationError(ctx, result.Error())
		return
	}

	apiResponse.Created(ctx, gin.H{"success": "Signup Succesfully Proccesed an Email Will Be Send to Activate your Account"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var requestlogin RequestLogin

	if err := ctx.ShouldBindBodyWithJSON(&requestlogin); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
	}

	if err := c.validator.Struct(&requestlogin); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	loginCommand := requestlogin.ToCommand()

	result := c.authCommandBus.Dispatch(loginCommand)
	if !result.IsSuccess() {
		apiResponse.ApplicationError(ctx, result.Error())
		return
	}

	apiResponse.Success(ctx, result.Session)
}

func (c *AuthController) Logout(ctx *gin.Context) {
	var requestLogout RequestLogout

	if err := ctx.ShouldBindBodyWithJSON(&requestLogout); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&requestLogout); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	logoutCommand := requestLogout.ToCommand()

	result := c.authCommandBus.Dispatch(logoutCommand)
	if !result.IsSuccess() {
		apiResponse.ApplicationError(ctx, result.Error())
		return
	}

	apiResponse.Success(ctx, result.Message)
}

func (c *AuthController) LogoutAll(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		apiResponse.Unauthorized(ctx, fmt.Errorf("authorization header required"))
		return
	}

	command := command.LogoutAllCommand{
		UserId: 0,
		CTX:    ctx.Request.Context(),
	}

	result := c.authCommandBus.Dispatch(command)
	if !result.IsSuccess() {
		apiResponse.ApplicationError(ctx, result.Error())
		return
	}

	apiResponse.Success(ctx, result.Message())
}
