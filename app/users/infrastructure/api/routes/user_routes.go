package userRoutes

import (
	userController "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, controller *userController.UserAdminController) {
	// User Query Routes
	path := "api/v2/admin/users"
	router.GET(path, controller.SearchUsers)
	router.GET(path+"/:id", controller.GetUserById)
	router.POST(path, controller.CreateUser)
	router.PATCH(path+"/:id/ban", controller.BanUser)
	router.PATCH(path+"/:id/unban", controller.UnBanUser)
}

func ProfileRoutes(router *gin.Engine, profileController *userController.ProfileController) {
	// Profile Routes
	path := "api/v2/users/profiles"

	router.GET(path, profileController.GetUserProfile)
	router.PATCH(path, profileController.UpdateUserProfile)

}
