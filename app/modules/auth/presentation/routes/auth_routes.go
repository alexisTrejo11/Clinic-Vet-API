// Package routes contains all authentication-related route definitions
package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/auth/presentation/controller"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	AuthController  *controller.AuthController
	TokenController *controller.TokenController
}

func NewAuthRoutes(authController *controller.AuthController, tokenController *controller.TokenController) *AuthRoutes {
	return &AuthRoutes{
		AuthController:  authController,
		TokenController: tokenController,
	}
}

func (r *AuthRoutes) RegisterRegistAndLoginRoutes(app *gin.RouterGroup, authMiddleware *middleware.AuthMiddleware) {
	authGroup := app.Group("/auth")
	authGroup.POST("/register/customer", r.AuthController.CustomerSignup)
	authGroup.POST("/register/employee", r.AuthController.EmployeeSignup)
	authGroup.POST("/login", r.AuthController.Login)

}

func (r *AuthRoutes) RegisterSessionRoutes(app *gin.RouterGroup, authMiddleware *middleware.AuthMiddleware) {
	authGroup := app.Group("/auth")

	refresh := authGroup.Use(authMiddleware.OptionalAuth())
	refresh.POST("/refresh-token", r.TokenController.RefreshSession)

	logoutGroup := app.Group("/auth")
	logoutGroup.Use(authMiddleware.Authenticate())
	logoutGroup.DELETE("/logout", r.AuthController.Logout)
	logoutGroup.DELETE("/logout-all", r.AuthController.LogoutAll)

	//authGroup.POST("/forgot-password", r.AuthControler.ForgotPassword)
	//authGroup.POST("/reset-password", r.AuthControler.ResetPassword)
	//authGroup.POST("/verify-email", r.AuthControler.VerifyEmail)
	//authGroup.POST("/resend-verification", r.AuthControler.ResendVerification)
}
