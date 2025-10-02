package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/account/auth/two_factor/presentation/controller"

	"github.com/gin-gonic/gin"
)

type TwoFARoutes struct {
	TwoFAAuthController *controller.TwoFAAuthController
	app                 *gin.RouterGroup
	authMiddleware      *middleware.AuthMiddleware
}

func NewTwoFARoutes(twoFAAuthController *controller.TwoFAAuthController, app *gin.RouterGroup, authMiddleware *middleware.AuthMiddleware) *TwoFARoutes {
	return &TwoFARoutes{
		TwoFAAuthController: twoFAAuthController,
		app:                 app,
		authMiddleware:      authMiddleware,
	}
}

func (r *TwoFARoutes) RegisterRoutes() {
	twoFaGroup := r.app.Group("auth/2FA")
	twoFaGroup.Use(r.authMiddleware.Authenticate())

	twoFaGroup.POST("/send-token", r.TwoFAAuthController.Send2FAToken)
	twoFaGroup.POST("/login", r.TwoFAAuthController.TwoFaLogin)
	twoFaGroup.POST("/disable", r.TwoFAAuthController.Disable2FA)
	twoFaGroup.POST("/enable", r.TwoFAAuthController.Enable2FA)
}
