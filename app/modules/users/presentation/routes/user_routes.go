package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/users/presentation/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(appGroup *gin.RouterGroup, controller *controller.UserAdminController, authMiddleware *middleware.AuthMiddleware) {
	// User Query Routes
	userGroup := appGroup.Group("/users")
	//userGroup.Use(authMiddleware.Authenticate(), authMiddleware.RequireAnyRole("admin", "superadmin"))
	userGroup.GET("", controller.SearchUsers)
	userGroup.GET("/:id", controller.GetUserByID)
	userGroup.GET("/email/:email", controller.GetUserByEmail)
	userGroup.GET("/by-role/:role", controller.FindByRole)

	userGroup.POST("", controller.CreateUser)
	userGroup.PATCH("/:id/ban", controller.BanUser)
	userGroup.PATCH("/:id/unban", controller.UnbanUser)
}

func ProfileRoutes(appGroup *gin.RouterGroup, profileController *controller.ProfileController, authMiddleware *middleware.AuthMiddleware) {
	// Profile Routes
	profileGroup := appGroup.Group("/users/profile")
	profileGroup.Use(authMiddleware.Authenticate())

	profileGroup.GET("", profileController.GetUserProfile)
	profileGroup.PATCH("", profileController.UpdateUserProfile)
}
