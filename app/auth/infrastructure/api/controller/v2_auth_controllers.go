package authController

import (
	"fmt"

	authDto "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/dtos"
	authUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/usecase"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	validator        *validator.Validate
	signupUseCase    authUsecase.SignUpUseCase
	loginUseCase     authUsecase.LoginUseCase
	logoutUseCase    authUsecase.LogoutUseCase
	logoutAllUseCase authUsecase.LogoutAllUseCase
}

func NewAuthController(
	validator *validator.Validate,
	signupUseCase authUsecase.SignUpUseCase,
	loginUseCase authUsecase.LoginUseCase,
	logoutUseCase authUsecase.LogoutUseCase,
	logoutAllUseCase authUsecase.LogoutAllUseCase,
) *AuthController {
	return &AuthController{
		validator:        validator,
		signupUseCase:    signupUseCase,
		loginUseCase:     loginUseCase,
		logoutUseCase:    logoutUseCase,
		logoutAllUseCase: logoutAllUseCase,
	}
}

func (c *AuthController) Signup(ctx *gin.Context) {
	var singup authDto.RequestSignup

	if err := ctx.ShouldBindBodyWithJSON(&singup); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
	}

	if err := c.validator.Struct(&singup); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.signupUseCase.Execute(singup); err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	apiResponse.Created(ctx, gin.H{"success": "Signup Succesfully Proccesed an Email Will Be Send to Activate your Account"})
}
func (c *AuthController) Login(ctx *gin.Context) {
	var requestlogin authDto.RequestLogin

	if err := ctx.ShouldBindBodyWithJSON(&requestlogin); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
	}

	if err := c.validator.Struct(&requestlogin); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	session, err := c.loginUseCase.Execute(requestlogin, "", "")
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	apiResponse.Created(ctx, session)
}

func (c *AuthController) Logout(ctx *gin.Context) {
	var requestLogout authDto.RequestLogout

	if err := ctx.ShouldBindBodyWithJSON(&requestLogout); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&requestLogout); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.logoutUseCase.Execute(requestLogout); err != nil {
		apiResponse.AppError(ctx, err)
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

	if err := c.logoutAllUseCase.Execute(1); err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	apiResponse.NoContent(ctx)
}
