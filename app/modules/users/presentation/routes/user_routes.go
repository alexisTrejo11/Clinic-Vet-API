package routes

import (
	"clinic-vet-api/app/modules/users/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, controller *controller.UserAdminController) {
	// User Query Routes
	path := "api/v2/admin/users"
	router.GET(path, controller.SearchUsers)
	router.GET(path+"/:id", controller.GetUserByID)
	router.POST(path, controller.CreateUser)
	router.PATCH(path+"/:id/ban", controller.BanUser)
	router.PATCH(path+"/:id/unban", controller.UnBanUser)
}

func ProfileRoutes(router *gin.Engine, profileController *controller.ProfileController) {
	// Profile Routes
	path := "api/v2/users/profiles"

	router.GET(path, profileController.GetUserProfile)
	router.PATCH(path, profileController.UpdateUserProfile)
}
