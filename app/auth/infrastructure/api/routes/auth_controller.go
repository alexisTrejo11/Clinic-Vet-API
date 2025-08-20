package authRoutes

import (
	authController "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(app *gin.Engine, authController authController.AuthController) {
	authClientV2 := app.Group("/v2/api/auth")
	authClientV2.POST("/signup", authController.Signup)
	authClientV2.POST("/login", authController.Login)
	authClientV2.DELETE("/logout", authController.Logout)
	authClientV2.DELETE("/loug-all", authController.LogoutAll)

}
