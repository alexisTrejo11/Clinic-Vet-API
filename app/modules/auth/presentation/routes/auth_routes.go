// Package routes contains all authentication-related route definitions
package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/auth/presentation/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(app *gin.RouterGroup, authController controller.AuthController, authMiddleware *middleware.AuthMiddleware) {
	authGroup := app.Group("/auth")
	authGroup.POST("/register/customer", authController.CustomerSignup)
	authGroup.POST("/register/employee", authController.EmployeeSignup)
	authGroup.POST("/login", authController.Login)

	logoutGroup := app.Group("/logout")
	logoutGroup.Use(authMiddleware.Authenticate())
	logoutGroup.DELETE("/logout", authController.Logout)
	logoutGroup.DELETE("/loug-all", authController.LogoutAll)

	//authGroup.POST("/refresh-token", authController.RefreshToken)
	//authGroup.POST("/forgot-password", authController.ForgotPassword)
	//authGroup.POST("/reset-password", authController.ResetPassword)
	//authGroup.POST("/verify-email", authController.VerifyEmail)
	//authGroup.POST("/resend-verification", authController.ResendVerification)
}
