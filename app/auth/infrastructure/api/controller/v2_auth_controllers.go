package authController

import (
	"fmt"

	authCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/command"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	validator        *validator.Validate
	singupHandler    authCmd.SignupHandler
	loginHandler     authCmd.LoginHandler
	logoutHandler    authCmd.LogoutHandler
	logoutAllHandler authCmd.LogoutAllHandler
}

func NewAuthController(
	validator *validator.Validate,
	singupHandler authCmd.SignupHandler,
	loginHandler authCmd.LoginHandler,
	logoutHandler authCmd.LogoutHandler,
	logoutAllHandler authCmd.LogoutAllHandler,
) *AuthController {
	return &AuthController{
		validator:        validator,
		singupHandler:    singupHandler,
		loginHandler:     loginHandler,
		logoutHandler:    logoutHandler,
		logoutAllHandler: logoutAllHandler,
	}
}

func (c *AuthController) Signup(ctx *gin.Context) {
	var singupRequest RequestSignup

	if err := ctx.ShouldBindBodyWithJSON(&singupRequest); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
	}

	if err := c.validator.Struct(&singupRequest); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	result := c.singupHandler.Handle(singupRequest.ToCommand())
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, result.Error)
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

	session, err := c.loginHandler.Handle(requestlogin.ToCommand())
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Created(ctx, session)
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

	if err := c.logoutHandler.Handle(*requestLogout.ToCommand()); err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.NoContent(ctx)
}

func (c *AuthController) LogoutAll(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		apiResponse.Unauthorized(ctx, fmt.Errorf("authorization header required"))
		return
	}

	command := authCmd.LogoutAllCommand{
		UserId: 0,
		CTX:    ctx.Request.Context(),
	}
	if err := c.logoutAllHandler.Handle(command); err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.NoContent(ctx)
}
