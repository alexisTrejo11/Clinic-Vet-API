// Package routes contains all authentication-related route definitions
package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/account/auth/local/presentation/controller"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	AuthController *controller.AuthController
	app            *gin.RouterGroup
	authMiddleware *middleware.AuthMiddleware
}

func NewAuthRoutes(authController *controller.AuthController, app *gin.RouterGroup, authMiddleware *middleware.AuthMiddleware) *AuthRoutes {
	return &AuthRoutes{
		AuthController: authController,
		app:            app,
		authMiddleware: authMiddleware,
	}
}

func (r *AuthRoutes) RegisterRoutes() {
	authGroup := r.app.Group("/auth")
	authGroup.POST("/register/customer", r.AuthController.CustomerSignup)
	authGroup.POST("/register/employee", r.AuthController.EmployeeSignup)
	authGroup.POST("/login", r.AuthController.Login)

}
