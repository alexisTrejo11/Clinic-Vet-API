package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/account/auth/session/presentation/controller"

	"github.com/gin-gonic/gin"
)

type SessionRoutes struct {
	AuthMiddleware *middleware.AuthMiddleware
	Controller     *controller.SessionController
	RouterGroup    *gin.RouterGroup
}

func NewSessionRoutes(
	authMiddleware *middleware.AuthMiddleware,
	controller *controller.SessionController,
	routerGroup *gin.RouterGroup,
) *SessionRoutes {
	return &SessionRoutes{
		AuthMiddleware: authMiddleware,
		Controller:     controller,
		RouterGroup:    routerGroup,
	}
}

func (r *SessionRoutes) RegisterRoutes() {
	authGroup := r.RouterGroup.Group("/auth")

	refresh := authGroup.Use(r.AuthMiddleware.OptionalAuth())
	refresh.POST("/refresh-token", r.Controller.RefreshSession)

	logoutGroup := authGroup.Group("/logout")
	logoutGroup.Use(r.AuthMiddleware.Authenticate())
	logoutGroup.DELETE("/revoke", r.Controller.RevokeToken)
	logoutGroup.DELETE("/logout-all", r.Controller.RevokeAllTokens)

}
