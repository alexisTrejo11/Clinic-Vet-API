// Package routes contains all authentication-related route definitions
package routes

import (
	"clinic-vet-api/app/modules/auth/presentation/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(app *gin.Engine, authController controller.AuthController) {
	authClientV2 := app.Group("/v2/api/auth")
	authClientV2.POST("/signup/customer", authController.CustomerSignup)
	authClientV2.POST("/signup/employee", authController.EmployeeSignup)
	authClientV2.POST("/login", authController.Login)
	authClientV2.DELETE("/logout", authController.Logout)
	authClientV2.DELETE("/loug-all", authController.LogoutAll)
}
